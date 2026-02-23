package main

import (
	"os"

	"github.com/basilysf1709/golos/cli"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
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
