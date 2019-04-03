# session-lock-manager

This code provides you with a session lock manager for Linux desktop environments.

## Features

The session lock manager acts as a service running in the background and listening to USB security token
is inserted and rejected events. It does lock the current user session when the device is ejected and unlock
when it is inserted again.

## Supported hardware

  - all Yubikey having an USB interface (with challenge-response configured slot)

## Development

### Building the application

    go build github.com/elvetemedve/session-lock-manager

### Running tests

    go test -v github.com/elvetemedve/session-lock-manager/device
    go test -v github.com/elvetemedve/session-lock-manager/authentication

### Running the application

    go run github.com/elvetemedve/session-lock-manager <service-name>

    where service name is the appropriate filename in the pam.d directory

### Architecture

![Architecture diagram](./docs/images/architecture-diagram.svg)

## Configuration

### Yubikey

Create a file like `/etc/pam.d/session-locker` with the content below:

    auth		required	pam_yubico.so mode=challenge-response
Now use the Yubikey configuration tool to setup a slot for challenge-response authentication without user presence.

## Known issues

   1. The application cannot be run as regular user on systems where polkit service is installed, because session locking, unlocking requires root privileges. See `/usr/share/polkit-1/actions/org.freedesktop.login1.policy` **org.freedesktop.login1.lock-sessions** section.

   To fix this, edit the policy file and find the XML node `<action id="org.freedesktop.login1.lock-sessions">`. Under that node change the `<default>` node to be like below:
   ```
   <defaults>
        <allow_any>yes</allow_any>
        <allow_inactive>auth_admin_keep</allow_inactive>
        <allow_active>yes</allow_active>
    </defaults>
    ```
