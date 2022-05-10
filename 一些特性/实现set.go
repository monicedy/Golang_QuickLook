package main

import (
	"fmt"
)

func main() {

	set0 := map[int]struct{}{}
	fmt.Println(set0)

	set := map[int]struct{}{}
	set[0] = struct{}{}
	set[1] = struct{}{}
	_, ok := set[0]
	fmt.Println(set)
	fmt.Println(ok)
}
