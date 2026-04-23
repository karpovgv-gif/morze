package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convert(str string) (res string, err error) {
	hasCode := strings.ContainsAny(str, ".-")
	hasCirillica := strings.ContainsAny(str, "АБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯабвгдеёжзийклмнопрстуфхцчшщъыьэюя")

	if hasCode && !hasCirillica {
		res = morse.ToText(str)
		if res == "" {
			return "", fmt.Errorf("Неправильное декодирование кода морзе: %q\n", str)
		}
		return res, nil
	} else {
		res = morse.ToMorse(str)
		if res == "" {
			return "", fmt.Errorf("Неправильное декодирование текста: %q\n", str)
		}
		return res, nil
	}
}
