// Tests to the objects package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

// evaluate if this test shouldn't be splitted into two different tests
func TestUpdateSnakeBasic(t *testing.T) {
	/*
	 Test the UpdateSnake function when:
	   1. the snake is moving to the right;
	   2. the snake is moving to the up;
	*/

	test_snake_vertical := objects.Snake{
		Body:      [][]int{{0, 1}, {0, 2}, {0, 3}},
		Direction: objects.Up,
	}
	test_snake_horizontal := objects.Snake{
		Body:      [][]int{{4, 3}, {3, 3}, {3, 2}},
		Direction: objects.Right,
	}

	test_snake_vertical = objects.UpdateSnake(test_snake_vertical)
	test_snake_horizontal = objects.UpdateSnake(test_snake_horizontal)

	expected_snake_vertical := objects.Snake{
		Body:      [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}},
		Direction: objects.Up,
	}
	expected_snake_horizontal := objects.Snake{
		Body:      [][]int{{5, 3}, {4, 3}, {3, 3}, {3, 2}},
		Direction: objects.Right,
	}

	assert.Equal(t, expected_snake_vertical, test_snake_vertical)
	assert.Equal(t, expected_snake_horizontal, test_snake_horizontal)
}
