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
	ecgHandler := NewECGHandler()

	t.Run("IF handler receives a file", func(t *testing.T) {
		t.Run("IF passed a valid CSV file, THEN it returns 201 status code", func(t *testing.T) {
			filePath, _ := os.Getwd()
			filePath += "/fixtures/Workbook2.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("IF passed a valid PRN file, THEN it returns 201 status code", func(t *testing.T) {

		})

		t.Run("IF the service fails downstream, THEN it returns 500 error code", func(t *testing.T) {

		})

	})

	t.Run("IF handler receives a file with invalid file extension, THEN returns 400 response code", func(t *testing.T) {
		filePath, _ := os.Getwd()
		filePath += "/fixtures/Workbook2.pdf"
		bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", multipartFormDatatype)

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	t.Run("IF handler receives a csv file with invalid data format, THEN returns 400 response code", func(t *testing.T) {
		filePath, _ := os.Getwd()
		filePath += "/fixtures/Workbook2_invalid.csv"
		bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", multipartFormDatatype)

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("IF handler receives no file, THEN returns 400 response code", func(t *testing.T) {
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", nil)
		req.Header.Set("Content-Type", "multipart/form-data")

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	t.Run("IF handler a valid file with no data, THEN returns 201 response code", func(t *testing.T) {
		filePath, _ := os.Getwd()
		filePath += "/fixtures/Workbook2_empty.csv"
		bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", multipartFormDatatype)

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
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
