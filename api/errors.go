package api

import "fmt"

// MissingParam missing param error
type MissingParam string

// Error get missing param error info data, format: Missing [<data>]
func (e MissingParam) Error() string {
	return fmt.Sprintf("Missing [%s]", string(e))
}

// BadParam bad param error
type BadParam string

// Error get bad param error info data, format: BadParam [<data>]
func (e BadParam) Error() string {
	return fmt.Sprintf("BadParam [%s]", string(e))
}

// NotFound not found error
type NotFound string

// Error get not found error info data, format: <data> not found
func (e NotFound) Error() string {
	return fmt.Sprintf("%s not found", string(e))
}

// Timeout timeout error
type Timeout struct{}

// Error get timeout error info data, format: timeout
func (e Timeout) Error() string {
	return "timeout"
}
