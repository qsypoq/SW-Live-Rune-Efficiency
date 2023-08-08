package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
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
	generate_tmp_imgs()
	rune_name, rune_stats, rune_subs := get_rune_infos()

	fmt.Print("\nRune Stats:\n", rune_stats)
	fmt.Print("Rune Name:\n", rune_name)
	fmt.Print("Rune Efficiency: ", get_efficiency(get_hit_number(rune_subs, rune_stats)), "%", "\n")
	// fmt.Print("Rune Inate Efficiency: ", get_efficiency(get_hit_number(rune_name)), "%", "\n")
	// // a := app.New()
	// w := a.NewWindow("SW Live Rune Analyzer")

	// hello := widget.NewLabel("SW Live Rune Analyzer")
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Select your game window", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))
	// w.Resize(fyne.NewSize(640, 460))
	// w.ShowAndRun()
}
