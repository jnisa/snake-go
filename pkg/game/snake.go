// Script that will contain game operations around the snake

package game

import (
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
)

type turnKey struct {
	prev, curr objects.Direction
}

var turnCombinations = map[turnKey]objects.Direction{
	/*
	 Variable containing all the possible movements that can be exectured by
	 the snake.
	*/

	{objects.Down, objects.Right}: objects.Right,
	{objects.Down, objects.Left}:  objects.Up,
	{objects.Left, objects.Up}:    objects.Right,
	{objects.Left, objects.Down}:  objects.Down,
	{objects.Up, objects.Left}:    objects.Left,
	{objects.Up, objects.Right}:   objects.Down,
	{objects.Right, objects.Down}: objects.Left,
	{objects.Right, objects.Up}:   objects.Up,
}

func GetSnakeParts(snake objects.Snake) map[[2]int]struct {
	Direction objects.Direction
	Image     string
} {
	/*
	 Taking into account the turning points of the snake, get the multiple parts that
	 compose the snake's body and it's direction so we can adjust the snake's body
	 rendering accordingly.

	 The response of this function will be map that will have the snake's body coordinates
	 as keys and a struct composed by the direction of such body part and the image that
	 must be rendered for such body part.

	 To ease the process of appending new body parts to the response map, a nested function
	 is defined and passed as argument to the current function.

	 :param snake: list of the snake's body
	 :return: map with the snake's body parts and its direction and image
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
		if idx == 0 {
			// snake's head
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

				if previous_direction, ok_prev := previous_direction.(objects.Direction); ok_prev {
					if current_direction, ok_curr := current_direction.(objects.Direction); ok_curr {
						responseAppender(response, turnCombinations[turnKey{previous_direction, current_direction}], coord, "curveBody")
					}
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
