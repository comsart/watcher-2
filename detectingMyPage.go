package main

import (
	"github.com/kbinani/screenshot"
	"image/color"
	"log"
)

func detectPage() bool {
	targetColor := color.RGBA{R: 255, G: 250, B: 240, A: 255} // flowery white

	// Capture the screenshot of the primary display
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		log.Fatalf("Failed to capture screenshot: %v", err)
	}

	// Check if the target color exists in the screenshot (skanuje kolumnami zeby bylo szybciej)
	pixelsFound := 0
	for x := 0; x < img.Bounds().Dx(); x++ {
		for y := 0; y < img.Bounds().Dy(); y++ {
			if img.At(x, y) == targetColor {
				pixelsFound++
				if pixelsFound > 500 {
					return true
				}
			}
		}
	}
	//dumpErrorToFile("The specified color is not found in the screenshot.")
	return false
}
