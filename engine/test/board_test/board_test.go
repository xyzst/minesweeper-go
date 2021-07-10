package board_test

import (
	b "github.com/xyzst/minesweeper-go/engine/internal/board"
	"testing"
)

/*
 * Minesweeper Engine
 *
 * Unit tests for board package
 *
 * @author Darren Rambaud
 */

func TestBoardAdjacentCells(t *testing.T) {
	cases := []struct {
		rows    int
		columns int
		mines   int
	}{
		{3, 3, 3},
		{10, 10, 10},
		{100, 100, 1000},
		{3, 10, 10},
	}

	for cs, c := range cases {
		board, err := b.Generate(c.rows, c.columns, c.mines)
		if err != nil {
			t.Errorf("🔴: unexpected error during test (rows: %d, columns: %d, mines: %d)", c.rows, c.columns, c.mines)
		}
		passed := true
		for row := 0; row < c.rows; row++ {
			for col := 0; col < c.columns; col++ {
				if board[row][col].Val != -1 && board[row][col].Val > 0 {
					actual := board[row][col].Val
					if row-1 >= 0 && col-1 >= 0 && board[row-1][col-1].Val == -1 { // top left
						actual--
					}
					if row-1 >= 0 && board[row-1][col].Val == -1 { // top
						actual--
					}
					if row-1 >= 0 && col+1 < c.columns && board[row-1][col+1].Val == -1 { // top right
						actual--
					}

					if col-1 >= 0 && board[row][col-1].Val == -1 { // middle left
						actual--
					}
					if col+1 < c.columns && board[row][col+1].Val == -1 { // middle right
						actual--
					}

					if row+1 < c.rows && col-1 >= 0 && board[row+1][col-1].Val == -1 { // bottom left
						actual--
					}
					if row+1 < c.rows && board[row+1][col].Val == -1 { // bottom
						actual--
					}
					if row+1 < c.rows && col+1 < c.columns && board[row+1][col+1].Val == -1 { // bottom right
						actual--
					}
					if actual != 0 {
						passed = false
						t.Errorf("🔴: number of adjacent mines mismatch at (%d,%d) for case %d", row, col, cs)
					}
				}
			}
		}
		if passed {
			t.Logf("🟢: case %d (rows: %d, columns: %d, mines: %d) passed", cs, c.rows, c.columns, c.mines)
		} else {
			t.Errorf("🔴: one or more adjacent mine mismatches in case %d", cs)
		}
	}
}

func TestMineInsertion(t *testing.T) {
	cases := []struct {
		rows    int
		columns int
		mines   int
		want    int
	}{
		{10, 10, 10, 10},
		{10, 5, 25, 25},
		{1, 1, 1, 1},
		{100, 100, 5000, 5000},
	}
	for _, c := range cases {
		board, err := b.Generate(c.rows, c.columns, c.mines)
		if err != nil {
			t.Errorf("🔴: unexpected error during test (rows: %d, columns: %d, mines: %d)", c.rows, c.columns, c.mines)
		}
		count := 0
		for row := 0; row < c.rows; row++ {
			for col := 0; col < c.columns; col++ {
				if board[row][col].Val == -1 {
					count++
				}
			}
		}
		if count != c.want {
			t.Errorf("🔴: Found %d mines in matrix, but wanted %d", count, c.want)
		} else {
			t.Logf("🟢: Expected: %d, Wanted: %d", count, c.want)
		}
	}
}

func TestMineInsertionNegative(t *testing.T) {
	cases := []struct {
		rows    int
		columns int
		mines   int
		want    error
	}{
		{-1, 1, 1, b.RowError{Err: "rows is < 0"}},
		{1, -1, 1, b.ColumnError{Err: "columns < 0"}},
		{1, 1, -1, b.LogicalError{Err: "mines < 0"}},
		{1, 1, 0, b.LogicalError{Err: "mines == 0"}},
		{1, 1, 2, b.LogicalError{Err: "mines > rows * columns"}},
	}
	for _, c := range cases {
		_, err := b.Generate(c.rows, c.columns, c.mines)
		if err == nil {
			t.Errorf("🔴: expected %v, but did not get any", c.want)
		} else {
			t.Logf("🟢: received the expected error (%v)", err)
		}
	}
}
