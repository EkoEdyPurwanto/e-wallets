package model

import (
	"github.com/google/uuid"
	"time"
)

type AccountStatus string

const (
	Active   AccountStatus = "aktif"
	Inactive AccountStatus = "nonaktif"
	Blocked  AccountStatus = "diblokir"
	Other    AccountStatus = "Other"
)

type UsersAccount struct {
	Id            uuid.UUID
	UserName      string
	Password      string
	Email         string
	PhoneNumber   string
	AccountStatus AccountStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
