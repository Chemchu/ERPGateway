package types

type APIResponse struct {
	Data *string `json:"data,omitempty"`
}

type APIData struct {
	Message    *string `json:"message,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
}
