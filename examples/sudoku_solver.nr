/*########### Backtracking Algorithm ##############
            
                By @VictorKariuki

        https://github.com/VictorKariuki

NURU program to solve Sudoku using Backtracking Algorithm

The sudoku puzzle is represented as a 2D array. The empty 
cells are represented by 0. The algorithm works by trying
out all possible numbers for an empty cell. If the number 
is valid, it is placed in the cell. If the number is invalid,
the algorithm backtracks to the previous cell and tries 
another number. The algorithm terminates when all cells 
are filled. The algorithm is implemented in the solveSudoku
function. The isValid function checks kama a number is
valid in a given cell. The printSudoku function prints
the sudoku puzzle. The solveSudoku function solves the
sudoku puzzle. The main function initializes the sudoku
puzzle and calls the solveSudoku function.

#################################################*/


fanya printing = unda(sudoku) {
    fanya row = 0
    wakati (row < 9){
        andika(sudoku[row])
        row++
    }
}

fanya sudoku = [[3, 0, 6, 5, 0, 8, 4, 0, 0],[5, 2, 0, 0, 0, 0, 0, 0, 0],[0, 8, 7, 0, 0, 0, 0, 3, 1],[0, 0, 3, 0, 1, 0, 0, 8, 0],[9, 0, 0, 8, 6, 3, 0, 0, 5],[0, 5, 0, 0, 9, 0, 6, 0, 0],[1, 3, 0, 0, 0, 0, 2, 5, 0],[0, 0, 0, 0, 0, 0, 0, 7, 4],[0, 0, 5, 2, 0, 6, 3, 0, 0]]



fanya isSafe = unda(grid, row, col, num) {
    kwa x ktk [0,1,2,3,4,5,6,7,8] {
        kama (grid[row][x] == num) {
            rudisha sikweli
        }
    }

    kwa x ktk [0,1,2,3,4,5,6,7,8] {
        kama (grid[x][col] == num) {
            rudisha sikweli
        }
    }

    fanya startRow = row - row % 3
    fanya startCol = col - col % 3

    kwa i ktk [0, 1, 2] {
        kwa j ktk [0, 1, 2] {
            kama (grid[i + startRow][j + startCol] == num) {
                rudisha sikweli
            }
        }
    }

    rudisha kweli
}

fanya solveSudoku = unda(grid, row, col) {
    kama (row == 8 && col == 9) {
        rudisha kweli
    }

    kama (col == 9) {
        row += 1
        col = 0
    }

    kama (grid[row][col] > 0) {
        rudisha solveSudoku(grid, row, col + 1)
    }

    kwa num ktk [1,2,3,4,5,6,7,8,9] {
        kama (isSafe(grid, row, col, num)) {
            grid[row][col] = num
            kama (solveSudoku(grid, row, col + 1)) {
                rudisha kweli
            }
        }

        grid[row][col] = 0
    }

    rudisha sikweli
}
andika()
andika("----- PUZZLE TO SOLVE -----")
printing(sudoku)
kama (solveSudoku(sudoku, 0, 0)){
    andika()
    andika("--------- SOLUTION --------")
    printing(sudoku)
    andika()
} sivyo {
    andika("imeshindikana")
}