package main

import (
	"fmt"
	"github.com/dly/task-cli/core"
	"os"
	"strconv"
	"strings"
)

const dataDirectory = "data/tasks.json"

func main() {
	checkFileExistence()
	data := loadDataFromFile()
	user := core.App{}
	user.LoadData(data)
	commandStr, args := parseUserInput(os.Args)
	if commandStr == "" {
		fmt.Println("Welcome to task-cli. -h for help")
		return
	}
	handleCommandStr(commandStr, args, &user)
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
		if err!= nil{
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

func checkFileExistence() error {
	file, err := os.OpenFile(dataDirectory, os.O_CREATE, 0644)
	if err != nil {
		if os.IsExist(err) {
			return nil
		}
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

func parseUserInput(input []string) (string, []string) {
	// case where if we call the executable with no command, error
	if len(input) == 1 {
		return "", nil

	}
	return input[1], input[2:]
}
