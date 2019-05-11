package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jD91mZM2/gtable"
	"github.com/jD91mZM2/stdutil"
)

const name = "Xedo"
const version = "0.1.0"

func main() {
	cmd := os.Args

	// If no argument was provided, check if
	// a list exists and display it.
	if len(cmd) == 1 {
		fmt.Println("Your Xedo list:")
		testarr := []string{"think", "thonk", "thank"}
		table := gtable.NewStringTable()
		table.AddStrings("ID", "Title", "Description")

		i := 1
		for _, thing := range testarr {
			table.AddRow()
			table.AddStrings(strconv.Itoa(i), thing+" name", thing+" desc")
			i+=1
		}

		fmt.Println(table.String())
		return
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
	case "rename":
		if nargs < 2 {
			argerr("rename <id> <new title>")
			return
		}
		id := args[0]
		newtitle := strings.Join(args[1:], " ")
		fmt.Println(id)
		fmt.Println(newtitle)
	case "renamedesc":
		if nargs < 2 {
			argerr("renamedesc <id> <new description>")
			return
		}
		id := args[0]
		newdesc := strings.Join(args[1:], " ")
		fmt.Println(id)
		fmt.Println(newdesc)
	case "help":
		printhelp()
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
	help = append(help, "\tadd \"<title>\" \"[description]\"\tAdd a new entry with an optional description.")
	help = append(help, "\tremove <id>\t\t\tRemove an entry from the list.")
	help = append(help, "\tclear\t\t\t\tDeletes the whole list permanently.")
	help = append(help, "\tmove <id> <up/down/top/bottom>\tMove an entry to a new place.")
	help = append(help, "\trename <id> <title>\t\tRename an entry's title.")
	help = append(help, "\trenamedesc <id> <description>\tRename an entry's description.")

	fmt.Println(strings.Join(help, "\n"))
}

func argerr(cmderr string) {
	stdutil.PrintErr("Invalid arguments. Usage: `"+strings.ToLower(name)+" "+cmderr+"`.", nil)
}