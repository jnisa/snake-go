// // Main Heart of the snake game

// package main

// import (
// 	"fmt"
// 	"time"

// 	"github.com/jnisa/snake-go/pkg/auxiliars"
// 	"github.com/jnisa/snake-go/pkg/inputs"
// 	"github.com/jnisa/snake-go/pkg/moves"
// 	"github.com/jnisa/snake-go/pkg/objects"
// 	"github.com/jnisa/snake-go/pkg/states"
// )

// func main() {

// 	var board objects.Board

// 	snake := objects.Snake{
// 		Body:      [][]int{{12, 16}, {11, 16}, {10, 16}},
// 		Direction: objects.Right,
// 		Score:     0,
// 	}

// 	ticker := time.NewTicker(500 * time.Millisecond) // Snake moves every 500ms
// 	defer ticker.Stop()

// 	done := make(chan bool)
// 	exit := make(chan bool)

// 	go func() {

// 		defer close(done)

// 		// get a new food position and add it to the board
// 		food_x, food_y := auxiliars.GetRandomPosition(snake, board)
// 		board.Cells[food_x][food_y] = 1

// 		for {

// 			select {
// 			case <-exit:
// 				fmt.Println("Exiting gamehelp...")
// 				return
// 			case <-ticker.C:

// 				// update snake's body in case there's any ingestion
// 				snake, board = states.SnakeIngestionUpdate(snake, board)

// 				// check if the game is over
// 				if states.IsGameOver(snake, board) {
// 					fmt.Println("Game Over")
// 					break
// 				}

// 				// check if the board has any food
// 				if !auxiliars.BoardContainsValue(1, board) {
// 					food_x, food_y = auxiliars.GetRandomPosition(snake, board)
// 					board.Cells[food_x][food_y] = 1
// 				}

// 				move := inputs.GetSnakeMove(inputs.InputReader(nil))

// 				// if a key was pressed by the user, then move the snake in the given direction
// 				// if no key was pressed by the user, then move the snake in the last direction
// 				switch move {
// 				case "up":
// 					snake = moves.MoveUp(snake)
// 				case "down":
// 					snake = moves.MoveDown(snake)
// 				case "left":
// 					snake = moves.MoveLeft(snake)
// 				case "right":
// 					snake = moves.MoveRight(snake)
// 				case "esc":
// 					exit <- true
// 				case "":
// 					snake = moves.UpdateSnake(snake)
// 				}
// 			}
// 		}
// 	}()

// 	<-done
// }

package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct{}

func (g *Game) Update() error {
	// Check if the arrow keys are pressed
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		fmt.Println("Left arrow key pressed")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		fmt.Println("Right arrow key pressed")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		fmt.Println("Up arrow key pressed")
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		fmt.Println("Down arrow key pressed")
	}

	// Check if the space key was just pressed
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		fmt.Println("Space key just pressed")
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// You can draw your game objects here
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // return the logical screen size
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Key Press Detection Example")

	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
