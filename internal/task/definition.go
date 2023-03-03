package task

import (
	"time"

	"github.com/google/uuid"
)

// All the base properties for all kinds of task types.
type DefinitionProperties struct {
	// Unique identifier for for a defined task
	ID uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	// The type of the defined task, which determines
	Type TaskType `sql:"type:ENUM('JS', 'HTTP_REQ')" gorm:"column:task_type"`
	// Name of the definition is required
	Name string
	// Optional description
	Description *string
	Created     time.Time `gorm:"default:NOW()"`
	Modified    time.Time `gorm:"default:NOW()"`
}

// All tasks must implement this interface for them to be included in a workflow.
type Definition interface {
	// Get a function to run the task
	Get() func(data *DataStore) error
}
