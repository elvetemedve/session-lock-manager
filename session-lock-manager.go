package main

import (
    "github.com/elvetemedve/session-lock-manager/device"
    "github.com/elvetemedve/session-lock-manager/authentication"
    "github.com/jochenvg/go-udev"
    "os"
)

func main() {
    scanner := &device.UdevScanner{&udev.Udev{}}
    presence := &device.Presence{
        authentication.AuthenticateCurrentUserAction("session-locker"),
        func(){
        }}
    _, done := presence.Scan(os.Stdout, scanner)
    <-done
}
