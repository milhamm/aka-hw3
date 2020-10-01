package main

import "fmt"


func partition(arr []int, low, high int) int{
	pivot := arr[high]

	i := low - 1

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			temp := arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	temp2 := arr[i + 1]
	arr[i + 1] = arr[high]
	arr[high] = temp2

	return i + 1
}

func quickSort(arr []int, low, high int){
	if(low < high){
		pi := partition(arr, low, high)

		quickSort(arr, low, pi - 1)
		quickSort(arr, pi + 1, high)
	}
}

func main(){
	var arr = []int{4, 2, 1, 6, 7, 5, 9, 0, 3, 8}

	quickSort(arr, 0, len(arr) - 1)

	for _, val := range arr{
		fmt.Printf("%d ", val)
	}
}