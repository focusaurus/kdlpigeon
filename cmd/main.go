package main

import (
	"fmt"
	"log"
	"os"

	kdl "github.com/focusaurus/kdlpigeon"
)

func main() {
	nodes, err := kdl.ParseFile(os.Args[1])
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Printf("%s\n", nodes)
	fmt.Printf("%+v\n", nodes)
}
