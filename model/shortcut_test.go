package model

import (
	"errors"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexKeyval(t *testing.T) {
	samples := []struct {
		title    string
		test     string
		expected bool
	}{
		{
			title:    "only lowercase",
			test:     "validkey",
			expected: true,
		},
		{
			title:    "only uppercase",
			test:     "VALIDKEY",
			expected: true,
		},
		{
			title:    "only number",
			test:     "12345678",
			expected: true,
		},
		{
			title:    "alphanumberical",
			test:     "VaL1dK3y",
			expected: true,
		},
		{
			title:    "with underscore",
			test:     "VaL1d_K3y",
			expected: true,
		},
		{
			title:    "with a dash",
			test:     "VaL1d-K3y",
			expected: true,
		},
		{
			title:    "with a space",
			test:     "not valid key",
			expected: false,
		},
		{
			title:    "with a symbol",
			test:     "notval=+)('\"dkey",
			expected: false,
		},
	}

	keyRegex := regexp.MustCompile(keyval)

	for _, sample := range samples {
		t.Run(sample.title, func(t *testing.T) {
			assert.Equal(t, sample.expected, keyRegex.MatchString(sample.test))
		})
	}

}

func TestValidateShortLink(t *testing.T) {
	samples := []struct {
		title    string
		test     ShortLink
		expected error
	}{
		{
			title: "valid shortlink model",
			test: ShortLink{
				SourceKey:        "validkey",
				DestinationValue: "https://google.com",
			},
			expected: nil,
		},
		{
			title: "valid shortlink model without key",
			test: ShortLink{
				DestinationValue: "https://google.com",
			},
			expected: nil,
		},
		{
			title: "invalid shortlink model (without destination value)",
			test: ShortLink{
				SourceKey: "validkey",
			},
			expected: errors.New("No destination url found"),
		},
		{
			title: "invalid shortlink model (invalid destination value)",
			test: ShortLink{
				SourceKey:        "validkey",
				DestinationValue: "google.com",
			},
			expected: errors.New("Invalid destination url"),
		},
	}

	for _, sample := range samples {
		t.Run(sample.title, func(t *testing.T) {
			assert.Equal(t, sample.expected, sample.test.Validate())
		})
	}
}
