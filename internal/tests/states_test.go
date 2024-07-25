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
	test_snake_limits := [][]int{{-1, 0}, {0, 0}, {1, 0}}
	test_snake_bite := [][]int{{0, 0}, {0, 1}, {0, 0}}

	assert.True(t, states.IsGameOver(test_snake_limits, test_board))
	assert.True(t, states.IsGameOver(test_snake_bite, test_board))
}

func TestIsGameOverAdvanced(t *testing.T) {
	/*
	 Test the isGameOver function when there's food in a position that
	 overlaps with the body of the snake - shouldn't be game over.
	*/

	var test_board objects.Board
	test_snake := [][]int{{9, 10}, {10, 10}, {11, 10}}

	test_board.Cells[10][10] = 1

	assert.False(t, states.IsGameOver(test_snake, test_board))

}

func TestIsIngestionBasic(t *testing.T) {
	/*
	 Test the isIngestion function when:
	   1. the snake reaches a position where there's food;
	   2. the snake reaches a position where there's something that's not food.
	*/

	var test_board objects.Board
	test_snake := [][]int{{9, 10}, {10, 10}, {11, 10}}

	test_board.Cells[10][10] = 1
	test_board.Cells[11][10] = 2

	assert.True(t, states.IsIngestion(test_snake, test_board))
	assert.True(t, states.IsIngestion(test_snake, test_board))
}
