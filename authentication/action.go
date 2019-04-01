package authentication

import "fmt"

var AuthenticateCurrentUserAction = func (serviceName string) func() {
    return func() {
        pamAuthenticator := &PamAuthenticator{serviceName}
        userProvider := &CurrentUserProvider{}
        authenticator := &OsCurrentUserAuthenticator{pamAuthenticator, userProvider}
        if (authenticator.IsAuthenticated()) {
            fmt.Println("Authentication success.")
        } else {
            fmt.Println("Authentication failed.")
        }
    }
}
