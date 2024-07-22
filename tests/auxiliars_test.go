// Test the functions from the auxiliar package.

package tests

import (
	"testing"
	"github.com/jnisa/snake-go/src"
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

	assert.Equal(t, 0, src.Get(0, 0, test_board))
	assert.Equal(t, 1, src.Get(1, 1, test_board))
	assert.Equal(t, -1, src.Get(3, 3, test_board))
}

// func TestGet_exception(t *testing.T) {
// 	/*
// 	 Test the Get function when no board is provided.
// 	*/

// 	test_board := [][]int{
// 		{0, 0, 0},
// 		{0, 1, 0},
// 		{0, 0, 0},
// 	}

// 	assert.Equal(t, 0, auxiliars.Get(0, 0, test_board))
// 	assert.Equal(t, 1, auxiliars.Get(1, 1, test_board))
// 	assert.Equal(t, -1, auxiliars.Get(3, 3, test_board))
// }