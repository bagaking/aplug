package methodology

import (
	"fmt"

	"github.com/khicago/got/util/delegate"
)

// Set 更新或新增一个技巧
func (table *Methodologies) Set(key string, m *Methodology) {
	table.mu.Lock()
	defer table.mu.Unlock()

	table.data[key] = m
}

// Delete 删除一个技巧
func (table *Methodologies) Delete(key string) {
	table.mu.Lock()
	defer table.mu.Unlock()

	delete(table.data, key)
}

// UpdatedBy 更新一个技巧
func (table *Methodologies) UpdatedBy(
	key string,
	fn delegate.Convert[*Methodology, *Methodology],
) error {
	table.mu.Lock()
	defer table.mu.Unlock()

	method, ok := table.data[key]
	if !ok {
		return fmt.Errorf("methodology not found: %v", key)
	}
	table.data[key] = fn(method)
	return nil
}
