package gohubspot

type Vid struct {
	VID   int  `json:"vid"`
	IsNew bool `json:"isNew"`
}

type ContactDeleteResult struct {
	VID     int    `json:"vid"`
	Deleted bool   `json:"deleted"`
	Reason  string `json:"reason"`
}
