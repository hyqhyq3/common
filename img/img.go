package img

import (
	"image"
	"image/png"
	"os"
)

func LoadImg(LoadImg string) (img image.Image, err error) {
	f, err := os.OpenFile(LoadImg, os.O_RDONLY, 0)
	if err != nil {
		return
	}
	defer f.Close()
	img, _, err = image.Decode(f)
	return
}

func SaveImg(filename string, img image.Image) (err error) {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}
	defer f.Close()
	return png.Encode(f, img)
}
