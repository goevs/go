package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	// Place your code here.
	var phrase string = "Hello, OTUS!"

	fmt.Println(stringutil.Reverse(phrase))
}
