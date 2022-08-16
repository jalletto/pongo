package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// Set default text style
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	width, height := s.Size()

	ball := Ball{
		X:      5,
		Y:      10,
		Xspeed: 1,
		Yspeed: 1,
	}

	player1 := Player{
		Score: 0,
		Paddle: Paddle{
			width:  1,
			height: 6,
			Y:      3,
			X:      5,
			Yspeed: 3,
		},
	}

	player2 := Player{
		Score: 0,
		Paddle: Paddle{
			width:  1,
			height: 6,
			Y:      3,
			X:      width - 5,
			Yspeed: 3,
		},
	}

	game := Game{
		Screen:  s,
		Ball:    ball,
		Player1: player1,
		Player2: player2,
	}

	go game.Run()

	for {

		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:
			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				game.Player2.Paddle.MoveUp()
			} else if event.Key() == tcell.KeyDown {
				game.Player2.Paddle.MoveDown(height)
			} else if event.Rune() == 'w' {
				game.Player1.Paddle.MoveUp()
			} else if event.Rune() == 's' {
				game.Player1.Paddle.MoveDown(height)
			}
		}
	}

}
