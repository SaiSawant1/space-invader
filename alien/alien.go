package alien

import rl "github.com/gen2brain/raylib-go/raylib"

var Images [3]rl.Texture2D

type Alien struct {
	Type     int
	Position rl.Vector2
}

func NewAlien(alienType int, position rl.Vector2) *Alien {

	if Images[0].ID == 0 {
		Images[0] = rl.LoadTexture("Graphics/alien_1.png")
		Images[1] = rl.LoadTexture("Graphics/alien_2.png")
		Images[2] = rl.LoadTexture("Graphics/alien_3.png")
	}
	return &Alien{
		Type:     alienType,
		Position: position,
	}
}

func (a *Alien) Draw() {
	rl.DrawTexture(Images[a.Type-1], int32(a.Position.X), int32(a.Position.Y), rl.White)
}

func (a *Alien) GetType() int {
	return a.Type
}

func (a *Alien) Update(direction int) {
	a.Position.X += float32(direction)

}
