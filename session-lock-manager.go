package main

import (
    "github.com/elvetemedve/session-lock-manager/device"
    "github.com/elvetemedve/session-lock-manager/authentication"
    "github.com/elvetemedve/session-lock-manager/session"
    "github.com/jochenvg/go-udev"
    "os"
    "fmt"
)

func main() {
    serviceName := parseArguments()
    scanner := &device.UdevScanner{&udev.Udev{}}
    presence := &device.Presence{
        authentication.AuthenticateCurrentUserAction(serviceName, session.Unlock, func(){
                fmt.Fprintln(os.Stderr, "Authentication failed.")
            }),
        session.Lock}
    _, done := presence.Scan(os.Stdout, scanner)
    <-done
}

func parseArguments() (string) {
    if len(os.Args) != 2 {
        fmt.Fprintln(os.Stderr, "Insufficient arguments given.")
        os.Exit(1)
    }

    return os.Args[1]
}
