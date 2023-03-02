package model

import (
	"time"

	"github.com/google/uuid"
)

type StampLog struct {
	ID        uuid.UUID
	Spot      *Spot
	User      *User
	StampCard *StampCard
	CreatedAt time.Time
}
