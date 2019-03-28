package main

import (
    "fmt"
    "context"
    "github.com/jochenvg/go-udev"
)

const YUBIKEY_DEVICE_ID = "1050/407/511";

func main() {
    // Create Udev and Monitor
    u := udev.Udev{}
    m := u.NewMonitorFromNetlink("udev")

    // Add filters to monitor
    m.FilterAddMatchSubsystemDevtype("usb", "usb_device")

    // Create a context
    ctx, _ := context.WithCancel(context.Background())

    // Start monitor goroutine and get receive channel
    ch, _ := m.DeviceChan(ctx)

    fmt.Println("Started listening on channel")
    for d := range ch {
        switch d.Action() {
        case "add":
            if (YUBIKEY_DEVICE_ID == d.Properties()["PRODUCT"]) {
                fmt.Println("USB device inserted.")
            }
        case "remove":
            if (YUBIKEY_DEVICE_ID == d.Properties()["PRODUCT"]) {
                fmt.Println("USB device ejected.")
            }
        }
    }
    fmt.Println("Channel closed")
}
