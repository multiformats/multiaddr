package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintf(os.Stderr, "run 'go test github.com/multiformats/multiaddr/test' instead\n")
	os.Exit(1)
}
