// Functions that will be used to get the input from the user

package inputs

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

type InputReader interface {
	PollEvent() termbox.Event
}

type MoveOptions struct {
	Up    termbox.Key
	Down  termbox.Key
	Left  termbox.Key
	Right termbox.Key
	Esc   termbox.Key
}

func DefineMoveOptions() MoveOptions {
	/*
	 Define the options that the user can select using the keyboard.

	 The options are defined by the arrow keys and the Esc key.

	 :return: a MoveOptions struct with the values of the arrow keys and the Esc key
	*/

	return MoveOptions{
		Up:    65517,
		Down:  65516,
		Left:  65515,
		Right: 65514,
		Esc:   27,
	}
}

func GetSnakeMove(inputReader InputReader) string {
	/*
	 Get the move that the user wants to apply to the snake.

	 The moves must be selected using the arrows on the keyboard.
	 If there's values outside the arrows, the function will not return any value.
	 Additionally, if the user presses the Esc key, the function will return an
	 empty string.

	 :return: a string with the selected value by the user
	*/

	moves := DefineMoveOptions()

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	fmt.Println("Press arrow keys or Esc to quit")

	for {
		ev := inputReader.PollEvent()

		switch ev.Key {
		case moves.Up:
			return "up"
		case moves.Down:
			return "down"
		case moves.Left:
			return "left"
		case moves.Right:
			return "right"
		case moves.Esc:
			return "quit"
		}
	}
}
