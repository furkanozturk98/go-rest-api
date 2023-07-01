package main

import "fmt"

func Run() error {
	return nil
}

func main() {
	fmt.Println("test")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
