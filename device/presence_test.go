package device

import (
    "testing"
    "strings"
    "bytes"
    "time"
)

func TestMonitoringIsStarted(t *testing.T) {
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}

    _, done := Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "Started listening on channel."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}

func TestMonitoringCanBeCancelled(t *testing.T) {
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

    cancel, done := Scan(buffer, scanner)
    cancel()
    <-done

    close(stopChannel)
}

func TestDeviceInsertIsDetected(t *testing.T) {
    insertEvent := SecurityTokenEvent{insert, "1050/407/511"}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}
    scanner.add(insertEvent)

    _, done := Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "USB device inserted."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}

func TestDeviceEjectIsDetected(t *testing.T) {
    ejectEvent := SecurityTokenEvent{eject, "1050/407/511"}
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{make([]SecurityTokenEvent, 0), true}
    scanner.add(ejectEvent)

    _, done := Scan(buffer, scanner)
    <-done

    got := buffer.String()
    want := "USB device ejected."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}
