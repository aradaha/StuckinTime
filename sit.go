package main

import (
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

type Game2 struct {
	image     *ebiten.Image
	coins     []*Coin
	player    *Player
	goblin    *Goblin
	tilesheet *Tilesheet
}

func (g *Game2) Init() {
	n := 3

	g.image, _, _ = ebitenutil.NewImageFromFile("assets/Tilesheet/colored_packed.png")
	g.tilesheet = NewTilesheet(g.image, tileSize)
	g.player = NewPlayer(vec2{128, 128}, g.tilesheet.GetTile(19, 7), vec2{320, 240})
	g.goblin = NewGoblin(vec2{64, 64}, g.tilesheet.GetTile(29, 2), vec2{320, 240})
	for i := 0; i < n; i++ {
		g.coins = append(g.coins, NewCoin(vec2{float64(rand.Intn(80) + 20), float64(rand.Intn(80) + 20)}, g.tilesheet.GetTile(41, 3)))
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
