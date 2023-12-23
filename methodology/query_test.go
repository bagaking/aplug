package methodology

import (
	"testing"
	"time"
)

func assertStringSliceNil(t *testing.T, field string, got []string) {
	t.Helper()
	if got != nil {
		t.Errorf("%s = %#v, want nil", field, got)
	}
}

func assertStringSliceEmpty(t *testing.T, field string, got []string) {
	t.Helper()
	if got == nil {
		t.Errorf("%s = nil, want non-nil empty slice", field)
	}
	if len(got) != 0 {
		t.Errorf("%s length = %d, want %d", field, len(got), 0)
	}
}

func assertMethodologyStringSlices(t *testing.T, got *Methodology, check func(*testing.T, string, []string)) {
	t.Helper()
	for _, field := range []struct {
		name string
		got  []string
	}{
		{name: "Scenario", got: got.Scenario},
		{name: "Strategy", got: got.Strategy},
		{name: "Steps", got: got.Steps},
		{name: "Examples", got: got.Examples},
	} {
		t.Run(field.name, func(t *testing.T) {
			check(t, field.name, field.got)
		})
	}
}

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

func TestMGetPreservesSliceNilAndEmptySemantics(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"nil-slices": {
				ID: "nil-slices",
			},
			"empty-slices": {
				ID:       "empty-slices",
				Scenario: []string{},
				Strategy: []string{},
				Steps:    []string{},
				Examples: []string{},
			},
		},
	}

	got := table.MGet("nil-slices", "empty-slices")
	if len(got) != 2 {
		t.Fatalf("MGet(%q, %q) length = %d, want %d", "nil-slices", "empty-slices", len(got), 2)
	}

	t.Run("nil slices", func(t *testing.T) {
		assertMethodologyStringSlices(t, got[0], assertStringSliceNil)
	})
	t.Run("empty slices", func(t *testing.T) {
		assertMethodologyStringSlices(t, got[1], assertStringSliceEmpty)
	})
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

func TestGetByScenePreservesSliceNilAndEmptySemantics(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"nil-slices": {
				ID:       "nil-slices",
				Scenario: []string{"debugging"},
			},
			"empty-slices": {
				ID:       "empty-slices",
				Scenario: []string{"debugging"},
				Strategy: []string{},
				Steps:    []string{},
				Examples: []string{},
			},
		},
	}

	got := table.GetByScene("debugging")
	if len(got) != 2 {
		t.Fatalf("GetByScene(%q) length = %d, want %d", "debugging", len(got), 2)
	}
	byID := map[string]*Methodology{
		got[0].ID: got[0],
		got[1].ID: got[1],
	}

	t.Run("nil slices", func(t *testing.T) {
		assertStringSliceNil(t, "Strategy", byID["nil-slices"].Strategy)
		assertStringSliceNil(t, "Steps", byID["nil-slices"].Steps)
		assertStringSliceNil(t, "Examples", byID["nil-slices"].Examples)
	})
	t.Run("empty slices", func(t *testing.T) {
		assertStringSliceEmpty(t, "Strategy", byID["empty-slices"].Strategy)
		assertStringSliceEmpty(t, "Steps", byID["empty-slices"].Steps)
		assertStringSliceEmpty(t, "Examples", byID["empty-slices"].Examples)
	})
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

func TestSetAndQueryPreserveSliceNilAndEmptySemantics(t *testing.T) {
	tests := []struct {
		name  string
		input *Methodology
		check func(*testing.T, string, []string)
	}{
		{
			name: "nil slices",
			input: &Methodology{
				ID: "nil-slices",
			},
			check: func(t *testing.T, field string, got []string) {
				t.Helper()
				if got != nil {
					t.Errorf("%s = %#v, want nil", field, got)
				}
			},
		},
		{
			name: "empty slices",
			input: &Methodology{
				ID:       "empty-slices",
				Scenario: []string{},
				Strategy: []string{},
				Steps:    []string{},
				Examples: []string{},
			},
			check: func(t *testing.T, field string, got []string) {
				t.Helper()
				if got == nil {
					t.Errorf("%s = nil, want non-nil empty slice", field)
				}
				if len(got) != 0 {
					t.Errorf("%s length = %d, want %d", field, len(got), 0)
				}
			},
		},
	}

	checkFields := func(t *testing.T, got *Methodology, check func(*testing.T, string, []string)) {
		t.Helper()
		for _, field := range []struct {
			name string
			got  []string
		}{
			{name: "Scenario", got: got.Scenario},
			{name: "Strategy", got: got.Strategy},
			{name: "Steps", got: got.Steps},
			{name: "Examples", got: got.Examples},
		} {
			t.Run(field.name, func(t *testing.T) {
				t.Helper()
				check(t, field.name, field.got)
			})
		}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := &Methodologies{
				data: make(map[string]*Methodology),
			}
			table.Set(tt.input.ID, tt.input)

			got := table.TryGet(tt.input.ID)
			if got == nil {
				t.Fatalf("TryGet(%q) after Set = nil, want methodology", tt.input.ID)
			}
			checkFields(t, got, tt.check)

			list := table.List()
			if len(list) != 1 {
				t.Fatalf("List() length = %d, want %d", len(list), 1)
			}
			checkFields(t, list[0], tt.check)
		})
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

func TestSetNilRemovesMethodology(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {ID: "debug", Usage: "inspect a failure"},
		},
	}

	table.Set("debug", nil)

	if got := table.TryGet("debug"); got != nil {
		t.Fatalf("TryGet(%q) after Set nil = %#v, want nil", "debug", got)
	}
	if got := table.List(); len(got) != 0 {
		t.Fatalf("List() after Set nil length = %d, want %d", len(got), 0)
	}
}

func TestZeroValueContainerQueryBoundaries(t *testing.T) {
	var table Methodologies

	tests := []struct {
		name string
		got  func() []*Methodology
	}{
		{
			name: "List",
			got:  table.List,
		},
		{
			name: "MGet",
			got: func() []*Methodology {
				return table.MGet("missing")
			},
		},
		{
			name: "GetByScene",
			got: func() []*Methodology {
				return table.GetByScene("missing")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.got()
			if len(got) != 0 {
				t.Errorf("%s() on zero-value container length = %d, want %d", tt.name, len(got), 0)
			}
		})
	}

	if got := table.TryGet("missing"); got != nil {
		t.Errorf("TryGet(%q) on zero-value container = %#v, want nil", "missing", got)
	}
}

func TestQuerySkipsNilMethodologies(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {
				ID:       "debug",
				Usage:    "inspect a failure",
				Scenario: []string{"debugging"},
			},
			"nil": nil,
		},
	}

	gotList := table.List()
	if len(gotList) != 1 {
		t.Fatalf("List() with nil methodology length = %d, want %d", len(gotList), 1)
	}
	if gotList[0].ID != "debug" {
		t.Errorf("List() with nil methodology returned ID = %q, want %q", gotList[0].ID, "debug")
	}

	gotMGet := table.MGet("nil", "debug")
	if len(gotMGet) != 1 {
		t.Fatalf("MGet(%q, %q) length = %d, want %d", "nil", "debug", len(gotMGet), 1)
	}
	if gotMGet[0].ID != "debug" {
		t.Errorf("MGet(%q, %q)[0].ID = %q, want %q", "nil", "debug", gotMGet[0].ID, "debug")
	}

	gotScene := table.GetByScene("debugging")
	if len(gotScene) != 1 {
		t.Fatalf("GetByScene(%q) with nil methodology length = %d, want %d", "debugging", len(gotScene), 1)
	}
	if gotScene[0].ID != "debug" {
		t.Errorf("GetByScene(%q)[0].ID = %q, want %q", "debugging", gotScene[0].ID, "debug")
	}
}

func TestMissingKeyModifyBoundaries(t *testing.T) {
	tests := []struct {
		name    string
		table   Methodologies
		act     func(*Methodologies) error
		wantErr bool
	}{
		{
			name: "Delete zero-value missing key",
			act: func(table *Methodologies) error {
				table.Delete("missing")
				return nil
			},
		},
		{
			name: "Delete initialized missing key",
			table: Methodologies{
				data: map[string]*Methodology{
					"debug": {ID: "debug", Usage: "inspect a failure"},
				},
			},
			act: func(table *Methodologies) error {
				table.Delete("missing")
				return nil
			},
		},
		{
			name: "UpdatedBy zero-value missing key",
			act: func(table *Methodologies) error {
				return table.UpdatedBy("missing", func(m *Methodology) *Methodology {
					return m
				})
			},
			wantErr: true,
		},
		{
			name: "UpdatedBy initialized missing key",
			table: Methodologies{
				data: map[string]*Methodology{
					"debug": {ID: "debug", Usage: "inspect a failure"},
				},
			},
			act: func(table *Methodologies) error {
				return table.UpdatedBy("missing", func(m *Methodology) *Methodology {
					return m
				})
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			table := tt.table
			err := tt.act(&table)
			if gotErr := err != nil; gotErr != tt.wantErr {
				t.Errorf("%s error = %v, want error presence = %t", tt.name, err, tt.wantErr)
			}
			if got := table.TryGet("debug"); tt.table.data != nil && got == nil {
				t.Errorf("%s removed existing key %q, want it retained", tt.name, "debug")
			}
		})
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

func TestUpdatedByPreservesSliceNilAndEmptySemantics(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"nil-slices": {
				ID: "nil-slices",
			},
			"empty-slices": {
				ID:       "empty-slices",
				Scenario: []string{},
				Strategy: []string{},
				Steps:    []string{},
				Examples: []string{},
			},
		},
	}

	err := table.UpdatedBy("nil-slices", func(m *Methodology) *Methodology {
		assertMethodologyStringSlices(t, m, assertStringSliceNil)
		return m
	})
	if err != nil {
		t.Fatalf("UpdatedBy(%q) error = %v, want nil", "nil-slices", err)
	}
	assertMethodologyStringSlices(t, table.TryGet("nil-slices"), assertStringSliceNil)

	err = table.UpdatedBy("empty-slices", func(m *Methodology) *Methodology {
		assertMethodologyStringSlices(t, m, assertStringSliceEmpty)
		return m
	})
	if err != nil {
		t.Fatalf("UpdatedBy(%q) error = %v, want nil", "empty-slices", err)
	}
	assertMethodologyStringSlices(t, table.TryGet("empty-slices"), assertStringSliceEmpty)
}

func TestUpdatedByNilRemovesMethodology(t *testing.T) {
	table := &Methodologies{
		data: map[string]*Methodology{
			"debug": {ID: "debug", Usage: "inspect a failure"},
		},
	}

	err := table.UpdatedBy("debug", func(m *Methodology) *Methodology {
		if m == nil {
			t.Fatal("UpdatedBy callback received nil, want existing methodology copy")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("UpdatedBy(%q) error = %v, want nil", "debug", err)
	}
	if got := table.TryGet("debug"); got != nil {
		t.Fatalf("TryGet(%q) after UpdatedBy nil = %#v, want nil", "debug", got)
	}
	if got := table.MGet("debug"); len(got) != 0 {
		t.Fatalf("MGet(%q) after UpdatedBy nil length = %d, want %d", "debug", len(got), 0)
	}
}

func TestNewContainerTakesLockedSnapshot(t *testing.T) {
	originalBeforeDefaultSnapshotRLock := beforeDefaultSnapshotRLock
	t.Cleanup(func() {
		beforeDefaultSnapshotRLock = originalBeforeDefaultSnapshotRLock
	})

	defaults := DefaultContainer()

	readyToLock := make(chan struct{})
	beforeDefaultSnapshotRLock = func() {
		close(readyToLock)
	}
	defaults.mu.Lock()
	locked := true
	defer func() {
		if locked {
			defaults.mu.Unlock()
		}
	}()

	done := make(chan struct{})
	go func() {
		defer close(done)
		NewContainer()
	}()

	select {
	case <-readyToLock:
	case <-time.After(time.Second):
		t.Fatal("NewContainer() did not reach DefaultContainer read lock attempt")
	}

	select {
	case <-done:
		t.Fatal("NewContainer() returned while DefaultContainer lock was held, want locked snapshot")
	default:
	}

	defaults.mu.Unlock()
	locked = false

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("NewContainer() did not return after DefaultContainer lock was released")
	}
}
