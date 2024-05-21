package main

import (
	"fmt"
	"os"

	gosubst "github.com/ToshihitoKon/go-subst"
)

func main() {
	c := &gosubst.CLI{}
	if err := c.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "command failed: %s", err)
		os.Exit(1)
	}
}
