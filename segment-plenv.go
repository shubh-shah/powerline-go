package main

import (
	"os"
	"path"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentPlEnv(p *powerline) []pwl.Segment {
	var env string
	if env == "" {
		env, _ = os.LookupEnv("PLENV_VERSION")
	}
	if env != "" {
		envName := path.Base(env)
		return []pwl.Segment{{
			Name:       "plenv",
			Content:    envName,
			Foreground: p.theme.PlEnvFg,
			Background: p.theme.PlEnvBg,
		}}
	}
	return []pwl.Segment{}
}
