package main

import (
	"fmt"
	"github.com/elvetemedve/session-lock-manager/authentication"
	"github.com/elvetemedve/session-lock-manager/device"
	"github.com/elvetemedve/session-lock-manager/session"
	"github.com/jochenvg/go-udev"
	"os"
)

func main() {
	serviceName := parseArguments()
	scanner := &device.UdevScanner{Udev: &udev.Udev{}}
	presence := &device.Presence{
		OnInsert: authentication.AuthenticateCurrentUserAction(serviceName, session.Unlock, func() {
			fmt.Fprintln(os.Stderr, "Authentication failed.")
		}),
		OnEject: session.Lock,
	}
	_, done := presence.Scan(os.Stdout, scanner)
	<-done
}

func parseArguments() string {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Insufficient arguments given.")
		os.Exit(1)
	}

	return os.Args[1]
}
