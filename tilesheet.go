package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tilesheet struct {
	image    *ebiten.Image
	tilesize int
	tiles    [][]*ebiten.Image
}

func NewTilesheet(i *ebiten.Image, t int) *Tilesheet {
	ts := &Tilesheet{tilesize: t, image: i}
	w, h := i.Size()
	xframes := w / ts.tilesize
	yframes := h / ts.tilesize
	for x := 0; x < xframes; x++ {
		ts.tiles = append(ts.tiles, make([]*ebiten.Image, 0))
		for y := 0; y < yframes; y++ {
			tile := ts.image.SubImage(image.Rect(x*ts.tilesize, y*ts.tilesize, (x+1)*ts.tilesize, (y+1)*ts.tilesize)).(*ebiten.Image)
			ts.tiles[x] = append(ts.tiles[x], tile)
		}

	}
	return ts

}

func (ts Tilesheet) GetTile(x int, y int) *ebiten.Image {
	return ts.tiles[x][y]
}
