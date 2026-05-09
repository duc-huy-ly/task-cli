package core

import "time"

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
		if u.tasks[i].GetId()== taskId {
			index = i
		}	
	}
	u.tasks = append(u.tasks[:index], u.tasks[index+1:]...)
	

}
