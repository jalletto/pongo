package main

import (
	"fmt"
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

	eventChannel := make(chan string)
	width, _ := s.Size()

	ball := Ball{
		X:      5,
		Y:      10,
		Xspeed: 1,
		Yspeed: 1,
	}

	player1 := Paddle{
		width:  1,
		height: 5,
		Y:      10,
		X:      3,
		Yspeed: 3,
	}

	player2 := Paddle{
		width:  1,
		height: 5,
		Y:      10,
		X:      width - 3,
		Yspeed: 3,
	}

	game := Game{
		Screen:       s,
		Ball:         ball,
		Player1:      player1,
		Player2:      player2,
		eventChannel: eventChannel,
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
