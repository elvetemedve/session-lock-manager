package authentication

import (
    "testing"
)

func TestItAuthenticatesTheCurrentUser(t *testing.T) {
    currentUsername := "geza"
    fakeUserProvider := &FakeUserProvider{currentUsername}
    fakeAuthenticator := &FakeAuthenticator{[]string{currentUsername}}
    authenticator := &OsCurrentUserAuthenticator{fakeAuthenticator, fakeUserProvider}

    if !authenticator.IsAuthenticated() {
        t.Errorf("The current user is %s and should be authenticated.", currentUsername)
    }

    currentUsername = "alice"
    fakeUserProvider.setCurrentUsername(currentUsername)

    if authenticator.IsAuthenticated() {
        t.Errorf("The current user is %s and should not be authenticated.", currentUsername)
    }
}
