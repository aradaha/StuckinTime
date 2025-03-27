package main

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Goblin struct {
	image       *ebiten.Image
	pos         vec2
	defeated    bool
	bounds      vec2
	originalpos vec2
}

func NewGoblin(p vec2, q *ebiten.Image, b vec2) *Goblin {
	arst := &Goblin{}
	arst.pos = p
	arst.image = q
	arst.bounds = b

	return arst
	//return &Goblin{pos: p, image: q}
}

func (g *Goblin) Update(playerpos vec2) {
	var delta vec2
	direction := rand.Intn(4)

	if playerpos != g.originalpos {
		if direction == 0 {
			if g.pos.X < g.bounds.X-tileSize {
				delta.X += tileSize
			}
			// goes right
		}
		if direction == 1 {
			if g.pos.X > 0 {
				delta.X -= tileSize
			}
			// goes left
		}
		if direction == 2 {
			if g.pos.Y > 0 {
				delta.Y -= tileSize
			}
			// goes up
		}
		if direction == 3 {
			if g.pos.Y < g.bounds.Y-tileSize {
				delta.Y += tileSize
			}
			// goes down
		}
		g.pos = delta.Added(g.pos)
		if g.pos.DistanceTo(playerpos) < tileSize && !g.defeated {
			g.defeated = true
		}
	}

	// save player position
	// check player position against saved value
	// move if different

	g.originalpos = playerpos
}

func (g *Goblin) Draw(screen *ebiten.Image) {
	if !g.defeated {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(1, 1)
		op.GeoM.Translate(g.pos.X, g.pos.Y)
		screen.DrawImage(g.image, op)
	}

}
