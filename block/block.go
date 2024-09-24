package block

import rl "github.com/gen2brain/raylib-go/raylib"

type Block struct {
	position rl.Vector2
}

func NewBlock(position rl.Vector2) *Block {
	return &Block{
		position: position,
	}
}

func (b *Block) Draw() {
	color := rl.Color{243, 216, 63, 255}
	rl.DrawRectangle(int32(b.position.X), int32(b.position.Y), 3, 3, color)
}

func (b *Block) GetRect() rl.Rectangle {
	rect := rl.Rectangle{X: b.position.X, Y: b.position.Y, Height: 3, Width: 3}
	return rect
}
