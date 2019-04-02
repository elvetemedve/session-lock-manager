package authentication

var AuthenticateCurrentUserAction = func (serviceName string, success func(), failure func()) func() {
    return func() {
        pamAuthenticator := &PamAuthenticator{serviceName}
        userProvider := &CurrentUserProvider{}
        authenticator := &OsCurrentUserAuthenticator{pamAuthenticator, userProvider}
        if (authenticator.IsAuthenticated()) {
            success()
        } else {
            failure()
        }
    }
}
