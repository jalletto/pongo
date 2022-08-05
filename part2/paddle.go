package main

import "strings"

type Paddle struct {
	Body Body
}

func (p *Paddle) MoveUp() {
	p.Body.Y -= p.Body.Yspeed
}

func (p *Paddle) MoveDown() {
	p.Body.Y += p.Body.Yspeed
}

func (p *Paddle) Display() string {
	return strings.Repeat("\u2588", p.Body.height)
}
