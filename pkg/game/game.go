// Script that contain function that will affect the game

package game

// import (
// 	"fmt"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// 	"github.com/jnisa/snake-go/pkg/moves"
// 	"github.com/jnisa/snake-go/pkg/objects"
// )

// // TODO. check if these dimensions are ok
// const (
// 	ScreenWidth  = 480
// 	ScreenHeight = 480
// )

// type Game struct {
// 	board      objects.Board
// 	snake      objects.Snake
// 	boardImage *ebiten.Image
// }

// func (g *Game) Update() error {
// 	/*
// 	 Function that will update the board game whenever there's a key pressed by the
// 	 user.

// 	 :param g: current game state
// 	*/

// 	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
// 		moves.MoveLeft(g.snake)
// 		// fmt.Println("Left arrow pressed!")
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
// 		moves.MoveRight(g.snake)
// 		// fmt.Println("Right arrow pressed!")
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
// 		moves.MoveUp(g.snake)
// 		// fmt.Println("Up arrow pressed!")
// 	}
// 	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
// 		moves.MoveDown(g.snake)
// 		// fmt.Println("Down arrow pressed!")
// 	}
// 	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
// 		fmt.Println("Quitting the game...")

// 	}

// 	// TODO. if no key is pressed then update the board with the last Direction
// 	// TODO. not sure if other functions should be added here as well (e.g. IsGameOver, etc...)

// 	return nil
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	/*
// 	 Add a description for this function here....

// 	 :param outsideWidth: ADD A DESCRIPTION
// 	 :param outsideHeight: ADD A DESCRIPTION
// 	 :return: ADD A DESCRIPTION
// 	*/

// 	return ScreenWidth, ScreenHeight
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	/*
// 	 Add a description for this function here...

// 	 :param g: ADD A DESCRIPTION
// 	 :param screen: ADD A DESCRIPITON
// 	*/

// 	if g.boardImage == nil {
// 		// TODO. revisit the following line
// 		g.boardImage = ebiten.NewImage(g.board.Cells)
// 	}

// 	screen.Fill(BackgroundColor)
// 	g.board.Draw(g.boardImage)
// 	op := &ebiten.DrawImageOptions{}

// 	sw, sh := screen.Bounds().Dx(), screen.Bounds().Dy()
// 	bw, bh := g.boardImage.Bounds().Dx(), g.boardImage.Bounds().Dy()
// 	x := (sw - bw) / 2
// 	y := (sh - bh) / 2
// 	op.GeoM.Translate(float64(x), float64(y))

// 	screen.DrawImage(g.boardImage, op)
// }
