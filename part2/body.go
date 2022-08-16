package main

type Body struct {
	width  int
	height int
	X      int
	Y      int
	Xspeed int
	Yspeed int
}

func (b *Body) reverseX() {
	b.Xspeed *= -1
}

func (b *Body) reverseY() {
	b.Yspeed *= -1
}

func intersects(thisBody Body, thatBody Body) bool {
	return thatBody.X >= thisBody.X && thatBody.X <= thisBody.X+thisBody.width && thatBody.Y >= thisBody.Y && thatBody.Y <= thisBody.Y+thisBody.height
}

func (thisBody *Body) collides(thatBody Body) bool {
	return intersects(*thisBody, thatBody) || intersects(thatBody, *thisBody)
}
