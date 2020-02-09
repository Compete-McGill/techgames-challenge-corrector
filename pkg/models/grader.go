package models

// LivenessResponse is the format of a response from /api/status
type LivenessResponse struct {
	Status string `json:"status"`
}
