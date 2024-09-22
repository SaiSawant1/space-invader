package spaceship

import (
	"time"

	"github.com/SaiSawant1/space-invader/laser"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Spaceship struct {
	image        rl.Texture2D
	position     rl.Vector2
	Lasers       []*laser.Laser
	lastFireTime int64
}

func NewSpaceship() *Spaceship {
	image := rl.LoadTexture("Graphics/spaceship.png")
	return &Spaceship{
		image: image,
		position: rl.Vector2{X: float32((rl.GetScreenWidth() - int(image.Width)) / 2),
			Y: float32((rl.GetScreenHeight() - int(image.Height)))},
		lastFireTime: 0.0,
	}
}

func (sp *Spaceship) Draw() {
	rl.DrawTextureV(sp.image, sp.position, rl.White)
}

func (sp *Spaceship) MoveLeft() {
	sp.position.X = sp.position.X - 10
	if sp.position.X < 0 {
		sp.position.X = 0
	}
}
func (sp *Spaceship) MoveRight() {
	sp.position.X = sp.position.X + 10
	if sp.position.X > (float32(rl.GetScreenWidth()) - float32(sp.image.Width)) {
		sp.position.X = float32(rl.GetScreenWidth()) - float32(sp.image.Width)
	}
}
func (sp *Spaceship) MoveUp() {
	sp.position.Y = sp.position.Y - 10
	if sp.position.Y < 0 {
		sp.position.Y = 0
	}
}
func (sp *Spaceship) MoveDown() {
	sp.position.Y = sp.position.Y + 10
	if sp.position.Y > float32(rl.GetScreenHeight())-float32(sp.image.Height) {
		sp.position.Y = float32(rl.GetScreenHeight()) - float32(sp.image.Height)
	}
}

func (sp *Spaceship) FireLaser() {
	currentTime := time.Now().UnixMilli()

	if currentTime-sp.lastFireTime >= 100 {
		sp.Lasers = append(sp.Lasers, laser.NewLaser(rl.Vector2{X: sp.position.X + float32(sp.image.Width)/2, Y: sp.position.Y}, -6))
		sp.lastFireTime = currentTime
	}
}
