package device

import (
    "testing"
    "strings"
    "bytes"
)

func TestMonitoringIsStarted(t *testing.T) {
    buffer := &bytes.Buffer{}
    scanner := &FakeScanner{}

    cancel, done := Scan(buffer, scanner)
    cancel()
    <-done

    got := buffer.String()
    want := "Started listening on channel."

    if !strings.Contains(got, want) {
        t.Errorf("The output does not contain the text: '%s' Got: '%s'", want, got)
    }
}

// func TestDeviceInsertIsDetected(t *testing.T) {
// TODO implement testing logic.
// }
