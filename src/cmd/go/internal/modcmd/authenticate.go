// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modcmd

import (

	// "errors" // do I need this?
	"cmd/go/internal/base"
	"context"
	"fmt"
	// "golang.org/x/mod/module"
	// "golang.org/x/mod/sumdb/dirhash"
)

var cmdAuthenticate = &base.Command{
	UsageLine: "go mod authenticate [args]",
	Short:     "check hashes from checksumdb matches with the local git libraries",
	Long:      `Authenticate checks whether checksumdb provides/serves the original hash for each git tag of a library`,
	Run:       runAuthenticate,
}

// do I need this?
// func init() {
// 	base.AddChdirFlag(&cmdVerify.Flag)
// 	base.AddModCommonFlags(&cmdVerify.Flag)
// }

func runAuthenticate(ctx context.Context, cmd *base.Command, args []string) {
	// localDir := args[0]
	fmt.Printf("local directory to be checked: %s \n", args)
}
