package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0:])
}

// go run /Users/younlee/work/gopl.io/me/ch1/echo3.go 1 2 3
