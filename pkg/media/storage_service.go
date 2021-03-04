package media

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/log"
	"strconv"
	"time"
)

var allowedOriginalDocumentMimeTypes = []string{
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

var allowedPreviewDocumentMimeTypes = []string{
	"application/pdf",
}

var allowedImageMimeTypes = []string{
	"image/jpeg",
	"image/png",
	"image/webp",
	"image/svg+xml",
}

const (
	documentSizeLimit = 32 << 20
	imageSizeLimit    = 15 << 20
)

type Service interface {
	GetUrl(ctx context.Context, key string) (string, error)
	GetImageLink(key string) string
	GetFile(ctx context.Context, key string) (io.ReadCloser, error)
	UploadDocumentOriginal(ctx context.Context, req *FileUploadOriginalReq) (*UploadResult, error)
	UploadDocumentPreview(ctx context.Context, req *FileUploadPreviewReq) (*UploadResult, error)
	UploadImage(ctx context.Context, req *FileUploadOriginalReq) (*UploadResult, error)
}

type S3StorageService struct {
	s3        *s3.S3
	s3manager *s3manager.Uploader
	config    *config.Config
	proxy     ImageProxyService
}

func NewStorageService(config *config.Config, proxy ImageProxyService) Service {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")}))
	return &S3StorageService{
		s3:        s3.New(sess),
		s3manager: s3manager.NewUploader(sess),
		config:    config,
		proxy:     proxy,
	}
}

func (s S3StorageService) UploadDocumentOriginal(ctx context.Context, req *FileUploadOriginalReq) (*UploadResult, error) {
	if req.Size > documentSizeLimit {
		return nil, apperror.New(apperror.ErrInvalid, "file too large", nil)
	}
	m, originalReader, err := validateMime(req.File, allowedOriginalDocumentMimeTypes)
	if err != nil {
		return nil, err
	}
	return s.upload(ctx, originalReader, map[string]*string{
		"userId":       aws.String(strconv.Itoa(req.User.Id)),
		"originalName": aws.String(req.Name),
	}, m)
}

func (s S3StorageService) UploadDocumentPreview(ctx context.Context, req *FileUploadPreviewReq) (*UploadResult, error) {
	m, originalReader, err := validateMime(req.File, allowedPreviewDocumentMimeTypes)
	if err != nil {
		return nil, err
	}
	return s.upload(ctx, originalReader, map[string]*string{
		"userId":       aws.String(strconv.Itoa(req.User.Id)),
		"originalName": aws.String(req.Name),
	}, m)
}

func (s S3StorageService) UploadImage(ctx context.Context, req *FileUploadOriginalReq) (*UploadResult, error) {
	m, originalReader, err := validateMime(req.File, allowedImageMimeTypes)
	if err != nil {
		return nil, err
	}
	return s.upload(ctx, originalReader, map[string]*string{
		"userId":       aws.String(strconv.Itoa(req.User.Id)),
		"originalName": aws.String(req.Name),
	}, m)
}

func (s S3StorageService) GetUrl(ctx context.Context, key string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
		return s.generatePresignUrl(key)
	}
}

func (s S3StorageService) GetFile(ctx context.Context, key string) (io.ReadCloser, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
		object, err := s.s3.GetObjectWithContext(ctx,
			&s3.GetObjectInput{
				Key:    aws.String(key),
				Bucket: aws.String(s.config.MediaBucket),
			})
		if err != nil {
			return nil, err
		}
		return object.Body, err
	}
}

func (s S3StorageService) GetImageLink(key string) string {
	return s.proxy.GetLink(key)
}

func (s S3StorageService) generatePresignUrl(key string) (string, error) {
	req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.config.MediaBucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}

func (s *S3StorageService) upload(ctx context.Context, stream io.Reader, metadata map[string]*string, m *mimetype.MIME) (*UploadResult, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	key := id.String() + m.Extension()
	output, err := s.s3manager.UploadWithContext(ctx, &s3manager.UploadInput{
		ACL:         aws.String(s3.ObjectCannedACLPrivate),
		Body:        stream,
		Bucket:      aws.String(s.config.MediaBucket),
		Key:         aws.String(key),
		ContentType: aws.String(m.String()),
		Metadata:    metadata,
	})
	if err != nil {
		return nil, err
	}
	log.Logger.Info("upload file completed", zap.Any("output", output))
	return &UploadResult{Key: key}, nil
}

func validateMime(r io.Reader, allowedMimes []string) (*mimetype.MIME, io.Reader, error) {
	in := make([]byte, 3072)
	n, err := io.ReadFull(r, in)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, nil, err
	}
	in = in[:n]
	m := mimetype.Detect(in)
	isAllowed := false
	for i := range allowedMimes {
		if m.Is(allowedMimes[i]) {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return nil, nil, apperror.New(
			apperror.ErrInvalid,
			fmt.Sprintf("file type not accepted: %v", m.String()),
			nil,
		)
	}

	concatHeaderAndReader := io.MultiReader(bytes.NewReader(in), r)
	return m, concatHeaderAndReader, nil
}
