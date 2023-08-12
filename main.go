package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func ocr(path string) string {
	cmdArgs := []string{path, "stdout", "--psm", "6"}
	cmd := exec.Command("C:\\Program Files\\Tesseract-OCR\\tesseract.exe", cmdArgs...)

	var out bytes.Buffer
	multi := io.MultiWriter(&out)
	cmd.Stdout = multi

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	return out.String()
}

func get_rune_infos() (string, string, string) {
	name := ocr("files/tmp/tmp_name.png")
	stats := ocr("files/tmp/tmp_stats.png")
	subs := ocr("files/tmp/tmp_subs.png")
	return name, stats, subs
}

// func image_to_byte(img image.Image) []byte {
// 	buf := new(bytes.Buffer)
// 	err := png.Encode(buf, img)
// 	if err != nil {
// 		fmt.Println("failed to create buffer", err)
// 	}
// 	img_byte := buf.Bytes()
// 	return img_byte
// }

// func byte_to_txt(img_byte []byte) string {
// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	client.SetImageFromBytes(img_byte)
// 	setimg_txt, _ := client.Text()
// 	return setimg_txt
// }

// func native_ocr(name_img image.Image, stats_img image.Image, subs_img image.Image) (string, string, string) {
// 	name := byte_to_txt(image_to_byte(name_img))
// 	stats := byte_to_txt(image_to_byte(stats_img))
// 	subs := byte_to_txt(image_to_byte(subs_img))
// 	return name, stats, subs
// }

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")

	title := widget.NewLabel("SW Live Rune Analyzer")
	efficiency_str := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title_item := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer())
	rune_efficience_item := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), efficiency_str, layout.NewSpacer())
	scan := false
	start_button := widget.NewButton("Scan Rune", func() {})
	stop_button := widget.NewButton("Stop Scan", func() {
		scan = false
		w.SetContent(
			container.NewVBox(title_item,
				start_button,
			))
		w.Content().Refresh()
	})

	start_button = widget.NewButton("Scan Rune", func() {
		scan = true
		w.SetContent(
			container.NewVBox(title_item,
				rune_efficience_item, stop_button,
			))
		w.Content().Refresh()
		inf_run := func() {
			for scan == true {
				rune_name, _, _, current_efficiency := get_efficiency()
				efficiency_str.SetText(rune_name + "\n" + current_efficiency + "%")
				time.Sleep(time.Millisecond * 250)
			}
		}
		go inf_run()
	})
	w.SetContent(
		container.NewVBox(title_item,
			rune_efficience_item, start_button,
		))
	w.Content().Refresh()
	w.Resize(fyne.NewSize(225, 170))
	w.ShowAndRun()
}
