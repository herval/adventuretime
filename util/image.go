package util

import (
	"bytes"
	"image/png"
	"image"
	"encoding/base64"
)

func ImageToBytes(img *image.RGBA) ([]byte, error) {
	buff := new(bytes.Buffer)

	// encode image to buffer
	err := png.Encode(buff, img)

	return buff.Bytes(), err
}

func ImageToBase64(img *image.RGBA) (string, error) {
	data, err := ImageToBytes(img)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(data), nil
}