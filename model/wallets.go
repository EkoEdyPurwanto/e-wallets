package model

import (
	"github.com/google/uuid"
	"time"
)

type Wallets struct {
	Id           uuid.UUID
	UserId       uuid.UUID
	RekeningUser string
	Balance      int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
