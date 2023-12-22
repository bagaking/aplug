package methodology

import "testing"

func TestListReturnsFullCopies(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				MainIdea: "isolate the smallest failing case",
				Scenario: []string{"debugging"},
				Strategy: []string{"reproduce"},
				Steps:    []string{"run focused test"},
				Examples: []string{"go test ./..."},
			},
		},
	}

	got := table.List()
	if len(got) != 1 {
		t.Fatalf("List() length = %d, want %d", len(got), 1)
	}
	if got[0].MainIdea != "isolate the smallest failing case" {
		t.Errorf("List()[0].MainIdea = %q, want %q", got[0].MainIdea, "isolate the smallest failing case")
	}
	if got[0].Strategy[0] != "reproduce" {
		t.Errorf("List()[0].Strategy[0] = %q, want %q", got[0].Strategy[0], "reproduce")
	}
	if got[0].Steps[0] != "run focused test" {
		t.Errorf("List()[0].Steps[0] = %q, want %q", got[0].Steps[0], "run focused test")
	}
	if got[0].Examples[0] != "go test ./..." {
		t.Errorf("List()[0].Examples[0] = %q, want %q", got[0].Examples[0], "go test ./...")
	}

	got[0].Scenario[0] = "changed"
	got[0].Strategy[0] = "changed"
	got[0].Steps[0] = "changed"
	got[0].Examples[0] = "changed"

	again := table.List()
	if len(again) != 1 {
		t.Fatalf("List() after mutation length = %d, want %d", len(again), 1)
	}
	if again[0].Scenario[0] != "debugging" {
		t.Errorf("List()[0].Scenario[0] after mutation = %q, want %q", again[0].Scenario[0], "debugging")
	}
	if again[0].Strategy[0] != "reproduce" {
		t.Errorf("List()[0].Strategy[0] after mutation = %q, want %q", again[0].Strategy[0], "reproduce")
	}
	if again[0].Steps[0] != "run focused test" {
		t.Errorf("List()[0].Steps[0] after mutation = %q, want %q", again[0].Steps[0], "run focused test")
	}
	if again[0].Examples[0] != "go test ./..." {
		t.Errorf("List()[0].Examples[0] after mutation = %q, want %q", again[0].Examples[0], "go test ./...")
	}
}

func TestTryGetReturnsCopy(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				MainIdea: "isolate the smallest failing case",
				Scenario: []string{"debugging"},
				Strategy: []string{"reproduce"},
				Steps:    []string{"run focused test"},
				Examples: []string{"go test ./..."},
			},
		},
	}

	got := table.TryGet("debug")
	if got == nil {
		t.Fatalf("TryGet(%q) = nil, want methodology", "debug")
	}

	got.Scenario[0] = "changed"
	got.Strategy[0] = "changed"
	got.Steps[0] = "changed"
	got.Examples[0] = "changed"

	again := table.TryGet("debug")
	if again.Scenario[0] != "debugging" {
		t.Errorf("TryGet(%q).Scenario[0] after mutation = %q, want %q", "debug", again.Scenario[0], "debugging")
	}
	if again.Strategy[0] != "reproduce" {
		t.Errorf("TryGet(%q).Strategy[0] after mutation = %q, want %q", "debug", again.Strategy[0], "reproduce")
	}
	if again.Steps[0] != "run focused test" {
		t.Errorf("TryGet(%q).Steps[0] after mutation = %q, want %q", "debug", again.Steps[0], "run focused test")
	}
	if again.Examples[0] != "go test ./..." {
		t.Errorf("TryGet(%q).Examples[0] after mutation = %q, want %q", "debug", again.Examples[0], "go test ./...")
	}
}

func TestMGetSkipsMissingKeys(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {ID: "debug", Usage: "inspect a failure"},
			"plan":  {ID: "plan", Usage: "sequence work"},
		},
	}

	got := table.MGet("debug", "missing", "plan")
	if len(got) != 2 {
		t.Fatalf("MGet(%q, %q, %q) length = %d, want %d", "debug", "missing", "plan", len(got), 2)
	}
	if got[0].ID != "debug" {
		t.Errorf("MGet(%q, %q, %q)[0].ID = %q, want %q", "debug", "missing", "plan", got[0].ID, "debug")
	}
	if got[1].ID != "plan" {
		t.Errorf("MGet(%q, %q, %q)[1].ID = %q, want %q", "debug", "missing", "plan", got[1].ID, "plan")
	}
}

func TestMGetReturnsCopies(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				MainIdea: "isolate the smallest failing case",
				Scenario: []string{"debugging"},
				Strategy: []string{"reproduce"},
				Steps:    []string{"run focused test"},
				Examples: []string{"go test ./..."},
			},
		},
	}

	got := table.MGet("debug")
	if len(got) != 1 {
		t.Fatalf("MGet(%q) length = %d, want %d", "debug", len(got), 1)
	}

	got[0].Scenario[0] = "changed"
	got[0].Strategy[0] = "changed"
	got[0].Steps[0] = "changed"
	got[0].Examples[0] = "changed"

	again := table.MGet("debug")
	if len(again) != 1 {
		t.Fatalf("MGet(%q) after mutation length = %d, want %d", "debug", len(again), 1)
	}
	if again[0].Scenario[0] != "debugging" {
		t.Errorf("MGet(%q)[0].Scenario[0] after mutation = %q, want %q", "debug", again[0].Scenario[0], "debugging")
	}
	if again[0].Strategy[0] != "reproduce" {
		t.Errorf("MGet(%q)[0].Strategy[0] after mutation = %q, want %q", "debug", again[0].Strategy[0], "reproduce")
	}
	if again[0].Steps[0] != "run focused test" {
		t.Errorf("MGet(%q)[0].Steps[0] after mutation = %q, want %q", "debug", again[0].Steps[0], "run focused test")
	}
	if again[0].Examples[0] != "go test ./..." {
		t.Errorf("MGet(%q)[0].Examples[0] after mutation = %q, want %q", "debug", again[0].Examples[0], "go test ./...")
	}
}

func TestGetBySceneReturnsCopies(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				Scenario: []string{"debugging", "testing"},
			},
			"plan": {
				ID:       "plan",
				Usage:    "sequence work",
				Scenario: []string{"planning"},
			},
		},
	}

	got := table.GetByScene("debugging")
	if len(got) != 1 {
		t.Fatalf("GetByScene(%q) length = %d, want %d", "debugging", len(got), 1)
	}
	if got[0].ID != "debug" {
		t.Errorf("GetByScene(%q)[0].ID = %q, want %q", "debugging", got[0].ID, "debug")
	}

	got[0].Scenario[0] = "changed"
	again := table.GetByScene("debugging")
	if again[0].Scenario[0] != "debugging" {
		t.Errorf("GetByScene(%q)[0].Scenario[0] after mutation = %q, want %q", "debugging", again[0].Scenario[0], "debugging")
	}
}

func TestSetStoresCopy(t *testing.T) {
	table := &Methodologies{
		data: make(map[string]*Methodology),
	}
	input := &Methodology{
		ID:       "debug",
		Usage:    "inspect a failure",
		MainIdea: "isolate the smallest failing case",
		Scenario: []string{"debugging"},
		Strategy: []string{"reproduce"},
		Steps:    []string{"run focused test"},
		Examples: []string{"go test ./..."},
	}

	table.Set("debug", input)
	input.Scenario[0] = "changed"
	input.Strategy[0] = "changed"
	input.Steps[0] = "changed"
	input.Examples[0] = "changed"

	got := table.TryGet("debug")
	if got == nil {
		t.Fatalf("TryGet(%q) after Set = nil, want methodology", "debug")
	}
	if got.Scenario[0] != "debugging" {
		t.Errorf("TryGet(%q).Scenario[0] after input mutation = %q, want %q", "debug", got.Scenario[0], "debugging")
	}
	if got.Strategy[0] != "reproduce" {
		t.Errorf("TryGet(%q).Strategy[0] after input mutation = %q, want %q", "debug", got.Strategy[0], "reproduce")
	}
	if got.Steps[0] != "run focused test" {
		t.Errorf("TryGet(%q).Steps[0] after input mutation = %q, want %q", "debug", got.Steps[0], "run focused test")
	}
	if got.Examples[0] != "go test ./..." {
		t.Errorf("TryGet(%q).Examples[0] after input mutation = %q, want %q", "debug", got.Examples[0], "go test ./...")
	}
}

func TestSetInitializesZeroValueContainer(t *testing.T) {
	var table Methodologies
	input := &Methodology{
		ID:       "debug",
		Usage:    "inspect a failure",
		Scenario: []string{"debugging"},
	}

	table.Set("debug", input)

	got := table.TryGet("debug")
	if got == nil {
		t.Fatalf("TryGet(%q) after zero-value Set = nil, want methodology", "debug")
	}
	if got.ID != "debug" {
		t.Errorf("TryGet(%q).ID after zero-value Set = %q, want %q", "debug", got.ID, "debug")
	}
	if got.Scenario[0] != "debugging" {
		t.Errorf("TryGet(%q).Scenario[0] after zero-value Set = %q, want %q", "debug", got.Scenario[0], "debugging")
	}
}

func TestUpdatedByUsesCopies(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				MainIdea: "isolate the smallest failing case",
				Scenario: []string{"debugging"},
				Strategy: []string{"reproduce"},
				Steps:    []string{"run focused test"},
				Examples: []string{"go test ./..."},
			},
		},
	}

	original := table.data["debug"]
	var returned *Methodology
	err := table.UpdatedBy("debug", func(m *Methodology) *Methodology {
		if m == original {
			t.Errorf("UpdatedBy(%q) passed stored methodology pointer to callback, want copy", "debug")
		}
		m.Scenario[0] = "callback changed"
		m.Strategy[0] = "callback changed"
		m.Steps[0] = "callback changed"
		m.Examples[0] = "callback changed"

		returned = m
		return m
	})
	if err != nil {
		t.Fatalf("UpdatedBy(%q) error = %v, want nil", "debug", err)
	}

	returned.Scenario[0] = "changed after update"
	returned.Strategy[0] = "changed after update"
	returned.Steps[0] = "changed after update"
	returned.Examples[0] = "changed after update"

	got := table.TryGet("debug")
	if got.Scenario[0] != "callback changed" {
		t.Errorf("TryGet(%q).Scenario[0] after returned mutation = %q, want %q", "debug", got.Scenario[0], "callback changed")
	}
	if got.Strategy[0] != "callback changed" {
		t.Errorf("TryGet(%q).Strategy[0] after returned mutation = %q, want %q", "debug", got.Strategy[0], "callback changed")
	}
	if got.Steps[0] != "callback changed" {
		t.Errorf("TryGet(%q).Steps[0] after returned mutation = %q, want %q", "debug", got.Steps[0], "callback changed")
	}
	if got.Examples[0] != "callback changed" {
		t.Errorf("TryGet(%q).Examples[0] after returned mutation = %q, want %q", "debug", got.Examples[0], "callback changed")
	}
}
