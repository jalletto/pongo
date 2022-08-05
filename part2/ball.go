package main

type Ball struct {
	Body Body
}

func (b *Ball) Display() string {
	return "\u25CF"
}

func (b *Ball) Update() {
	b.Body.X += b.Body.Xspeed
	b.Body.Y += b.Body.Yspeed
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	// if b.Body.X <= 0 || b.Body.X >= maxWidth {
	// 	b.Body.reverseX()
	// }

	if b.Body.Y <= 0 || b.Body.Y >= maxHeight {
		b.Body.reverseY()
	}
}
