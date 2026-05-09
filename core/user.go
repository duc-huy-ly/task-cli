package core

import (
	"errors"
	"fmt"
	"time"
)

type IUser interface {
	AddTask(name string)
	RemoveTask(id int)
	UpdateTaskName(id int, name string)
	ChangeStatus(id int, s TaskStatus)
	ListDone()
	ListToDo()
	ListInProgress()
}

type User struct {
	tasks []Task
}

func (u *User) AddTask(_name string) {
	task := &TaskImpl{
		name:      _name,
		id:        generateID(),
		status:    Empty,
		createdAt: time.Now().Truncate(time.Minute),
	}
	u.tasks = append(u.tasks, task)

}

func (u *User) RemoveTask(taskId int) {
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

func (u *User) UpdateTaskName(id int, newName string) {
	toUpdate, err := getTaskById(id, u)
	if err != nil {
		return
	}
	toUpdate.SetName(newName)

}

func (u *User) ChangeStatus(id int, s TaskStatus) {
	toUpdate, err := getTaskById(id, u)
	if err != nil {
		return
	}
	toUpdate.SetStatus(s)
}

func (u *User) ListDone() {
	for _, task := range u.tasks {
		if task.GetStatus() == Done {
			fmt.Println(task.ToString())
		}
	}
}

func (u *User) ListToDo() {
	for _, task := range u.tasks {
		if task.GetStatus() == Empty {
			fmt.Println(task.ToString())
		}
	}
}

func (u *User) ListInProgress() {
	for _, task := range u.tasks {
		if task.GetStatus() == InProgress {
			fmt.Println(task.ToString())
		}
	}
}

func getTaskById(value int, u *User) (Task, error) {
	for i := 0; i < len(u.tasks); i++ {
		if u.tasks[i].GetId() == value {
			return u.tasks[i], nil
		}
	}
	return nil, errors.New("task not found")
}
