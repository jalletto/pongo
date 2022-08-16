package main

import "strings"

type Paddle struct {
	width  int
	height int
	X      int
	Y      int
	Yspeed int
}

func (p *Paddle) Display() string {
	return strings.Repeat("\u2588", p.height)
}

func (p *Paddle) MoveUp() {

	if p.Y > 0 {
		p.Y -= p.Yspeed
	}
}

func (p *Paddle) MoveDown(windowHeight int) {
	if p.Y < windowHeight-p.height {
		p.Y += p.Yspeed
	}
}
