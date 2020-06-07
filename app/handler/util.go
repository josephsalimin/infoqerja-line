package handler

import (
	"regexp"
	"strings"
)

// IsValidCommand : Function to check wether user inputs is a command or not
func IsValidCommand(message string) bool {
	re := regexp.MustCompile("^!")
	return re.FindString(message) != ""
}

// GetCommand : get the type of command from user inputs
func GetCommand(command string) string {
	co := strings.TrimSpace(command)
	return co[1:]
}
