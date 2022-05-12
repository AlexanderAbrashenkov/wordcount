package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	src := readInput()
	var res int
	if src == "" {
		res = 0
	} else {
		words := strings.Split(src, " ")
		res = len(words)
	}

	fmt.Println(res)
}

func readInput() (src string) {
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}
