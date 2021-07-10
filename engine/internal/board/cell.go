package board

/*
 * Minesweeper Engine
 *
 * @author Darren Rambaud
 */

//Cell represents the state of a single cell or field within the mine field. Val has the following 3 states:
//
// 1. -1 indicates a mine;
//
// 2.  0 indicates an empty or open cell; and
//
// 3.  Val > 0 indicates the number of mines adjacent to this Cell
type Cell struct {
	Row int
	Col int
	Val int
}

func (c Cell) NewCell(row int, column int, val int) Cell {
	c.Row = row
	c.Col = column
	c.Val = val
	return c
}
