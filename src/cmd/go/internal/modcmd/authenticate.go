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
	GOPROXY=direct dev-go mod authenticate golang.org/x/text`,
	Run: runAuthenticate,
}

// do I need this?
// func init() {
// 	base.AddChdirFlag(&cmdVerify.Flag)
// 	base.AddModCommonFlags(&cmdVerify.Flag)
// }

func runAuthenticate(ctx context.Context, cmd *base.Command, args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: go mod authenticate <repo_path>")
		return
	}

	repoPath := args[0]
	fmt.Printf("Local directory: %s\n", repoPath)

	r := modfetch.Lookup(ctx, "direct", repoPath)
	fmt.Println(r.ModulePath())

	v, _ := r.Versions(ctx, "v")
	for _, vrs := range v.List {
		fmt.Println(vrs)
	}

	latestTag, _ := r.Latest(ctx)

	fmt.Printf("Latest semantic version tag: %s\n", latestTag.Version)

	// Authenticate using the latest semantic version tag
	mod := module.Version{Path: repoPath, Version: latestTag.Version}
	_, err := modfetch.DownloadZip(ctx, mod)
	if err != nil {
		fmt.Printf("Error downloading zip: %v\n", err)
		return
	}

	fmt.Println("Successfully authenticated the repository")
}
