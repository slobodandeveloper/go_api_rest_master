package click

import "gitlab.com/menuxd/api-rest/pkg/model"

// Click is a count to ad's clicks.
type Click struct {
	model.Model
	Type   uint `json:"type"`
	TypeID uint `json:"type_id"`
}

// Click type.
const (
	Ad uint = iota
	Promotion
	Suggested
)

// Clicks is an alias to a slice of clicks.
type Clicks []Click
