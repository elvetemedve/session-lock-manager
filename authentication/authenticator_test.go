package authentication

type FakeAuthenticator struct {
    acceptedUsernames []string
}

func (authenticator *FakeAuthenticator) IsAuthenticated(username string) (bool) {
    for _, acceptedUsername := range authenticator.acceptedUsernames {
        if acceptedUsername == username {
            return true
        }
    }

    return false
}
