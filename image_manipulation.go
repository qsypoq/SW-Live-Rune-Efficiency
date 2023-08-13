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

func crop_custom(img image.Image, ratiow float64, ratioh float64, anchor gift.Anchor) image.Image {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	h := gift.New(gift.CropToSize(int(float64(width)*ratiow), int(float64(height)*ratioh), anchor))
	resized := image.NewRGBA(h.Bounds(img.Bounds()))
	h.Draw(resized, img)
	return resized
}

func generate_rune_name(img image.Image) image.Image {
	finaldst := crop_custom(img, 0.8803, 0.7778, gift.LeftAnchor)
	generated := crop_custom(finaldst, 0.88, 0.2143, gift.TopRightAnchor)
	return generated
}

func generate_rune_stats(img image.Image) image.Image {
	finaldst := crop_custom(img, 0.8028, 0.5556, gift.TopRightAnchor)
	generated := crop_custom(finaldst, 0.5263, 0.5, gift.BottomLeftAnchor)
	return generated
}

func generate_rune_subs(img image.Image) image.Image {
	finaldst := crop_custom(img, 0.5352, 0.80, gift.LeftAnchor)
	generated := crop_custom(finaldst, 1, 0.4167, gift.BottomRightAnchor)
	return generated
}

func generate_rune() (string, string, string) {
	img, _ := capture("Summoners War - MuMu Player")
	adjusted__img := adjust_brightness(img)
	name := get_text(generate_rune_name(adjusted__img))
	stats := get_text(generate_rune_stats(adjusted__img))
	subs := get_text(generate_rune_subs(adjusted__img))
	return name, stats, subs
}
