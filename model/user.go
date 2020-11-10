package model

import (
	"errors"
	"html"
	"regexp"
	"strings"
	"time"
)

const emailregexp = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

type User struct {
	ID     uint   `gorm:"primaryKey"`
	Email  string `gorm:"unique" json:"email"`
	APIKEY string `gorm:"unique"`

	CreateAt  time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (user *User) Validate() error {

	user.Prepare()

	emailregexp := regexp.MustCompile(emailregexp)
	if user.Email == "" {
		return errors.New("Email is required")
	} else if !emailregexp.MatchString(user.Email) {
		return errors.New("Invalid Email Format")
	}
	return nil
}

func (user *User) Prepare() {
	user.Email = strings.ToLower(html.EscapeString(strings.TrimSpace(user.Email)))
}
