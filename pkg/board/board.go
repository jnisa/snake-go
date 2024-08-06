// Script that contain all the operations related with the board

package board

// import (
// 	"github.com/hajimehoshi/ebiten"
// 	"github.com/jnisa/snake-go/pkg/game/colors"
// 	"github.com/jnisa/snake-go/pkg/objects"
// )

// func (b *objects.Board) Draw(boardImage *ebiten.Image) {
// 	boardImage.Fill(colors.FrameColor)

// 	for j := 0; j < b.size; j++ {
// 		for i := 0; i < b.size; i++ {
// 			v := 0
// 			op := &ebiten.DrawImageOptions{}
// 			x := i*tileSize + (i+1)*tileMargin
// 			y := j*tileSize + (j+1)*tileMargin
// 			op.GeoM.Translate(float64(x), float64(y))
// 			op.ColorScale.ScaleWithColor(tileBackgroundColor(v))
// 			boardImage.DrawImage(tileImage, op)
// 		}
// 	}

// 	animatingTiles := map[*Tile]struct{}{}
// 	nonAnimatingTiles := map[*Tile]struct{}{}

// 	for t := range b.tiles {
// 		if t.IsMoving() {
// 			animatingTiles[t] = struct{}{}
// 		} else {
// 			nonAnimatingTiles[t] = struct{}{}
// 		}
// 	}

// 	for t := range nonAnimatingTiles {
// 		t.Draw(boardImage)
// 	}

// 	for t := range animatingTiles {
// 		t.Draw(boardImage)
// 	}
// }
