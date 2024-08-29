// Script that contain all the operations related with the multiple states of the game

package states

import (
	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
)

func IsGameOver(snake objects.Snake, board objects.Board) bool {
	/*
	 Check if the game is over.

	 There's two options that might be considered as game over: 1) collision
	 of the snake with one of the walls and 2) collision of the snake with
	 itself.

	 :param snake: list of the snake's body
	 :param board: matrix representing the board
	 :return: boolean indicating if the game is over
	*/

	checkCollision := func(snake objects.Snake, board objects.Board) bool {
		/*
		 Check if the snake collides with one of the walls.

		 :param snake: list of the snake's body
		 :param board: matrix representing the board
		 :return: boolean indicating if the snake collides with one of the walls
		*/

		if snake.Body[0][0] < 0 || snake.Body[0][0] >= len(board.Cells) || snake.Body[0][1] < 0 || snake.Body[0][1] >= len(board.Cells[0]) {
			return true
		}

		return false
	}

	checkBiteItself := func(snake objects.Snake) bool {
		/*
		 Check if the snake collides with itself.

		 This can be done by removing the duplicates from the snake's body. If the size
		 of the snake is the different after removing duplicates, then the snake has bitten
		 itself.

		 :param snake: list of the snake's body
		 :return: boolean indicating if the snake collides with itself
		*/

		return len(snake.Body) != len(auxiliars.RemoveDuplicates(snake.Body))
	}

	return checkCollision(snake, board) || checkBiteItself(snake)
}

func SnakeIngestionUpdate(snake *objects.Snake, board *objects.Board) {
	/*
	 If there's an ingestion of food, update the snake's body accordingly.

	 :param snake: list of the snake's body
	 :param board: matrix representing the board
	 :return: the snake updated with a new tile in the body
	*/

	isIngestion := func(snake objects.Snake, board objects.Board) bool {
		/*
			Check if the snake has reached a position where there's food.

			:param snake: list of the snake's body
			:param board: matrix representing the board
			:return: boolean indicating if the snake has reached a position where
			there's food
		*/

		return snake.Body[0][0] == board.Food[0] && snake.Body[0][1] == board.Food[1]
	}

	addNewTile := func(snake *objects.Snake) {
		/*
		 Add a new tile to the snake's body.

		 The new tile will be added to the end of the snake's body. This is done
		 by getting the last tile of the snake's body and adding a new tile to
		 the last position of the snake's body.

		 To make that happen, the Direction of the snake must be taken into account.

		 :param snake: list of the snake's body
		 :return: the snake updated with a new tile in the body
		*/

		snakeTail := snake.Body[len(snake.Body)-1]
		snakeDirection := snake.Direction
		snake.Score += 1

		switch snakeDirection {
		case objects.Up:
			snake.Body = append(snake.Body, [2]int{snakeTail[0], snakeTail[1] - 1})
		case objects.Down:
			snake.Body = append(snake.Body, [2]int{snakeTail[0], snakeTail[1] + 1})
		case objects.Left:
			snake.Body = append(snake.Body, [2]int{snakeTail[0] - 1, snakeTail[1]})
		case objects.Right:
			snake.Body = append(snake.Body, [2]int{snakeTail[0] + 1, snakeTail[1]})
		}
	}

	if isIngestion(*snake, *board) {
		var newFoodPosition []int
		board.Food = newFoodPosition
		addNewTile(snake)
	}
}
