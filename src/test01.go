package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hi, Charles")
	var arr [4]int = [4]int{0, 1, 2, 3}
	fmt.Printf("%d", arr[1:4])
}
