// Script containing all the movements allowed in the game

package moves

import (
	"github.com/jnisa/snake-go/pkg/objects"
)

func UpdateSnake(snake objects.Snake) objects.Snake {
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

	var newCoordinate [][]int

	switch snake.Direction {
	case objects.Up:
		newCoordinate = [][]int{{snake.Body[0][0], snake.Body[0][1] - 1}}
	case objects.Down:
		newCoordinate = [][]int{{snake.Body[0][0], snake.Body[0][1] + 1}}
	case objects.Left:
		newCoordinate = [][]int{{snake.Body[0][0] - 1, snake.Body[0][1]}}
	case objects.Right:
		newCoordinate = [][]int{{snake.Body[0][0] + 1, snake.Body[0][1]}}
	}

	snake.Body = append(newCoordinate, snake.Body...)

	return snake
}

func MoveRight(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake to the right.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the right of the first position of the snake.

	 Note: if the snake is moving left and the user tries to move right, the snake
	 will continue moving left. The only way to change the direction is to move
	 up or down.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the right
	*/

	if snake.Direction != objects.Left {

		snakeHead := snake.Body[0]

		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][]int{{snakeHead[0] + 1, snakeHead[1]}}, snake.Body...)

		snake.Direction = objects.Right

	}

	return snake
}

func MoveLeft(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake to the left.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the left of the first position of the snake.

	 Note: if the snake is moving right and the user tries to move left, the snake
	 will continue moving right. The only way to change the direction is to move
	 up or down.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the left
	*/

	if snake.Direction != objects.Right {

		snakeHead := snake.Body[0]

		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][]int{{snakeHead[0] - 1, snakeHead[1]}}, snake.Body...)

		snake.Direction = objects.Left

	}

	return snake
}

func MoveUp(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake up.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position above the first position of the snake.

	 Note: if the snake is moving down and the user tries to move up, the snake
	 will continue moving down. The only way to change the direction is to move
	 left or right.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving up
	*/

	if snake.Direction != objects.Down {

		snakeHead := snake.Body[0]

		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][]int{{snakeHead[0], snakeHead[1] + 1}}, snake.Body...)

		snake.Direction = objects.Up

	}

	return snake
}

func MoveDown(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake down.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position below the first position of the snake.

	 Note: if the snake is moving up and the user tries to move down, the snake
	 will continue moving up. The only way to change the direction is to move
	 left or right.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving down
	*/

	if snake.Direction != objects.Up {

		snakeHead := snake.Body[0]

		snake.Body = snake.Body[:len(snake.Body)-1]
		snake.Body = append([][]int{{snakeHead[0], snakeHead[1] - 1}}, snake.Body...)

		snake.Direction = objects.Down

	}

	return snake
}
