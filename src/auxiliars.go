// Functions that will be auxiliars to other main functions
// These will be used throughout other modules of the code

package src

func Get(x, y int, board [][]int) int {
	/*
	 Get the value of the board at the given position.

	 :param board: matrix representing the board
	 :param x: x coordinate
	 :param y: y coordinate
	 :return int: value of the board at the given position
	*/

	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return -1
	}

	return board[x][y]
}

// TODO. another function that can be probably added here is the one that will be
// responsible for picking a random position for the apple. Might be needed to import
// a package for that.