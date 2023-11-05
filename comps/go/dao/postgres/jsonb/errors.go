package jsonb

import "strings"

// isMissingColumn function to return if the error is related to missing column
func isMissingColumn(err error) bool {
	return strings.Contains(err.Error(), "column \"") &&
		strings.Contains(err.Error(), "\" does not exist")
}

// isAuthenticationError function to return if the error is related Authentication for the database
func isAuthenticationError(err error) bool {
	return strings.Contains(err.Error(), "authentication failed")
}

// isDuplicatedFieldError function to return if the error is related duplicate Field
func isDuplicatedFieldError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key Value violates unique constraint")
}
