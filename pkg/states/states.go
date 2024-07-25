// Script that contain all the operations related with the multiple states of the game

package states

import (
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
)

func IsGameOver(snake [][]int, board objects.Board) bool {
	/*
	 Check if the game is over.

	 There's two options that might be considered as game over: 1) collision
	 of the snake with one of the walls and 2) collision of the snake with
	 itself.

	 :param snake: list of the snake's body
	 :param board: matrix representing the board
	 :return: boolean indicating if the game is over
	*/

	checkCollision := func(snake [][]int, board objects.Board) bool {
		/*
		 Check if the snake collides with one of the walls.

		 :param snake: list of the snake's body
		 :param board: matrix representing the board
		 :return: boolean indicating if the snake collides with one of the walls
		*/

		if snake[0][0] < 0 || snake[0][0] >= len(board.Cells) || snake[0][1] < 0 || snake[0][1] >= len(board.Cells[0]) {
			return true
		}

		return false
	}

	checkBiteItself := func(snake [][]int) bool {
		/*
		 Check if the snake collides with itself.

		 This can be done by removing the duplicates from the snake's body. If the size
		 of the snake is the different after removing duplicates, then the snake has bitten
		 itself.

		 :param snake: list of the snake's body
		 :return: boolean indicating if the snake collides with itself
		*/

		return len(snake) != len(auxiliars.RemoveDuplicates(snake))
	}

	return checkCollision(snake, board) || checkBiteItself(snake)
}
