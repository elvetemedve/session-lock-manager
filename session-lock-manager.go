package main

import (
    "github.com/elvetemedve/session-lock-manager/device"
    "os"
    "github.com/jochenvg/go-udev"
)

func main() {
    scanner := &device.UdevScanner{&udev.Udev{}}
    _, done := device.Scan(os.Stdout, scanner)
    <-done
}
