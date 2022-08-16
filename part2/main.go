package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

func main() {

	// initialize our tcell screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	// set default text style (text, background color etc.)
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	s.SetStyle(defStyle)

	// Set up and run our game
	eventChannel := make(chan string)
	width, height := s.Size()

	ball := Ball{
		Body{
			X:      width / 2,
			Y:      height / 2,
			Xspeed: 1,
			Yspeed: 1,
		},
	}

	player1 := Player{
		Score: 0,
		Paddle: Paddle{
			Body{

				width:  1,
				height: 6,
				Y:      3,
				X:      5,
				Yspeed: 3,
			},
		},
	}

	player2 := Player{
		Score: 0,
		Paddle: Paddle{
			Body{
				width:  1,
				height: 6,
				Y:      3,
				X:      width - 5,
				Yspeed: 3,
			},
		},
	}

	game := Game{
		Screen:       s,
		Ball:         ball,
		Player1:      player1,
		Player2:      player2,
		eventChannel: eventChannel,
	}

	go game.Run()

	// Event Loop to listen for user input
	for {

		switch event := game.Screen.PollEvent().(type) {
		case *tcell.EventResize:
			game.Screen.Sync()
		case *tcell.EventKey:

			if event.Key() == tcell.KeyEscape || event.Key() == tcell.KeyCtrlC {
				game.Screen.Fini()
				os.Exit(0)
			} else if event.Key() == tcell.KeyUp {
				eventChannel <- "up"
			} else if event.Key() == tcell.KeyDown {
				eventChannel <- "down"
			} else if event.Rune() == 'w' {
				eventChannel <- "w"
			} else if event.Rune() == 's' {
				eventChannel <- "s"
			}

		}

	}

}
