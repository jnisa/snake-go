// Script containing all the functions affecting the dynamics of the game

package main

import (
	"fmt"
	"image/color"
	_ "image/png" // Register PNG format
	"math"
	"sync"
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
	ScreenWidth  = 550
	ScreenHeight = 550
	BoardHeight  = 22
	boardWidth   = 22
	CellSize     = 25
)

type GameState int

const (
	StateStartScreen GameState = iota
	StateRunning
	StateGameOver
)

type Game struct {
	board      objects.Board
	snake      objects.Snake
	lastUpdate time.Time
	state      GameState
	snakeMutex sync.RWMutex
}

var (
	Down  float64 = math.Pi / 2
	Up    float64 = -math.Pi / 2
	Right float64 = 0
	Left  float64 = math.Pi
)

var imageDirection map[objects.Direction]float64 = map[objects.Direction]float64{
	objects.Up:    Up,
	objects.Down:  Down,
	objects.Left:  Left,
	objects.Right: Right,
}

var imageMap map[string]*ebiten.Image = map[string]*ebiten.Image{
	"head":      game.ImageLoader("snake_elements/head.png"),
	"tail":      game.ImageLoader("snake_elements/tail.png"),
	"body":      game.ImageLoader("snake_elements/body.png"),
	"curveBody": game.ImageLoader("snake_elements/curve.png"),
	"apple":     game.ImageLoader("snake_elements/apple.png"),
}

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

	if now.Sub(g.lastUpdate) >= time.Millisecond*100 {

		g.snakeMutex.Lock()
		defer g.snakeMutex.Unlock()

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

func (g *Game) DrawSnake(screen *ebiten.Image) {
	/*
	 Render the snake on the board game.

	 :param screen: the screen on which to draw the snake.
	*/

	g.snakeMutex.RLock()
	defer g.snakeMutex.RUnlock()

	// TODO. Error mentioned on the documentation - this print proves it
	if len(game.GetSnakeParts(g.snake)) < len(g.snake.Body) {
		fmt.Println("Something is wrong")
	}

	for coord, features := range game.GetSnakeParts(g.snake) {

		image := imageMap[features.Image]
		direction := features.Direction

		x, y := coord[0]*CellSize, coord[1]*CellSize

		imageOperations := game.AdjustImage(image, x, y, CellSize, imageDirection[direction])
		screen.DrawImage(image, imageOperations)
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

		foodImage := imageMap["apple"]

		op := &ebiten.DrawImageOptions{}
		scaleX := float64(CellSize) / float64(foodImage.Bounds().Dx())
		scaleY := float64(CellSize) / float64(foodImage.Bounds().Dy())
		op.GeoM.Scale(scaleX, scaleY)
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(foodImage, op)
	}
}

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
			Body:      [][2]int{{12, 16}, {11, 16}, {10, 16}},
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
