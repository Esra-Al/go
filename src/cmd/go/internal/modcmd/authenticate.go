// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modcmd

import (
	"cmd/go/internal/base"
	"cmd/go/internal/modfetch"
	"context"
	"fmt"

	"golang.org/x/mod/module"
)

var cmdAuthenticate = &base.Command{
	UsageLine: "go mod authenticate [args]",
	Short:     "check hashes from checksumdb matches with the local git libraries",
	Long: `Authenticate checks whether checksumdb provides/serves the original hash for each git tag of a library
	use it as 
	GOPROXY=direct dev-go mod authenticate golang.org/x/text v0.16.0 `,
	Run: runAuthenticate,
}

// do I need this?
// func init() {
// 	base.AddChdirFlag(&cmdVerify.Flag)
// 	base.AddModCommonFlags(&cmdVerify.Flag)
// }

func runAuthenticate(ctx context.Context, cmd *base.Command, args []string) {

	if len(args) < 2 {
		fmt.Println("Usage: go mod authenticate <string1> <string2>")
		return
	}
	path := args[0]
	version := args[1]
	fmt.Printf("local directory: %s, and its prefix: %s\n", path, version)
	mod := module.Version{Path: path, Version: version}
	_, err := modfetch.DownloadZip(ctx, mod)
	fmt.Println("err", err)
}
