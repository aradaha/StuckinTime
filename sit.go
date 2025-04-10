package main

import (
	"bytes"
	_ "embed"
	"image"
	"log"
	"math/rand"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tileSize = 16
)

var coinscollected int

//go:embed assets/Tilesheet/colored_packed.png
var tilesheetpng []byte

type Game2 struct {
	backgroundtiles [][]*ebiten.Image
	tiles           *ebiten.Image
	image           *ebiten.Image
	coins           []*Coin
	player          *Player
	goblin          *Goblin
	tilesheet       *Tilesheet
}

func (g *Game2) Init() {
	n := 3
	img, _, _ := image.Decode(bytes.NewReader(tilesheetpng))
	g.tiles = ebiten.NewImageFromImage(img)

	g.tilesheet = NewTilesheet(g.tiles, tileSize)
	g.player = NewPlayer(vec2{128, 128}, g.tilesheet.GetTile(19, 7), vec2{320, 240})
	g.goblin = NewGoblin(vec2{64, 64}, g.tilesheet.GetTile(29, 2), vec2{320, 240})
	for i := 0; i < n; i++ {
		g.coins = append(g.coins, NewCoin(vec2{float64(rand.Intn(80) + 20), float64(rand.Intn(80) + 20)}, g.tilesheet.GetTile(41, 3)))
	}

	for x := 0; x < 20; x++ {
		g.backgroundtiles = append(g.backgroundtiles, make([]*ebiten.Image, 0))
		for y := 0; y < 15; y++ {
			tile := g.tilesheet.GetTile(rand.Intn(7), 0)
			//tile := ts.image.SubImage(image.Rect(x*g.tilesize, y*g.tilesize, (x+1)*ts.tilesize, (y+1)*ts.tilesize)).(*ebiten.Image)
			g.backgroundtiles[x] = append(g.backgroundtiles[x], tile)
		}
	}
}

func (g *Game2) Update() error {

	g.player.Update()
	g.goblin.Update(g.player.pos)
	for _, v := range g.coins {
		v.Update(g.player.pos)
	}

	return nil
}

func (g *Game2) Draw(screen *ebiten.Image) {

	for column, v := range g.backgroundtiles {
		for row, u := range v {
			DrawAt(screen, u, vec2{float64(column) * tileSize, float64(row) * tileSize})
		}
	}

	for _, v := range g.coins {
		v.Draw(screen)
	}
	g.goblin.Draw(screen)
	g.player.Draw(screen)
	ebitenutil.DebugPrint(screen, strconv.Itoa(coinscollected))
}

func (g *Game2) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	g := &Game2{}
	g.Init()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
