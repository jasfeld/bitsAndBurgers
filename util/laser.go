package util

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

const (
	laserSpeedPerSecond = 350.0
)

var laser = MustLoadImage("assets/laser.png")

type Laser struct {
	position Vector
	bild     *ebiten.Image
}

func NewLaser(playerPos Vector, playerBounds image.Rectangle) *Laser {
	halfW := float64(playerBounds.Dx()) / 2
	halfH := float64(playerBounds.Dy()) / 2

	playerPos.X += halfW
	playerPos.Y -= halfH

	laserBounds := laser.Bounds()
	halfW = float64(laserBounds.Dx()) / 2
	halfH = float64(laserBounds.Dy()) / 2

	playerPos.X -= halfW
	playerPos.Y -= halfH

	b := &Laser{
		position: playerPos,
		bild:     laser,
	}

	return b
}

func (b *Laser) Update() {
	speed := laserSpeedPerSecond / float64(ebiten.TPS())
	b.position.Y += -speed
}

func (b *Laser) Draw(screen *ebiten.Image) {
	bounds := b.bild.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(b.position.X+halfW, b.position.Y+halfH)

	screen.DrawImage(b.bild, op)
}

func (b *Laser) KollisionsRechteck() Rect {
	bounds := b.bild.Bounds()

	return NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
