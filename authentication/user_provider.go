package authentication

import (
    "log"
    "os"
    "os/user"
)

type UserProvider interface {
    GetCurrentUsername() (string)
}

type CurrentUserProvider struct {}

func (provider *CurrentUserProvider) GetCurrentUsername() (string) {
    currentUser, error := user.Current()

    if (error != nil) {
        log.Fatalf("Current user cannot be determined.")
        os.Exit(4)
    }

    return currentUser.Username
}
