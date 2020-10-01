package main 

import "fmt"

func towerOfHanoi(n int, from string, to string, temp string){
	if n == 1{
		fmt.Println("Move ", n, " From ", from, " to ", to)
		return
	}
	towerOfHanoi(n - 1, from, temp , to)
	fmt.Println("Move ", n, " From ", from, " to ", to)
	towerOfHanoi(n - 1, temp, to , from)
}

func main(){
	towerOfHanoi(4, "A", "C", "B")
}