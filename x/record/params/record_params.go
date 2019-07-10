package params

// Param for record
type RecordParams struct {
	Name               string  `json:"name"`
	Author             string  `json:"author"`
	Hash               string  `json:"hash"`
	RecordNo           string  `json:"record_number"`
	RecordType         string  `json:"record_type"`
	Description        string  `json:"description"`
}
