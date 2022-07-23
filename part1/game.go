package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen tcell.Screen
	Ball   Ball
}

func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()

		width, height := s.Size()

		g.Ball.CheckEdges(width, height)
		g.Ball.Update()
		s.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)

		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}
