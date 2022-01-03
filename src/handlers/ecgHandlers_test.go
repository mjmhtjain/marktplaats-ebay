package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gocarina/gocsv"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/mjmhtjain/marktplaats-ebay/src/services"
	"github.com/stretchr/testify/assert"
)

func TestECGHandler_GetAll(t *testing.T) {
	t.Run("IF success scenario, THEN expect 200 code response, and some creditors data", func(t *testing.T) {
		expCreditors := readCreditorsFromCSV(t, "/../../resources/Workbook2_small.csv")
		fakeService := fakeCreditService{expCreditors: expCreditors}
		ecgHandler := handlerWithFakeService(&fakeService)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ecgHandler.GetAll(w, req)

		actualCreditors := []models.Creditor{}
		err := json.Unmarshal(w.Body.Bytes(), &actualCreditors)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expCreditors, actualCreditors)
	})

	t.Run("IF there is a downstream failure, THEN expect 500 code response", func(t *testing.T) {
		expErr := errors.New("Some error")
		fakeService := fakeCreditService{err: expErr}
		ecgHandler := handlerWithFakeService(&fakeService)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/", nil)

		ecgHandler.GetAll(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}

func TestECGHandler_UploadHandler(t *testing.T) {
	fakeService := fakeCreditService{}
	ecgHandler := handlerWithFakeService(&fakeService)

	t.Run("IF handler receives a valid file", func(t *testing.T) {
		fakeService.err = nil

		t.Run("IF file is CSV, THEN it returns 201 status code", func(t *testing.T) {
			filePath := "/../../resources/Workbook2.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusOK, w.Code)
		})

		t.Run("IF file has no data, THEN returns 201 response code", func(t *testing.T) {
			filePath := "/../../resources/Workbook2_empty.csv"
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
			filePath := "/../../resources/Workbook2.csv"
			bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
			req.Header.Set("Content-Type", multipartFormDatatype)

			ecgHandler.UploadHandler(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})

	})

	t.Run("IF handler receives a file with invalid file extension, THEN returns 400 response code", func(t *testing.T) {
		filePath := "/../../resources/Workbook2.pdf"
		bytesBuffer, multipartFormDatatype := uploadFile(t, filePath)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/upload", bytesBuffer)
		req.Header.Set("Content-Type", multipartFormDatatype)

		ecgHandler.UploadHandler(w, req)

		assert.Equal(t, w.Code, http.StatusBadRequest)
	})

	t.Run("IF handler receives a csv file with invalid data format, THEN returns 400 response code", func(t *testing.T) {
		// filePath, _ := os.Getwd()
		filePath := "/../../resources/Workbook2_invalid.csv"
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

func readCreditorsFromCSV(t *testing.T, relativePath string) []models.Creditor {
	wd, _ := os.Getwd()
	filePath := wd + relativePath
	file, err := os.Open(filePath)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()

	creditors := []models.Creditor{}
	if err := gocsv.UnmarshalFile(file, &creditors); err != nil {
		t.Error(err)
	}

	return creditors
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
	err          error
	expCreditors []models.Creditor
}

func (ecg *fakeCreditService) UploadCreditorInfo(creditors []models.Creditor) ([]models.Creditor, error) {
	if ecg.err != nil {
		return nil, ecg.err
	}
	return creditors, nil
}

func (ecg *fakeCreditService) GetCreditors() ([]models.Creditor, error) {
	if ecg.err != nil {
		return nil, ecg.err
	}

	if ecg.expCreditors != nil {
		return ecg.expCreditors, nil
	}

	return nil, nil
}
