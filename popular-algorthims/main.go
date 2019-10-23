package main

import "fmt"

var a string

func main() {
	a = "ssss"

	fmt.Println(a)
	printA()
}

func printA(){
	fmt.Println(a)
}


