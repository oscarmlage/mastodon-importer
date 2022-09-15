package models

import (
	"gorm.io/gorm"
	"time"
)

// statuses (id, uri, text, created_at, updated_at, sensitive, visibility,
// spoiler_text, local, account_id, application_id)

type Status struct {
	gorm.Model
	Id             uint
	Uri            string
	Text           string
	Created_at     time.Time
	Updated_at     time.Time
	Sensitive      string
	Visibility     uint
	Spoiler_text   string
	Local          string
	Account_id     uint
	Application_id uint
}

// tags (name, created_at, updated_at)

type Tag struct {
	Name       string
	Created_at time.Time
	Updated_at time.Time
}

// statuses_tags (status_id, tag_id)

type StatusTag struct {
	Status_id uint
	Tag_id    uint
}
