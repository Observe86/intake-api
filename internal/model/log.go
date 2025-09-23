package model

type Log struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	Level     string `json:"level,omitempty"`
}
