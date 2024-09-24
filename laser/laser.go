package laser

import rl "github.com/gen2brain/raylib-go/raylib"

type Laser struct {
	position rl.Vector2
	speed    int
	IsActive bool
}

func NewLaser(position rl.Vector2, speed int) *Laser {
	return &Laser{
		position: position,
		speed:    speed,
		IsActive: true,
	}
}

func (l Laser) Draw() {
	yello := rl.Color{243, 216, 63, 255}
	rl.DrawRectangle(int32(l.position.X), int32(l.position.Y), 4, 15, yello)
}

func (l *Laser) Update() {
	l.position.Y += float32(l.speed)
	if l.position.Y < 0 || l.position.Y > float32(rl.GetScreenHeight()) {
		l.IsActive = false
	}
}

func (l *Laser) GetRect() rl.Rectangle {
	return rl.Rectangle{
		X:      l.position.X,
		Y:      l.position.Y,
		Height: 15,
		Width:  4,
	}
}
