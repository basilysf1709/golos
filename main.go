package main

import (
	"os"

	"github.com/basilysf1709/golos/cli"
)

func main() {
	// golos stop
	if len(os.Args) > 1 && os.Args[1] == "stop" {
		cli.Stop()
		return
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
