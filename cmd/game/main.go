package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
	boxSize      = 50
	moveSpeed    = 5
)

type Game struct {
	boxX float64
	boxY float64
}

func NewGame() *Game {
	return &Game{
		boxX: screenWidth/2 - boxSize/2, // Start in the middle of the screen
		boxY: screenHeight/2 - boxSize/2,
	}
}

func (g *Game) Update() error {
	// Handle WASD movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.boxY -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.boxY += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.boxX -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.boxX += moveSpeed
	}

	// Keep the box within screen bounds
	if g.boxX < 0 {
		g.boxX = 0
	}
	if g.boxX > screenWidth-boxSize {
		g.boxX = screenWidth - boxSize
	}
	if g.boxY < 0 {
		g.boxY = 0
	}
	if g.boxY > screenHeight-boxSize {
		g.boxY = screenHeight - boxSize
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Create a new image for the box
	box := ebiten.NewImage(boxSize, boxSize)
	box.Fill(color.RGBA{R: 255, G: 255, B: 0, A: 255}) // Red box

	// Draw options for positioning the box
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.boxX, g.boxY)

	// Draw the box on the screen
	screen.DrawImage(box, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("ComixZone - WASD to move")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
