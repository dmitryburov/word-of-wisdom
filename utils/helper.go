package utils

import "log"

const (
	ApplicationClientName = "client"
	ApplicationServerName = "server"
)

// FatalApplication global error application
func FatalApplication(msg string, err error) {
	log.Fatalf("%s > %s\n", msg, err)
}
