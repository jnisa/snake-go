// Test the functions from the auxiliar package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/auxiliars"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

func TestGetRandomPositionBasic(t *testing.T) {
	/*
	 Test the GetRandomPosition function when the snake is a single cell and the
	 board is a 3x3 matrix.
	*/

	test_snake := objects.Snake{
		Body:      [][2]int{{1, 1}},
		Direction: objects.Up,
	}

	var test_board objects.Board

	x, y := auxiliars.GetRandomPosition(test_snake, test_board)

	assert.NotEqual(t, 1, x)
	assert.NotEqual(t, 1, y)
}

func TestGetRandomPositionException(t *testing.T) {
	/*
	 Test the GetRandomPosition function when the board is empty.
	*/

	test_snake := objects.Snake{
		Body:      [][2]int{},
		Direction: objects.Up,
	}
	var test_board objects.Board

	x, y := auxiliars.GetRandomPosition(test_snake, test_board)

	assert.True(t, x >= 0 && x < len(test_board.Cells) && y >= 0 && y < len(test_board.Cells[0]))
}

func TestRemoveDuplicatesBasic(t *testing.T) {
	/*
	 Remove duplicates from a list of elements.
	*/

	test_elements := [][2]int{
		{1, 1}, {1, 1}, {1, 2},
		{1, 2}, {1, 3}, {1, 3},
	}

	actual := auxiliars.RemoveDuplicates(test_elements)
	expected := [][2]int{{1, 1}, {1, 2}, {1, 3}}

	assert.Equal(t, expected, actual)
}

func TestRemoveDuplicatesException(t *testing.T) {
	/*
	 When an empty list is provided.
	*/

	test_elements := [][2]int{}

	actual := auxiliars.RemoveDuplicates(test_elements)
	expected := [][2]int{}

	assert.Equal(t, actual, expected)
}

func TestIsIn(t *testing.T) {
	/*
	 Check if the ContainsValue function performs in two cases:
	  1. when there's a value in the nested list;
	  2. when there's not a value in the nested list;
	  3. when the nested list is empty.
	*/

	var empty_list [][2]int
	non_empty_lst := [][2]int{{1, 2}, {3, 4}, {5, 6}}

	assert.True(t, auxiliars.IsIn([2]int{1, 2}, non_empty_lst))
	assert.False(t, auxiliars.IsIn([2]int{1, 1}, non_empty_lst))
	assert.False(t, auxiliars.IsIn([2]int{1, 1}, empty_list))
}
