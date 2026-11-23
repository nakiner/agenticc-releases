package task

import (
	"strconv"
)

func ConfigPort(text string, fallback int) (int, error) {
	if text == "" {
		return fallback, nil
	}
	return strconv.Atoi(text)
}
