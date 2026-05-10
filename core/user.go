package core

import (
	"errors"
	"fmt"
	"time"
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
	tasks []Task
}

func (u *App) AddTask(_name string) {
	task := &TaskImpl{
		name:      _name,
		id:        generateID(),
		status:    Empty,
		createdAt: time.Now().Truncate(time.Minute),
	}
	u.tasks = append(u.tasks, task)

}

func (u *App) RemoveTask(taskId int) {
	var index = -1
	for i := 0; i < len(u.tasks); i++ {
		if u.tasks[i].GetId() == taskId {
			index = i
		}
	}
	if index == -1 {
		// nothing to remove
		return
	}
	u.tasks = append(u.tasks[:index], u.tasks[index+1:]...)

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
	for _, task := range u.tasks {
		if task.GetStatus() == Done {
			fmt.Println(task.ToString())
		}
	}
}

func (u *App) ListToDo() {
	for _, task := range u.tasks {
		if task.GetStatus() == Empty {
			fmt.Println(task.ToString())
		}
	}
}

func (u *App) ListInProgress() {
	for _, task := range u.tasks {
		if task.GetStatus() == InProgress {
			fmt.Println(task.ToString())
		}
	}
}

func getTaskById(value int, u *App) (Task, error) {
	for i := 0; i < len(u.tasks); i++ {
		if u.tasks[i].GetId() == value {
			return u.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
