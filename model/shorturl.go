package model

import (
	"errors"
	"math/rand"
	"net/url"
	"regexp"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	keyval      = "^[a-zA-Z0-9_-]*$"
)

// ShortLink is the main model that contain
// the source url to redirect to
type ShortLink struct {
	ID               uint   `gorm:"primaryKey"`
	SourceKey        string `json:"source_key" gorm:"unique"`
	DestinationValue string `json:"destination_value"`

	CreateAt  time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (sl *ShortLink) Validate() error {

	if sl.SourceKey != "" {
		re := regexp.MustCompile(keyval)
		if !re.MatchString(sl.SourceKey) {
			return errors.New("invalid source string key")
		}
	} else {
		sl.GenerateSourceKey(8)
	}

	if sl.DestinationValue == "" {
		return errors.New("No destination url found")
	}

	u, err := url.ParseRequestURI(sl.DestinationValue)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("Invalid destination url")
	}

	return nil
}

func (sl *ShortLink) GenerateSourceKey(n int) {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	sl.SourceKey = string(b)
}

func (sl *ShortLink) ToPublic() map[string]interface{} {
	return map[string]interface{}{
		"id":                sl.ID,
		"source key":        sl.SourceKey,
		"destination value": sl.DestinationValue,
	}
}
