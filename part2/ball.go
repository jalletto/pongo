package main

type Ball struct {
	Body Body
}

func (b *Ball) Display() string {
	return "\u25CF"
}

func (b *Ball) Reset(x int, y int, xSpeed int, ySpeed int) {
	b.Body.X = x
	b.Body.Y = y
	b.Body.Xspeed = xSpeed
	b.Body.Yspeed = ySpeed

}

func (b *Ball) Update() {
	b.Body.X += b.Body.Xspeed
	b.Body.Y += b.Body.Yspeed
}

func (b *Ball) CheckEdges(maxWidth int, maxHeight int) {

	if b.Body.Y <= 0 || b.Body.Y >= maxHeight {
		b.Body.reverseY()
	}
}
