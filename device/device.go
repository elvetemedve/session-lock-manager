package device

type Device struct {
    deviceId string
}

type Devices struct {
    devices []Device
}

var devices Devices

func init() {
    yubikey := Device{"1050/407/511"}
    devices = Devices{ []Device{ yubikey } }
}

func (devices *Devices) IsSupported(deviceId string) (bool) {
    for _, device := range devices.devices {
        if device.deviceId == deviceId {
            return true
        }
    }

    return false
}
