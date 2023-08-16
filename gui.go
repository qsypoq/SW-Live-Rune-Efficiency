package main

import (
	"image/color"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"golang.org/x/sys/windows"
)

func GetWindowHandleByWindowName(window_name string) uintptr {
	user32dll := windows.MustLoadDLL("user32.dll")
	enumwindows := user32dll.MustFindProc("EnumWindows")

	var the_handle uintptr
	window_byte_name := []byte(window_name)

	// Windows will loop over this function for each window.
	wndenumproc_function := syscall.NewCallback(func(hwnd uintptr, lparam uintptr) uintptr {
		// Allocate 100 characters so that it has something to write to.
		var filename_data [100]uint16
		max_chars := uintptr(100)

		getwindowtextw := user32dll.MustFindProc("GetWindowTextW")
		getwindowtextw.Call(hwnd, uintptr(unsafe.Pointer(&filename_data)), max_chars)

		// If there's a match, save the value and return 0 to stop the iteration.
		if strings.Contains(string(windows.UTF16ToString([]uint16(filename_data[:]))), string(window_byte_name)) {
			the_handle = hwnd
			return 0
		}

		return 1
	})

	// Call the above looping function.
	enumwindows.Call(wndenumproc_function, uintptr(0))

	return the_handle
}

const SWP_NOSIZE = uintptr(0x0001)
const SWP_NOMOVE = uintptr(0x0002)

// This is dumb but Go doesn't like the inline conversion (see above image).
func IntToUintptr(value int) uintptr {
	return uintptr(value)
}

func setontop(windows_name string) {
	time.Sleep(time.Millisecond * 100)
	SetWindowAlwaysOnTop(GetWindowHandleByWindowName(windows_name))
}

func SetWindowAlwaysOnTop(hwnd uintptr) {
	user32dll := windows.MustLoadDLL("user32.dll")
	setwindowpos := user32dll.MustFindProc("SetWindowPos")
	setwindowpos.Call(hwnd, IntToUintptr(-1), 0, 0, 100, 100, SWP_NOSIZE|SWP_NOMOVE)
}

func gen_txt(content string, color color.Color, style fyne.TextStyle, size float32) *canvas.Text {
	newtxt := canvas.NewText(content, color)
	newtxt.Alignment = fyne.TextAlignCenter
	newtxt.TextStyle = style
	newtxt.TextSize = size
	return newtxt
}

func customize_container(target_container *fyne.Container, items []fyne.CanvasObject) {
	target_container.RemoveAll()
	target_container.Objects = items
	target_container.Refresh()
}
