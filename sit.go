package main

import (
	"fmt"
	"image"

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
	coin      *Coin
	Playerpos vec2
}

func (g *Game2) Init() {
	g.image, _, _ = ebitenutil.NewImageFromFile("assets/Tilesheet/colored_packed.png")
	g.character = g.image.SubImage(image.Rect(19*tileSize, 7*tileSize, 20*tileSize, 8*tileSize)).(*ebiten.Image)
	coinimage := g.image.SubImage(image.Rect(41*tileSize, 3*tileSize, 42*tileSize, 4*tileSize)).(*ebiten.Image)
	g.coin = NewCoin(vec2{20, 20}, coinimage)
	fmt.Println(g.coin)

}

func (g *Game2) Update() error {
	speed := 32.0 / ebiten.ActualTPS()
	var delta vec2
	//WASD
	if ebiten.IsKeyPressed(ebiten.KeyRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		delta.X += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		delta.X -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		delta.Y -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		delta.Y += 1
	}

	if delta.X != 0 || delta.Y != 0 {
		delta = delta.Normalized()
		delta = delta.Multiply(speed)
		fmt.Println(delta)
		g.Playerpos = delta.Added(g.Playerpos)
	}
	g.coin.Update(g.Playerpos)

	return nil
}

func (g *Game2) Draw(screen *ebiten.Image) {
	g.coin.Draw(screen)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(g.Playerpos.X, g.Playerpos.Y)
	screen.DrawImage(g.character, op)

}

func (g *Game2) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (v *vec2) Add2(a vec2) {
	v.X += a.X
	v.Y += a.Y
	fmt.Println("Inside:", v)
}

func main() {
	w := &vec2{100, 100}
	fmt.Println("Before: ", w)
	w.Add2(vec2{50, 50})
	fmt.Println("After: ", w)
	g := &Game2{Playerpos: vec2{100, 100}}
	g.Init()
	fmt.Println(g.image)

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
