package main

import (
	"game/util"
	"github.com/hajimehoshi/ebiten/v2"
	_ "image/png"
)

const (
	Spielbreite = 800
	Spielhoehe  = 600
)

var PlayerBild = util.MustLoadImage("assets/player.png")

type Vector struct {
	X float64
	Y float64
}

type Meteorit struct {
}

type Player struct {
}

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
}

func (g *Game) Draw(screen *ebiten.Image) {
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return Spielbreite, Spielhoehe
}

func main() {
	g := &Game{}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
