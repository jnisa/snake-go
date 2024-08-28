// Script containing all the movements allowed in the game

package moves

import (
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
)

// TODO. when updating the snake body we need to check if there's isn't coordinates that
// should be removed from the turning points list
func UpdateSnake(snake *objects.Snake) {
	/*
	 Add a new coordinate to the snake's body when no movement is detected.

	 Adding a new coordinate to the snake's body will be higly dependent on the
	 current direction of the snake. The new coordinate will be added to the
	 beginning of the snake's body.

	 I.e.
	 - If the snake is moving up, the new coordinate will be the one above the
	 first body cell;
	 - If the snake is moving down, the new coordinate will be the one below the
	 first body cell;
	 - If the snake is moving left, the new coordinate will be the one to the left
	 of the first body cell;
	 - If the snake is moving right, the new coordinate will be the one to the right
	 of the first body cell.

	 :param snake: list of the snake's body
	 :return: the snake updated
	*/

	removeIndex := func(index int, list []map[string]interface{}) []map[string]interface{} {
		/*
		 Remove an element from a list of elements from a given index.

		 If the index provided is out of bounds, then the list will be returned as is.

		 :param index: index of the element to be removed
		 :param list: list of elements
		 :return: list of elements with the element removed
		*/

		if index < 0 || index >= len(list) {
			return list
		} else {
			return append(list[:index], list[index+1:]...)
		}
	}

	isTurningPointOutFunc := func(snake *objects.Snake) {
		/*
		 Check if the turning points are still part of the snake's body.

		 Whenever the snake changes its direction, the head of the snake is added to
		 the turning points list. This function will check if the turning points are
		 still part of the snake's body. If they are not, then they will be removed
		 from the list.

		 :param snake: list of the snake's body
		 :return: the snake updated
		*/

		for idx, turningPoint := range snake.TurningPoints {

			turningPointPosition := turningPoint["position"].([2]int)

			if !auxiliars.IsIn(turningPointPosition, snake.Body) {
				snake.TurningPoints = removeIndex(idx, snake.TurningPoints)
			}
		}
	}

	var newCoordinate [][2]int

	switch snake.Direction {
	case objects.Up:
		newCoordinate = [][2]int{{snake.Body[0][0], snake.Body[0][1] - 1}}
	case objects.Down:
		newCoordinate = [][2]int{{snake.Body[0][0], snake.Body[0][1] + 1}}
	case objects.Left:
		newCoordinate = [][2]int{{snake.Body[0][0] - 1, snake.Body[0][1]}}
	case objects.Right:
		newCoordinate = [][2]int{{snake.Body[0][0] + 1, snake.Body[0][1]}}
	}

	snake.Body = append(newCoordinate, snake.Body...)
	snake.Body = snake.Body[:len(snake.Body)-1]

	isTurningPointOutFunc(snake)

	// Just for testing purposes remove afterwards
	// fmt.Println("Number of turning points: ", len(snake.TurningPoints))
	// for _, turningPoint := range snake.TurningPoints {
	// 	if point, ok := turningPoint["position"].([]int); ok {
	// 		fmt.Println(point)
	// 	}
	// }
}

func MoveRight(snake *objects.Snake) {
	/*
	 Turn the snake to the right.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the right of the first position of the snake.

	 Note: if the snake is moving left and the user tries to move right, the snake
	 will continue moving left. The only way to change the direction is to move
	 up or down.

	 Additionally, the head of the snake will be added to the turning points list.
	 This will be usefukl to determine the image that will be applied to the snake's
	 body.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the right
	*/

	if snake.Direction != objects.Left {

		// get the coordinates of the snake's head
		snakeHead := snake.Body[0]

		// remove the last position of the snake
		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][2]int{{snakeHead[0] + 1, snakeHead[1]}}, snake.Body...)

		// add the head to the turning points list
		snake.TurningPoints = append(snake.TurningPoints, map[string]interface{}{
			"position":           snakeHead,
			"previous_direction": snake.Direction,
			"current_direction":  objects.Right,
		})

		// update the snake's direction
		snake.Direction = objects.Right
	}
}

func MoveLeft(snake *objects.Snake) {
	/*
	 Turn the snake to the left.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the left of the first position of the snake.

	 Note: if the snake is moving right and the user tries to move left, the snake
	 will continue moving right. The only way to change the direction is to move
	 up or down.

	 Additionally, the head of the snake will be added to the turning points list.
	 This will be usefukl to determine the image that will be applied to the snake's
	 body.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the left
	*/

	if snake.Direction != objects.Right {

		// get the coordinates of the snake's head
		snakeHead := snake.Body[0]

		// remove the last position of the snake
		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][2]int{{snakeHead[0] - 1, snakeHead[1]}}, snake.Body...)

		// add the head to the turning points list
		snake.TurningPoints = append(snake.TurningPoints, map[string]interface{}{
			"position":           snakeHead,
			"previous_direction": snake.Direction,
			"current_direction":  objects.Left,
		})

		// update the snake's direction
		snake.Direction = objects.Left
	}
}

func MoveUp(snake *objects.Snake) {
	/*
	 Turn the snake up.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position above the first position of the snake.

	 Note: if the snake is moving down and the user tries to move up, the snake
	 will continue moving down. The only way to change the direction is to move
	 left or right.

	 Additionally, the head of the snake will be added to the turning points list.
	 This will be usefukl to determine the image that will be applied to the snake's
	 body.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving up
	*/

	if snake.Direction != objects.Down {

		// get the coordinates of the snake's head
		snakeHead := snake.Body[0]

		// remove the last position of the snake
		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][2]int{{snakeHead[0], snakeHead[1] - 1}}, snake.Body...)

		// add the head to the turning points list
		snake.TurningPoints = append(snake.TurningPoints, map[string]interface{}{
			"position":           snakeHead,
			"previous_direction": snake.Direction,
			"current_direction":  objects.Up,
		})

		// update the snake's direction
		snake.Direction = objects.Up
	}
}

func MoveDown(snake *objects.Snake) {
	/*
	 Turn the snake down.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position below the first position of the snake.

	 Note: if the snake is moving up and the user tries to move down, the snake
	 will continue moving up. The only way to change the direction is to move
	 left or right.

	 Additionally, the head of the snake will be added to the turning points list.
	 This will be usefukl to determine the image that will be applied to the snake's
	 body.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving down
	*/

	if snake.Direction != objects.Up {

		// get the coordinates of the snake's head
		snakeHead := snake.Body[0]

		// remove the last position of the snake
		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][2]int{{snakeHead[0], snakeHead[1] + 1}}, snake.Body...)

		// add the head to the turning points list
		snake.TurningPoints = append(snake.TurningPoints, map[string]interface{}{
			"position":           snakeHead,
			"previous_direction": snake.Direction,
			"current_direction":  objects.Down,
		})

		// update the snake's direction
		snake.Direction = objects.Down
	}
}
