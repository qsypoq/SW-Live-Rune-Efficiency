package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os/exec"
)

func ocrbis(img []byte) string {
	cmdArgs := []string{"stdin", "stdout", "--psm", "6"}
	cmd := exec.Command("C:\\Program Files\\Tesseract-OCR\\tesseract.exe", cmdArgs...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}

	go func() {
		defer stdin.Close()
		if _, err := stdin.Write(img); err != nil {
			panic(err)
		}
	}()

	var out bytes.Buffer
	multi := io.MultiWriter(&out)
	cmd.Stdout = multi

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	return out.String()
}

func image_to_byte(img image.Image) []byte {
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		fmt.Println("failed to create buffer", err)
	}
	img_byte := buf.Bytes()
	return img_byte
}

func get_text(img image.Image) string {
	imgbyte := image_to_byte(img)
	return ocrbis(imgbyte)
}
