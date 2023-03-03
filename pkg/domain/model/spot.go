package model

import (
	"time"

	"github.com/google/uuid"
)

type Spot struct {
	ID        uuid.UUID
	Name      string
	Detail    string
	Like      uint
	ImgURL    string
	Address   string
	CreatedAt time.Time
}
