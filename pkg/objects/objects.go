// Script containing the definition of all the game objects

package objects

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Board struct {
	Cells [32][32]int
	Food  []int
}

type Snake struct {
	Direction Direction
	Body      [][]int
	Score     int
}
