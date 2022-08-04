package main

import (
	"github.com/gdamore/tcell/v2"
	"time"
)

type Game struct {
	Screen       tcell.Screen
	Ball         Ball
	Player1      Paddle
	Player2      Paddle
	eventChannel chan string
}

func drawSprite(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	row := y1
	col := x1
	for _, r := range []rune(text) {
		s.SetContent(col, row, r, nil, style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}

func (g *Game) Run() {
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)

	for {

		s.Clear()

		//ball
		width, height := s.Size()
		g.Ball.CheckEdges(width, height)
		g.Ball.Update()

		// s.SetContent(g.Ball.X, g.Ball.Y, g.Ball.Display(), nil, defStyle)
		drawSprite(s, g.Ball.X, g.Ball.Y, g.Ball.X, g.Ball.Y, defStyle, g.Ball.Display())

		// Paddle
		drawSprite(s, g.Player1.X, g.Player1.Y, g.Player1.X+g.Player1.width, g.Player1.Y+g.Player1.height, defStyle, g.Player1.Display())
		drawSprite(s, g.Player2.X, g.Player2.Y, g.Player2.X+g.Player2.width, g.Player2.Y+g.Player2.height, defStyle, g.Player2.Display())

		//channel

		select {
		case msg := <-g.eventChannel:
			switch msg {
			case "up":
				g.Player2.MoveUp()
			case "down":
				g.Player2.MoveDown()
			case "w":
				g.Player1.MoveUp()
			case "s":
				g.Player1.MoveDown()
			}

		default:
			width = 0
		}

		time.Sleep(40 * time.Millisecond)
		s.Show()
	}

}
