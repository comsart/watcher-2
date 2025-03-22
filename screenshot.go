package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/kbinani/screenshot"
	"image/jpeg"
)

func takeScreenshot() (string, error) {
	// Capture the screen
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		fmt.Println("Failed to capture screen:", err)
		return "", fmt.Errorf("in capture screen: %v", err)
	}

	// Encode the image to JPEG
	var imgBytes bytes.Buffer
	err = jpeg.Encode(&imgBytes, img, nil)
	if err != nil {

		return "", fmt.Errorf("in encode image: %v", err)
	}

	// Encode the JPEG bytes to base64
	base64Str := base64.StdEncoding.EncodeToString(imgBytes.Bytes())

	return base64Str, nil
}
