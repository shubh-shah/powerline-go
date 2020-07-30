package main

import (
	"os"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentUser(p *powerline) []pwl.Segment {
	var userPrompt string
	if *p.args.Shell == "bash" {
		userPrompt = "\\u"
	} else if *p.args.Shell == "zsh" {
		userPrompt = "%n"
	} else {
		userPrompt = p.username
	}

	var background uint8
	if os.Getuid() == 0 {
		background = p.theme.UsernameRootBg
	} else {
		background = p.theme.UsernameBg
	}

	return []pwl.Segment{{
		Name:       "user",
		Content:    "\uf31c " + userPrompt,
		Foreground: p.theme.UsernameFg,
		Background: background,
	}}
}
