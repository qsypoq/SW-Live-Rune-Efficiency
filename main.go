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

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")

	title := widget.NewLabel("SW Live Rune Analyzer")
	efficiency_str := widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	title_item := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer())
	rune_efficience_item := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), efficiency_str, layout.NewSpacer())
	w.SetContent(
		container.NewVBox(title_item,
			rune_efficience_item,
			widget.NewButton("Scan Rune", func() {
				for {
					rune_name, _, _, current_efficiency := get_efficiency()
					efficiency_str.SetText(rune_name + "\n" + current_efficiency + "%")
					time.Sleep(time.Millisecond * 250)
				}
			}),
		))
	w.Resize(fyne.NewSize(225, 170))
	w.ShowAndRun()
}
