package device

import (
    "context"
    "github.com/jochenvg/go-udev"
)

type UdevScanner struct {
    *udev.Udev
}

func (udev *UdevScanner) Scan() (<-chan *SecurityTokenEvent, context.CancelFunc) {
    // Create Monitor
    monitor :=  udev.NewMonitorFromNetlink("udev")
    // Add filters to monitor
    monitor.FilterAddMatchSubsystemDevtype("usb", "usb_device")

    // Create a context
    ctx, cancel := context.WithCancel(context.Background())

    // Start monitor goroutine and get receive channel
    channel, _ := monitor.DeviceChan(ctx)

    eventChannel := make(chan *SecurityTokenEvent)
    go udev.readEvents(eventChannel, channel)

    return eventChannel, cancel
}

func (udev *UdevScanner) readEvents(eventChannel chan *SecurityTokenEvent, channel <-chan *udev.Device) {
    for device := range channel {
        event := SecurityTokenEvent{}
        event.deviceId = device.Properties()["PRODUCT"]
        switch device.Action() {
        case insert.String():
            event.action = insert
        case eject.String():
            event.action = eject
        default:
            continue
        }
        eventChannel <- &event
    }
}
