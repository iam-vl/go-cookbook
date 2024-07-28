package main

import (
	"regexp"
)

var bookDetailsPattern = regexp.MustCompile(`Title: (.+), Author: (.+), Pages: (\d+)`)
