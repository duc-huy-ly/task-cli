package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dly/task-cli/core"
)

const dataDirectory = "data/tasks.json"

func main() {
	checkFileExistence()
	data := loadDataFromFile() 
	user := core.App{}
	user.LoadData(data)
	commandStr, args := parseUserInput(os.Args)
	if commandStr== "" {
		fmt.Println("No command given")
		return
	}
	handleCommandStr(commandStr, args, &user)
	user.SaveData(dataDirectory)
}

func handleCommandStr(commandStr string, args []string, user *core.App) {
	action := strings.ToLower(commandStr)
	if action == "add" && len(args) != 0{
		user.AddTask(strings.Join(args, " "))
	} else if action == "remove" && len(args) != 0{
		// user.RemoveTask(int)
		fmt.Printf("delete %v\n", args[0])
	}else{
		fmt.Printf("Action not recognized\n")
	}
}

func checkFileExistence()  error {
	file, err := os.OpenFile(dataDirectory, os.O_CREATE, 0644)
	if err != nil {
		if os.IsExist(err){
			return nil
		}
		return err
	} 
	defer file.Close()
	return nil
}

func loadDataFromFile() []byte{
	data, err := os.ReadFile(dataDirectory)
	if err != nil {
		fmt.Printf("Error reading file : %v\n", err)
	}
	return data
}

func parseUserInput(input []string) (string, []string){
	// case where if we call the executable with no command, error
	if len(input) ==1  {
		return "", nil

	}	
	return input[1], input[2:] 
}