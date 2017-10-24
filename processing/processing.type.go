package processing

type ProcessingType string

const (
	DONE         ProcessingType = "DONE"
	REFRESHING   ProcessingType = "REFRESHING"
	INITIALIZING ProcessingType = "INITIALIZING"
	PROCESSING   ProcessingType = "PROCESSING"
)
