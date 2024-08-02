### **Snake**

**TODO** 
- Use [ebitengine](https://ebitengine.org/) to play the game;
- Create a Makefile to automatically install the dependencies;

**Changes on the current version of the code**
- All the functions that envolve manipulating the board should be modified to leverage pointers;
- The inputs from the keyboard should be detected by using the `ebitengine` package;
- Check if there's functions that can be excluded cause they're not being used;
- Check if there's function that can be changed from the following:

```
func GetPositionValue(b Board, x_coord int, y_coord int) int {
    return Board.Cells[x_coord][y_coord]
}
```

to

```
func (b Board) GetPositionValue (x_coord int, y_coord int) int{
    return Board.Cells[x_coord][y_coord]
}
```

they both do the same thing, but we should get use to the Golang notation;
- [DONE] Reconsider the structure of the project - most of the repos I've checked online would have the 
`auxiliars.go`, `board.go`, `inputs.go`, etc scripts inside one directory only;
`Note`: according to this article the structure that we have atm is the one that follows the best practises.

**Documentation with Go best-practises**
- [Golang official documentation](https://go.dev/doc/effective_go);
- 


**Learning resources**

1. [Pointers in Golang: when, why, and how to use them?](https://www.youtube.com/watch?v=3WsEDZRif6U)
