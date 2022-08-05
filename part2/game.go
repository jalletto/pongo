package main

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

type Game struct {
	Screen       tcell.Screen
	Ball         Ball
	Player1      Player
	Player2      Player
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
		// clear the screen so we can update it
		s.Clear()

		// calculate collision
		if g.Ball.Body.intersects(g.Player1.Paddle.Body) || g.Ball.Body.intersects(g.Player2.Paddle.Body) {
			g.Ball.Body.reverseX()
			g.Ball.Body.reverseY()
		}

		// update the ball
		width, height := s.Size()
		g.Ball.CheckEdges(width, height)
		g.Ball.Update()
		drawSprite(s, g.Ball.Body.X, g.Ball.Body.Y, g.Ball.Body.X, g.Ball.Body.Y, defStyle, g.Ball.Display())

		// update the players
		drawSprite(s, g.Player1.Paddle.Body.X,
			g.Player1.Paddle.Body.Y,
			g.Player1.Paddle.Body.X+g.Player1.Paddle.Body.width,
			g.Player1.Paddle.Body.Y+g.Player1.Paddle.Body.height,
			defStyle,
			g.Player1.Paddle.Display())

		drawSprite(s, g.Player2.Paddle.Body.X,
			g.Player2.Paddle.Body.Y,
			g.Player2.Paddle.Body.X+g.Player2.Paddle.Body.width,
			g.Player2.Paddle.Body.Y+g.Player2.Paddle.Body.height,
			defStyle,
			g.Player2.Paddle.Display())

		// update the score
		// TODO

		//Determine if the game is over
		// TODO

		// Update the screen
		time.Sleep(50 * time.Millisecond)
		s.Show()

		//read from input channel
		select {
		case msg := <-g.eventChannel:
			switch msg {
			case "up":
				g.Player2.Paddle.MoveUp()
			case "down":
				g.Player2.Paddle.MoveDown()
			case "w":
				g.Player1.Paddle.MoveUp()
			case "s":
				g.Player1.Paddle.MoveDown()
			}
		default:
			continue
		}

	}

}
