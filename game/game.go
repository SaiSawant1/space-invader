package game

import (
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	sp *spaceship.Spaceship
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
	if rl.IsKeyDown(rl.KeySpace) {
		g.sp.FireLaser()
	}
}

func (g *Game) Update() {
	for _, laser := range g.sp.Lasers {
		laser.Update()
	}
	g.deleteInactiveLasers()
}

func (g *Game) Draw() {
	g.sp.Draw()

	for _, laser := range g.sp.Lasers {
		laser.Draw()
	}
}

func (g *Game) deleteInactiveLasers() {

	for index, laser := range g.sp.Lasers {
		if laser.IsActive == false {
			g.sp.Lasers = append(g.sp.Lasers[:index], g.sp.Lasers[index+1:]...)
		}
	}

}
