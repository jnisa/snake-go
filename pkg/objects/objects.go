// Script containing the definition of all the game objects

package objects

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Board struct {
	Cells [24][32]int
}

type Snake struct {
	Direction Direction
	Body      [][]int
}

func UpdateSnake(snake Snake) Snake {
	/*
		Add a new coordinate to the snake's body.

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
	case Up:
		newCoordinate = [][]int{{snake.Body[0][0], snake.Body[0][1] - 1}}
	case Down:
		newCoordinate = [][]int{{snake.Body[0][0], snake.Body[0][1] + 1}}
	case Left:
		newCoordinate = [][]int{{snake.Body[0][0] - 1, snake.Body[0][1]}}
	case Right:
		newCoordinate = [][]int{{snake.Body[0][0] + 1, snake.Body[0][1]}}
	}

	snake.Body = append(newCoordinate, snake.Body...)

	return snake
}
