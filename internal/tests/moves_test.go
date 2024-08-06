// Tests performed to the moves package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/moves"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

// TODO. split this tests into multiple tests
func TestUpdateSnakeBasic(t *testing.T) {
	/*
	 Test the UpdateSnake function when:
	   1. the snake is moving to the right;
	   2. the snake is moving to the up;
	*/

	actual_snake_vertical := objects.Snake{
		Body:      [][]int{{0, 1}, {0, 2}, {0, 3}},
		Direction: objects.Up,
	}
	actual_snake_horizontal := objects.Snake{
		Body:      [][]int{{4, 3}, {3, 3}, {3, 2}},
		Direction: objects.Right,
	}

	moves.UpdateSnake(&actual_snake_vertical)
	moves.UpdateSnake(&actual_snake_horizontal)

	expected_snake_vertical := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Up,
	}
	expected_snake_horizontal := objects.Snake{
		Body:      [][]int{{5, 3}, {4, 3}, {3, 3}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake_vertical, actual_snake_vertical)
	assert.Equal(t, expected_snake_horizontal, actual_snake_horizontal)
}

func TestMoveRight_basic(t *testing.T) {
	/*
	 Test the MoveRight function when the snake is moving Down.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	moves.MoveRight(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveRight_exception(t *testing.T) {
	/*
	 Test the MoveRight function when the snake is moving to the left.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Left,
	}

	moves.MoveRight(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Left,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveLeft_basic(t *testing.T) {
	/*
	 Test the MoveLeft function when the snake is moving Down.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	moves.MoveLeft(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{-1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Left,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveLeft_exception(t *testing.T) {
	/*
	 Test the MoveLeft function when the snake is moving to the right.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Right,
	}

	moves.MoveLeft(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveUp_basic(t *testing.T) {
	/*
	 Test the MoveUp function when the snake is moving Left.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{15, 10}, {16, 10}, {17, 10}},
		Direction: objects.Left,
	}

	moves.MoveUp(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{15, 9}, {15, 10}, {16, 10}},
		Direction: objects.Up,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveUp_exception(t *testing.T) {
	/*
	 Test the MoveUp function when the snake is moving down.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	moves.MoveUp(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveDown_basic(t *testing.T) {
	/*
	 Test the MoveDown function when the snake is moving to the right.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{1, 0}, {2, 0}, {3, 0}},
		Direction: objects.Left,
	}

	moves.MoveDown(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{1, 1}, {1, 0}, {2, 0}},
		Direction: objects.Down,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveDown_exception(t *testing.T) {
	/*
	 Test the MoveDown function when the snake is moving Up.
	*/

	actual_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Up,
	}

	moves.MoveDown(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Up,
	}

	assert.Equal(t, expected_snake, actual_snake)
}
