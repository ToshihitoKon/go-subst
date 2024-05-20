package main

import (
	"context"

	gosubst "github.com/ToshihitoKon/go-subst"
)

func main() {
	ctx := context.Background()
	c := &gosubst.CLI{}
	c.Run(ctx)
}
