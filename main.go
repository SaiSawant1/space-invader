package main

import (
	"github.com/SaiSawant1/space-invader/game"
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	grey := rl.Color{R: 29, G: 29, B: 27, A: 255}
	yellow := rl.Color{R: 243, G: 216, B: 63, A: 255}

	screenWidth := int32(750)
	screenHeight := int32(700)
	offSet := 50

	rl.InitWindow(screenWidth+int32(offSet), screenHeight+2*int32(offSet), "Space Invader")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	g := game.NewGame(spaceship.NewSpaceship())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		g.HandleInput()
		g.Update()

		rl.ClearBackground(grey)
		rl.DrawRectangleRoundedLines(rl.NewRectangle(10, 10, 780, 780), 0.18, 20, 2, yellow)
		rl.DrawLineEx(rl.Vector2{X: 25, Y: 730}, rl.Vector2{X: 775, Y: 730}, 3, yellow)

		g.Draw()

		rl.EndDrawing()
	}

}
