package models

// ErrorResponse is a custom made error response struct to be used in the whole app
type ErrorResponse struct {
	ApplicationMessage string
	UserMessage        string
}
