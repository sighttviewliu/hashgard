package params

// Param for record
type RecordParams struct {
	Name               string  `json:"name"`
	Author             string  `json:"author"`
	Type               string  `json:"type"`
	Hash               string  `json:"hash"`
	RecordNo           string  `json:"record_number"`
	Description        string  `json:"description"`
}
