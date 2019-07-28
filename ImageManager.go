package gfull

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

// ImageManager model manager image
type ImageManager struct {
}

func main() {
	var img = new(ImageManager)
	var imagen, _ = img.ConverteBase64ToImage("")
	img.SaveImagePNG(imagen, "imagen.png")
}

// ConverteBase64ToImage convert base64 to Image
func (ImageManager) ConverteBase64ToImage(bs64 string) (image.Image, error) {
	coI := strings.Index(bs64, ",")
	rawImage := bs64[coI+1:]
	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(rawImage)
	res := bytes.NewReader(unbased)
	switch strings.TrimSuffix(bs64[5:coI], ";base64") {
	case "image/png":
		return png.Decode(res)
	case "image/jpeg":
		return jpeg.Decode(res)
	}
	return nil, errors.New("Image mime not detect")
}

// SaveImagePNG to format png
func (ImageManager) SaveImagePNG(src image.Image, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return png.Encode(f, src)
}

// SaveImageJPEG to format JPEG
func (ImageManager) SaveImageJPEG(src image.Image, file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	return jpeg.Encode(f, src, &jpeg.Options{Quality: 100})
}
