// Unitary tests to the inputs package

package tests

import (
	"testing"

	"github.com/jnisa/snake-go/pkg/inputs"
	"github.com/nsf/termbox-go"
	"github.com/stretchr/testify/assert"
)

type MockInputReader struct {
	Events []termbox.Event
	index  int
}

func (m *MockInputReader) PollEvent() termbox.Event {
	if m.index < len(m.Events) {
		event := m.Events[m.index]
		m.index++
		return event
	}
	return termbox.Event{}
}

func TestDefineMoveOptions(t *testing.T) {
	/*
	 Test the DefineMoveOptions function.
	*/

	actual := inputs.DefineMoveOptions()
	expected := inputs.MoveOptions{
		Up:    65517,
		Down:  65516,
		Left:  65515,
		Right: 65514,
		Esc:   27,
	}

	assert.Equal(t, actual, expected)
}

func TestGetSnakeMove(t *testing.T) {
	/*
	 Test if the GetSnakeMove is able to capture all the possible moves.
	*/

	test_set := []struct {
		name       string
		mockEvents []termbox.Event
		expected   string
	}{
		{"test - move up", []termbox.Event{{Key: termbox.KeyArrowUp}}, "up"},
		{"test - move down", []termbox.Event{{Key: termbox.KeyArrowDown}}, "down"},
		{"test - move left", []termbox.Event{{Key: termbox.KeyArrowLeft}}, "left"},
		{"test - move right", []termbox.Event{{Key: termbox.KeyArrowRight}}, "right"},
		{"test - escape", []termbox.Event{{Key: termbox.KeyEsc}}, "quit"},
	}

	actual := []string{}
	expected := []string{"up", "down", "left", "right", "quit"}

	for _, test := range test_set {
		t.Run(test.name, func(t *testing.T) {
			mockInputReader := &MockInputReader{Events: test.mockEvents}
			actual = append(actual, inputs.GetSnakeMove(mockInputReader))
		})
	}

	assert.Equal(t, actual, expected)

}
