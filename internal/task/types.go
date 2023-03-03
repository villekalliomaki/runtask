package task

// Every task type has an unique string identifier.
// The type determines what task type's definition the server will query from the database.
type TaskType string

const (
	JS       TaskType = "js"
	HTTP_REQ TaskType = "http_req"
)
