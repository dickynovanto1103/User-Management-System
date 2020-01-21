package fileuploader

import (
	"io"
	"log"
	"net/http"
	"os"
)

const MaxFileSize = 10 << 20

func UploadFile(r *http.Request, fieldname string, username string) (string, error) {
	r.ParseMultipartForm(MaxFileSize)
	file, handler, err := r.FormFile(fieldname)
	log.Println("file, handler, err: ", file, handler, err)
	defer file.Close()
	if err != nil {
		log.Println("error retrieving file with err:", err)
		return "", err
	}

	filename := username + handler.Filename

	f, err := os.OpenFile("assets/"+username+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error in opening file: ", err)
		return "", err
	}
	io.Copy(f, file)
	return filename, nil
}
