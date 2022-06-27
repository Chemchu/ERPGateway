package types

type APIData struct {
	Data *string `json:"data,omitempty"`
}

type APIResponse struct {
	Message    *string `json:"message,omitempty"`
	Successful *bool   `json:"successful,omitempty"`
}

type GraphQLQuery struct {
	Query     string `json:"query,omitempty"`
	Variables string `json:"variables,omitempty"`
}
