package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
	"github.com/mjmhtjain/marktplaats-ebay/src/services"
)

type ECGHandler struct {
	creditService services.CreditService
}

func NewECGHandler() *ECGHandler {
	return &ECGHandler{
		creditService: services.NewCreditService(),
	}
}

// UploadHandler accepts multipart/form-data file upload of .csv and .prn extension
// TODO: logger
func (ecg *ECGHandler) UploadHandler(w http.ResponseWriter, req *http.Request) {
	multipartFile, fileHeader, err := req.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer multipartFile.Close()

	creditors, err := readFile(multipartFile, fileHeader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uploadedCreaditors, err := ecg.creditService.UploadCreditorInfo(creditors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(uploadedCreaditors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func readFile(multipartFile multipart.File, fileHeader *multipart.FileHeader) ([]models.Creditor, error) {
	var err error
	var creditors []models.Creditor

	buff := new(bytes.Buffer)
	io.Copy(buff, multipartFile)

	if strings.HasSuffix(fileHeader.Filename, ".csv") {
		creditors, err = readCSVFile(buff)
	} else if strings.HasSuffix(fileHeader.Filename, ".prn") {
		creditors, err = readPRNFile(buff)
	} else {
		return nil, errors.New("invalid extension")
	}

	if err != nil {
		return nil, errors.New("parsing error")
	}
	return creditors, nil
}

func readCSVFile(buff *bytes.Buffer) ([]models.Creditor, error) {
	creditors := []models.Creditor{}
	err := gocsv.UnmarshalBytes(buff.Bytes(), &creditors)
	if err != nil {
		return nil, err
	}

	return creditors, nil
}

func readPRNFile(buff *bytes.Buffer) ([]models.Creditor, error) {
	return nil, nil
}
