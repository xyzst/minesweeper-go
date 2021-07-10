package board

import (
	"math/rand"
	"time"
)

/*
 * Minesweeper Engine
 *
 * https://en.wikipedia.org/wiki/Minesweeper_%28video_game%29 [wikipedia]
 *
 * @author Darren Rambaud
 */

//Generate returns a 2-dimensional (rows x columns) matrix representing a minesweeper board. Mines (as indicated)
// by a -1 sentinel value, will be randomly plotted on the 2D matrix and it's immediate neighbors are incremented
// by 1.
// If given the following parameters, rows = 3, columns = 3, mines = 3, the board will be represented as:
//      [0]  [1]  [2]
//
// [0]  -1    2    1
//
// [1]   2   -1    2
//
// [2]   1    2   -1
//
//
// BUG(xyzst): currently using a weak or predictable seed for rng
func Generate(rows int, columns int, mines int) ([][]Cell, error) {
	if rows < 0 {
		return nil, &RowError{Err: "rows cannot be <= 0"}
	}
	if columns < 0 {
		return nil, &ColumnError{Err: "columns cannot be <= 0"}
	}
	if mines <= 0 {
		return nil, &LogicalError{Err: "mines cannot be <= 0"}
	}
	if rows*columns < mines {
		return nil, &LogicalError{Err: "mines > (rows * columns)"}
	}

	board := make([][]Cell, rows)

	for row := 0; row < rows; row++ {
		board[row] = make([]Cell, columns)
		for col := 0; col < columns; col++ {
			board[row][col] = board[row][col].NewCell(row, col, 0)
		}
	}

	placed := 0

	for placed < mines {
		rand.Seed(time.Now().UnixNano()) // bad seed, it's predictable
		row := rand.Intn(rows)
		rand.Seed(time.Now().UnixNano()) // bad seed, it's predictable
		col := rand.Intn(columns)

		current := board[row][col]
		if current.Val != -1 {
			board[row][col].Val = -1
			if row-1 >= 0 && col-1 >= 0 && board[row-1][col-1].Val != -1 { // top left
				board[row-1][col-1].Val++
			}
			if row-1 >= 0 && board[row-1][col].Val != -1 { // top
				board[row-1][col].Val++
			}
			if row-1 >= 0 && col+1 < columns && board[row-1][col+1].Val != -1 { // top right
				board[row-1][col+1].Val++
			}

			if col-1 >= 0 && board[row][col-1].Val != -1 { // middle left
				board[row][col-1].Val++
			}
			if col+1 < columns && board[row][col+1].Val != -1 { // middle right
				board[row][col+1].Val++
			}

			if row+1 < rows && col-1 >= 0 && board[row+1][col-1].Val != -1 { // bottom left
				board[row+1][col-1].Val++
			}
			if row+1 < rows && board[row+1][col].Val != -1 { // bottom
				board[row+1][col].Val++
			}
			if row+1 < rows && col+1 < columns && board[row+1][col+1].Val != -1 { // bottom right
				board[row+1][col+1].Val++
			}
			placed += 1
		}
	}

	return board, nil
}
