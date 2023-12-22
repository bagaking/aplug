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
	for key, value := range DefaultContainer().data {
		container.data[key] = value
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
