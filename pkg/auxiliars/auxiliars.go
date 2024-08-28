// Functions that will be auxiliars to other main functions
// These will be used throughout other modules of the code

package auxiliars

import (
	"math/rand"
	"reflect"

	"github.com/jnisa/snake-go/pkg/objects"
)

func Get(x, y int, board objects.Board) int {
	/*
	 Get the value of the board at the given position.

	 :param board: matrix representing the board
	 :param x: x coordinate
	 :param y: y coordinate
	 :return int: value of the board at the given position
	*/

	if x < 0 || x >= len(board.Cells) || y < 0 || y >= len(board.Cells[0]) {
		return -1
	}

	return board.Cells[x][y]
}

func GetRandomPosition(snake objects.Snake, board objects.Board) (int, int) {
	/*
	 Get a random position in the board.

	 Whenever a new positions is generated, one of the things that must be
	 held into account is that the new position must not be in the snake's
	 body. That's is why the snake's body is passed as an argument.

	 :param snake: list of the snake's body
	 :param board: matrix representing the board
	 :return (int, int): random position in the board
	*/

	x := rand.Intn(len(board.Cells))
	y := rand.Intn(len(board.Cells[0]))

	for _, pos := range snake.Body {
		if x == pos[0] && y == pos[1] {
			return GetRandomPosition(snake, board)
		}
	}

	return x, y
}

func RemoveDuplicates(elements [][2]int) [][2]int {
	/*
	 Remove duplicates from a list of elements.

	 The way this function works is essentially by iterating over the list of
	 elements and checking if the element is already in the list of non duplicates.
	 If it is not, then it is added to the list of non duplicates.

	 :param elements: list of elements
	 :return [][]int: list of elements without duplicates
	*/

	contains := func(list [][2]int, element [2]int) bool {
		/*
		 Check if the list of elements contains the given element.

		 :param list: list of values
		 :param element: element whose existence is to be checked
		 :return bool: boolean indicating if the element is in the list
		*/

		for _, e := range list {
			if e[0] == element[0] && e[1] == element[1] {
				return true
			}
		}

		return false
	}

	non_duplicates := [][2]int{}

	for _, element := range elements {
		if !contains(non_duplicates, element) {
			non_duplicates = append(non_duplicates, element)
		}
	}

	return non_duplicates
}

func IsIn(targetValue [2]int, searchingList [][2]int) bool {
	/*
	 Function that evaluates if a given list is present in a given nested list.

	 :param targetValue: value to be searched
	 :param searchingList: list that will be targeted by the search
	 :return: a boolean indicating whether the target values is present in the nested list
	*/

	for _, val := range searchingList {
		if reflect.DeepEqual(targetValue, val) {
			return true
		}
	}

	return false
}
