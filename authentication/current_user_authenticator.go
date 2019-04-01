package authentication

type OsCurrentUserAuthenticator struct {
    Authenticator Authenticator
    UserProvider UserProvider
}

func (currentUserAuthenticator *OsCurrentUserAuthenticator) IsAuthenticated() (bool) {
    return currentUserAuthenticator.Authenticator.IsAuthenticated(currentUserAuthenticator.UserProvider.GetCurrentUsername())
}
