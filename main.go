package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/basilysf1709/golos/cli"
)

func init() {
	// Pin the main goroutine to macOS thread 0 so that
	// dispatch_get_main_queue() and AppKit UI calls work correctly.
	runtime.LockOSThread()
}

// Set by goreleaser ldflags.
var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version", "--version", "-v":
			fmt.Printf("golos %s (commit %s, built %s)\n", version, commit, date)
			return
		case "stop":
			cli.Stop()
			return
		case "add":
			cli.DictAdd(os.Args[2:])
			return
		case "delete":
			cli.DictDelete(os.Args[2:])
			return
		case "list":
			cli.DictList()
			return
		case "import":
			cli.DictImport(os.Args[2:])
			return
		}
	}

	// Parse our own flags before flag.Parse() sees them
	detach := false
	background := false
	filtered := []string{os.Args[0]}
	for _, a := range os.Args[1:] {
		switch a {
		case "-d", "--detach":
			detach = true
		case "--background":
			background = true
		default:
			filtered = append(filtered, a)
		}
	}
	os.Args = filtered

	if background {
		cli.Run(false)
		return
	}

	cli.Run(detach)
}
