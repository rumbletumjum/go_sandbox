package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)
	typeFlag := newCmd.String("type", "widget", "type of thing")

	switch os.Args[1] {
	case "new":
		newCmd.Parse(os.Args[2:])
		fmt.Println("type: ", *typeFlag)
	}
}
