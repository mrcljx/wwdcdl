package main

import (
  "path"
  "regexp"
  "strings"
)

// Inspired by github.com/asaskevich/govalidator
func SafeFileName(name string) string {
	name = path.Clean(name)
	name = strings.Trim(name, " ")
	separators, err := regexp.Compile(`[&=+:/]`)
	if err == nil {
		name = separators.ReplaceAllString(name, "-")
	}
	legal, err := regexp.Compile(`[^[:alnum:]\s-.()]`)
	if err == nil {
		name = legal.ReplaceAllString(name, "")
	}
	for strings.Contains(name, "--") {
		name = strings.Replace(name, "--", "-", -1)
	}
	return name
}
