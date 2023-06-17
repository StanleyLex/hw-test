package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	hello := "Hello, OTUS!"
	fmt.Println(stringutil.Reverse(hello))
}
