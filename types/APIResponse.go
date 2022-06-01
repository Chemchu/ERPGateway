package types

type APIResponse struct {
	Message    *string `json:"message,omitempty"`
	Data       *string `json:"data,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
}
