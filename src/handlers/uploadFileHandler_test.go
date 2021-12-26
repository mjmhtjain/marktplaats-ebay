package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadHandler(t *testing.T) {
	t.Run("IF handler receives a file, THEN it does something", func(t *testing.T) {
		filePath, _ := os.Getwd()
		filePath += "/fixtures/Workbook2.csv"
		bytesBuffer, datatype := uploadFile(t, filePath)
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", datatype)

		UploadHandler(w, req)

		assert.Equal(t, w.Code, http.StatusOK)
	})
}

func uploadFile(t *testing.T, filePath string) (*bytes.Buffer, string) {
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Error(err)
	}
	io.Copy(part, file)

	return body, writer.FormDataContentType()
}
