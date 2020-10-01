package main

import "fmt"

const N int = 6;

func isSafe(board [][]int, row, col int) bool{
	var i, j int;

	i = 0
	for i < col {
		if board[row][i] == 1{
			return false
		}
		i++
	}

	i = row
	j = col
	for i >= 0 && j >= 0{
		if board[i][j] == 1{
			return false
		}
		i--
		j--
	}


	i = row 
	j = col
	for j >= 0 && i < N{
		if board[i][j] == 1{
			return false;
		}
		i++
		j--
	}

	return true;
}

func nQueens(board [][]int, col int) bool{
	if col >= N{
		return true
	}

	for i:=0; i < N; i++{
		if isSafe(board, i, col) {
			board[i][col] = 1

			if nQueens(board, col + 1) {
				return true
			}

			board[i][col] = 0
		}
	}

	return false
}

func fillBoard()[][]int{
	var arr = make([][]int, N)

	for i:=0; i < N; i++{
		arr[i] = make([]int, N)
		for j:=0; j < N; j++{
			arr[i][j] = 0;
		}
	}

	return arr 
}

func main(){
	
	var arr [][]int = fillBoard()

	if !nQueens(arr, 0){
		fmt.Println("Solution not exist")
	} else {
		for _, r := range arr {
			for _, val := range r{
				fmt.Printf("%d ", val)
			}

			fmt.Printf("\n")
		}
	}
}