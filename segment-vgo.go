package main

import (
	"os"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentVirtualGo(p *powerline) []pwl.Segment {
	var env string
	if env == "" {
		env, _ = os.LookupEnv("VIRTUALGO")
	}
	segments := []pwl.Segment{}
	if env != "" {
		segments = append(segments, pwl.Segment{
			Name:       "vgo",
			Content:    env,
			Foreground: p.theme.VirtualGoFg,
			Background: p.theme.VirtualGoBg,
		})
	}
	return segments
}
