package device

type Action int

const (
    insert Action = iota
    eject Action = iota
)

func (a Action) String() string {
    return [...]string{"add", "remove"}[a]
}

type SecurityTokenEvent struct {
    action Action
    deviceId string
}
