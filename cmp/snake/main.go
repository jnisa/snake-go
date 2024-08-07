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

const (
	ScreenWidth  = 640
	ScreenHeight = 640
	BoardHeight  = 32
	boardWidth   = 32
	CellSize     = 20
)

type Game struct {
	board      objects.Board
	snake      objects.Snake
	lastUpdate time.Time
}

func (g *Game) Update() error {
	/*
	 Function that will update the board game whenever there's a key pressed by the
	 user.

	 :param g: current game state
	*/

	// TODO. add a press to start screen here

	now := time.Now()

	// TODO. change the values that is multiplied by the time.Millisecond
	// depending on the level of the game
	if now.Sub(g.lastUpdate) >= time.Millisecond*100 {
		switch {
		case ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
			moves.MoveLeft(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowRight):
			moves.MoveRight(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowUp):
			moves.MoveUp(&g.snake)
		case ebiten.IsKeyPressed(ebiten.KeyArrowDown):
			moves.MoveDown(&g.snake)
		case inpututil.IsKeyJustPressed(ebiten.KeyEscape):
			return fmt.Errorf("the user has quit the game")
		default:
			moves.UpdateSnake(&g.snake)
		}

		if len(g.board.Food) == 0 {
			x, y := auxiliars.GetRandomPosition(g.snake, g.board)
			g.board.Food = []int{x, y}
		}
		if states.IsGameOver(g.snake, g.board) {
			return fmt.Errorf("Game Over")
		}

		states.SnakeIngestionUpdate(&g.snake, &g.board)
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

	 // TODO. this function needs to consider that there might be food in the board

	 :param screen: The screen on which to draw the board.
	*/

	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < boardWidth; x++ {
			ebitenutil.DrawRect(
				screen,
				float64(x*CellSize),
				float64(y*CellSize),
				float64(CellSize),
				float64(CellSize),
				color.RGBA{116, 116, 116, 255},
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

	if (len(g.board.Food)) != 0 {
		x, y := g.board.Food[0]*CellSize, g.board.Food[1]*CellSize
		ebitenutil.DrawRect(
			screen,
			float64(x),
			float64(y),
			CellSize,
			CellSize,
			color.RGBA{214, 212, 59, 5},
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
