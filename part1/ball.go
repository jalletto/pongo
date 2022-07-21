package main

type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (b *Ball) Display() rune {
	return '\u25CF'
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {
	if b.X <= 0 || b.X >= maxWidth {
		b.Xspeed *= -1
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.Yspeed *= -1
	}
}
