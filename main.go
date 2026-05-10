package main

import (
	"fmt"
	"os"

	"github.com/dly/task-cli/core"
)

const dataDirectory = "data/tasks.json"

func main() {
	fmt.Println("Hello app")
	checkFileExistence()
	data := loadDataFromFile() 
	user := core.App{}
	user.LoadData(data)
	user.AddTask("hello")
	user.SaveData(dataDirectory)
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
