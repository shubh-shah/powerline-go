package main

import (
	"os"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentSSH(p *powerline) []pwl.Segment {
	sshClient, _ := os.LookupEnv("SSH_CLIENT")
	var networkIcon string
	if *p.args.SshAlternateIcon {
		networkIcon = p.symbolTemplates.NetworkAlternate
	} else {
		networkIcon = p.symbolTemplates.Network
	}

	if sshClient != "" {
		return []pwl.Segment{{
			Name:       "ssh",
			Content:    networkIcon,
			Foreground: p.theme.SSHFg,
			Background: p.theme.SSHBg,
		}}
	}
	return []pwl.Segment{}
}
