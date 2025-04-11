package switchbot

import "fmt"

type ControlRequest struct {
	Command     string `json:"command"`
	Parameter   string `json:"parameter"`
	CommandType string `json:"commandType"`
}

func (device *BotDevice) TurnOn() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *BotDevice) TurnOff() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *BotDevice) Press() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "press",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CurtainDevice) SetPosition(mode string, position int) (*CommonResponse, error) {
	if mode != "0" && mode != "1" && mode != "ff" {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}

	request := ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		// MEMO: The "index0" parameter is unclear, so it is fixed to 0 for now.
		Parameter: fmt.Sprintf("0,%s,%d", mode, position),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CurtainDevice) TurnOn() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CurtainDevice) TurnOff() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CurtainDevice) Pause() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "pause",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *LockDevice) Lock() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "lock",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *LockDevice) Unlock() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "unlock",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (InfraredRemoteDevice *InfraredRemoteDevice) TurnOn() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
	return InfraredRemoteDevice.Client.SendCommand(InfraredRemoteDevice.DeviceID, request)
}

func (InfraredRemoteDevice *InfraredRemoteDevice) TurnOff() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
	return InfraredRemoteDevice.Client.SendCommand(InfraredRemoteDevice.DeviceID, request)
}

func (client *Client) SendCommand(deviceId string, request ControlRequest) (*CommonResponse, error) {
	return client.PostRequest("/devices/"+deviceId+"/commands", request)
}
