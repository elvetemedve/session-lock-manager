package session

import (
	"fmt"
	"github.com/godbus/dbus"
	"log"
)

func Lock() {
	error := callDbusMethod("LockSessions")
	if error != nil {
		fmt.Println(fmt.Sprintf("Failed to lock session: %s", error))
		return
	}

	log.Println("Locking session.")
}

func Unlock() {
	error := callDbusMethod("UnlockSessions")
	if error != nil {
		fmt.Println(fmt.Sprintf("Failed to lock session: %s", error))
		return
	}

	log.Println("Unlocking session.")
}

func callDbusMethod(methodName string) error {
	connection, error := dbus.SystemBus()
	if error != nil {
		return fmt.Errorf("failed to connectionect to session bus: %w", error)
	}

	object := connection.Object("org.freedesktop.login1", "/org/freedesktop/login1")
	call := object.Call(fmt.Sprintf("org.freedesktop.login1.Manager.%s", methodName), 0)
	if call.Err != nil {
		return call.Err
	}
	return nil
}
