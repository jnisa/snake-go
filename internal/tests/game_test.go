// Tests performed to the game package.

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/game"
	"github.com/jnisa/snake-go/pkg/objects"
	"github.com/stretchr/testify/assert"
)

var (
	noCurveSnake = objects.Snake{
		Body:          [][2]int{{10, 9}, {10, 10}, {10, 11}, {10, 12}},
		Direction:     objects.Down,
		TurningPoints: []map[string]interface{}{},
	}

	oneCurveSnake = objects.Snake{
		Body:      [][2]int{{10, 9}, {10, 10}, {10, 11}, {11, 11}, {12, 11}},
		Direction: objects.Down,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{10, 11},
				"previous_direction": objects.Left,
				"current_direction":  objects.Down,
			},
		},
	}

	twoCurveSnake = objects.Snake{
		Body:      [][2]int{{10, 9}, {10, 10}, {10, 11}, {11, 11}, {12, 11}, {12, 10}, {12, 9}},
		Direction: objects.Down,
		TurningPoints: []map[string]interface{}{
			{
				"position":           [2]int{12, 11},
				"previous_direction": objects.Up,
				"current_direction":  objects.Left,
			},
			{
				"position":           [2]int{10, 11},
				"previous_direction": objects.Left,
				"current_direction":  objects.Down,
			},
		},
	}
)

var (
	expected_1 = map[[2]int]struct {
		Direction objects.Direction
		Image     string
	}{
		{10, 9}:  {Direction: objects.Down, Image: "head"},
		{10, 10}: {Direction: objects.Down, Image: "body"},
		{10, 11}: {Direction: objects.Down, Image: "body"},
		{10, 12}: {Direction: objects.Down, Image: "tail"},
	}

	expected_2 = map[[2]int]struct {
		Direction objects.Direction
		Image     string
	}{
		{10, 9}:  {Direction: objects.Down, Image: "head"},
		{10, 10}: {Direction: objects.Down, Image: "body"},
		{10, 11}: {Direction: objects.Left, Image: "curveBody"},
		{11, 11}: {Direction: objects.Left, Image: "body"},
		{12, 11}: {Direction: objects.Left, Image: "tail"},
	}

	expected_3 = map[[2]int]struct {
		Direction objects.Direction
		Image     string
	}{
		{10, 9}:  {Direction: objects.Down, Image: "head"},
		{10, 10}: {Direction: objects.Down, Image: "body"},
		{10, 11}: {Direction: objects.Left, Image: "curveBody"},
		{11, 11}: {Direction: objects.Left, Image: "body"},
		{12, 11}: {Direction: objects.Up, Image: "curveBody"},
		{12, 10}: {Direction: objects.Up, Image: "body"},
		{12, 9}:  {Direction: objects.Up, Image: "tail"},
	}
)

func TestGetSnakeTurningPoints(t *testing.T) {
	/*
	 Test the GetSnakeParts function agains the following scenarios:
	 1. when the snake has no curves
	 2. when the snake has one curve
	 3. when the snake has two curves
	*/

	actual_1 := game.GetSnakeParts(noCurveSnake)
	actual_2 := game.GetSnakeParts(oneCurveSnake)
	actual_3 := game.GetSnakeParts(twoCurveSnake)

	assert.Equal(t, expected_1, actual_1)
	assert.Equal(t, expected_2, actual_2)
	assert.Equal(t, expected_3, actual_3)
}
