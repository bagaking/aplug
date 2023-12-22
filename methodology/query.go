package methodology

import (
	_ "embed"
)

// List 返回所有的技巧名称和适用场景，操作步骤和例子
func (table *Methodologies) List() []*Methodology {
	table.mu.RLock()
	defer table.mu.RUnlock()

	ret := make([]*Methodology, 0, len(table.data))
	for _, value := range table.data {
		// 深拷贝
		details := &Methodology{
			ID:       value.ID,
			Usage:    value.Usage,
			Scenario: append([]string{}, value.Scenario...),
		}
		ret = append(ret, details)
	}
	return ret
}

// TryGet 查询具体某一个技巧的操作步骤和例子
func (table *Methodologies) TryGet(key string) *Methodology {
	table.mu.RLock()
	defer table.mu.RUnlock()

	if value, ok := (table.data)[key]; ok {
		// 深拷贝
		details := &Methodology{
			ID:       value.ID,
			Usage:    value.Usage,
			MainIdea: value.MainIdea,
			Scenario: append([]string{}, value.Scenario...),
			Strategy: append([]string{}, value.Strategy...),
			Steps:    append([]string{}, value.Steps...),
			Examples: append([]string{}, value.Examples...),
		}
		return details
	}
	return nil
}

// MGet 批量获取技巧信息
func (table *Methodologies) MGet(keys ...string) []*Methodology {
	table.mu.RLock()
	defer table.mu.RUnlock()

	ret := make([]*Methodology, 0, len(keys))
	for _, key := range keys {
		details := table.TryGet(key)
		if details != nil {
			ret = append(ret, details)
		}
	}
	return ret
}

// GetByScene 根据场景获取技巧
func (table *Methodologies) GetByScene(scene string) []*Methodology {
	table.mu.RLock()
	defer table.mu.RUnlock()

	ret := make([]*Methodology, 0, len(table.data))
	// 遍历所有的技巧
	for key, value := range table.data {
		if value.hasScenario(scene) {
			details := table.TryGet(key)
			if details != nil {
				ret = append(ret, details)
			}
		}
	}
	return ret
}
