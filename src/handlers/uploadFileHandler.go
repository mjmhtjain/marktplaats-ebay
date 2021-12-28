package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/mjmhtjain/marktplaats-ebay/src/models"
)

// UploadHandler accepts multipart/form-data file upload of .csv and .prn extension
// TODO: logger
func UploadHandler(w http.ResponseWriter, req *http.Request) {
	multipartFile, fileHeader, err := req.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer multipartFile.Close()

	// check invalid file extension
	if !(strings.HasSuffix(fileHeader.Filename, ".csv") ||
		strings.HasSuffix(fileHeader.Filename, ".prn")) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	creditor, err := readFile(multipartFile, fileHeader)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := json.Marshal(creditor)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func readFile(multipartFile multipart.File, fileHeader *multipart.FileHeader) ([]models.Creditor, error) {
	creditor := []models.Creditor{}
	buff := new(bytes.Buffer)
	io.Copy(buff, multipartFile)
	fmt.Println(buff.String())

	err := gocsv.UnmarshalBytes(buff.Bytes(), &creditor)
	// err := json.Unmarshal(buff.Bytes(), &creditor)
	if err != nil {
		return nil, err
	}

	return creditor, nil
}
