package service

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type coreService struct {
}

//nolint:gochecknoglobals
var CoreService = &coreService{}

func (s *coreService) GenerateAvatar(name string, size int, dark bool) ([]byte, string, error) {
	var background, foreground color.Color
	if dark {
		background = color.Black
		foreground = color.White
	} else {
		background = color.White
		foreground = color.Black
	}

	img := image.NewRGBA(
		image.Rectangle{
			Min: image.Point{},
			Max: image.Point{
				X: size,
				Y: size,
			},
		},
	)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			img.Set(x, y, background)
		}
	}

	point := fixed.Point26_6{X: fixed.I(0), Y: fixed.I(size / 2)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(foreground),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(name)

	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, "", fmt.Errorf("failed to encode png: %w", err)
	}

	return buf.Bytes(), "image/png", nil
}
