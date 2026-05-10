package core

import (
	"encoding/json"
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
	id         int
	name       string
	status     TaskStatus
	createdAt  time.Time
	modifiedAt time.Time
}

func NewTask(_name string) TaskImpl {
	return TaskImpl{
		id:        generateID(),
		name:      _name,
		createdAt: time.Now().Truncate(time.Minute),
	}
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

func (task TaskImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Name       string
		Id         int
		Status     TaskStatus
		CreatedAt  time.Time
		ModifiedAt time.Time
	}{
		Name:       task.name,
		Id:         task.id,
		Status:     task.status,
		CreatedAt:  task.createdAt,
		ModifiedAt: task.modifiedAt,
	})
}

func (task *TaskImpl) UnmarshalJSON(data []byte) error {
	// 1. Define a temporary "shadow" struct with public fields
	// that matches the JSON structure
	type Alias struct {
		Name       string    `json:"Name"`
		Id         int       `json:"Id"`
		Status     TaskStatus `json:"Status"`
		CreatedAt  time.Time `json:"CreatedAt"`
		ModifiedAt time.Time `json:"ModifiedAt"`
	}

	temp := &Alias{}

	// 2. Unmarshal the JSON into the shadow struct
	if err := json.Unmarshal(data, temp); err != nil {
		return err
	}

	// 3. Manually assign the public values to your private fields
	task.name = temp.Name
	task.id = temp.Id
	task.status = temp.Status
	task.createdAt = temp.CreatedAt
	task.modifiedAt = temp.ModifiedAt

	return nil
}

func (task *TaskImpl) ToString() string {
	var statusStr string
	switch task.status {
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
		task.id,
		task.name,
		statusStr,
		task.createdAt.Format("2006-01-02 15:04:05"),
		task.modifiedAt.Format("2006-01-02 15:04:05"))
}
