package main

import (
	"image"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func loadImageFromFile(imgPath string) image.Image {
	imageFile, _ := os.Open(imgPath)
	defer imageFile.Close()
	img, _, _ := image.Decode(imageFile)
	return img
}

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")
	logo, _ := fyne.LoadResourceFromPath("./logo.png")
	w.SetIcon(logo)
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
				layout.NewSpacer(), start_button,
			))
		w.Content().Refresh()
	})

	start_button = widget.NewButton("Scan Rune", func() {
		scan = true
		w.SetContent(
			container.NewVBox(title_item,
				rune_efficience_item, layout.NewSpacer(), stop_button,
			))
		w.Content().Refresh()
		inf_run := func() {
			for scan == true {
				rune_name, _, _, current_efficiency := get_efficiency()
				efficiency_str.SetText(rune_name + "\n" + current_efficiency + "%")
				time.Sleep(time.Millisecond * 200)
			}
		}
		go inf_run()
	})
	w.SetContent(
		container.NewVBox(title_item,
			rune_efficience_item, layout.NewSpacer(), start_button,
		))
	w.Content().Refresh()
	w.Resize(fyne.NewSize(225, 170))
	w.ShowAndRun()
}
