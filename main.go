package main

import (
	"fmt"
	"os"
	"strings"

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
	case "remove":
		if nargs != 1 {
			argerr("remove <id>")
			return
		}
		fmt.Println("Something worked!")
	default:
		fmt.Println("Unknown argument.\n")
		printhelp()
	}

	os.Exit(0)
}

func printhelp() {
	fmt.Println(name+" ("+version+"), the todo list manager by Martin Persson <mnpn03@gmail.com>")
	help := make([]string, 0)
	help = append(help, "USAGE:")
	help = append(help, "\tadd \"<title>\" \"[description]\"\tAdd a new todo list entry with an optional description.")
	help = append(help, "\tremove <id>\t\t\tRemove an entry from the todo list.")
	help = append(help, "\tclear\t\t\t\tDeletes the whole list permanently.")
	help = append(help, "\tmove <id> <up/down/top/bottom>\tMove an entry to a new place.")
	help = append(help, "\trename <id> <title>\t\tRename an entry's title.")
	help = append(help, "\trenamedesc <id> <description>\tRename an entry's description.")

	fmt.Println(strings.Join(help, "\n"))
}

func argerr(cmderr string) {
	stdutil.PrintErr("Invalid arguments. Usage: `"+name+" "+cmderr+"`.", nil)
}