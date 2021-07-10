package main

import (
	"fmt"
	b "github.com/xyzst/minesweeper-go/engine/internal/board"
)

func main() {
	var i, _ = b.Generate(10, 10, 5)

	for x := 0; x < len(i); x++ {
		for y := 0; y < 10; y++ {
			v := i[x][y].Val
			fmt.Printf(" %d ", v)
		}
		fmt.Print("\n")
	}
}