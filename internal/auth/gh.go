package auth

import (
	"os/exec"
	"strings"
)

type GHStatus struct {
	Installed     bool
	Authenticated bool
	Username      string
}

func CheckGH() GHStatus {
	s := GHStatus{}

	if _, err := exec.LookPath("gh"); err != nil {
		return s
	}
	s.Installed = true

	out, err := exec.Command("gh", "auth", "status").CombinedOutput()
	if err != nil {
		return s
	}
	s.Authenticated = true

	for _, line := range strings.Split(string(out), "\n") {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "Logged in to github.com account") {
			fields := strings.Fields(line)
			for i, f := range fields {
				if f == "account" && i+1 < len(fields) {
					s.Username = fields[i+1]
					break
				}
			}
		}
	}

	return s
}
