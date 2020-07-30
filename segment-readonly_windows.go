// +build windows

package main

import (
	"os"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentPerms(p *powerline) []pwl.Segment {
	cwd := p.cwd
	const W_USR = 0002
	// Check user's permissions on directory in a portable but probably slower way
	fileInfo, _ := os.Stat(cwd)
	if fileInfo.Mode()&W_USR != W_USR {
		return []pwl.Segment{{
			Name:       "perms",
			Content:    p.symbolTemplates.Lock,
			Foreground: p.theme.ReadonlyFg,
			Background: p.theme.ReadonlyBg,
		}}
	}
	return []pwl.Segment{}
}
