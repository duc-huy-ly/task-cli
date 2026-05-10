package core

import "testing"

func TestAdd(t *testing.T) {
	u := App{}
	u.AddTask("go running")
	if len(u.tasks) != 1 {
		t.Fatalf("adding task didn't work")
	}
	if u.tasks[0].GetName() != "go running" {
		t.Fatalf("naming went wrong in task, got %v", u.tasks[0].GetName())
	}
}

func TestId(t *testing.T) {
	u := App{}
	u.AddTask("first")
	u.AddTask("Second")
	u.AddTask("third")

	if u.tasks[0].GetId() != 1 {
		t.Fatalf("expected first task ID 1, got %d", u.tasks[0].GetId())
	}
	if u.tasks[1].GetId() != 2 {
		t.Fatalf("expected second task ID 2, got %d", u.tasks[1].GetId())
	}
	if u.tasks[2].GetId() != 3 {
		t.Fatalf("expected third task ID 3, got %d", u.tasks[2].GetId())
	}
}

func TestRemove(t *testing.T) {
	u := App{}
	u.AddTask("first")
	u.AddTask("Second")
	u.RemoveTask(1)
	if len(u.tasks) != 1 {
		t.Fatalf("expected to have removed task from slice. Slice size is %v", len(u.tasks))
	}
	if u.tasks[0].GetId() != 2 {
		t.Fatalf("id must be unchagned. got %v", u.tasks[0].GetId())
	}
}

func TestGetTaskWithInexistentId(t *testing.T) {
	u := App{}
	u.AddTask("first")
	u.AddTask("second")

	// Try to get a task with ID 999 that doesn't exist
	_, err := getTaskById(999, &u)
	if err == nil {
		t.Fatalf("expected error when getting task with inexistent ID, got nil")
	}
}

func TestChangeNameUser(t *testing.T) {
	u := App{}
	u.AddTask("my name")
	u.UpdateTaskName(1, "new name")
	if u.tasks[0].GetName() != "new name" {
		t.Fatalf("expected name to have changed to 'new name'. Got %v", u.tasks[0].GetName())
	}
}
