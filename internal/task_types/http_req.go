package tasktypes

import (
	"database/sql"
	"runtask/internal/task"
)

// Simple HTTP request to a target. Results can be saved as JSON
// or the respose body to a key to the DataStore.
type HttpReq struct {
	task.DefinitionProperties
	// The target URL
	Target string
	// One of the HTTP methods
	Method     string
	JsonOut    bool
	BodyOutKey sql.NullString
	Headers    []string
}
