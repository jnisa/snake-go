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
	 Test the isGameOver function when the snake collides witht one of the walls.
	*/

	var test_board objects.Board
	test_snake := [][]int{{-1, 0}, {0, 0}, {1, 0}}

	assert.True(t, states.IsGameOver(test_snake, test_board))
}

// TODO. Uncommennt the following tests after knowing how to run tests
// func TestIsGameOverAdvanced(t *testing.T) {
// 	/*
// 	 Test the isGameOver function when the snake collides with itself.
// 	*/
// }

// func TestIsGameOverException(t *testing.T) {
// 	/*
// 	 Test the isGameOver function when the snake is in a valid position.
// 	*/
// }
