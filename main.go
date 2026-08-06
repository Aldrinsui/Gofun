package main

import (
	"fmt"
)

func main() {
	var name string
	var age int64

	fmt.Scan(&name)
	fmt.Scan(&age)

	fmt.Printf("Hi, %s! You are %d years old.\n", name, age)

}
