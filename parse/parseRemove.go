package parse

import (
	"strings"

	"github.com/docker/engine-api/types"
)

// RemoveAction takes the argument string that was passed with the --remove flag,
// parses it, and updates the Seccomp config accordingly
func RemoveAction(arguments string, config *types.Seccomp) {
	var syscallsToRemove []string
	if strings.Contains(arguments, ",") {
		syscallsToRemove = strings.Split(arguments, ",")
	} else {
		syscallsToRemove = append(syscallsToRemove, arguments)
	}

	for _, syscall := range syscallsToRemove {
		for counter, syscallStruct := range config.Syscalls {
			if syscallStruct.Name == syscall {
				config.Syscalls = append(config.Syscalls[:counter], config.Syscalls[counter+1:]...)
			}
		}
	}
}
