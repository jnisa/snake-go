// Script that contain all the operations related with the appearance of the game elements

package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func AdjustImage(image *ebiten.Image, x, y int, pixelSize int, rotation float64) *ebiten.DrawImageOptions {
	/*
	 Function that is responsible for adjusting the image to the desired size.

	 It acts on two dimensions of the image: the size and position and the rotation.
	 Meaning that the image size will be adjusted according to the pixelSize and, if
	 it needs to be rotated, it will be rotated.

	 :param image: the image to be rescaled
	 :param x: x coordinate position
	 :param y: y coordinate position
	 :param pixelSize: size of the pixel to which the image will be rescaled
	 :param rotation: rotation of the image
	 :return: the rescaled image and draw options
	*/

	width, height := image.Bounds().Dx(), image.Bounds().Dy()

	rotateImage := func(rotation float64, imageOps *ebiten.DrawImageOptions) *ebiten.DrawImageOptions {
		/*
		 Rotate the image by the given angle.

		 :param image: the image to be rotated
		 :param rotation: the angle of rotation
		 :return: the rotated image
		*/

		imageOps.GeoM.Translate(-float64(width)/2, -float64(height)/2)
		imageOps.GeoM.Rotate(rotation)

		return imageOps
	}

	scaleImage := func(pixelSize int, imageOps *ebiten.DrawImageOptions) *ebiten.DrawImageOptions {
		/*
		 Scale the image to the desired size.

		 :param image: the image to be rescaled
		 :param pixelSize: the size of the pixel
		 :return: the rescaled image
		*/

		scaleX := float64(pixelSize) / float64(width)
		scaleY := float64(pixelSize) / float64(height)

		imageOps.GeoM.Scale(scaleX, scaleY)
		imageOps.GeoM.Translate(float64(x), float64(y))

		return imageOps
	}

	op := &ebiten.DrawImageOptions{}

	op = scaleImage(pixelSize, op)

	if rotation != 0 {
		op = rotateImage(rotation, op)
	}

	return op
}
