package main

import (
	"fmt"
	"os"
	"stef/sort"
	"strconv"
)

func main() {
	//test interface
	in := make([]int, 0)
	for _, arg := range os.Args[1:] {
		i, _ := strconv.Atoi(arg)
		in = append(in, i)
	}
	sort.QuickSort(in)
	fmt.Println(in)

}
