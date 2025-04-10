package main

import "github.com/hajimehoshi/ebiten/v2"

func DrawTo(screen *ebiten.Image, image *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(0, 0)
	screen.DrawImage(image, op)
}

func DrawAt(screen *ebiten.Image, image *ebiten.Image, v vec2) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(v.X,v.Y)
	screen.DrawImage(image, op)
}
