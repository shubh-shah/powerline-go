package main

import (
	"time"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentTime(p *powerline) []pwl.Segment {
	return []pwl.Segment{{
		Name:       "time",
		Content:    time.Now().Format("15:04:05"),
		Foreground: p.theme.TimeFg,
		Background: p.theme.TimeBg,
	}}
}
