package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	data := bufio.NewScanner(os.Stdin)
	data.Scan()
	line := data.Text()
	fmt.Println(strings.ToUpper(line))
}
