package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	pos   vec2
}

func NewPlayer(p vec2, q *ebiten.Image) *Player {
	return &Player{pos: p, image: q}
}

func (p *Player) Update() {
	speed := 32.0 / ebiten.ActualTPS()
	var delta vec2
	//WASD
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y += 1
	}

	if delta.X != 0 || delta.Y != 0 {
		delta = delta.Normalized()
		delta = delta.Multiply(speed)
		fmt.Println(delta)
		p.pos = delta.Added(p.pos)
	}

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(p.pos.X, p.pos.Y)
	screen.DrawImage(p.image, op)
}
