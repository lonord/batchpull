package main

import (
	"fmt"
	"os"
)

func error(s string) {
	fmt.Fprintln(os.Stderr, s)
}

func errorExit(s string) {
	error(s)
	os.Exit(1)
}
