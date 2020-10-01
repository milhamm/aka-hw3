package main 

import "fmt"

const NMAX = 100

func rBinarySearch(arr [NMAX]int, searched int, l int, r int) int {
	if r >= 1 {
		mid := l + (r - l) / 2

		if arr[mid] == searched {
			return mid
		}

		if arr[mid] > searched {
			return rBinarySearch(arr, searched, l, mid-1)
		}

		return rBinarySearch(arr, searched, mid+1, r)
	}

	return -1
}

func main(){
	var arr = [NMAX]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Harus Sorted
	var search int = rBinarySearch(arr, 8, 0 , NMAX - 1)
	fmt.Println()
	fmt.Println(search)
}