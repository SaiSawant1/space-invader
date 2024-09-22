package game

import (
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	sp *spaceship.Spaceship
}

func (g *Game) Draw() {
	g.sp.Draw()
}

func NewGame(sp *spaceship.Spaceship) *Game {
	return &Game{
		sp: sp,
	}
}

func (g *Game) HandleInput() {
	if rl.IsKeyDown(rl.KeyLeft) {
		g.sp.MoveLeft()
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.sp.MoveRight()
	}
	if rl.IsKeyDown(rl.KeyUp) {
		g.sp.MoveUp()
	}
	if rl.IsKeyDown(rl.KeyDown) {
		g.sp.MoveDown()
	}
}
