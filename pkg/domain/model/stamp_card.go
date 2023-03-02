package model

import (
	"time"

	"github.com/google/uuid"
)

type StampCard struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
	Spots     []*Spot
}
