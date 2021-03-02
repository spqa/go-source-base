package media

import (
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
	"time"
)

var allowedOriginalDocumentMimeTypes = []string{
	"application/msword",
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

var allowedPreviewDocumentMimeTypes = []string{
	"application/pdf",
}

type StorageService interface {
	UploadDocumentOriginal(stream io.Reader, metadata map[string]*string) (*UploadResult, error)
	UploadDocumentPreview(stream io.Reader, metadata map[string]*string) (*UploadResult, error)
}

type S3StorageService struct {
	s3        *s3.S3
	s3manager *s3manager.Uploader
	config    *config.Config
}

func NewStorageService(config *config.Config) StorageService {
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("ap-southeast-1")}))
	return &S3StorageService{
		s3manager: s3manager.NewUploader(sess),
		config:    config,
	}
}

func (s S3StorageService) GeneratePresignUrl(bucket, key string) (string, error) {
	req, _ := s.s3.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}
	return urlStr, nil
}

func (s *S3StorageService) uploadDocument(stream io.Reader, metadata map[string]*string, m *mimetype.MIME) (*UploadResult, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	key := id.String() + m.Extension()
	output, err := s.s3manager.Upload(&s3manager.UploadInput{
		ACL:         aws.String(s3.ObjectCannedACLPrivate),
		Body:        stream,
		Bucket:      aws.String(s.config.DocumentBucket),
		Key:         aws.String(key),
		ContentType: aws.String(m.String()),
		Metadata:    metadata,
	})
	fmt.Println(output)
	if err != nil {
		return nil, err
	}
	log.Logger.Info("upload file completed", zap.Any("output", output))
	return &UploadResult{Key: key}, nil
}

func (s S3StorageService) UploadDocumentOriginal(stream io.Reader, metadata map[string]*string) (*UploadResult, error) {
	m, err := mimetype.DetectReader(stream)
	if err != nil {
		return nil, err
	}
	isAllowed := false
	for i := range allowedOriginalDocumentMimeTypes {
		if m.Is(allowedOriginalDocumentMimeTypes[i]) {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return nil, apperror.New(apperror.ErrInvalid, "file type not accepted", nil)
	}
	return s.uploadDocument(stream, metadata, m)
}

func (s S3StorageService) UploadDocumentPreview(stream io.Reader, metadata map[string]*string) (*UploadResult, error) {
	m, err := mimetype.DetectReader(stream)
	if err != nil {
		return nil, err
	}
	isAllowed := false
	for i := range allowedPreviewDocumentMimeTypes {
		if m.Is(allowedOriginalDocumentMimeTypes[i]) {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return nil, apperror.New(apperror.ErrInvalid, "file type not accepted", nil)
	}
	return s.uploadDocument(stream, metadata, m)
}
