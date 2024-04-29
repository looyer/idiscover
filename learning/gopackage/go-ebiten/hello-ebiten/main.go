package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func (p *Game) Update() error {
	return nil
}

func (p *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world!")
}

func (p *Game) Layout(w, h int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ltime)
	log.Printf("~~~ hello-ebiten ~~~")

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, world!")

	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
