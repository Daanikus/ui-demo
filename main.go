package main

import "fmt"

func main() {
	fmt.Println("UI Demo!")
	Exported()
	this_is_a_bad_name()
}

func Exported() {
	fmt.Println("This func should have a comment")
}

func this_is_a_bad_name() {
	fmt.Println("Truly awful")
}
