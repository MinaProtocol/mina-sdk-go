package mina

import (
	"fmt"
	"strings"
)

// GraphQLError is returned when the GraphQL endpoint returns an error response.
type GraphQLError struct {
	Errors    []GraphQLErrorEntry
	QueryName string
}

// GraphQLErrorEntry represents a single error from a GraphQL response.
type GraphQLErrorEntry struct {
	Message string `json:"message"`
}

func (e *GraphQLError) Error() string {
	messages := make([]string, len(e.Errors))
	for i, entry := range e.Errors {
		messages[i] = entry.Message
	}
	return fmt.Sprintf("GraphQL error in %s: %s", e.QueryName, strings.Join(messages, "; "))
}

// ConnectionError is returned when the client cannot connect to the daemon after retries.
type ConnectionError struct {
	QueryName string
	Retries   int
	LastError error
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("failed to execute %s after %d attempts: %v", e.QueryName, e.Retries, e.LastError)
}

func (e *ConnectionError) Unwrap() error {
	return e.LastError
}

// AccountNotFoundError is returned when the requested account does not exist.
type AccountNotFoundError struct {
	PublicKey string
}

func (e *AccountNotFoundError) Error() string {
	return fmt.Sprintf("account not found: %s", e.PublicKey)
}
