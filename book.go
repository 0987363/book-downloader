package main

import (
	"regexp"
)

type Rule struct {
	Site string				`json:"site"`
	DirectoryBegin string	`json:"directory_begin"`
	DirectoryEnd string		`json:"directory_end"`
	TextBegin string		`json:"text_begin"`
	TextEnd string			`json:"text_end"`
	Class []string			`json:"class"`
}

func FindRule(rules []*Rule, url string) *Rule {
	for _, rule := range rules {
		var rr = regexp.MustCompile(url)
		if rr.MatchString(url) {
			return rule
		}
	}

	return nil
}

