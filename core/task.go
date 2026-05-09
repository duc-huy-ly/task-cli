package core

import (
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
	GetName() string
	GetStatus() TaskStatus
	GetCreationTime() time.Time
	GetModifiedTime() time.Time
	SetName(name string)
	SetStatus(s TaskStatus)
	SetModified(t time.Time)
}

type TaskImpl struct {
	id         int
	name       string
	status     TaskStatus
	createdAt  time.Time
	modifiedAt time.Time
}

func (t *TaskImpl) GetId() int {
	return t.id
}

func (t *TaskImpl) GetName() string {
	return t.name
}

func (t *TaskImpl) GetStatus() TaskStatus {
	return t.status
}

func (t *TaskImpl) GetCreationTime() time.Time {
	return t.createdAt
}

func (t *TaskImpl) GetModifiedTime() time.Time {
	return t.modifiedAt
}

func (t *TaskImpl) SetName(n string) {
	t.name = n
}

func (t *TaskImpl) SetStatus(s TaskStatus) {
	t.status = s
}

func (task *TaskImpl) SetModified(t time.Time) {
	task.modifiedAt = t
}
