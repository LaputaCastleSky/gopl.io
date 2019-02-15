package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
	for i, arg := range os.Args {
		fmt.Println("i:", i, "arg:", arg)
	}
}
