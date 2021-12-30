package handlers

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/mjmhtjain/marktplaats-ebay/src/services"
	"github.com/stretchr/testify/assert"
)

func TestUploadHandler(t *testing.T) {
	fakeService := fakeCreditService{}
	ecgHandler := handlerWithFakeService(&fakeService)

	t.Run("IF handler receives a valid file", func(t *testing.T) {
		fakeService.err = nil

		t.Run("IF file is CSV, THEN it returns 201 status code", func(t *testing.T) {
			filePath := "/fixtures/Workbook2.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("IF file has no data, THEN returns 201 response code", func(t *testing.T) {
			filePath := "/fixtures/Workbook2_empty.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("IF file is PRN file, THEN it returns 201 status code", func(t *testing.T) {

		})

		t.Run("IF there is a downstream service error, THEN it returns 500 error code", func(t *testing.T) {
			fakeService.err = errors.New("Un expected error")
			filePath := "/fixtures/Workbook2.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

	})

	t.Run("IF handler receives a file with invalid file extension, THEN returns 400 response code", func(t *testing.T) {
		filePath := "/fixtures/Workbook2.pdf"
		bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", multipartFormDatatype)

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	t.Run("IF handler receives a csv file with invalid data format, THEN returns 400 response code", func(t *testing.T) {
		// filePath, _ := os.Getwd()
		filePath := "/fixtures/Workbook2_invalid.csv"
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

}

func uploadFile(t *testing.T, relativePath string) (*bytes.Buffer, string) {
	wd, _ := os.Getwd()
	filePath := wd + relativePath
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

func handlerWithFakeService(fakeService services.CreditService) *ECGHandler {
	return &ECGHandler{
		creditService: fakeService,
	}
}

type fakeCreditService struct {
	err error
}

func (ecg *fakeCreditService) UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error) {
	if ecg.err != nil {
		return nil, ecg.err
	}
	return creditors, nil
}

func (ecg *fakeCreditService) GetCreditors() {

}
