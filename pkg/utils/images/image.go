package images

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"os"
)

func Read(path string) (image.Image, error) {
	open, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	img, _, err := image.Decode(open)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func Compress(img image.Image) (bytes.Buffer, error) {
	newImg := resize.Resize(200, 0, img, resize.Lanczos3)
	var imgData bytes.Buffer
	if err := jpeg.Encode(&imgData, newImg, nil); err != nil {
		return imgData, err
	}
	return imgData, nil
}
