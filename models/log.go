package models

// Log - structure for saving and processing
// logs from the monitors
type Log struct {
	Model

	ResponseTime float64 `json:"response_time"`
	IsSuccess    bool    `json:"is_success"`
	ErrorReason  string  `json:"error_reason"`
	ErrorMessage string  `json:"error_message"`
	Response     string  `json:"response"`
}
