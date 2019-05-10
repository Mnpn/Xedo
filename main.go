package main

import (
	"fmt"
	"os"

	"github.com/jD91mZM2/stdutil"
)

const name = "Xedo"
const version = "0.1.0"

func main() {
	cmd := os.Args

	if len(cmd) == 1 {
		printhelp()
		os.Exit(0)
	}

	args := cmd[2:]
	nargs := len(args)
	
	switch cmd[1] {
	case "add":
		if nargs < 1 {
			argerr("add \"<title>\" \"[description]\"")
			return
		}
		fmt.Println("Something worked!")
	case "test":
		fmt.Println("Something worked!")
	default:
		fmt.Println("Unknown argument.\n")
		printhelp()
	}

	os.Exit(0)
}

func printhelp() {
	fmt.Println(name+" ("+version+"), the todo list manager by Martin Persson <mnpn03@gmail.com>")
	fmt.Println("USAGE:")
}

func argerr(cmderr string) {
	stdutil.PrintErr("Invalid arguments. Usage: `xedo "+cmderr+"`", nil)
}