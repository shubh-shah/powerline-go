package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	pwl "github.com/shubh-shah/powerline-go/powerline"
)

func segmentJobs(p *powerline) []pwl.Segment {
	nJobs := -1

	ppid := os.Getppid()
	if *p.args.Shell == "bash" {
		pppidOut, _ := exec.Command("ps", "-p", strconv.Itoa(ppid), "-oppid=").Output()
		pppid, _ := strconv.ParseInt(strings.TrimSpace(string(pppidOut)), 10, 64)
		ppid = int(pppid)
	}

	out, _ := exec.Command("ps", "-oppid=").Output()
	processes := strings.Split(string(out), "\n")
	for _, processPpidStr := range processes {
		processPpid, _ := strconv.ParseInt(strings.TrimSpace(processPpidStr), 10, 64)
		if int(processPpid) == ppid {
			nJobs++
		}
	}

	if nJobs > 0 {
		return []pwl.Segment{{
			Name:       "jobs",
			Content:    fmt.Sprintf("%d", nJobs),
			Foreground: p.theme.JobsFg,
			Background: p.theme.JobsBg,
		}}
	}
	return []pwl.Segment{}
}
