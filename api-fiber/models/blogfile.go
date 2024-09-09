package models

import (
	"time"
)

type BlogFile struct {
	Id           int      `json:"id"`
	Filename     string   `json:"filename"`
	CreatedDate time.Time `json:"created_date"`
	UpdatedDate time.Time `json:"updated_date"`
}
