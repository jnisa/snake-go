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
		Direction: objects.Right,
	}

	test_snake = moves.MoveRight(test_snake)

	expected_snake := objects.Snake{
		Body:      [][]int{{1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, test_snake)
}

func TestMoveLeft(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Right,
	}

	test_snake = moves.MoveLeft(test_snake)

	expected_snake := objects.Snake{
		Body:      [][]int{{-1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, test_snake)
}

func TestMoveUp(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 2}, {0, 1}, {0, 0}},
		Direction: objects.Right,
	}

	test_snake = moves.MoveUp(test_snake)

	expected_snake := objects.Snake{
		Body:      [][]int{{0, 3}, {0, 2}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, test_snake)
}

func TestMoveDown(t *testing.T) {
	/*
	 Test the moveRight function when the snake is moving to the right.
	*/

	test_snake := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}},
		Direction: objects.Right,
	}

	test_snake = moves.MoveDown(test_snake)

	expected_snake := objects.Snake{
		Body:      [][]int{{0, -1}, {0, 0}, {0, 1}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake, test_snake)
}
