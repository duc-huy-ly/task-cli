package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/dly/task-cli/core"
)

var dataDirectory = "data/tasks.json"

func main() {
	// Define flags
	helpFlag := flag.Bool("h", false, "Display help message")
	flag.Parse()

	// Check if help flag was provided
	if *helpFlag {
		displayHelp()
		return
	}

	checkFileExistence()
	data := loadDataFromFile()
	user := core.App{}
	user.LoadData(data)

	// Get remaining arguments after flag parsing
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("Welcome to task-cli. -h for help")
		return
	}

	commandStr := args[0]
	commandArgs := args[1:]
	handleCommandStr(commandStr, commandArgs, &user)
	user.SaveData(dataDirectory)
}

func handleCommandStr(commandStr string, args []string, user *core.App) {
	switch strings.ToLower(commandStr) {
	case "add":
		if len(args) == 0 {
			fmt.Printf("Specify argument to add a new task.\n")
			return
		}
		user.AddTask(strings.Join(args, " "))
	case "rm":
		fallthrough
	case "remove":
		if len(args) == 0 {
			fmt.Printf("Give index of task to remove\n")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error converting argument to string : %v", err)
			return
		}
		user.RemoveTask(index)
	case "update":
		if len(args) < 2 {
			fmt.Printf("Update must be followed with index + new name of the task")
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Str to int conversion err : %v\n", err)
			return
		}
		user.UpdateTaskName(index, strings.Join(args[1:], " "))
	case "l":
		fallthrough
	case "list":
		if len(args) == 0 {
			user.ListAll()
			return
		}
		switch args[0] {
		case "done":
			user.ListDone()
		case "todo":
			user.ListToDo()
		case "in-progress":
			user.ListInProgress()
		default:
			raiseInvalidArgument()
			return
		}
	case "mark-in-progress":
		if len(args) == 0 {
			raiseInvalidArgument()
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error converting str to int, %v\n", err)
			return
		}
		user.ChangeStatus(index, core.InProgress)
	case "mark-done":
		if len(args) == 0 {
			raiseInvalidArgument()
			return
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error converting str to int, %v\n", err)
			return
		}
		user.ChangeStatus(index, core.Done)
	default:
		raiseInvalidArgument()
		return
	}
}

func raiseInvalidArgument() {
	fmt.Printf("Unknown command\n")
}

func displayHelp() {
	fmt.Println("Task CLI - A simple task management tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  task-cli <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  add <task>              Add a new task")
	fmt.Println("  rm, remove <index>      Remove a task by index")
	fmt.Println("  update <index> <name>   Update task name")
	fmt.Println("  list [status]           List tasks (optionally filtered by status)")
	fmt.Println("    - list done           Show completed tasks")
	fmt.Println("    - list todo           Show incomplete tasks")
	fmt.Println("    - list in-progress    Show in-progress tasks")
	fmt.Println("  mark-done <index>       Mark a task as done")
	fmt.Println("  mark-in-progress <index> Mark a task as in-progress")
	fmt.Println("  -h, --help              Display this help message")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  task-cli add \"Buy groceries\"")
	fmt.Println("  task-cli list")
	fmt.Println("  task-cli mark-done 1")
	fmt.Println("  task-cli remove 2")
}

func checkFileExistence() error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(dataDirectory)
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	// Create file if it doesn't exist
	file, err := os.OpenFile(dataDirectory, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func loadDataFromFile() []byte {
	data, err := os.ReadFile(dataDirectory)
	if err != nil {
		fmt.Printf("Error reading file : %v\n", err)
	}
	return data
}

