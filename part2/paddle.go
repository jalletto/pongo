package main

type Paddle struct {
	width  int
	height int
	Y      int
	X      int
	Xspeed int
	Yspeed int
}

func (p *Paddle) MoveUp() {
	p.Y -= p.Yspeed
}

func (p *Paddle) MoveDown() {
	p.Y += p.Yspeed
}

func (p *Paddle) Display() string {
	return "\u2588\u2588\u2588\u2588"
}
