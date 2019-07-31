package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"os"

	"github.com/fatih/color"
	"github.com/jD91mZM2/gtable"
	"github.com/jD91mZM2/stdutil"
	"github.com/OpenPeeDeeP/xdg"
)

type ListItem struct {
    Title string
    Description string
}

var output []ListItem;

const name = "Xedo"
const version = "0.2.0"
const author = "Mnpn"

var xdgDir = xdg.New(author, name)
var dataFile = xdgDir.DataHome()+"/list.json"

func main() {
	cmd := os.Args

	// If the folder doesn't exist
	if _, err := os.Stat(dataFile); os.IsNotExist(err) {
		// Make the folder
		derr := os.MkdirAll(xdgDir.DataHome(), 0755) // Needs rwx (7) or else it errors
		if derr != nil {
			stdutil.PrintErr("Directory creation failed", derr)
			return
		}

		// Make the file
		cfile, ferr := os.Create(dataFile)
		if ferr != nil {
			stdutil.PrintErr("File creation failed", ferr)
			return
		}

		// Add basic stuff to the file
		_, werr := cfile.Write([]byte("[]"))
		if werr != nil {
			stdutil.PrintErr("Failed to write to file", werr)
		}
	}

	// Read the contents of the file
	data, err := ioutil.ReadFile(dataFile)
	if err != nil {
		stdutil.PrintErr("Error opening list", err)
		return
	}

	dataerr := json.Unmarshal([]byte(data), &output)
	if dataerr != nil {
		stdutil.PrintErr("Failed to unmarshal", dataerr)
		return
	}

	// If no argument was provided, check if
	// a list exists and display it.
	if len(cmd) == 1 {
		listPrint(output)
		return
	}

	args := cmd[2:]
	nargs := len(args)
	
	switch cmd[1] {
	case "add":
		if nargs < 1 {
			argErr("add \"<title>\" \"[description]\"")
			return
		}

		if nargs > 2 {
			fmt.Println("Pro tip! Use \"quotes\" to have several words.\nExample: `"+
			strings.ToLower(name)+" add \"long title\" \"long description\"`.\n")
		}

		d := ""
		if len(args) > 1 { d = args[1] }
		output = append(output, ListItem{args[0], d})

		listdata, err := json.Marshal(output)
		if err != nil {
			stdutil.PrintErr("Failed to marshal", err)
			return
		}

		// Add add new item to the file
		jfile, _ := os.Create(dataFile)
		_, werr := jfile.Write(listdata)
		if werr != nil {
			stdutil.PrintErr("Failed to write to file", werr)
			return
		}

		listPrint(output)
	case "remove":
		if nargs != 1 {
			argErr("remove <id>")
			return
		}
		fmt.Println("Something worked!")
	case "rename":
		if nargs < 2 {
			argErr("rename <id> <new title>")
			return
		}
		id := args[0]
		newtitle := strings.Join(args[1:], " ")
		fmt.Println(id)
		fmt.Println(newtitle)
	case "renamedesc":
		if nargs < 2 {
			argErr("renamedesc <id> <new description>")
			return
		}
		id := args[0]
		newdesc := strings.Join(args[1:], " ")
		fmt.Println(id)
		fmt.Println(newdesc)
	case "clear":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Clearing your list is permanent. Please confirm your decision. [y/N]:")
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')

		// It returns y\n, remove it so we can compare.
		// Of course Windows has to be special.
		if runtime.GOOS == "windows" {
			text = strings.TrimRight(text, "\r\n")
		} else {
			text = strings.TrimRight(text, "\n")
		}

		if strings.ToLower(text) == "y" {
			// Delete the file
			delerr := os.Remove(dataFile)
			if delerr != nil {
				stdutil.PrintErr("File deletion failed", delerr)
				return
			}
			fmt.Println("Your list has been cleared.")
		} else {
			fmt.Println("Abort.")
		}
	case "help":
		printHelp()
	case "version":
		fmt.Println(name+" "+version+", running on "+runtime.GOOS+" ("+runtime.GOARCH+").")
	default:
		fmt.Println("Unknown argument.\n")
		printHelp()
	}
}

func listPrint(output []ListItem) {
		headers := make([]string, 0)
		descriptions := make([]string, 0)
		for _, item := range output {
			headers = append(headers, item.Title)
			descriptions = append(descriptions, item.Description)
		}
		printList(headers, descriptions)
}

func printList(titles []string, descriptions []string) {
	color.Set(color.FgBlue, color.Bold)
	fmt.Println("Your Xedo list:")
	color.Unset()
	table := gtable.NewStringTable()
	table.AddStrings("ID", "Title", "Description")

	for i, _ := range titles {
		table.AddRow()
		table.AddStrings(strconv.Itoa(i+1), titles[i], descriptions[i])
	}

	fmt.Println(table.String())
	return
}

func printHelp() {
	fmt.Println(name+" ("+version+"), the todo list manager by Martin Persson <mnpn03@icloud.com>")
	help := make([]string, 0)
	help = append(help, "USAGE:")
	help = append(help, "\tadd \"<title>\" \"[description]\"\tAdd a new entry with an optional description.")
	help = append(help, "\tremove <id>\t\t\tRemove an entry from the list.")
	help = append(help, "\tclear\t\t\t\tDeletes the whole list permanently.")
	help = append(help, "\tmove <id> <up/down/top/bottom>\tMove an entry to a new place.")
	help = append(help, "\trename <id> <title>\t\tRename an entry's title.")
	help = append(help, "\trenamedesc <id> <description>\tRename an entry's description.")
	help = append(help, "\tversion\t\t\t\tPrint the version and OS/arch.")

	fmt.Println(strings.Join(help, "\n"))
}

func argErr(cmderr string) {
	stdutil.PrintErr("Invalid arguments. Usage: `"+strings.ToLower(name)+" "+cmderr+"`.", nil)
}