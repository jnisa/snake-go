// Script that contain all the operations related with the appearance of the game elements

package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func AdjustImage(image *ebiten.Image, x, y int, pixelSize int, rotation float64) *ebiten.DrawImageOptions {
	/*
	 Function that adjusts the image size, position, and rotation.

	 It resizes the image according to the pixelSize, positions it at (x, y),
	 and rotates it if the rotation value is not zero.

	 :param image: the image to be adjusted
	 :param x: x coordinate position
	 :param y: y coordinate position
	 :param pixelSize: size of the pixel to which the image will be scaled
	 :param rotation: rotation angle of the image in radians
	 :return: draw options containing the transformations
	*/

	width, height := image.Bounds().Dx(), image.Bounds().Dy()

	op := &ebiten.DrawImageOptions{}

	// // Calculate scale factors
	scaleX := float64(pixelSize) / float64(width)
	scaleY := float64(pixelSize) / float64(height)

	// Set up the translation to position the image
	op.GeoM.Translate(float64(-width)/2, float64(-height)/2) // Center it first
	op.GeoM.Rotate(rotation)                                 // Rotate around the center
	op.GeoM.Translate(float64(width)/2, float64(height)/2)   // Translate back to the top-left corner
	op.GeoM.Scale(scaleX, scaleY)                            // Apply scaling
	op.GeoM.Translate(float64(x), float64(y))                // Move to the desired position

	return op
}

func ImageLoader(filepath string) *ebiten.Image {
	/*
	 Load the image from the given file path.

	 This function is also responsible for logging any errors that might occur.

	 :param filepath: the path to the image file
	 :return: the image
	*/

	image, _, err := ebitenutil.NewImageFromFile(filepath)

	if err != nil {
		log.Fatalf("could not read the image file: %v", err)
	}

	return image
}
