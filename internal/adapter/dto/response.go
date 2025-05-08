package dto

type Response struct {
	Status    int         `json:"status,omitempty"`
	Message   interface{} `json:"message,omitempty"`
	MessageTh interface{} `json:"message_th,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	Total     int         `json:"total,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
