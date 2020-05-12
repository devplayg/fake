package fake

import (
	"bytes"
	"image"
	"testing"
)

func TestGif(t *testing.T) {
	data := Gif(100, 100)
	_, got, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	want := "gif"
	if got != want {
		t.Errorf("Gif() = %s, got = %s", want, got)
	}
}

func TestImage(t *testing.T) {

}

func TestJpeg(t *testing.T) {
	data := Jpeg(100, 100)
	_, got, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	want := "jpeg"
	if got != want {
		t.Errorf("Jpeg() = %s, got = %s", want, got)
	}
}

func TestPng(t *testing.T) {
	data := Png(100, 100)
	_, got, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	want := "png"
	if got != want {
		t.Errorf("Png() = %s, got = %s", want, got)
	}
}
