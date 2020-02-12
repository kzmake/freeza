package freeza

// Request ...
type Request struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Body    string            `json:"body,omitempty"`
	Headers map[string]string `json:"header,omitempty"`
}
