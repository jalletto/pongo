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

func (thisBody *Body) intersects(thatBody Body) bool {

	if thatBody.X >= thisBody.X && thatBody.X <= thisBody.X+thisBody.width && thatBody.Y >= thisBody.Y && thatBody.Y <= thisBody.Y+thisBody.height {
		return true
	} else if thisBody.X >= thatBody.X && thisBody.X <= thatBody.X+thatBody.width && thisBody.Y >= thatBody.Y && thisBody.Y <= thatBody.Y+thatBody.height {
		return true
	}
	return false
}
