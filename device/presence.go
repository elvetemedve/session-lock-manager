package device

import (
    "fmt"
    "context"
    "io"
)

type Presence struct {
    OnInsert func()
    OnEject func()
}

func (presence *Presence) Scan(out io.Writer, scanner Scanner) (context.CancelFunc, chan bool) {
    eventChannel, cancel := scanner.Scan()
    done := make(chan bool)
    go presence.listen(out, eventChannel, done)

    return cancel, done
}

func (presence *Presence) listen(out io.Writer, eventChannel <-chan *SecurityTokenEvent, done chan<- bool) {
    fmt.Fprintln(out, "Started listening on channel.")

    // Enumerate device events
    for event := range eventChannel {
        if (!devices.IsSupported(event.deviceId)) {
            continue
        }

        switch event.action {
        case insert:
            fmt.Fprintln(out, "USB device inserted.")
            presence.OnInsert()
        case eject:
            fmt.Fprintln(out, "USB device ejected.")
            presence.OnEject()
        }
    }

    fmt.Fprintln(out, "Channel closed.")
    done <- true
}
