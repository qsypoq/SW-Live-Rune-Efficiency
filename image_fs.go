package main

import (
	"image"
	"image/png"
	"os"
)

// func loadImageFromFile(imgPath string) image.Image {
// 	imageFile, _ := os.Open(imgPath)
// 	defer imageFile.Close()
// 	img, _, _ := image.Decode(imageFile)
// 	return img
// }

func save_screen(img image.Image, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
	f.Close()
}

func generate_tmp_imgs() {
	img, _ := capture("Summoners War - MuMu Player")
	save_screen(img, "files/tmp/tmp.png")
	generate_rune(img)
}
