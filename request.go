package namec

// Request ...
type Request struct {
	Method string              `json:"method"`
	URL    string              `json:"url"`
	Body   string              `json:"body,omitempty"`
	Header map[string][]string `json:"header,omitempty"`
}
