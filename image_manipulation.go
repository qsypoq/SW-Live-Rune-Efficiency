package main

import (
	"image"

	"github.com/anthonynsimon/bild/adjust"
	"github.com/disintegration/gift"
)

func adjust_brightness(img image.Image) image.Image {
	result := adjust.Contrast(img, +0.5)
	grayscale := gift.New(gift.Grayscale())
	dst := image.NewNRGBA(grayscale.Bounds(result.Bounds()))
	grayscale.Draw(dst, result)
	return dst
}

func crop_custom(img image.Image, width int, height int, anchor gift.Anchor) image.Image {
	h := gift.New(gift.CropToSize(width, height, anchor))
	resized := image.NewRGBA(h.Bounds(img.Bounds()))
	h.Draw(resized, img)
	return resized
}

func generate_rune_name(img image.Image) image.Image {
	finaldst := crop_custom(img, 625, 350, gift.LeftAnchor)
	generated := crop_custom(finaldst, 550, 75, gift.TopRightAnchor)
	return generated
}

func generate_rune_stats(img image.Image) image.Image {
	finaldst := crop_custom(img, 570, 250, gift.TopRightAnchor)
	generated := crop_custom(finaldst, 300, 125, gift.BottomLeftAnchor)
	return generated
}

func generate_rune_subs(img image.Image) image.Image {
	finaldst := crop_custom(img, 380, 360, gift.LeftAnchor)
	generated := crop_custom(finaldst, 705, 150, gift.BottomRightAnchor)
	return generated
}

func generate_rune(img image.Image) {
	adjusted__img := adjust_brightness(img)
	// native_ocr(generate_rune_name(adjusted__img), generate_rune_stats(adjusted__img), generate_rune_subs(adjusted__img))
	save_screen(generate_rune_name(adjusted__img), "files/tmp/tmp_name.png")
	save_screen(generate_rune_stats(adjusted__img), "files/tmp/tmp_stats.png")
	save_screen(generate_rune_subs(adjusted__img), "files/tmp/tmp_subs.png")
}
