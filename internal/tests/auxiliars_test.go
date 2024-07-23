// Test the functions from the auxiliar package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/stretchr/testify/assert"
)

func TestGet_basic(t *testing.T) {
	/*
	 Test the Get function when the board is a 3x3 matrix and the coordinates
	 provided are both within and outside the board.
	*/

	test_board := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	assert.Equal(t, 0, auxiliars.Get(0, 0, test_board))
	assert.Equal(t, 1, auxiliars.Get(1, 1, test_board))
	assert.Equal(t, -1, auxiliars.Get(3, 3, test_board))
}

func TestGet_exception(t *testing.T) {
	/*
	 Test the Get function when no board is provided.
	*/

	test_board := [][]int{}

	assert.Equal(t, -1, auxiliars.Get(1, 1, test_board))
}

func TestGetRandomPosition_basic(t *testing.T) {
	/*
	 Test the GetRandomPosition function when the snake is a single cell and the
	 board is a 3x3 matrix.
	*/

	test_snake := [][]int{
		{1, 1},
	}

	test_board := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	x, y := auxiliars.GetRandomPosition(test_snake, test_board)

	assert.NotEqual(t, 1, x)
	assert.NotEqual(t, 1, y)
}

func TestGetRandomPosition_exception(t *testing.T) {
	/*
	 Test the GetRandomPosition function when the board is empty.
	*/

	test_snake := [][]int{}
	test_board := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	x, y := auxiliars.GetRandomPosition(test_snake, test_board)

	assert.True(t, x >= 0 && x < len(test_board) && y >= 0 && y < len(test_board[0]))
}

func TestRemoveDuplicates_basic(t *testing.T) {
	/*
	 Remove duplicates from a list of elements.
	*/

	test_elements := [][]int{
		{1, 1}, {1, 1}, {1, 2},
		{1, 2}, {1, 3}, {1, 3},
	}

	actual := auxiliars.RemoveDuplicates(test_elements)
	expected := [][]int{{1, 1}, {1, 2}, {1, 3}}

	assert.Equal(t, expected, actual)
}

func TestRemoveDuplicates_exception(t *testing.T) {
	/*
	 When an empty list is provided.
	*/

	test_elements := [][]int{}

	actual := auxiliars.RemoveDuplicates(test_elements)
	expected := [][]int{}

	assert.Equal(t, actual, expected)
}
