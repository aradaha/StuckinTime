package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Goblin struct {
	image *ebiten.Image
	pos   vec2
}

func NewGoblin(p vec2, q *ebiten.Image) *Goblin {
	return &Goblin{pos: p, image: q}
}

func (g *Goblin) Update() {

}

func (g *Goblin) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(g.pos.X, g.pos.Y)
	screen.DrawImage(g.image, op)
}
