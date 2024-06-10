package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	commandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
)

func run(args []string) int {
	max := commandLine.Int("max", 255, "max value")
	name := commandLine.String("name", "", "my name")
	if err := commandLine.Parse(args); err != nil {
		fmt.Fprintf(os.Stderr, "cannot parse flags: %v\n", err)
	}

	if *max > 999 {
		fmt.Fprintf(os.Stderr, "invalid max value: %v\n", *max)
		return 1
	}

	if *name == "" {
		fmt.Fprintln(os.Stderr, "name must be provided")
		return 1
	}

	fmt.Println(*max, *name)

	return 0
}

func main() {
	os.Exit(run(os.Args[1:]))
}
