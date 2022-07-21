package main

import (
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
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

	ball := Ball{
		X:      5,
		Y:      10,
		Xspeed: 1,
		Yspeed: 1,
	}

	game := Game{
		Screen: s,
		Ball:   ball,
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
			}
		}
	}

}
