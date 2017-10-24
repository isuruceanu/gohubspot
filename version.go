package gohubspot

type Version struct {
	Value       interface{} `json:"value"`
	SourceType  string      `json:"source-type"`
	SourceID    string      `json:"source-id"`
	SourceLabel string      `json:"source-label"`
	Timestamp   UnixTime    `json:"timestamp"`
	Selected    bool        `json:"selected"`
}

type Versions []Version
