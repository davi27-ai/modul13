package main

import "fmt"

func main() {
	var number int
	var continueLoop bool
	for continueLoop = true; continueLoop; {
		fmt.Scan(&number)
		continueLoop = number <= 0
	}
	fmt.Println("%d adalah bilangan bualat positif\n", number)
}
