// // Script that will contain game operations around the snake

package game

import (
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
)

func GetSnakeParts(snake objects.Snake) map[[2]int]struct {
	Direction objects.Direction
	Image     string
} {
	/*
	 Taking into account the turning points of the snake, get the multiple parts that
	 compose the snake's body and it's direction so we can adjust the snake's body
	 rendering accordingly.

	 ADD A CLEAR DESCRIPTION OF THE OPERATIONS THAT ARE BEING PERFORMED HERE

	 :param snake: list of the snake's body
	 :return: map with the snake's body parts and it's directions
	*/

	getTurningPointsList := func(snake objects.Snake) [][2]int {
		/*
		 Collect all the TurningPoints from a given snake object.

		 :param snake: list of the snake's body
		 :return: nested list that will contain all the snake's turning points
		*/

		var turningPointsList [][2]int

		for _, turningPoint := range snake.TurningPoints {
			if point, ok := turningPoint["position"].([2]int); ok {
				turningPointsList = append(turningPointsList, point)
			}
		}

		return turningPointsList
	}

	isTurningPoint := func(snakePortion [2]int) bool {
		/*
		 Check if the current part of the snake is a turning point.

		 :param snake: list of the snake's body
		 :param index: index of the current part
		 :return: bool indicating if the current part is a turning point
		*/

		return auxiliars.IsIn(snakePortion, getTurningPointsList(snake))
	}

	responseAppender := func(
		response map[[2]int]struct {
			Direction objects.Direction
			Image     string
		},
		direction objects.Direction,
		coord [2]int,
		image string,
	) {
		/*
		 Append a new body part to the response map which will be the output of the
		 parent function.

		 :param direction: direction of the current part
		 :param coord: coordinate of the current part
		 :param image: image to attach to the current part
		*/

		response[coord] = struct {
			Direction objects.Direction
			Image     string
		}{
			Direction: direction,
			Image:     image,
		}
	}

	turningIndex := -1
	var currentDirection objects.Direction

	response := make(map[[2]int]struct {
		Direction objects.Direction
		Image     string
	})

	for idx, coord := range snake.Body {
		// snake's head
		if idx == 0 {
			responseAppender(response, snake.Direction, coord, "head")
			currentDirection = snake.Direction
		} else if idx == len(snake.Body)-1 {
			// snake's tail
			if len(snake.TurningPoints) > 0 {
				tailDirection := snake.TurningPoints[0]["previous_direction"]
				if tailDirection, ok := tailDirection.(objects.Direction); ok {
					responseAppender(response, tailDirection, coord, "tail")
				}
			} else {
				responseAppender(response, currentDirection, coord, "tail")
			}
		} else {
			// snake's body
			if isTurningPoint(coord) {
				// snake's turning point
				previous_direction := snake.TurningPoints[len(snake.TurningPoints)+turningIndex]["previous_direction"]
				current_direction := snake.TurningPoints[len(snake.TurningPoints)+turningIndex]["current_direction"]

				switch {
				case previous_direction == objects.Down && current_direction == objects.Right:
					responseAppender(response, objects.Right, coord, "curveBody")
				case previous_direction == objects.Down && current_direction == objects.Left:
					responseAppender(response, objects.Up, coord, "curveBody")
				case previous_direction == objects.Left && current_direction == objects.Up:
					responseAppender(response, objects.Right, coord, "curveBody")
				case previous_direction == objects.Left && current_direction == objects.Down:
					responseAppender(response, objects.Down, coord, "curveBody")
				case previous_direction == objects.Up && current_direction == objects.Left:
					responseAppender(response, objects.Left, coord, "curveBody")
				case previous_direction == objects.Up && current_direction == objects.Right:
					responseAppender(response, objects.Down, coord, "curveBody")
				case previous_direction == objects.Right && current_direction == objects.Down:
					responseAppender(response, objects.Left, coord, "curveBody")
				case previous_direction == objects.Right && current_direction == objects.Up:
					responseAppender(response, objects.Up, coord, "curveBody")
				}

				if previous_direction, ok := snake.TurningPoints[len(snake.TurningPoints)+turningIndex]["previous_direction"].(objects.Direction); ok {
					currentDirection = previous_direction
				}

				turningIndex--

			} else {
				// snake's straight body
				responseAppender(response, currentDirection, coord, "body")
			}
		}
	}

	return response
}
