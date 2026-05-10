package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type IApp interface {
	AddTask(name string)
	RemoveTask(id int)
	UpdateTaskName(id int, name string)
	ChangeStatus(id int, s TaskStatus)
	ListDone()
	ListToDo()
	ListInProgress()
}

type App struct {
	Tasks []TaskImpl
}

func (u *App) AddTask(_name string) {
	task := NewTask(_name)
	u.Tasks = append(u.Tasks, task)
}

func (u *App) LoadData(data []byte) {
	if len(data) == 0 {
		fmt.Println("nohting to do")
		return
	}
	// decode if the file is not empty
	err := json.Unmarshal(data, &u.Tasks)
	if err != nil {
		fmt.Printf("Unmarshaling not successuful. %v\n", err)
	}
}

func (u *App) SaveData(destination string) {
	encoded, err := json.Marshal(u.Tasks)
	if err != nil {
		fmt.Printf("Error marshaling tasks: %v\n", err)
	}
	os.WriteFile(destination, encoded, 0644)
}
func (u *App) RemoveTask(taskId int) {
	var index = -1
	for i := 0; i < len(u.Tasks); i++ {
		if u.Tasks[i].GetId() == taskId {
			index = i
		}
	}
	if index == -1 {
		// nothing to remove
		return
	}
	u.Tasks = append(u.Tasks[:index], u.Tasks[index+1:]...)

}

func (u *App) UpdateTaskName(id int, newName string) {
	toUpdate, err := getTaskById(id, u)
	if err != nil {
		return
	}
	toUpdate.SetName(newName)

}

func (u *App) ChangeStatus(id int, s TaskStatus) {
	toUpdate, err := getTaskById(id, u)
	if err != nil {
		return
	}
	toUpdate.SetStatus(s)
}

func (u *App) ListDone() {
	for _, task := range u.Tasks {
		if task.GetStatus() == Done {
			fmt.Println(task.ToString())
		}
	}
}

func (u *App) ListToDo() {
	for _, task := range u.Tasks {
		if task.GetStatus() == Empty {
			fmt.Println(task.ToString())
		}
	}
}

func (u *App) ListInProgress() {
	for _, task := range u.Tasks {
		if task.GetStatus() == InProgress {
			fmt.Println(task.ToString())
		}
	}
}

func getTaskById(value int, u *App) (*TaskImpl, error) {
	for i := 0; i < len(u.Tasks); i++ {
		if u.Tasks[i].GetId() == value {
			return &u.Tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
