package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	image *ebiten.Image
	pos   vec2
}

func Player(p vec2, q *ebiten.Image) *Player {
	return &Player{pos: p, image: q}
}

func (c *Player) Update(playerpos vec2) {
	if c.pos.DistanceTo(playerpos) < 16 && !c.collected {
		fmt.Println("You collected a Player. Yay ")
		c.collected = true
	}
}

func (c *Player) Draw(screen *ebiten.Image) {
	if !c.collected {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(c.pos.X, c.pos.Y)
		screen.DrawImage(c.image, op)
	}
}
