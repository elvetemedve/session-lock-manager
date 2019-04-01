package authentication

type FakeUserProvider struct {
    currentUsername string
}

func (userProvider *FakeUserProvider) GetCurrentUsername() (string) {
    return userProvider.currentUsername
}

func (userProvider *FakeUserProvider) setCurrentUsername(username string) {
    userProvider.currentUsername = username
}
