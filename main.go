package main

import (
	"image/color"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/sys/windows"
)

func gen_txt(content string, color color.Color) *canvas.Text {
	newtxt := canvas.NewText(content, color)
	newtxt.Alignment = fyne.TextAlignCenter
	newtxt.TextStyle = fyne.TextStyle{Bold: true}
	return newtxt
}

func customize_container(target_container *fyne.Container, items []fyne.CanvasObject) {
	target_container.RemoveAll()
	target_container.Objects = items
	// for _, v := range items {
	// 	target_container.Add(v)
	// }
	target_container.Refresh()
}
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

func setpontop(windows_name string) {
	time.Sleep(time.Millisecond * 100)
	SetWindowAlwaysOnTop(GetWindowHandleByWindowName(windows_name))
}

func SetWindowAlwaysOnTop(hwnd uintptr) {
	user32dll := windows.MustLoadDLL("user32.dll")
	setwindowpos := user32dll.MustFindProc("SetWindowPos")
	setwindowpos.Call(hwnd, IntToUintptr(-1), 0, 0, 100, 100, SWP_NOSIZE|SWP_NOMOVE)
}
func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")
	logo, _ := fyne.LoadResourceFromPath("./logo.png")
	w.SetIcon(logo)
	title_item := gen_txt("SW Live Rune Analyzer", color.White)
	scan := false
	rune_name_txt := gen_txt("Rune Name", color.White)
	rune_eff_txt := gen_txt("Rune Eff", color.White)
	rune_tier_txt := gen_txt("Rune Tier", color.White)
	rune_maxeff_txt := gen_txt("Rune Max Eff", color.White)
	rune_maxtier_txt := gen_txt("Rune Max Tier", color.White)
	start_button := widget.NewButton("Scan Rune", func() {})
	custom_container := container.NewVBox()
	stop_button := widget.NewButton("Stop Scan", func() {
		scan = false
		customize_container(custom_container, []fyne.CanvasObject{title_item, layout.NewSpacer(), start_button})
		w.Content().Refresh()
	})

	start_button = widget.NewButton("Scan Rune", func() {
		scan = true
		customize_container(custom_container, []fyne.CanvasObject{title_item, rune_name_txt, rune_name_txt, rune_eff_txt, rune_tier_txt, rune_maxeff_txt, rune_maxtier_txt, layout.NewSpacer(), stop_button})
		w.Content().Refresh()
		inf_run := func() {
			for scan {
				rune_name, _, _, current_efficiency, max_efficiency := get_efficiency()
				rune_name_txt.Text = rune_name
				rune_eff_txt.Text = current_efficiency + "%"
				rune_tier_txt.Text, rune_tier_txt.Color = get_tier(current_efficiency)
				if max_efficiency == current_efficiency {
					rune_maxeff_txt.Text = " "
					rune_maxtier_txt.Text = " "
				} else {
					rune_maxeff_txt.Text = max_efficiency + "%"
					rune_maxtier_txt.Text, rune_maxtier_txt.Color = get_tier(max_efficiency)
				}
				custom_container.Refresh()
				w.Content().Refresh()
				time.Sleep(time.Millisecond * 100)
			}
		}
		go inf_run()
	})
	customize_container(custom_container, []fyne.CanvasObject{title_item, layout.NewSpacer(), start_button})
	custom_container.Refresh()
	w.SetContent(custom_container)
	w.Content().Refresh()
	w.Resize(fyne.NewSize(225, 170))
	go setpontop("SW Live Rune Analyzer")
	w.ShowAndRun()
}
