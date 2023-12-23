package methodology

import "sync"

type (
	Methodologies struct {
		data map[string]*Methodology
		mu   sync.RWMutex
	}
)

var (
	onceLoadDefault  sync.Once
	defaultContainer *Methodologies
)

func NewContainer() *Methodologies {
	container := &Methodologies{
		data: make(map[string]*Methodology),
		mu:   sync.RWMutex{},
	}

	defaults := DefaultContainer()
	defaults.mu.RLock()
	defer defaults.mu.RUnlock()
	for key, value := range defaults.data {
		container.data[key] = cloneMethodology(value)
	}

	return container
}

func DefaultContainer() *Methodologies {
	onceLoadDefault.Do(func() {
		defaultContainer = &Methodologies{
			data: mustLoadDefaultConf(),
			mu:   sync.RWMutex{},
		}
	})
	return defaultContainer
}
