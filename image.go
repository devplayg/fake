package fake

import (
	"bytes"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"
)

func Image(width int, height int) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			img.Set(i, j, color.RGBA{
				R: Uint8() / 3,
				G: Uint8() / 3,
				B: Uint8() / 3,
				A: 0xff,
			})
		}
	}
	return img
}

func Jpeg(width int, height int) []byte {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, Image(width, height), nil)
	return buf.Bytes()
}

func Png(width int, height int) []byte {
	buf := new(bytes.Buffer)
	png.Encode(buf, Image(width, height))
	return buf.Bytes()
}

func Gif(width int, height int) []byte {
	buf := new(bytes.Buffer)
	gif.Encode(buf, Image(width, height), nil)
	return buf.Bytes()
}
