package main

import (
	"fmt"

	"github.com/pietrosparks/pluto/printer"
	"github.com/pietrosparks/pluto/tree"
)

func main() {

	t := &tree.BinaryTree{}
	datastorePrinter := printer.New(t)

	datastorePrinter.Build()
	datastorePrinter.DFS()
	datastorePrinter.BFS()

	fmt.Printf("\n Done Running \n")
}
