package main

import (
	"strconv"
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

func (g *Game) GameOver() bool {
	return g.Player1.Score == 5 || g.Player2.Score == 5
}

func (g *Game) DeclareWinner() string {
	if !g.GameOver() {
		return "Game Not Over. No Winner"
	}

	if g.Player1.Score > g.Player2.Score {
		return "Left Player"
	} else {
		return "Right Player"
	}
}

func (g *Game) Run() {
	// set up screen
	s := g.Screen
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	width, height := s.Size()

	for {
		// clear the screen so we can update it
		s.Clear()

		//Determine if the game is over
		if g.GameOver() {
			drawSprite(s, (width/2)-4, 7, (width/2)+5, 7, defStyle, "Game Over")
			drawSprite(s, (width/2)-8, 11, (width/2)+10, 11, defStyle, g.DeclareWinner()+" Wins!")
			s.Show()
		}

		// calculate collision
		if g.Ball.Body.collides(g.Player1.Paddle.Body) || g.Ball.Body.collides(g.Player2.Paddle.Body) {
			g.Ball.Body.reverseX()
			g.Ball.Body.reverseY()
		}

		// update the ball
		g.Ball.CheckEdges(width, height)
		g.Ball.Update()
		drawSprite(s,
			g.Ball.Body.X,
			g.Ball.Body.Y,
			g.Ball.Body.X,
			g.Ball.Body.Y,
			defStyle,
			g.Ball.Display())

		// update the players
		drawSprite(s,
			g.Player1.Paddle.Body.X,
			g.Player1.Paddle.Body.Y,
			g.Player1.Paddle.Body.X+g.Player1.Paddle.Body.width,
			g.Player1.Paddle.Body.Y+g.Player1.Paddle.Body.height,
			defStyle,
			g.Player1.Paddle.Display())

		drawSprite(s,
			g.Player2.Paddle.Body.X,
			g.Player2.Paddle.Body.Y,
			g.Player2.Paddle.Body.X+g.Player2.Paddle.Body.width,
			g.Player2.Paddle.Body.Y+g.Player2.Paddle.Body.height,
			defStyle,
			g.Player2.Paddle.Display())

		// update and display the score
		if g.Ball.Body.X <= 0 {
			g.Player2.Score++
			g.Ball.Reset(width/2, height/2, -1, 1)
		}

		if g.Ball.Body.X >= width {
			g.Player1.Score++
			g.Ball.Reset(width/2, height/2, 1, 1)
		}

		drawSprite(s, (width/2)-5, 1, 1, 1, defStyle, strconv.Itoa(g.Player1.Score))
		drawSprite(s, (width/2)+5, 1, 1, 1, defStyle, strconv.Itoa(g.Player2.Score))

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
