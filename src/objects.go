// Script containing the definition of all the game objects

// Things to take into consideration
// 1. the dimension of the board will be 32x24 tiles
// 2. the snake will be represented by a list of coordinates
// 3. a Position object should also be defined, suited to a bidimensional board

package src

type Board struct {
	Cells [24][32]int
}
