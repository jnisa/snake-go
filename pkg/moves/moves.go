// Script containing all the movements allowed in the game

package moves

import (
	"github.com/jnisa/snake-go/pkg/objects"
)

func MoveRight(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake to the right.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the right of the first position of the snake.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the right
	*/

	snakeHead := snake.Body[0]

	snake.Body = snake.Body[:len(snake.Body)-1]
	snake.Body = append([][]int{{snakeHead[0] + 1, snakeHead[1]}}, snake.Body...)

	snake.Direction = objects.Right

	return snake
}

func MoveLeft(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake to the left.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position to the left of the first position of the snake.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving to the left
	*/

	snakeHead := snake.Body[0]

	snake.Body = snake.Body[:len(snake.Body)-1]
	snake.Body = append([][]int{{snakeHead[0] - 1, snakeHead[1]}}, snake.Body...)

	snake.Direction = objects.Left

	return snake
}

func MoveUp(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake up.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position above the first position of the snake.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving up
	*/

	snakeHead := snake.Body[0]

	snake.Body = snake.Body[:len(snake.Body)-1]
	snake.Body = append([][]int{{snakeHead[0], snakeHead[1] + 1}}, snake.Body...)

	snake.Direction = objects.Up

	return snake
}

func MoveDown(snake objects.Snake) objects.Snake {
	/*
	 Turn the snake down.

	 In order to accomplish this movement, the following steps should be taken:
	   1. remove the last position of the snake;
	   2. add a new position below the first position of the snake.

	 :param snake: list of the snake's body
	 :return: the snake updated when moving down
	*/

	snakeHead := snake.Body[0]

	snake.Body = snake.Body[:len(snake.Body)-1]
	snake.Body = append([][]int{{snakeHead[0], snakeHead[1] - 1}}, snake.Body...)

	snake.Direction = objects.Down

	return snake
}
