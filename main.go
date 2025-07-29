package main

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
)

const (
	Spielbreite = 800
	Spielhoehe  = 600
)

//go:embed assets/*
var assets embed.FS

var PlayerBild = mustLoadImage("assets/player.png")

type Vector struct {
	X float64
	Y float64
}

type Meteorit struct {
}

type Player struct {
	bild     *ebiten.Image
	position Vector
}

type Game struct {
	player *Player
}

func (g *Game) Update() error {
	return nil
}

func NewPlayer() *Player {
	return &Player{
		position: Vector{X: 100, Y: 100},
		bild:     PlayerBild,
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.bild, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return Spielbreite, Spielhoehe
}

func main() {
	g := &Game{
		player: NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}

func mustLoadImage(name string) *ebiten.Image {
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
