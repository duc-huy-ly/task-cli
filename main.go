package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	help := flag.Bool("h", false, "show help")
	flag.Parse()
	if *help {
		printHelp()
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("Hello world")
		return
	}
	command := os.Args[1]
	possibleCommands := []string{"add", "delete", "update", "mark-in-progress", "mark-done", "list"}
	for _, m := range possibleCommands {
		if m == command {
			fmt.Println(m)
			return
		}
	}
	fmt.Println("command not not exist")

}

func printHelp() {
	fmt.Println("Help for the cli-app")
}

