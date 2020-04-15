package main

import "fmt"

var i int = 10

func main() {
	var (
		i int = 20
		test1 string = "test"
		t2 string = "t2"
	)
	fmt.Println(i)
	fmt.Printf("%v, %T", test1, test1)
	fmt.Printf("%v, %T", t2, t2)
}