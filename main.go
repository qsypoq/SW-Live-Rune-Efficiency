package main

import (
	"bytes"
	"io"
	"log"
	"os/exec"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
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

func get_rune_name() string {
	name, _, _ := get_rune_infos()
	return name
}

func main() {

	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")

	title := widget.NewLabel("SW Live Rune Analyzer")
	txtBound := binding.NewString()
	txtWid := widget.NewEntryWithData(txtBound)
	txtWid.MultiLine = true
	w.SetContent(container.NewVBox(
		title, txtWid,
		widget.NewButton("Scan Rune", func() {
			for {
				rune_name, _, _, current_efficiency := get_efficiency()
				txtBound.Set(rune_name + "\n" + "Efficiency: " + current_efficiency + "%")
				time.Sleep(time.Millisecond * 250)
			}
		}),
	))
	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
