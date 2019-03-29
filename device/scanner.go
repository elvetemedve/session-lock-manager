package device

import "context"

type Scanner interface {
    Scan() (<-chan *SecurityTokenEvent, context.CancelFunc)
}
