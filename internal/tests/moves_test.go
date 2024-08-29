// Tests performed to the moves package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/moves"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSnakeBasic(t *testing.T) {
	/*
		Test the UpdateSnake function when:
		1. the snake is moving to the right;
		2. the snake is moving to the up;
	*/

	// Variables containing the test snakes.
	var (
		test_snake_1 = objects.Snake{
			Body:      [][2]int{{0, 1}, {0, 2}, {0, 3}},
			Direction: objects.Up,
			TurningPoints: []map[string]interface{}{
				{
					"position":           [2]int{0, 3},
					"previous_direction": objects.Down,
					"current_direction":  objects.Up,
				},
			},
		}

		test_snake_2 = objects.Snake{
			Body:      [][2]int{{4, 3}, {3, 3}, {3, 2}},
			Direction: objects.Right,
		}
	)

	// Variables containing the expected snakes.
	var (
		expected_snake_1 = objects.Snake{
			Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
			Direction:     objects.Up,
			TurningPoints: []map[string]interface{}{},
		}

		expected_snake_2 = objects.Snake{
			Body:      [][2]int{{5, 3}, {4, 3}, {3, 3}},
			Direction: objects.Right,
		}
	)

	moves.UpdateSnake(&test_snake_1)
	moves.UpdateSnake(&test_snake_2)

	assert.Equal(t, expected_snake_1, test_snake_1)
	assert.Equal(t, expected_snake_2, test_snake_2)
}

func TestMoveRight_basic(t *testing.T) {
	/*
	 Test the MoveRight function when the snake is moving Down.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Down,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveRight(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][2]int{{1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Right,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{0, 0},
				"previous_direction": objects.Down,
				"current_direction":  objects.Right,
			},
		}}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveRight_exception(t *testing.T) {
	/*
	 Test the MoveRight function when the snake is moving to the left.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Left,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveRight(&actual_snake)
	expected_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Left,
		TurningPoints: []map[string]interface{}{},
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveLeft_basic(t *testing.T) {
	/*
	 Test the MoveLeft function when the snake is moving Down.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Down,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveLeft(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][2]int{{-1, 0}, {0, 0}, {0, 1}},
		Direction: objects.Left,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{0, 0},
				"previous_direction": objects.Down,
				"current_direction":  objects.Left,
			},
		}}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveLeft_exception(t *testing.T) {
	/*
	 Test the MoveLeft function when the snake is moving to the right.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Right,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveLeft(&actual_snake)
	expected_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Right,
		TurningPoints: []map[string]interface{}{},
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveUp_basic(t *testing.T) {
	/*
	 Test the MoveUp function when the snake is moving Left.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{15, 10}, {16, 10}, {17, 10}},
		Direction:     objects.Left,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveUp(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][2]int{{15, 9}, {15, 10}, {16, 10}},
		Direction: objects.Up,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{15, 10},
				"previous_direction": objects.Left,
				"current_direction":  objects.Up,
			},
		}}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveUp_exception(t *testing.T) {
	/*
	 Test the MoveUp function when the snake is moving down.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Down,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveUp(&actual_snake)
	expected_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Down,
		TurningPoints: []map[string]interface{}{},
	}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveDown_basic(t *testing.T) {
	/*
	 Test the MoveDown function when the snake is moving to the right.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{1, 0}, {2, 0}, {3, 0}},
		Direction:     objects.Left,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveDown(&actual_snake)
	expected_snake := objects.Snake{
		Body:      [][2]int{{1, 1}, {1, 0}, {2, 0}},
		Direction: objects.Down,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{1, 0},
				"previous_direction": objects.Left,
				"current_direction":  objects.Down,
			},
		}}

	assert.Equal(t, expected_snake, actual_snake)
}

func TestMoveDown_exception(t *testing.T) {
	/*
	 Test the MoveDown function when the snake is moving Up.
	*/

	actual_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Up,
		TurningPoints: []map[string]interface{}{},
	}

	moves.MoveDown(&actual_snake)
	expected_snake := objects.Snake{
		Body:          [][2]int{{0, 0}, {0, 1}, {0, 2}},
		Direction:     objects.Up,
		TurningPoints: []map[string]interface{}{},
	}

	assert.Equal(t, expected_snake, actual_snake)
}
