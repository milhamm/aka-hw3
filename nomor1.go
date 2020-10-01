package main

import "fmt"

func recursiveFunc(n int) int{
	if(n <= 1){
		return n
	}

	return recursiveFunc(n - 1) + (2 * n - 1)
}

func main(){
	fmt.Println(recursiveFunc(3))
}