package handlers

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

// UploadHandler accepts multipart/form-data file upload of .csv and .prn extension
// TODO: logger
func UploadHandler(w http.ResponseWriter, req *http.Request) {
	multipartFile, _, err := req.FormFile("file")
	if err != nil {
		// panic ??
		panic(err)
	}

	readFile(multipartFile)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("File Uploaded successfully"))
}

func readFile(multipartFile multipart.File) {
	defer multipartFile.Close()

	buff := new(bytes.Buffer)
	io.Copy(buff, multipartFile)
	fmt.Println(buff.String())
}
