package game

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/SaiSawant1/space-invader/alien"
	"github.com/SaiSawant1/space-invader/laser"
	mysteryship "github.com/SaiSawant1/space-invader/mystery-ship"
	"github.com/SaiSawant1/space-invader/obstacle"
	"github.com/SaiSawant1/space-invader/spaceship"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	space_ship              *spaceship.Spaceship
	mystery_ship            *mysteryship.MysteryShip
	obstacles               []*obstacle.Obstacle
	aliens                  []*alien.Alien
	aliensDirection         int
	alienLasers             []*laser.Laser
	alienLaserShootInterval int64
	lastAlienFired          int64
	mysteryshipSpawIntercal int64
	mysteryShipLastSpawn    int64
}

func NewGame(sp *spaceship.Spaceship) *Game {
	g := Game{}
	var obstacles []*obstacle.Obstacle
	var aliens []*alien.Alien

	obstacles = g.createObstacles()
	aliens = g.createAliens()
	mp := mysteryship.NewMysteryShip()
	mp.Spawn()

	return &Game{
		space_ship:              sp,
		mystery_ship:            mp,
		obstacles:               obstacles,
		aliens:                  aliens,
		aliensDirection:         1,
		lastAlienFired:          0,
		alienLaserShootInterval: 350,
		mysteryshipSpawIntercal: 20000,
		mysteryShipLastSpawn:    0,
	}
}

func (g *Game) HandleInput() {
	if rl.IsKeyDown(rl.KeyLeft) {
		g.space_ship.MoveLeft()
	}
	if rl.IsKeyDown(rl.KeyRight) {
		g.space_ship.MoveRight()
	}
	if rl.IsKeyDown(rl.KeySpace) {
		g.space_ship.FireLaser()
	}
}

func (g *Game) Update() {
	g.CheckForCollisions()
	currentTime := time.Now().UnixMilli()
	if currentTime-g.mysteryShipLastSpawn > g.mysteryshipSpawIntercal {
		g.mystery_ship.Spawn()
		g.mysteryShipLastSpawn = currentTime
		max := big.NewInt(20000)
		randomInterval, err := rand.Int(rand.Reader, max)
		if err != nil {
			g.mysteryshipSpawIntercal = 10000

		}
		g.mysteryshipSpawIntercal = randomInterval.Int64() + 10000
	}
	for _, laser := range g.space_ship.Lasers {
		laser.Update()
	}
	for _, laser := range g.alienLasers {
		laser.Update()
	}
	g.moveAliens()
	g.alienShotLaser()
	g.mystery_ship.Update()
	g.deleteInactiveLasers()
}

func (g *Game) Draw() {
	g.space_ship.Draw()
	g.mystery_ship.Draw()

	for _, laser := range g.space_ship.Lasers {
		laser.Draw()
	}
	for _, laser := range g.alienLasers {
		laser.Draw()
	}

	for _, obs := range g.obstacles {
		obs.Draw()
	}

	for _, a := range g.aliens {
		a.Draw()
	}
}

func (g *Game) deleteInactiveLasers() {

	for index, laser := range g.space_ship.Lasers {
		if laser.IsActive == false {
			g.space_ship.Lasers = append(g.space_ship.Lasers[:index], g.space_ship.Lasers[index+1:]...)
		}
	}
	for index, laser := range g.alienLasers {
		if laser.IsActive == false {
			g.alienLasers = append(g.alienLasers[:index], g.alienLasers[index+1:]...)
		}
	}

}

func (g *Game) createObstacles() []*obstacle.Obstacle {
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

func (g *Game) createAliens() []*alien.Alien {
	var aliens []*alien.Alien

	for i := 0; i < 5; i++ {
		for j := 0; j < 11; j++ {
			x := 75 + j*55
			y := 110 + i*55
			if i < 1 {
				aliens = append(aliens, alien.NewAlien(1, rl.Vector2{X: float32(x), Y: float32(y)}))
			} else if i < 3 {
				aliens = append(aliens, alien.NewAlien(2, rl.Vector2{X: float32(x), Y: float32(y)}))
			} else {
				aliens = append(aliens, alien.NewAlien(3, rl.Vector2{X: float32(x), Y: float32(y)}))

			}
		}

	}
	return aliens
}

func (g *Game) moveAliens() {

	for _, a := range g.aliens {
		if float64(a.Position.X+float32(alien.Images[a.Type-1].Width)) > float64(rl.GetScreenWidth()) {
			g.aliensDirection = -1
		} else if a.Position.X < 0 {
			g.aliensDirection = 1
			g.MoveDownAliens(4)
		}
		a.Update(g.aliensDirection)
	}

}

func (g *Game) MoveDownAliens(distance int) {
	for _, a := range g.aliens {
		a.Position.Y += float32(distance)
	}
}

func (g *Game) alienShotLaser() {
	currentTime := time.Now().UnixMilli()
	if currentTime-g.lastAlienFired > g.alienLaserShootInterval && len(g.aliens) > 0 {

		maxIndex := big.NewInt(int64(len(g.aliens)))
		randomIndex, err := rand.Int(rand.Reader, maxIndex)
		if err != nil {
			randomIndex = big.NewInt(1) // Fallback to 1 on error
		}
		a := g.aliens[randomIndex.Int64()]
		laserPosition := rl.Vector2{X: (a.Position.X + float32(alien.Images[a.Type-1].Width)/2),
			Y: a.Position.Y + float32(alien.Images[a.Type-1].Height)}
		g.alienLasers = append(g.alienLasers, laser.NewLaser(laserPosition, 6))
		g.lastAlienFired = currentTime
	}
}

func (g *Game) CheckForCollisions() {
	//space ship lasers and aliens
	for laserIndex := len(g.space_ship.Lasers) - 1; laserIndex >= 0; laserIndex-- {
		for alienIndex := len(g.aliens) - 1; alienIndex >= 0; alienIndex-- {
			laser := g.space_ship.Lasers[laserIndex]
			alien := g.aliens[alienIndex]
			if rl.CheckCollisionRecs(laser.GetRect(), alien.GetRect()) {
				g.space_ship.Lasers = append(g.space_ship.Lasers[:laserIndex], g.space_ship.Lasers[laserIndex+1:]...)
				g.aliens = append(g.aliens[:alienIndex], g.aliens[alienIndex+1:]...)
				break
			}
		}
	}
	//space ship lasers and obstacles{
	for laserIndex := len(g.space_ship.Lasers) - 1; laserIndex >= 0; laserIndex-- {
		for obstacleIndex := len(g.obstacles) - 1; obstacleIndex >= 0; obstacleIndex-- {
			obstacle := g.obstacles[obstacleIndex]
			for blockIndex := len(obstacle.Blocks) - 1; blockIndex >= 0; blockIndex-- {
				laser := g.space_ship.Lasers[laserIndex]
				block := obstacle.Blocks[blockIndex]
				if rl.CheckCollisionRecs(block.GetRect(), laser.GetRect()) {
					obstacle.Blocks = append(obstacle.Blocks[:blockIndex], obstacle.Blocks[blockIndex+1:]...)
					laser.IsActive = false
					break
				}
			}
		}
	}
	//alien lasers and spaceship
	for alienLaserIndex := len(g.alienLasers) - 1; alienLaserIndex >= 0; alienLaserIndex-- {
		laser := g.alienLasers[alienLaserIndex]
		if rl.CheckCollisionRecs(laser.GetRect(), g.space_ship.GetRect()) {
			laser.IsActive = false
			g.space_ship.Damage()
		}
	}

	// alien laser and obstacles
	for alienLaserIndex := len(g.alienLasers) - 1; alienLaserIndex >= 0; alienLaserIndex-- {
		for obstacleIndex := len(g.obstacles) - 1; obstacleIndex >= 0; obstacleIndex-- {
			obstacle := g.obstacles[obstacleIndex]
			for blockIndex := len(obstacle.Blocks) - 1; blockIndex >= 0; blockIndex-- {
				alienLaser := g.alienLasers[alienLaserIndex]
				block := obstacle.Blocks[blockIndex]
				if rl.CheckCollisionRecs(block.GetRect(), alienLaser.GetRect()) {
					obstacle.Blocks = append(obstacle.Blocks[:blockIndex], obstacle.Blocks[blockIndex+1:]...)
					alienLaser.IsActive = false
					break
				}
			}
		}
	}

}
