package main

import (
	"os"

	"github.com/iqbalyusuf/golos/cmd"
)

func main() {
	// golos stop
	if len(os.Args) > 1 && os.Args[1] == "stop" {
		cmd.Stop()
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
		cmd.Run(false)
		return
	}

	cmd.Run(detach)
}
