package main

import (
	"image"
	"math/rand"

	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	tileSize = 16
)

type Game2 struct {
	image     *ebiten.Image
	character *ebiten.Image
	coins     []*Coin
	player    *Player
}

func (g *Game2) Init() {
	n := 3

	g.image, _, _ = ebitenutil.NewImageFromFile("assets/Tilesheet/colored_packed.png")
	g.character = g.image.SubImage(image.Rect(19*tileSize, 7*tileSize, 20*tileSize, 8*tileSize)).(*ebiten.Image)
	coinimage := g.image.SubImage(image.Rect(41*tileSize, 3*tileSize, 42*tileSize, 4*tileSize)).(*ebiten.Image)
	g.player = NewPlayer(vec2{100, 100}, g.character)
	for i := 0; i < n; i++ {
		g.coins = append(g.coins, NewCoin(vec2{float64(rand.Intn(80) + 20), float64(rand.Intn(80) + 20)}, coinimage))
	}

}

func (g *Game2) Update() error {

	g.player.Update()
	for _, v := range g.coins {
		v.Update(g.player.pos)
	}

	return nil
}

func (g *Game2) Draw(screen *ebiten.Image) {

	for _, v := range g.coins {
		v.Draw(screen)
	}
	g.player.Draw(screen)
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
