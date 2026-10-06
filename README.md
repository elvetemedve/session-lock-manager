# session-lock-manager

This code provides you with a session lock manager for Linux desktop environments.

## Features

The session lock manager acts as a service running in the background and listening to USB security token
is inserted and rejected events. It does lock the current user session when the device is ejected and unlock
when it is inserted again.

## Supported hardware

  - all Yubikey having an USB interface (with challenge-response configured slot)

## Development

### Prerequisites

  - Go 1.22 or newer
  - a C compiler and `pkg-config` (the PAM and udev bindings use cgo)
  - PAM and libudev development headers, e.g. `libpam0g-dev libudev-dev` on Debian/Ubuntu
    or `pam-devel systemd-devel` on Fedora

### Building the application

    go build .

### Running tests

    go test ./...

### Running the application

    go run . <service-name>

where `<service-name>` is the name of the appropriate file in the `/etc/pam.d` directory
(see [Yubikey](#yubikey) below).

### Architecture

![Architecture diagram](./docs/images/architecture-diagram.svg)

## Configuration

### Yubikey

Create a file like `/etc/pam.d/session-locker` with the content below:

    auth		required	pam_yubico.so mode=challenge-response
Now use the Yubikey configuration tool to setup a slot for challenge-response authentication without user presence.

### Polkit

If your Linux desktop has Polkit installed to control system-wide privileges, then you need to configure it to allow locking/unlocking the session as regular user,
because the session lock manager should not to be run as root. The simplest way to achieve this is to add your user into a new group who is trusted to lock/unlock session.
The following example assumes that the user is added to the **wheel** group.

Create the file `/etc/polkit-1/rules.d/49-sesson-lock-manager.rules` with the content below:

    /* Allow members of the wheel group to execute the defined actions
     * without password authentication, similar to "sudo NOPASSWD:"
     */
    polkit.addRule(function(action, subject) {
        if (action.id == "org.freedesktop.login1.lock-sessions" &&
            subject.isInGroup("wheel"))
        {
            return polkit.Result.YES;
        }
    });
