package main

import (
	"image"
	"image/color"
	_ "image/gif"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 800
	screenHeight = 600
	moveSpeed    = 5
	frameWidth   = 46
	frameHeight  = 96
)

type Animation struct {
	frames []*ebiten.Image
	speed  int
}

type Game struct {
	playerX     float64
	playerY     float64
	animations  map[string]*Animation
	currentAnim string
	frameCount  int
	isFlipped   bool
}

func loadAnimation(sprite *ebiten.Image, frameCount int, speed int) *Animation {
	anim := &Animation{
		frames: make([]*ebiten.Image, frameCount),
		speed:  speed,
	}

	for i := 0; i < frameCount; i++ {
		sx := i * frameWidth
		sy := 0
		frame := sprite.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
		anim.frames[i] = frame
	}

	return anim
}

func NewGame() *Game {
	// Load player sprite sheet
	sprite, _, err := ebitenutil.NewImageFromFile("assets/player.gif")
	if err != nil {
		log.Fatal("Error loading image:", err)
	}

	// Create animations
	animations := make(map[string]*Animation)
	animations["idle"] = loadAnimation(sprite, 10, 60) // 10 frames, change every 60 updates
	// Add more animations as needed, e.g.:
	// animations["walk"] = loadAnimation(sprite, 8, 30)
	// animations["jump"] = loadAnimation(sprite, 4, 20)

	return &Game{
		playerX:     screenWidth/2 - frameWidth/2,
		playerY:     screenHeight/2 - frameHeight/2,
		animations:  animations,
		currentAnim: "idle",
	}
}

func (g *Game) Update() error {
	// Handle WASD movement
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += moveSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= moveSpeed
		g.isFlipped = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += moveSpeed
		g.isFlipped = false
	}

	// Keep the player within screen bounds
	if g.playerX < 0 {
		g.playerX = 0
	}
	if g.playerX > screenWidth-frameWidth {
		g.playerX = screenWidth - frameWidth
	}
	if g.playerY < 0 {
		g.playerY = 0
	}
	if g.playerY > screenHeight-frameHeight {
		g.playerY = screenHeight - frameHeight
	}

	// Update animation frame
	g.frameCount++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// background
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	// Get current animation
	anim := g.animations[g.currentAnim]
	if anim == nil {
		return
	}

	// Calculate current frame
	frameIndex := (g.frameCount / anim.speed) % len(anim.frames)
	currentFrame := anim.frames[frameIndex]

	// Draw options for positioning the player
	op := &ebiten.DrawImageOptions{}

	// Handle flipping
	if g.isFlipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(frameWidth, 0)
	}

	op.GeoM.Translate(g.playerX, g.playerY)

	// Draw the player on the screen
	screen.DrawImage(currentFrame, op)
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
