package parameter

import (
	"encoding/json"
	"time"
)

type Parameter struct {
	ID        uint            `json:"id"`
	Name      string          `json:"name"`
	Value     json.RawMessage `json:"value"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

type Parameters []*Parameter
