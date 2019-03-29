package device

import "context"

type FakeScanner struct {
    events []SecurityTokenEvent
}

func (scanner *FakeScanner) Scan() (<-chan *SecurityTokenEvent, context.CancelFunc) {
    eventChannel := make(chan *SecurityTokenEvent)
    cancel := func() {
        close(eventChannel)
    }

    go scanner.readEvents(eventChannel)

    return eventChannel, cancel
}

func (scanner *FakeScanner) readEvents(eventChannel chan *SecurityTokenEvent) {
    for _, event := range scanner.events {
        eventChannel <- &event
    }
}

func (scanner *FakeScanner) add(event SecurityTokenEvent) {
    scanner.events = append(scanner.events, event)
}
