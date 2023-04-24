package entities

import (
	"time"

	amqp "github.com/kaellybot/kaelly-amqp"
)

type TwitterAccount struct {
	Id         string        `gorm:"unique"`
	Locale     amqp.Language `gorm:"primaryKey"`
	LastUpdate time.Time     `gorm:"not null; default:current_timestamp"`
}
