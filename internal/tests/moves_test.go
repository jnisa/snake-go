// ADD A DESCRIPTION HERE

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/moves"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

func TestMoveRight(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	actual_snake := moves.MoveRight(test_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveLeft(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Down,
	}

	actual_snake := moves.MoveLeft(test_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{-1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Left,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveUp(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{15, 10}, {16, 10}, {17, 10}},
		Direction: objects.Left,
	}

	actual_snake := moves.MoveUp(test_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{15, 11}, {15, 10}, {16, 10}},
		Direction: objects.Up,
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveDown(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Right,
	}

	actual_snake := moves.MoveDown(test_snake)
	expected_snake := objects.Snake{
		Body:      [][]int{{0, -1}, {0, 0}, {0, 1}},
		Direction: objects.Down,
	}

	assert.Equal(t, expected_snake, actual_snake)
}
