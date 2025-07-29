package util

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"io/fs"
)

//go:embed assets/*
var assets embed.FS

func MustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func MustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = MustLoadImage(match)
	}

	return images
}
