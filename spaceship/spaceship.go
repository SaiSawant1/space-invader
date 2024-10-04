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
	health       int64
}

func NewSpaceship() *Spaceship {
	image := rl.LoadTexture("Graphics/spaceship.png")
	return &Spaceship{
		image: image,
		position: rl.Vector2{X: float32((rl.GetScreenWidth() - int(image.Width)) / 2),
			Y: float32((rl.GetScreenHeight() - int(image.Height) - 100))},
		lastFireTime: 0.0,
		health:       100,
	}
}

func (sp *Spaceship) Draw() {
	rl.DrawTextureV(sp.image, sp.position, rl.White)
}

func (sp *Spaceship) MoveLeft() {
	sp.position.X = sp.position.X - 10
	if sp.position.X < 25 {
		sp.position.X = 25
	}
}
func (sp *Spaceship) MoveRight() {
	sp.position.X = sp.position.X + 10
	if sp.position.X > (float32(rl.GetScreenWidth()) - float32(sp.image.Width) - 25) {
		sp.position.X = float32(rl.GetScreenWidth()) - float32(sp.image.Width) - 25
	}
}

func (sp *Spaceship) FireLaser() {
	currentTime := time.Now().UnixMilli()

	if currentTime-sp.lastFireTime >= 100 {
		sp.Lasers = append(sp.Lasers, laser.NewLaser(rl.Vector2{X: sp.position.X + float32(sp.image.Width)/2, Y: sp.position.Y}, -6))
		sp.lastFireTime = currentTime
	}
}

func (s *Spaceship) GetRect() rl.Rectangle {
	rect := rl.Rectangle{X: s.position.X, Y: s.position.Y, Width: float32(s.image.Width), Height: float32(s.image.Height)}
	return rect
}

func (s *Spaceship) Damage() {
	s.health -= 35
}
func (s *Spaceship) GetCurrentHealth() int64 {
	return s.health
}
