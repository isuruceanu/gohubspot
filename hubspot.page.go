package gohubspot

type Page struct {
	Offset  int  `json:"offset"`
	HasMore bool `json:"has-more"`
}
