package gohubspot

type Property struct {
	Property string      `json:"property"`
	Value    interface{} `json:"value"`
	Versions Versions    `json:"versions,omitempty"`
}

type Properties struct {
	Properties []Property `json:"properties"`
}

// AddProperty addes a new property to list
func (p *Properties) AddProperty(prop string, value interface{}) {
	p.Properties = append(p.Properties, Property{Property: prop, Value: value})
}
