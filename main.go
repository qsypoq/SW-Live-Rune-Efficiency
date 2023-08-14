package main

import (
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func gen_txt(content string, color color.Color) *canvas.Text {
	newtxt := canvas.NewText(content, color)
	newtxt.Alignment = fyne.TextAlignCenter
	newtxt.TextStyle = fyne.TextStyle{Bold: true}
	return newtxt
}

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")
	logo, _ := fyne.LoadResourceFromPath("./logo.png")
	w.SetIcon(logo)
	title := widget.NewLabel("SW Live Rune Analyzer")
	title_item := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), title, layout.NewSpacer())
	scan := false
	rune_name_txt := gen_txt("Rune Name", color.White)
	rune_eff_txt := gen_txt("Rune Eff", color.White)
	rune_tier_txt := gen_txt("Rune Tier", color.White)

	start_button := widget.NewButton("Scan Rune", func() {})
	stop_button := widget.NewButton("Stop Scan", func() {
		scan = false
		w.SetContent(
			container.NewVBox(title_item,
				layout.NewSpacer(), start_button,
			))
		w.Resize(fyne.NewSize(225, 170))
		w.Content().Refresh()
	})

	start_button = widget.NewButton("Scan Rune", func() {
		scan = true
		w.SetContent(
			container.NewVBox(title_item, rune_name_txt,
				rune_eff_txt, rune_tier_txt, layout.NewSpacer(), stop_button,
			))
		w.Resize(fyne.NewSize(225, 170))
		w.Content().Refresh()
		inf_run := func() {
			for scan == true {
				rune_name, _, _, current_efficiency := get_efficiency()
				rune_name_txt.Text = rune_name
				rune_eff_txt.Text = current_efficiency + "%"
				rune_tier_txt.Text, rune_tier_txt.Color = get_tier(current_efficiency)
				w.Content().Refresh()
				time.Sleep(time.Millisecond * 200)
			}
		}
		go inf_run()
	})
	w.SetContent(
		container.NewVBox(title_item,
			rune_name_txt, rune_eff_txt, rune_tier_txt, layout.NewSpacer(), start_button,
		))
	w.Content().Refresh()
	w.Resize(fyne.NewSize(225, 170))
	w.ShowAndRun()
}
