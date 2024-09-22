package main

import (
	"github.com/SaiSawant1/space-invader/game"
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	grey := rl.Color{R: 29, G: 29, B: 27, A: 255}

	screenWidth := int32(750)
	screenHeight := int32(700)

	rl.InitWindow(screenWidth, screenHeight, "Space Invader")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := game.NewGame(spaceship.NewSpaceship())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		g.HandleInput()
		g.Update()

		rl.ClearBackground(grey)

		g.Draw()

		rl.EndDrawing()
	}

}
