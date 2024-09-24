package mysteryship

import (
	"crypto/rand"
	"math/big"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type MysteryShip struct {
	image    rl.Texture2D
	position rl.Vector2
	speed    int32
	Alive    bool
}

func NewMysteryShip() *MysteryShip {
	image := rl.LoadTexture("Graphics/mystery.png")
	return &MysteryShip{
		image: image,
		Alive: false,
	}
}

func (m *MysteryShip) Spawn() {
	m.position.Y = 80
	maxIndex := big.NewInt(2)
	side, err := rand.Int(rand.Reader, maxIndex)
	if err != nil {
		side = big.NewInt(1) // Fallback to 1 on error
	}

	if side.Int64() == 0 {
		m.position.X = 0
		m.speed = 3
	} else {
		m.position.X = float32(rl.GetScreenWidth()) - float32(m.image.Width)
		m.speed = -3
	}
	m.Alive = true
}

func (m *MysteryShip) Update() {
	if m.Alive == true {
		m.position.X += float32(m.speed)
		if m.position.X > float32(rl.GetScreenWidth()-int(m.image.Width)) || m.position.X < 0 {
			m.Alive = false
		}
	}
}

func (m *MysteryShip) Draw() {
	if m.Alive {
		rl.DrawTexture(m.image, int32(m.position.X), int32(m.position.Y), rl.White)

	}
}

func (m *MysteryShip) GetRect() rl.Rectangle {
	if m.Alive {
		rect := rl.Rectangle{X: m.position.X, Y: m.position.Y, Width: float32(m.image.Width), Height: float32(m.image.Height)}
		return rect
	}
	rect := rl.Rectangle{X: m.position.X, Y: m.position.Y, Width: 0, Height: 0}
	return rect
}
