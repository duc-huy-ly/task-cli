package core

import "testing"

func TestTaskDefaults(t *testing.T) {
	tk := TaskImpl{}
	if tk.Name != "" {
		t.Fatalf("expected empty Dest, got %q", tk.Name)
	}
	if tk.Status != Empty {
		t.Fatalf("expected status Empty, got %v", tk.Status)
	}
}

func TestChangeNameTask(t *testing.T) {
	tk := TaskImpl{Name: "Ben"}
	tk.SetName("Tim")
	if tk.Name != "Tim" {
		t.Fatalf("expected name to be changed, got %v", tk.Name)
	}
}
