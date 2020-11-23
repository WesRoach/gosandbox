package main

import "fmt"

func main() {
	name := "johndoe"
	password := "verysecure"
	printString := fmt.Sprintf("'%s', '%s'\n", name, password)
	fmt.Println(printString)
}
