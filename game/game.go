package game

import (
	"github.com/SaiSawant1/space-invader/obstacle"
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	sp        *spaceship.Spaceship
	obstacles []*obstacle.Obstacle
}

func NewGame(sp *spaceship.Spaceship) *Game {
	g := Game{}
	obstacles := g.CreateObstacles()
	return &Game{
		sp:        sp,
		obstacles: obstacles,
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

	for _, obs := range g.obstacles {
		obs.Draw()
	}
}

func (g *Game) deleteInactiveLasers() {

	for index, laser := range g.sp.Lasers {
		if laser.IsActive == false {
			g.sp.Lasers = append(g.sp.Lasers[:index], g.sp.Lasers[index+1:]...)
		}
	}

}

func (g *Game) CreateObstacles() []*obstacle.Obstacle {
	obstacleWidth := len(obstacle.Grid[0]) * 3
	gap := (rl.GetScreenWidth() - (4 * obstacleWidth)) / 5

	for i := 0; i < 4; i++ {
		offsetX := (i+1)*gap + i*obstacleWidth

		position := rl.Vector2{X: float32(offsetX), Y: float32(rl.GetScreenHeight() - 100)}
		obs := obstacle.NewObstacle(position)
		g.obstacles = append(g.obstacles, obs)

	}
	return g.obstacles
}
