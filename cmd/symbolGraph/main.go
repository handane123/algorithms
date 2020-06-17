package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/handane123/algorithms/graph"
	"github.com/handane123/algorithms/stdin"
)

func main() {
	filename := os.Args[1]
	delimiter := os.Args[2]
	sg := graph.NewSymbolGraph(filename, delimiter)
	g := sg.Graph()
	stdin := stdin.NewStdInLine()
	for !stdin.IsEmpty() {
		source := strings.Trim(stdin.ReadString(), " ")
		if ok, err := sg.Contains(source); err != nil {
			fmt.Println(err)
		} else if ok {
			if s, err := sg.IndexOf(source); err != nil {
				fmt.Println(err)
			} else {
				for _, v := range g.Adj(s) {
					fmt.Println(" ", sg.NameOf(v))
				}
				fmt.Println()
			}
		} else {
			fmt.Printf("input not contain '%s'\n", source)
		}
	}
}