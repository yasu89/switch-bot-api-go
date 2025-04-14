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

func (device *CeilingLightDevice) TurnOn() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CeilingLightDevice) TurnOff() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CeilingLightDevice) Toggle() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "toggle",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CeilingLightDevice) SetBrightness(brightness int) (*CommonResponse, error) {
	if brightness < 0 || brightness > 100 {
		return nil, fmt.Errorf("invalid brightness: %d", brightness)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setBrightness",
		Parameter:   fmt.Sprintf("%d", brightness),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (device *CeilingLightDevice) SetColorTemperature(colorTemperature int) (*CommonResponse, error) {
	if colorTemperature < 2700 || colorTemperature > 6500 {
		return nil, fmt.Errorf("invalid colorTemperature: %d", colorTemperature)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setColorTemperature",
		Parameter:   fmt.Sprintf("%d", colorTemperature),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the InfraredRemoteDevice
func (device *InfraredRemoteDevice) TurnOn() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOff sends a command to turn off the InfraredRemoteDevice
func (device *InfraredRemoteDevice) TurnOff() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

type AirConditionerMode int
type AirConditionerFanMode int
type AirConditionerPowerState string

const (
	AirConditionerModeAuto      = AirConditionerMode(1)
	AirConditionerModeCool      = AirConditionerMode(2)
	AirConditionerModeDry       = AirConditionerMode(3)
	AirConditionerModeFan       = AirConditionerMode(4)
	AirConditionerModeHeat      = AirConditionerMode(5)
	AirConditionerFanModeAuto   = AirConditionerFanMode(1)
	AirConditionerFanModeLow    = AirConditionerFanMode(2)
	AirConditionerFanModeMedium = AirConditionerFanMode(3)
	AirConditionerFanModeHigh   = AirConditionerFanMode(4)
	AirConditionerPowerStateOn  = AirConditionerPowerState("on")
	AirConditionerPowerStateOff = AirConditionerPowerState("off")
)

// SetAll sends a command to configure all parameters of the InfraredRemoteAirConditionerDevice
func (device *InfraredRemoteAirConditionerDevice) SetAll(
	temperatureCelsius int, mode AirConditionerMode, fan AirConditionerFanMode, powerState AirConditionerPowerState,
) (*CommonResponse, error) {
	if temperatureCelsius < -10 || temperatureCelsius > 40 {
		return nil, fmt.Errorf("invalid temperatureCelsius: %d", temperatureCelsius)
	}
	if mode < 1 || mode > 5 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	if fan < 1 || fan > 4 {
		return nil, fmt.Errorf("invalid fan: %d", fan)
	}

	parameter := fmt.Sprintf("%d,%d,%d,%s", temperatureCelsius, mode, fan, powerState)
	request := ControlRequest{
		CommandType: "command",
		Command:     "setAll",
		Parameter:   parameter,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetChannel sends a command to set the channel of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) SetChannel(channel int) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "SetChannel",
		Parameter:   fmt.Sprintf("%d", channel),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// VolumeAdd sends a command to increase the volume of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) VolumeAdd() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "volumeAdd",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// VolumeSub sends a command to decrease the volume of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) VolumeSub() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "volumeSub",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// ChannelAdd sends a command to increase the channel of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) ChannelAdd() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "channelAdd",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetMute sends a command to mute/unmute the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) SetMute() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMute",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// FastForward sends a command to fast-forward the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) FastForward() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "FastForward",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Rewind sends a command to rewind the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Rewind() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Rewind",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Next sends a command to play the next track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Next() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Next",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Previous sends a command to play the previous track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Previous() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Previous",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Pause sends a command to pause the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Pause() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Pause",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Play sends a command to play/resume the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Play() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Play",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Stop sends a command to stop the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Stop() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "Stop",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// VolumeAdd sends a command to increase the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) VolumeAdd() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "volumeAdd",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// VolumeSub sends a command to decrease the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) VolumeSub() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "volumeSub",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Swing sends a command to enable/disable the swing feature of the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) Swing() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "swing",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// Timer sends a command to set the timer of the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) Timer() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "timer",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// LowSpeed sends a command to set the fan speed to low on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) LowSpeed() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "lowSpeed",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// MiddleSpeed sends a command to set the fan speed to middle on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) MiddleSpeed() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "middleSpeed",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// HighSpeed sends a command to set the fan speed to high on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) HighSpeed() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "highSpeed",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// BrightnessUp sends a command to increase the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) BrightnessUp() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "brightnessUp",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// BrightnessDown sends a command to decrease the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) BrightnessDown() (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "brightnessDown",
		Parameter:   "default",
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// CustomCommand sends a user-defined command to the InfraredRemoteOthersDevice
func (device *InfraredRemoteOthersDevice) CustomCommand(buttonName string) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "customize",
		Parameter:   buttonName,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

func (client *Client) SendCommand(deviceId string, request ControlRequest) (*CommonResponse, error) {
	return client.PostRequest("/devices/"+deviceId+"/commands", request)
}
