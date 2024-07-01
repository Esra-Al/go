// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// call go list -m

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
	GOPROXY=direct dev-go mod authenticate golang.org/x/text /Users/esra/golang-text/.git/
	GOPROXY=direct dev-go mod authenticate github.com/Esra-Al/cool-go /Users/esra/cool-go/.git/`,
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
	localdir := args[1]
	r := modfetch.Lookup(ctx, "direct", repoPath, localdir)

	v, _ := r.Versions(ctx, "v")
	for _, vrs := range v.List {
		fmt.Println(vrs)
	}

	vrs := v.List[len(v.List)-1]

	fmt.Printf("Latest semantic version tag: %s\n", vrs)

	// Authenticate using the latest semantic version tag
	mod := module.Version{Path: repoPath, Version: vrs, Localdir: localdir}
	_, err := modfetch.DownloadZip(ctx, mod)
	if err != nil {
		fmt.Printf("Error downloading zip: %v\n", err)
		return
	}

	fmt.Println("Successfully authenticated the repository")
}

// 	// Fetch all versions
// 	v, err := r.Versions(ctx, "")
// 	if err != nil {
// 		fmt.Printf("Error fetching versions: %v\n", err)
// 		return
// 	}
// 	for _, version := range v.List {
// 		fmt.Println("version", version)
// 	}
// 	// Create a wait group to synchronize goroutines
// 	var wg sync.WaitGroup
// 	// Create a channel to collect errors
// 	errChan := make(chan error, len(v.List))

// 	// Function to authenticate a version
// 	authenticateVersion := func(version string) {
// 		defer wg.Done()
// 		mod := module.Version{Path: repoPath, Version: version}
// 		_, err := modfetch.DownloadZip(ctx, mod)
// 		if err != nil {
// 			errChan <- fmt.Errorf("error authenticating version %s: %v", version, err)
// 			return
// 		}
// 		fmt.Printf("Successfully authenticated version %s\n", version)
// 	}

// 	// Start a goroutine for each version
// 	for _, version := range v.List {
// 		wg.Add(1)
// 		go authenticateVersion(version)
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()
// 	close(errChan)

// 	// Print any errors that occurred
// 	for err := range errChan {
// 		fmt.Println(err)
// 	}

// 	fmt.Println("Authentication process completed.")
// }

// package modcmd

// import (
// 	"cmd/go/internal/base"
// 	"cmd/go/internal/modfetch"
// 	"cmd/go/internal/modload"
// 	"context"
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"golang.org/x/mod/modfile"
// )

// var cmdAuthenticate = &base.Command{
// 	UsageLine: "go mod authenticate [args]",
// 	Short:     "check hashes from checksumdb matches with the local git libraries",
// 	Long: `Authenticate checks whether checksumdb provides/serves the original hash for each git tag of a library.
// 	Use it as:
// 	GOPROXY=direct dev-go mod authenticate golang.org/x/text`,
// 	Run: runAuthenticate,
// }

// func init() {
// 	base.AddChdirFlag(&cmdAuthenticate.Flag)
// 	base.AddModCommonFlags(&cmdAuthenticate.Flag)
// }

// func runAuthenticate(ctx context.Context, cmd *base.Command, args []string) {
// 	modload.InitWorkfile()

// 	var modDir, repoPath string

// 	if len(args) != 0 {
// 		// Use the provided module directory
// 		repoPath = args[0]
// 		fmt.Printf("Using provided module directory: %s\n", repoPath)
// 	} else {
// 		// Use the default module in the current directory
// 		modload.ForceUseModules = true
// 		modload.RootMode = modload.NeedRoot

// 		// Ensure we are in a module context
// 		if !modload.HasModRoot() {
// 			base.Fatalf("go: not in a module")
// 		}

// 		// Get the directory containing the go.mod file
// 		modFilePath := modload.ModFilePath()
// 		modDir = filepath.Dir(modFilePath)
// 		// Read the module path from the go.mod file
// 		modFilePath2 := filepath.Join(modDir, "go.mod")
// 		data, err := os.ReadFile(modFilePath2)
// 		if err != nil {
// 			fmt.Printf("Error reading go.mod file: %v\n", err)
// 			return
// 		}

// 		modFile, err := modfile.Parse("go.mod", data, nil)
// 		if err != nil {
// 			fmt.Printf("Error parsing go.mod file: %v\n", err)
// 			return
// 		}

// 		repoPath = modFile.Module.Mod.Path
// 	}

// 	fmt.Printf("Using module path: %s\n", repoPath)
// 	r := modfetch.Lookup(ctx, "direct", repoPath)

// 	// Fetch all versions
// 	v, err := r.Versions(ctx, "")
// 	if err != nil {
// 		fmt.Printf("Error fetching versions: %v\n", err)
// 		return
// 	}

// 	for _, version := range v.List {
// 		fmt.Println("version", version)
// 	}
// 	// // Create a wait group to synchronize goroutines
// 	// var wg sync.WaitGroup
// 	// // Create a channel to collect errors
// 	// errChan := make(chan error, len(v.List))

// 	// // Function to authenticate a version
// 	// authenticateVersion := func(version string) {
// 	// 	defer wg.Done()
// 	// 	mod := module.Version{Path: repoPath, Version: version}
// 	// 	_, err := modfetch.DownloadZip(ctx, mod)
// 	// 	if err != nil {
// 	// 		errChan <- fmt.Errorf("error authenticating version %s: %v", version, err)
// 	// 		return
// 	// 	}
// 	// 	fmt.Printf("Successfully authenticated version %s\n", version)
// 	// }

// 	// // Start a goroutine for each version
// 	// for _, version := range v.List {
// 	// 	wg.Add(1)
// 	// 	go authenticateVersion(version)
// 	// }

// 	// // Wait for all goroutines to finish
// 	// wg.Wait()
// 	// close(errChan)

// 	// // Print any errors that occurred
// 	// for err := range errChan {
// 	// 	fmt.Println(err)
// 	// }

// 	// fmt.Println("Authentication process completed.")
// }
