package authentication

import (
    "testing"
    "bytes"
    "strings"
    "github.com/msteinert/pam"
)

func TestItDoesNotAllowInteractiveAuthentication(t *testing.T) {
    handler := &SilentConversationHandler{}

    _, error := handler.RespondPAM(pam.PromptEchoOn, "Please enter your password")

    if error.Error() != "PAM interactive authentication is not supported by this application." {
        t.Errorf("Conversation handler should not allow user interaction.")
    }
}

func TestItDoesPrintErrorMessage(t *testing.T) {
    outputBuffer := &bytes.Buffer{}
    errorOutputBuffer := &bytes.Buffer{}
    handler := &SilentConversationHandler{outputBuffer, errorOutputBuffer}

    _, error := handler.RespondPAM(pam.ErrorMsg, "This is an intentional error message.")

    if error != nil {
        t.Errorf("Conversation handler should be able to print error messages.")
    }

    got := errorOutputBuffer.String()
    want := "This is an intentional error message."
    if !strings.Contains(got, want) {
        t.Errorf("An error message should have been printed on the standard error channel.")
    }
}

func TestItDoesPrintInformationalMessage(t *testing.T) {
    outputBuffer := &bytes.Buffer{}
    errorOutputBuffer := &bytes.Buffer{}
    handler := &SilentConversationHandler{outputBuffer, errorOutputBuffer}

    _, error := handler.RespondPAM(pam.TextInfo, "This is an informational message.")

    if error != nil {
        t.Errorf("Conversation handler should be able to print informational messages.")
    }

    got := outputBuffer.String()
    want := "This is an informational message."
    if !strings.Contains(got, want) {
        t.Errorf("An informational message should have been printed on the standard output channel.")
    }
}
