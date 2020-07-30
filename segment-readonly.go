// +build !windows

package main

import (
	pwl "github.com/shubh-shah/powerline-go/powerline"
	"golang.org/x/sys/unix"
)

func segmentPerms(p *powerline) []pwl.Segment {
	cwd := p.cwd
	if unix.Access(cwd, unix.W_OK) != nil {
		return []pwl.Segment{{
			Name:       "perms",
			Content:    p.symbolTemplates.Lock,
			Foreground: p.theme.ReadonlyFg,
			Background: p.theme.ReadonlyBg,
		}}
	}
	return []pwl.Segment{}
}
