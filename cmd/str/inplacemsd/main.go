package main

//   Sort an array of strings using in-place MSD radix sort.

//    % go run main.go < shells.txt
//    are
//    by
//    sea
//    seashells
//    seashells
//    sells
//    sells
//    she
//    she
//    shells
//    shore
//    surely
//    the
//    the

import (
	"fmt"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/handane123/algorithms/str"
)

func main() {
	a := stdin.ReadAllStrings()
	n := len(a)

	str.InplaceMsdSort(a)
	for i := 0; i < n; i++ {
		fmt.Println(a[i])
	}
}
