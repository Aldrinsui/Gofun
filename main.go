package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(square(n))

}

func square(n int) int {

	return n * n
}
