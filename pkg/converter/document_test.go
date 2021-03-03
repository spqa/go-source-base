package converter

import (
	"context"
	"mcm-api/config"
	"mcm-api/pkg/common"
	"mcm-api/pkg/media"
	"testing"
)

func TestGotenbergDocumentConverter_Convert(t *testing.T) {
	cfg := &config.Config{
		ConverterService: "http://localhost:3001",
		MediaBucket:      "spqa-personal",
	}
	converter := NewGotenbergDocumentConverter(cfg, media.NewStorageService(cfg))
	result, err := converter.Convert(context.TODO(), "94fb201b-4afe-4e29-9fc0-5d80351c7373.docx", common.LoggedInUser{
		Id:    1,
		Email: "",
		Name:  "",
		Role:  "",
	})
	if err != nil {
		t.Error(err)
	} else {
		t.Log(result)
	}
}
