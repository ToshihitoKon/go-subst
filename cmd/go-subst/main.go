package main

import (
	"fmt"
	"os"

	gosubst "github.com/ToshihitoKon/go-subst"
)

func main() {
	c := &gosubst.CLI{}
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
}
