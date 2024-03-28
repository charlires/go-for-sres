package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "", "The name to greet.")
	flag.Parse()

	if *name == "" {
		fmt.Printf("Usage: simple-cli -name carlos\n")
		fmt.Printf("Common actions for this command:\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	fmt.Printf("Hello, %s!\n", *name)
}
