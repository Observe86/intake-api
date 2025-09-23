package model

type Trace struct {
	TraceID  string `json:"trace_id"`
	Span     string `json:"span"`
	Duration int    `json:"duration_ms"`
}
