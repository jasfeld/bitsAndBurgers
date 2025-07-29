package main

import (
	"game/util"
	_ "image/png"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	Spielbreite = 800
	Spielhoehe  = 600
)

var PlayerBild = util.MustLoadImage("assets/player.png")
var MeteoritenBilder = util.MustLoadImages("assets/meteors/*.png")

type Meteorit struct {
	bild     *ebiten.Image
	position util.Vector
	bewegung util.Vector
}

type Player struct {
	bild          *ebiten.Image
	position      util.Vector
	shootCooldown *util.Timer
	lasers        []*util.Laser
}

type Game struct {
	player               *Player
	meteoriten           []*Meteorit
	meteoritenSpawnTimer *util.Timer
}

func NewPlayer() *Player {

	grenzen := PlayerBild.Bounds()
	halbeBreite := float64(grenzen.Dx()) / 2
	halbeHoehe := float64(grenzen.Dy()) / 2

	pos := util.Vector{
		X: Spielbreite/2 - halbeBreite,
		Y: Spielhoehe/2 - halbeHoehe,
	}

	return &Player{
		position:      pos,
		bild:          PlayerBild,
		shootCooldown: util.NewTimer(1 * time.Second),
	}
}

func NewMeteorit() *Meteorit {
	return &Meteorit{
		bild: MeteoritenBilder[rand.Intn(len(MeteoritenBilder))],
		position: util.Vector{
			X: float64(rand.Intn(Spielbreite)),
			Y: 0,
		},
		bewegung: util.Vector{
			X: 0,
			Y: float64(rand.Intn(3) + 1), // Random speed between 1 and 3
		},
	}
}

func (m *Meteorit) Update() {
	// Update meteorite position based on its bewegung vector
	m.position.X += m.bewegung.X
	m.position.Y += m.bewegung.Y

	// Check if the meteorite is off-screen and remove it if necessary
	if m.position.Y > Spielhoehe {
		// Remove logic can be added here if needed
	}
}

func (p *Player) Update() error {
	geschwindigkeit := 5.0
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.position.Y += geschwindigkeit
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.position.Y -= geschwindigkeit
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.position.X -= geschwindigkeit
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.position.X += geschwindigkeit
	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()
		p.lasers = append(p.lasers, util.NewLaser(p.position, p.bild.Bounds()))
	}

	return nil
}

func (g *Game) Update() error {
	g.player.Update()

	g.meteoritenSpawnTimer.Update()
	if g.meteoritenSpawnTimer.IsReady() {
		g.meteoritenSpawnTimer.Reset()
		g.meteoriten = append(g.meteoriten, NewMeteorit())
	}

	for _, laser := range g.player.lasers {
		laser.Update()
	}

	for _, m := range g.meteoriten {
		m.Update()
		if m.KollisionsRechteck().IstKollidiert(g.player.KollisionsRechteck()) {
			// Handle collision with player
			g.GameOver()
		}
		for _, laser := range g.player.lasers {
			if m.KollisionsRechteck().IstKollidiert(laser.KollisionsRechteck()) {
				// Remove the meteorite and the laser
				g.meteoriten = append(g.meteoriten[:0], g.meteoriten[1:]...) // Remove the first meteorite
			}
		}
	}

	return nil
}

func (g *Game) GameOver() {
	g.player = NewPlayer()
	g.meteoriten = nil
}

func (m *Meteorit) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.bild, op)
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.position.X, p.position.Y)
	screen.DrawImage(p.bild, op)
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, laser := range g.player.lasers {
		laser.Draw(screen)
	}

	for _, m := range g.meteoriten {
		m.Draw(screen)
	}
}

func (p *Player) KollisionsRechteck() util.Rect {
	bounds := p.bild.Bounds()
	return util.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (m *Meteorit) KollisionsRechteck() util.Rect {
	bounds := m.bild.Bounds()
	return util.NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return Spielbreite, Spielhoehe
}

func main() {
	g := &Game{
		meteoritenSpawnTimer: util.NewTimer(2 * time.Second), // Spawn meteoriten every 2 seconds
		player:               NewPlayer(),
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
