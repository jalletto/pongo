package main

type Ball struct {
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (b *Ball) Display() string {
	return "\u25CF"
}

func (b *Ball) Update() {
	b.X += b.Xspeed
	b.Y += b.Yspeed
}

func (b *Ball) intersects(p Paddle) bool {
	return b.X >= p.X && b.X <= p.X+p.width && b.Y >= p.Y && b.Y <= p.Y+p.height
}

func (b *Ball) reverseX() {
	b.Xspeed *= -1
}

func (b *Ball) reverseY() {
	b.Yspeed *= -1
}

func (b *Ball) Reset(x int, y int, xSpeed int, ySpeed int) {
	b.X = x
	b.Y = y
	b.Xspeed = xSpeed
	b.Yspeed = ySpeed

}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {

	if b.X <= 0 || b.X >= maxWidth {
		b.reverseY()
	}

	if b.Y <= 0 || b.Y >= maxHeight {
		b.reverseY()
	}
}
