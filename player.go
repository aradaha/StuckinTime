package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	image      *ebiten.Image
	pos        vec2
	keypressed bool
	bounds     vec2
}

func NewPlayer(p vec2, q *ebiten.Image, b vec2) *Player {
	return &Player{pos: p, image: q, bounds: b}
}

func (p *Player) Update() {
	// speed := 32.0 / ebiten.ActualTPS()
	var delta vec2

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Println("Attack")
	}

	//WASD
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		delta.X += tileSize
		if p.pos.X > p.bounds.X-2*tileSize {
			delta.X -= tileSize
			fmt.Println("ouch")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		delta.X -= tileSize
		if p.pos.X < tileSize {
			delta.X += tileSize
			fmt.Println("ouch")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		delta.Y -= tileSize
		if p.pos.Y < tileSize {
			delta.Y += tileSize
			fmt.Println("ouch")
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		delta.Y += tileSize
		if p.pos.Y > p.bounds.Y-2*tileSize {
			delta.Y -= tileSize
			fmt.Println("ouch")
		}
	}
	p.pos = delta.Added(p.pos)

	//if delta.X != 0 || delta.Y != 0 {
	//delta = delta.Normalized()
	//delta = delta.Multiply(speed)
	//fmt.Println(delta)

}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(p.pos.X, p.pos.Y)
	screen.DrawImage(p.image, op)
}
