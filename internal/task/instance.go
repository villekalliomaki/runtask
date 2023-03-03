package task

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Instance struct {
	ID      uuid.UUID `gorm:"primaryKey;default:uuid_generate_v4()"`
	Created time.Time `gorm:"default:NOW()"`
	Started sql.NullTime
}
