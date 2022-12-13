// Copyright 2016 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.
package main

import (
	"os"

	"github.com/juju/cmd/v3"

	"github.com/kian99/candid/cmd/candid/internal/admincmd"
)

func main() {
	ctxt := &cmd.Context{
		Dir:    ".",
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
	os.Exit(cmd.Main(admincmd.New(), ctxt, os.Args[1:]))
}
