package gohubspot

type Property struct {
	Property string      `json:"property"`
	Value    interface{} `json:"value"`
	Versions Versions    `json:"versions,omitempty"`
}

type Properties []Property
