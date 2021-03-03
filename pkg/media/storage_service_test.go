package media

import (
	"context"
	"mcm-api/config"
	"os"
	"testing"
)

func TestS3StorageService_UploadDocument(t *testing.T) {
	storageService := NewStorageService(&config.Config{MediaBucket: "spqa-personal"})
	file, err := os.Open("file-sample_100kB.docx")
	if err != nil {
		t.Error(err)
		return
	}
	stat, _ := file.Stat()
	_, err = storageService.UploadDocumentOriginal(context.Background(), &FileUploadOriginalReq{
		File: file,
		Size: stat.Size(),
		Name: "",
		User: nil,
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestS3StorageService_GeneratePresignUrl(t *testing.T) {
	storageService := NewStorageService(&config.Config{MediaBucket: "spqa-personal"})
	url, err := storageService.GetUrl(context.TODO(), "946017b1-f63f-4529-971c-7bbdabbab8e1.docx")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(url)
	}
}
