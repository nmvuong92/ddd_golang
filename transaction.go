package tavern

import (
	"github.com/google/uuid"
	"time"
)

// Transaction is a value object because it no identifier and is un-mutable
type Transaction struct {
	amount   int
	from     uuid.UUID //person
	to       uuid.UUID //person
	createAt time.Time
}
