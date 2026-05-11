package core

import (
	"fmt"
	"time"
)

type TaskStatus int

const (
	Empty TaskStatus = iota
	InProgress
	Done
)

type Task interface {
	GetId() int
	SetId(value int)
	GetName() string
	GetStatus() TaskStatus
	GetCreationTime() time.Time
	GetModifiedTime() time.Time
	SetName(name string)
	SetStatus(s TaskStatus)
	SetModified(t time.Time)
	ToString() string
}

type TaskImpl struct {
	Id         int
	Name       string
	Status     TaskStatus
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewTask(_name string) TaskImpl {
	return TaskImpl{
		Id:        generateID(),
		Name:      _name,
		CreatedAt: time.Now().Truncate(time.Minute),
	}
}

func (t *TaskImpl) GetId() int {
	return t.Id
}

func (t *TaskImpl) SetId(value int) {
	t.Id = value
}

func (t *TaskImpl) GetName() string {
	return t.Name
}

func (t *TaskImpl) GetStatus() TaskStatus {
	return t.Status
}

func (t *TaskImpl) GetCreationTime() time.Time {
	return t.CreatedAt
}

func (t *TaskImpl) GetModifiedTime() time.Time {
	return t.ModifiedAt
}

func (t *TaskImpl) SetName(n string) {
	t.Name = n
}

func (t *TaskImpl) SetStatus(s TaskStatus) {
	t.Status = s
}

func (task *TaskImpl) SetModified(t time.Time) {
	task.ModifiedAt = t
}

func (task *TaskImpl) ToString() string {
	var statusStr string
	switch task.Status {
	case Empty:
		statusStr = "Empty"
	case InProgress:
		statusStr = "InProgress"
	case Done:
		statusStr = "Done"
	default:
		statusStr = "Unknown"
	}

	return fmt.Sprintf("ID:%d Name:%s Status:%s Created:%s Modified:%s",
		task.Id,
		task.Name,
		statusStr,
		task.CreatedAt.Format("2006-01-02 15:04:05"),
		task.ModifiedAt.Format("2006-01-02 15:04:05"))
}
