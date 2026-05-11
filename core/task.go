package core

import (
	"time"
)

const minuteTimeLayout = "2006-01-02T15:04"

type TaskStatus int

const (
	Empty TaskStatus = iota
	InProgress
	Done
)

type Task struct {
	Id         int
	Name       string
	Status     TaskStatus
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewTask(_name string) Task {
	return Task{
		Id:        generateID(),
		Name:      _name,
		CreatedAt: time.Now().Truncate(time.Minute),
	}
}

func (t *Task) GetId() int {
	return t.Id
}

func (t *Task) SetId(value int) {
	t.Id = value
}

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) GetStatus() TaskStatus {
	return t.Status
}

func (t *Task) GetCreationTime() time.Time {
	return t.CreatedAt
}

func (t *Task) GetModifiedTime() time.Time {
	return t.ModifiedAt
}

func (t *Task) SetName(n string) {
	t.Name = n
}

func (t *Task) SetStatus(s TaskStatus) {
	t.Status = s
}

func (task *Task) SetModified(t time.Time) {
	task.ModifiedAt = t
}


