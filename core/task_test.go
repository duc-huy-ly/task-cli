package core

import "testing"

func TestTaskDefaults(t *testing.T) {
	tk := TaskImpl{}
	if tk.name != "" {
		t.Fatalf("expected empty Dest, got %q", tk.name)
	}
	if tk.status != Empty {
		t.Fatalf("expected status Empty, got %v", tk.status)
	}
}

func TestChangeName(t *testing.T) {
	tk := TaskImpl{name: "Ben"}
	tk.SetName("Tim")
	if tk.name != "Tim" {
		t.Fatalf("expected name to be changed, got %v", tk.name)
	}
}
