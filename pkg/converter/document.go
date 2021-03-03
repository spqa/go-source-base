package converter

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"io"
	"mcm-api/config"
	"mcm-api/pkg/apperror"
	"mcm-api/pkg/common"
	"mcm-api/pkg/log"
	"mcm-api/pkg/media"
	"mime/multipart"
	"net/http"
)

type DocumentConverter interface {
	Convert(ctx context.Context, key string, user common.LoggedInUser) (*ConvertResult, error)
}

type GotenbergDocumentConverter struct {
	cfg     *config.Config
	service media.Service
}

func NewGotenbergDocumentConverter(config *config.Config, service media.Service) DocumentConverter {
	return &GotenbergDocumentConverter{
		cfg:     config,
		service: service,
	}
}

func (r GotenbergDocumentConverter) Convert(ctx context.Context, key string, user common.LoggedInUser) (*ConvertResult, error) {
	url := fmt.Sprintf("%v/convert/office", r.cfg.ConverterService)
	file, err := r.service.GetFile(ctx, key)
	if err != nil {
		return nil, err
	}

	reader, writer := io.Pipe()
	w := multipart.NewWriter(writer)
	go func() {
		// TODO improve error handling in this goroutine
		fw, err := w.CreateFormFile("files", "file.docx")
		if err != nil {
			log.Logger.Error("create form file failed", zap.Error(err))
			return
		}
		written, err := io.Copy(fw, file)
		fmt.Println(written)
		if err != nil {
			log.Logger.Error("copy bytes to form writer failed", zap.Error(err))
		}
		defer writer.Close()
		defer w.Close()
		defer file.Close()
	}()
	response, err := http.Post(url, w.FormDataContentType(), reader)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		readAll, _ := io.ReadAll(response.Body)
		log.Logger.Error("error calling convert service", zap.ByteString("response", readAll))
		return nil, apperror.New(apperror.ErrInternal, "error calling convert service", nil)
	}
	result, err := r.service.UploadDocumentPreview(ctx, &media.FileUploadPreviewReq{
		File: response.Body,
		Name: key,
		User: user,
	})
	if err != nil {
		return nil, err
	}
	return &ConvertResult{
		Key: result.Key,
	}, nil
}
