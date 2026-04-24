package service

import (
	"fmt"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(str string) (res string, err error) {
	if isMorseCode(str) {
		result := morse.ToText(str)
		if result == "" {
			return "", fmt.Errorf("failed to decode morse: %q", str)
		}
		return result, nil
	} else {
		result := morse.ToMorse(str)
		if result == "" {
			return "", fmt.Errorf("failed to encode to morse: %q", str)
		}
		return result, nil
	}
}

func isMorseCode(s string) bool {
	if s == "" {
		return false
	}

	hasDotOrDash := false
	for _, r := range s {
		switch r {
		case ' ', '\t', '\n', '.', '-':
			hasDotOrDash = hasDotOrDash || r == '.' || r == '-'
		default:
			return false
		}
	}
	return hasDotOrDash
}
