package media

import (
	"mcm-api/config"
	"os"
	"testing"
)

func TestS3StorageService_UploadDocument(t *testing.T) {
	storageService := NewStorageService(&config.Config{DocumentBucket: "spqa-personal"})
	file, err := os.Open("file-sample_100kB.docx")
	if err != nil {
		t.Error(err)
		return
	}
	_, err = storageService.UploadDocumentOriginal(file, nil)
	if err != nil {
		t.Error(err)
		return
	}
}
