// Script containing all the functions affecting the dynamics of the game

//////////////////////////////////////////////////////////////////////////
// --TODO--
// There's a few things to do on this script:
// 1. Add a description to the functions
// 2. Reallocate some of the functions into a game dedicated script
// 3. Add a few more functions to the game script
//////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"image/color"
	_ "image/png" // Register PNG format
	"math"
	"reflect"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/game"
	"github.com/jnisa/snake-go/pkg/moves"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/jnisa/snake-go/pkg/states"
	"golang.org/x/image/font/basicfont"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 640
	BoardHeight  = 32
	boardWidth   = 32
	CellSize     = 20
)

type GameState int

const (
	StateStartScreen GameState = iota
	StateRunning
	StateGameOver
)

type Game struct {
	board              objects.Board
	snake              objects.Snake
	lastUpdate         time.Time
	state              GameState
	foodImage          *ebiten.Image
	snakeHead          *ebiten.Image
	snakeTail          *ebiten.Image
	snakeVerticalRight *ebiten.Image
}

var (
	apple *ebiten.Image = game.ImageLoader("snake_elements/apple.png")
	tail  *ebiten.Image = game.ImageLoader("snake_elements/tail.png")
	head  *ebiten.Image = game.ImageLoader("snake_elements/head.png")
	// straight       *ebiten.Image = game.ImageLoader("snake_elements/straight.png")
	verticalRight *ebiten.Image = game.ImageLoader("snake_elements/curve_1.png")
	// verticalLeft   *ebiten.Image = game.ImageLoader("snake_elements/curve_2.png")
	// horizontalDown *ebiten.Image = game.ImageLoader("snake_elements/curve_3.png")
	// horizontalUp   *ebiten.Image = game.ImageLoader("snake_elements/curve_4.png")

	// TODO. these variables are going to be used for the tail and head of the snake
	Down  float64 = math.Pi / 2
	Up    float64 = -math.Pi / 2
	Right float64 = 0
	Left  float64 = math.Pi
)

func (g *Game) Update() error {
	/*
	 Function that will update the board game whenever there's a key pressed by the
	 user.

	 :param g: current game state
	*/

	if g.state == StateStartScreen {
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = StateRunning
		}
		return nil
	}

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
		if states.IsGameOver(g.snake, g.board) || ebiten.IsKeyPressed(ebiten.KeyEscape) {
			g.state = StateGameOver
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

	switch {
	case g.state == StateStartScreen:
		g.DrawStartScreen(screen)
	case g.state == StateGameOver:
		g.DrawGameOverScreen(screen)
	default:
		g.DrawBoard(screen)
		g.DrawSnake(screen)
		g.DrawFood(screen)
	}
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
	 Render the board on the game screen.

	 :param screen: the screen on which to draw the board.
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

// TODO. the snake is currently a set of pixels and should leverage images
// TODO. this function needs some fine tunning to accomodate the images
// TODO. don't forget to ROTATE the images depending on the direction of the snake
func (g *Game) DrawSnake(screen *ebiten.Image) {
	/*
	 Render the snake on the board game.

	 :param screen: the screen on which to draw the snake.
	*/

	for idx, coord := range g.snake.Body {
		x, y := coord[0]*CellSize, coord[1]*CellSize

		switch idx {
		case 0:
			imageOperations := game.AdjustImage(g.snakeHead, x, y, CellSize, math.Pi/2)
			screen.DrawImage(g.snakeHead, imageOperations)
		case len(g.snake.Body) - 1:
			imageOperations := game.AdjustImage(g.snakeTail, x, y, CellSize, math.Pi/2)
			// imageOperations := game.AdjustImage(g.snakeTail, x, y, CellSize, 90)
			screen.DrawImage(g.snakeTail, imageOperations)
		default:
			for idx, point := range g.snake.TurningPoints {
				if reflect.DeepEqual(coord, point["position"]) {
					switch {
					case g.snake.TurningPoints[idx]["previous_direction"] == objects.Down && g.snake.TurningPoints[idx]["current_direction"] == objects.Right:
						imageOperations := game.AdjustImage(g.snakeVerticalRight, x, y, CellSize, 0)
						screen.DrawImage(g.snakeVerticalRight, imageOperations)
					}
				} else {
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
		}
	}
}

func (g *Game) DrawFood(screen *ebiten.Image) {
	/*
	 Add the food to the game board.

	 This function will add an image of an apple that will represent the food.
	 Additionally, the image will be rescalled to fit the size of the cell.

	 :param screen: the screen on which to draw the food.
	*/

	if (len(g.board.Food)) != 0 {
		x, y := g.board.Food[0]*CellSize, g.board.Food[1]*CellSize

		foodImage := g.foodImage

		op := &ebiten.DrawImageOptions{}
		scaleX := float64(CellSize) / float64(foodImage.Bounds().Dx())
		scaleY := float64(CellSize) / float64(foodImage.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(foodImage, op)
	}
}

// TODO. change the current font to something more arcade looking
func (g *Game) DrawStartScreen(screen *ebiten.Image) {
	/*
	 Render the start screen.

	 This screen will the first facing screen that the user will see when the
	 game is started.

	 :param screen: the screen on which to draw the start screen.
	*/

	msg := "Press SPACE to Start"
	bounds := text.BoundString(basicfont.Face7x13, msg)
	x := (ScreenWidth - bounds.Dx()) / 2
	y := (ScreenHeight - bounds.Dy()) / 2

	text.Draw(screen, msg, basicfont.Face7x13, x, y, color.White)
}

// TODO. change the current font to something more arcade looking
func (g *Game) DrawGameOverScreen(screen *ebiten.Image) {
	/*
	 Render the Game Over screen.

	 This screen will automatically show up when the snake either collides
	 with the walls or bites itself.

	 :param screen: the screen on which to draw the start screen.
	*/

	msg := "Game Over"
	bounds := text.BoundString(basicfont.Face7x13, msg)
	x := (ScreenWidth - bounds.Dx()) / 2
	y := (ScreenHeight - bounds.Dy()) / 2

	text.Draw(screen, msg, basicfont.Face7x13, x, y, color.White)
}

func main() {
	game := &Game{
		snake: objects.Snake{
			Body:      [][]int{{12, 16}, {11, 16}, {10, 16}},
			Direction: objects.Right,
			Score:     0,
		},
		foodImage:          apple,
		snakeHead:          head,
		snakeTail:          tail,
		snakeVerticalRight: verticalRight,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
