// Script containing all the functions affecting the dynamics of the game

//////////////////////////////////////////////////////////////////////////
// --TODO--
// There's a few things to do on this script:
// 1. Add a description to the functions
// 2. Reallocate some of the functions into a game dedicated script
// 3. Add a few more functions to the game script
// 4. Add a clock to the game
//////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/moves"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/jnisa/snake-go/pkg/states"
)

// TODO. check if these dimensions are ok
// TODO. there must be a disalignment the board array and the dimensions of the screen
const (
	ScreenWidth  = 1000
	ScreenHeight = 1000
	BoardHeight  = 100
	boardWidth   = 100
	CellSize     = 20
)

type Game struct {
	board      objects.Board
	snake      objects.Snake
	lastUpdate time.Time
	// boardImage *ebiten.Image
}

func (g *Game) Update() error {
	/*
	 Function that will update the board game whenever there's a key pressed by the
	 user.

	 :param g: current game state
	*/

	now := time.Now()

	if now.Sub(g.lastUpdate) >= time.Second {

		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			fmt.Println("Left arrow pressed!")
			moves.MoveLeft(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			fmt.Println("Right arrow pressed!")
			moves.MoveRight(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
			fmt.Println("Up arrow pressed!")
			moves.MoveUp(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
			fmt.Println("Down arrow pressed!")
			moves.MoveDown(&g.snake)
		case inpututil.IsKeyJustPressed(ebiten.KeyEscape):
			fmt.Println("Quitting the game...")
			return fmt.Errorf("the user has quit the game")

		// TODO. not sure if this should stay out of the switch
		case states.IsGameOver(g.snake, g.board):
			return fmt.Errorf("Game Over")
		default:
			moves.UpdateSnake(&g.snake)
		}

		g.lastUpdate = now
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*
	 Draw the multiple game components.

	 For now the components only include a snake and the board where that snake moves, eventually
	 in the future, we might add scores and level to this.

	 :param screen: ADD A DESCRIPITON HERE...
	*/

	g.DrawBoard(screen)
	g.DrawSnake(screen)
	g.DrawFood(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	/*
	 Add a description for this function here....

	 :param outsideWidth: ADD A DESCRIPTION
	 :param outsideHeight: ADD A DESCRIPTION
	 :return: ADD A DESCRIPTION
	*/

	return ScreenWidth, ScreenHeight
}

func (g *Game) DrawBoard(screen *ebiten.Image) {
	/*
	 Render the board on the front-end.

	 :param screen: ADD A DESCRIPTION HERE...
	*/

	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			ebitenutil.DrawRect(
				screen,
				float64(x*CellSize),
				float64(y*CellSize),
				float64(CellSize),
				float64(CellSize),
				color.RGBA{116, 116, 116, 0},
			)
		}
	}
}

func (g *Game) DrawSnake(screen *ebiten.Image) {
	/*
	 Render the snake on the front-end.

	 :param screen: ADD A DESCRIPTION HERE...
	*/

	for _, p := range g.snake.Body {
		x, y := p[0]*CellSize, p[1]*CellSize
		ebitenutil.DrawRect(
			screen,
			float64(x),
			float64(y),
			CellSize,
			CellSize,
			color.RGBA{255, 0, 0, 255},
		)
	}
}

func (g *Game) DrawFood(screen *ebiten.Image) {
	/*
	 Render the food on the front-end.

	 :param screen: ADD A DESCRIPTION HERE...
	*/

	if auxiliars.BoardContainsValue(1, g.board) {
		x, y := auxiliars.GetRandomPosition(g.snake, g.board)
		ebitenutil.DrawRect(
			screen,
			float64(x),
			float64(y),
			CellSize,
			CellSize,
			color.RGBA{183, 185, 5, 0},
		)
	}
}

func main() {

	game := &Game{
		snake: objects.Snake{
			Body:      [][]int{{12, 16}, {11, 16}, {10, 16}},
			Direction: objects.Right,
			Score:     0,
		},
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
