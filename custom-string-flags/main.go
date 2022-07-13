package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	lineStr := "push -dir=/Desktop/go-test"
	lineArgs := strings.Split(lineStr, " ")
	for _, a := range lineArgs {
		fmt.Print(a, ", ")
	}
	fmt.Println()
	line := flag.NewFlagSet(lineArgs[0], flag.ExitOnError)
	dir := line.String("dir", ".", "dir to push files in cloud")

	if err := line.Parse(lineArgs[1:]); err != nil {
		panic(err)
	}
	fmt.Println("args: ", line.Args())
	fmt.Println("dir: ", *dir)
}
