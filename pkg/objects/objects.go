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
	Cells [24][32]int
}

type Snake struct {
	Direction Direction
	Body      [][]int
	Score     int
}
