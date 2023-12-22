package methodology

// Methodology - Detailed information of a methodology.
type Methodology struct {
	// The ID of this methodology, that is, the name.
	ID string `json:"id"`

	// The Usage describes what problem this methodology solves, and roughly what situations it is used for.
	// You should decide whether to use this methodology to solve your problem based on the usage.
	Usage string `json:"usage"`

	// MainIdea describes the methodology's main idea and how it works.
	MainIdea string `json:"main_idea,omitempty"`

	// Scenario contains Some highly generalized tags that roughly describe the applicable scenarios
	// or contexts of a methodology.
	Scenario []string `json:"scenario,omitempty"`

	// Strategy describes the methodology's thinking strategy.
	Strategy []string `json:"strategy,omitempty"`

	// Specific execution Steps listed in order.
	Steps []string `json:"steps,omitempty"` // 具体的操作步骤

	// Examples of response situations and formats.
	Examples []string `json:"examples,omitempty"` // 举例
}

func (m Methodology) hasScenario(scenario string) bool {
	for _, s := range m.Scenario {
		if s == scenario {
			return true
		}
	}
	return false
}
