package measurement

type PathValue struct {
	Path    string            `json:"path,omitempty"`    // location to write data
	Value   interface{}       `json:"value,omitempty"`   // value to write
	Headers map[string]string `json:"headers,omitempty"` // optional header values to write
}

type ActionCommand struct {
	ID      string      `json:"id,omitempty"`      // Identify the command (optional)
	Write   []PathValue `json:"write,omitempty"`   // Write values
	Comment string      `json:"comment,omitempty"` // Write all values or
	From    string      `json:"from,omitempty"`    // optional say where the command was listed from
}
