package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Scanln(&cmd)
		if cmd == "help" {
			fmt.Println("This is help cmd!")
		} else if cmd == "quit" {
			break
		} else {
			fmt.Println("Wrong cmd!")
		}
	}
}

