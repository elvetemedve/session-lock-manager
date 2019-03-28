# session-lock-manager

This code provides you with a session lock manager for Linux desktop environments.

## Features

The session lock manager acts as a service running in the background and listening to USB security token
is inserted and rejected events. It does lock the current user session when the device is ejected and unlock
when it is inserted again.

## Supported hardware

  - all Yubikey having an USB interface (with challenge-response configured slot)
