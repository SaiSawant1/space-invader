package obstacle

import rl "github.com/gen2brain/raylib-go/raylib"

type Obstacle struct {
	image    rl.Texture2D
	position rl.Vector2
}

func NewObstacle() *Obstacle {
	image := rl.LoadTexture("Graphics/alien_1.png")
	return &Obstacle{
		image:    image,
		position: rl.Vector2{X: 100, Y: 100},
	}
}

func (o *Obstacle) Draw() {
	rl.DrawTexture(o.image, int32(o.position.X), int32(o.position.Y), rl.White)

}
