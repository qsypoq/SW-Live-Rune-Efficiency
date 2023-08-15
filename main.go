package main

import (
	"image/color"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	os.Setenv("FYNE_THEME", "dark")
	a := app.New()
	w := a.NewWindow("SW Live Rune Analyzer")
	logo, _ := fyne.LoadResourceFromPath("./logo.png")
	w.SetIcon(logo)
	title_item := gen_txt("SW Live Rune Analyzer", color.White, fyne.TextStyle{Bold: true})
	scan := false
	rune_name_txt := gen_txt("Rune Name", color.White, fyne.TextStyle{Bold: true})
	rune_eff_txt := gen_txt("Rune Eff", color.White, fyne.TextStyle{Bold: true})
	rune_tier_txt := gen_txt("Rune Tier", color.White, fyne.TextStyle{Bold: true})
	rune_maxeff_txt := gen_txt("Rune Max Eff", color.White, fyne.TextStyle{Bold: true})
	rune_maxtier_txt := gen_txt("Rune Max Tier", color.White, fyne.TextStyle{Bold: true})
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
				rune_tier_txt.Text, rune_tier_txt.Color = get_tier(current_efficiency)
				rune_eff_txt.Text = "Efficiency: " + current_efficiency + "%"
				rune_tier_txt.Text = "Tier: " + rune_tier_txt.Text
				if max_efficiency == current_efficiency {
					rune_maxeff_txt.Text = " "
					rune_maxtier_txt.Text = " "
				} else {
					rune_maxeff_txt.Text = "Potential Efficiency: " + max_efficiency + "%"
					rune_maxtier_txt.Text, rune_maxtier_txt.Color = get_tier(max_efficiency)
					rune_maxtier_txt.Text = "Potential Tier: " + rune_maxtier_txt.Text
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
	go setontop("SW Live Rune Analyzer")
	w.ShowAndRun()
}
