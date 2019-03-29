package device

import (
    "fmt"
    "context"
    "io"
)

const yubikeyDeviceId = "1050/407/511";

func Scan(out io.Writer, scanner Scanner) (context.CancelFunc, chan bool) {
    eventChannel, cancel := scanner.Scan()
    done := make(chan bool)
    go listen(out, eventChannel, done)

    return cancel, done
}

func listen(out io.Writer, eventChannel <-chan *SecurityTokenEvent, done chan<- bool) {
    fmt.Fprintln(out, "Started listening on channel.")

    // Enumerate device events
    for event := range eventChannel {
        if (yubikeyDeviceId != event.deviceId) {
            continue
        }

        switch event.action {
        case insert:
            fmt.Println("USB device inserted.")
        case eject:
            fmt.Println("USB device ejected.")
        }
    }

    fmt.Fprintln(out, "Channel closed.")
    done <- true
}