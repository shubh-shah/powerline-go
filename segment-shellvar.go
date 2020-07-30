package main

import (
	"os"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentShellVar(p *powerline) []pwl.Segment {
	shellVarName := *p.args.ShellVar
	varContent, varExists := os.LookupEnv(shellVarName)

	if varExists {
		if varContent != "" {
			return []pwl.Segment{{
				Name:       "shell-var",
				Content:    varContent,
				Foreground: p.theme.ShellVarFg,
				Background: p.theme.ShellVarBg,
			}}
		}
		warn("Shell variable " + shellVarName + " is empty.")

	} else {
		warn("Shell variable " + shellVarName + " does not exist.")
	}
	return []pwl.Segment{}
}
