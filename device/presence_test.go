package device

import (
    "testing"
    "strings"
    "bytes"
    "time"
)

func TestMonitoringIsStarted(t *testing.T) {
    presence := &Presence{}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}

    _, done := presence.Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "Started listening on channel."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}

func TestMonitoringCanBeCancelled(t *testing.T) {
    presence := &Presence{}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{}
    stopChannel := make(chan bool)

    go func() {
		time.Sleep(time.Second)
        select {
          default:
            panic("Scan operation has not been cancelled. Dying.")
        case <-stopChannel:
            return
        }
	}()

    cancel, done := presence.Scan(buffer, scanner)
    cancel()
    <-done

    close(stopChannel)
}

func TestDeviceInsertIsDetected(t *testing.T) {
    presence := &Presence{func(){}, func(){}}
    insertEvent := SecurityTokenEvent{insert, "1050/407/511"}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}
    scanner.add(insertEvent)

    _, done := presence.Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "USB device inserted."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}

func TestDeviceEjectIsDetected(t *testing.T) {
    presence := &Presence{func(){}, func(){}}
    ejectEvent := SecurityTokenEvent{eject, "1050/407/511"}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}
    scanner.add(ejectEvent)

    _, done := presence.Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "USB device ejected."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}
