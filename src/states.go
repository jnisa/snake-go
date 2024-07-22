// Script that contain all the operations related with the multiple states of the game

package src

// func isGameOver(snake [][]int, board [][]int) bool {
// 	/*
// 	 Check if the game is over.

// 	 There's two options that might be considered as game over: 1) collision
// 	 of the snake with one of the walls and 2) collision of the snake with
// 	 itself.

// 	 :param snake: list of the snake's body
// 	 :param board: matrix representing the board
// 	 :return: boolean indicating if the game is over
// 	*/

// 	checkCollision := func(snake, board) {
// 		/*
// 		 Check if the snake collides with one of the walls.

// 		 :param snake: list of the snake's body
// 		 :param board: matrix representing the board
// 		 :return: boolean indicating if the snake collides with one of the walls
// 		*/

// 		// understand if the following code is right
// 		if snake[0][0] < 0 || snake[0][0] >= len(board) || snake[0][1] < 0 || snake[0][1] >= len(board[0]) {
// 			return true
// 		}

// 	}

// 	checkBiteItself := func(snake) {
// 		/*
// 		 Check if the snake collides with itself.

// 		 :param snake: list of the snake's body
// 		 :return: boolean indicating if the snake collides with itself
// 		*/

// 		// understand if the following code is right
// 		for i := 1; i < len(snake); i++ {
// 			if snake[0] == snake[i] {
// 				return true
// 			}
// 		}
// 	}

// 	return checkCollision(snake, board) || checkBiteItself(snake)
// }

