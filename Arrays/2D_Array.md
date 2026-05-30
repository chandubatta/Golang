#2D Array in Golang

1. A 2D array (two-dimensional array) is an array whose elements are themselves arrays. It is used to store data in a rows-and-columns format, similar to a table, spreadsheet, chessboard, or matrix.

*Purpose*
Organize related data in tabular form.
Represent matrices and grids.
Store structured data where both row and column positions matter.
*When It's Commonly Used*
Game boards (Tic-Tac-Toe, Chess, Sudoku)
Matrix calculations
Seating arrangements
Image pixel representation
Tabular data processing

**2.Simple Code example**
package main

import "fmt"

func main() {
	// 3 rows and 4 columns
	var marks [3][4]int

	marks[0][0] = 85
	marks[0][1] = 90
	marks[0][2] = 78
	marks[0][3] = 88

	marks[1][0] = 92
	marks[1][1] = 81
	marks[1][2] = 75
	marks[1][3] = 89

	marks[2][0] = 80
	marks[2][1] = 95
	marks[2][2] = 85
	marks[2][3] = 91

	fmt.Println(marks)
}

Output:=[[85 90 78 88] [92 81 75 89] [80 95 85 91]]

*Iterating Through a 2D Array*
package main

import "fmt"

func main() {
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}

output: 1 2 3
        4 5 6

**3.Common Mistakes and Misconceptions**


