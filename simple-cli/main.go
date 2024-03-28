package main

import (
	"flag"
	"fmt"
	"os"
)

// git
func main() {
	name := flag.String("name", "", "The name to greet.")
	flag.Parse()

	if *name == "" {
		fmt.Printf("Common actions for this command:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Hello, %s!\n", *name)
}
