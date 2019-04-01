package authentication

import (
     "github.com/msteinert/pam"
     "errors"
     "io"
     "fmt"
     "os"
)

type SilentConversationHandler struct {
    stdout io.Writer
    stderr io.Writer
}

func (handler *SilentConversationHandler) RespondPAM(style pam.Style, message string) (string, error) {
    switch style {
	case pam.ErrorMsg:
        fmt.Fprintln(handler.stderr, message)
		return "", nil
    case pam.TextInfo:
		fmt.Fprintln(handler.stdout, message)
		return "", nil
	}
	return "", errors.New("PAM interactive authentication is not supported by this application.")
}

type PamAuthenticator struct {
    ServiceName string
}

func (authenticator *PamAuthenticator) IsAuthenticated(username string) (bool) {
    transaction, error := pam.Start(authenticator.ServiceName, username, &SilentConversationHandler{os.Stdout, os.Stderr})

    if error != nil {
        fmt.Fprintln(os.Stderr, "Cannot communicate with PAM.")
        os.Exit(3)
    }

    return nil == transaction.Authenticate(pam.Silent)
}
