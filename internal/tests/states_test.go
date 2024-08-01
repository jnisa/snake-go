// Tests to the states package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/jnisa/snake-go/pkg/states"
	"github.com/stretchr/testify/assert"
)

func TestIsGameOverBasic(t *testing.T) {
	/*
	 Test the isGameOver function when:
	   1. the snake collides with the limits of the board;
	   2. the snake bites itself.
	*/

	var test_board objects.Board

	test_snake_limits := objects.Snake{
		Body:      [][]int{{-1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Left,
	}
	test_snake_bite := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 0}},
		Direction: objects.Down,
	}

	assert.True(t, states.IsGameOver(test_snake_limits, test_board))
	assert.True(t, states.IsGameOver(test_snake_bite, test_board))
}

func TestIsGameOverAdvanced(t *testing.T) {
	/*
	 Test the isGameOver function when there's food in a position that
	 overlaps with the body of the snake - shouldn't be game over.
	*/

	var test_board objects.Board
	test_snake := objects.Snake{
		Body:      [][]int{{9, 10}, {10, 10}, {11, 10}},
		Direction: objects.Left,
	}

	test_board.Cells[10][10] = 1

	assert.False(t, states.IsGameOver(test_snake, test_board))

}

// TODO. check if the documentation needs to be updated afterwards
// TODO. add more tests later
func TestSnakeIngestionUpdate_Up(t *testing.T) {
	/*
	 Test the SnakeIngestionUpdate function - when the snake is moving up.
	*/

	var test_board objects.Board

	test_snake := objects.Snake{
		Body:      [][]int{{10, 11}, {10, 10}, {10, 9}},
		Direction: objects.Up,
		Score:     0,
	}

	// add food to the board
	test_board.Cells[10][11] = 1

	actual := states.SnakeIngestionUpdate(test_snake, test_board)
	expected := objects.Snake{
		Body:      [][]int{{10, 12}, {10, 11}, {10, 10}, {10, 9}},
		Direction: objects.Up,
		Score:     1,
	}

	assert.Equal(t, expected, actual)
}
