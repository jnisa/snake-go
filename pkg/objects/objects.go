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
	Cells [22][22]int
	Food  []int
}

// TODO. the snake must have a list of maps to collect all
// the turning points to each the images will be applied to
type Snake struct {
	Direction     Direction
	Body          [][2]int
	Score         int
	TurningPoints []map[string]interface{}
}
