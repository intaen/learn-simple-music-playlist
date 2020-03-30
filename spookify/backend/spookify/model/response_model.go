package model

// Resp represent standard response message
type Resp struct {
	Status string
	Message string
	Data    interface{} `json:"data,omitempty"`
}
