package handlers

import (
	"io"
	"net/http"
	"os"
)

// UploadHandler accepts multipart/form-data file upload of .csv and .prn extension
// TODO: logger
func UploadHandler(w http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("file")
	if err != nil {
		// panic ??
		panic(err)
	}
	defer file.Close()

	// fileName := req.FormValue("file_name")
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// panic ??
		panic(err)
	}
	defer f.Close()

	_, _ = io.WriteString(w, "File Uploaded successfully")
	_, _ = io.Copy(f, file)
	// w.Write([]byte("Hello, world!\n"))
}
