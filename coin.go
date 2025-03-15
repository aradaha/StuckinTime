package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Coin struct {
	image     *ebiten.Image
	pos       vec2
	collected bool
}

func NewCoin(p vec2, q *ebiten.Image) *Coin {
	return &Coin{pos: p, image: q}
}

func (c *Coin) Update(playerpos vec2) {
	if c.pos.DistanceTo(playerpos) < 16 && !c.collected {
		fmt.Println("You collected a coin. Yay ")
		c.collected = true
	}
}

func (c *Coin) Draw(screen *ebiten.Image) {
	if !c.collected {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(0.75, 0.75)
		op.GeoM.Translate(c.pos.X, c.pos.Y)
		screen.DrawImage(c.image, op)
	}
}
