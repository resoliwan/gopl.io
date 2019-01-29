package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for k, v := range os.Args[0:] {
		fmt.Println(strconv.Itoa(k) + " " + v)
	}
}

// go run /Users/younlee/work/gopl.io/me/ch1/echo4.go 10 20 30
