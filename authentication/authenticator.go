package authentication

type Authenticator interface {
    IsAuthenticated(username string) (bool)
}

type CurrentUserAuthenticator interface {
    IsAuthenticated() (bool)
}
