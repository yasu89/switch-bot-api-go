package switchbot

import (
	"fmt"
	"image/color"
	"regexp"
)

// SwitchableDevice represents a device that can be switched ON or OFF.
type SwitchableDevice interface {
	TurnOn() (*CommonResponse, error)
	TurnOff() (*CommonResponse, error)
}

type ControlRequest struct {
	Command     string      `json:"command"`
	Parameter   interface{} `json:"parameter"`
	CommandType string      `json:"commandType"`
}

// sendDefaultParameterCommand sends a command with the parameter set to "default"
func sendDefaultParameterCommand(client *Client, deviceID, command string) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     command,
		Parameter:   "default",
	}
	return client.SendCommand(deviceID, request)
}

// TurnOn sends a command to turn on the BotDevice
func (device *BotDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the BotDevice
func (device *BotDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Press sends a command to press the BotDevice
func (device *BotDevice) Press() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "press")
}

type CurtainPositionMode string

const (
	CurtainPositionModePerformance CurtainPositionMode = "0"
	CurtainPositionModeSilent      CurtainPositionMode = "1"
	CurtainPositionModeDefault     CurtainPositionMode = "ff"
)

// SetPosition sends a command to set the position of the CurtainDevice
func (device *CurtainDevice) SetPosition(mode CurtainPositionMode, position int) (*CommonResponse, error) {
	if mode != CurtainPositionModePerformance && mode != CurtainPositionModeSilent && mode != CurtainPositionModeDefault {
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

// TurnOn sends a command to turn on the CurtainDevice
func (device *CurtainDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the CurtainDevice
func (device *CurtainDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Pause sends a command to pause the CurtainDevice
func (device *CurtainDevice) Pause() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "pause")
}

// Lock sends a command to lock the LockDevice
func (device *LockDevice) Lock() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "lock")
}

// Unlock sends a command to unlock the LockDevice
func (device *LockDevice) Unlock() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "unlock")
}

type KeypadKey struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	Password  string `json:"password"`
	StartTime int64  `json:"startTime,omitempty"`
	EndTime   int64  `json:"endTime,omitempty"`
}

// NewKeypadKey creates a new KeypadKey instance with validation.
func NewKeypadKey(name string, keyType string, password string, startTime int64, endTime int64) (*KeypadKey, error) {
	if keyType != "permanent" && keyType != "timeLimit" && keyType != "disposable" && keyType != "urgent" {
		return nil, fmt.Errorf("invalid keyType: %s", keyType)
	}
	passwordRegexp := regexp.MustCompile(`^\d{6,12}$`)
	if !passwordRegexp.MatchString(password) {
		return nil, fmt.Errorf("invalid password: %s", password)
	}
	if keyType == "timeLimit" || keyType == "disposable" {
		if startTime <= 0 || endTime <= 0 {
			return nil, fmt.Errorf("invalid startTime or endTime: %d, %d", startTime, endTime)
		}
		if endTime <= startTime {
			return nil, fmt.Errorf("startTime must be less than endTime: %d >= %d", startTime, endTime)
		}
	}

	return &KeypadKey{
		Name:      name,
		Type:      keyType,
		Password:  password,
		StartTime: startTime,
		EndTime:   endTime,
	}, nil
}

// CreateKey sends a command to create a new key for the KeypadDevice.
// Note: The result of this request is not returned by this method but is asynchronously returned via a webhook.
func (device *KeypadDevice) CreateKey(keypadKey *KeypadKey) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "createKey",
		Parameter:   keypadKey,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// DeleteKey sends a command to delete a key from the KeypadDevice.
// Note: The result of this request is not returned by this method but is asynchronously returned via a webhook.
func (device *KeypadDevice) DeleteKey(id string) (*CommonResponse, error) {
	deleteKeyParameter := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "deleteKey",
		Parameter:   deleteKeyParameter,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the CeilingLightDevice
func (device *CeilingLightDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the CeilingLightDevice
func (device *CeilingLightDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the CeilingLightDevice
func (device *CeilingLightDevice) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

// SetBrightness sends a command to set the brightness of the CeilingLightDevice
func (device *CeilingLightDevice) SetBrightness(brightness int) (*CommonResponse, error) {
	if brightness < 1 || brightness > 100 {
		return nil, fmt.Errorf("invalid brightness: %d", brightness)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setBrightness",
		Parameter:   fmt.Sprintf("%d", brightness),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetColorTemperature sends a command to set the color temperature of the CeilingLightDevice
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

// TurnOn sends a command to turn on the PlugMiniDevice
func (device *PlugMiniDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the PlugMiniDevice
func (device *PlugMiniDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the PlugMiniDevice
func (device *PlugMiniDevice) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

// TurnOn sends a command to turn on the PlugDevice
func (device *PlugDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the PlugDevice
func (device *PlugDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// TurnOn sends a command to turn on the StripLightDevice
func (device *StripLightDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the StripLightDevice
func (device *StripLightDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the StripLightDevice
func (device *StripLightDevice) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

// SetBrightness sends a command to set the brightness of the StripLightDevice
func (device *StripLightDevice) SetBrightness(brightness int) (*CommonResponse, error) {
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

// SetColor sends a command to set the color of the StripLightDevice
func (device *StripLightDevice) SetColor(color color.RGBA) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "setColor",
		Parameter:   fmt.Sprintf("%d:%d:%d", color.R, color.G, color.B),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the ColorBulbDevice
func (device *ColorBulbDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the ColorBulbDevice
func (device *ColorBulbDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the ColorBulbDevice
func (device *ColorBulbDevice) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

// SetBrightness sends a command to set the brightness of the ColorBulbDevice
func (device *ColorBulbDevice) SetBrightness(brightness int) (*CommonResponse, error) {
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

// SetColor sends a command to set the color of the ColorBulbDevice
func (device *ColorBulbDevice) SetColor(color color.RGBA) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "setColor",
		Parameter:   fmt.Sprintf("%d:%d:%d", color.R, color.G, color.B),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetColorTemperature sends a command to set the color temperature of the ColorBulbDevice
func (device *ColorBulbDevice) SetColorTemperature(colorTemperature int) (*CommonResponse, error) {
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

// Start sends a command to start vacuuming the RobotVacuumCleanerDevice
func (device *RobotVacuumCleanerDevice) Start() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "start")
}

// Stop sends a command to stop vacuuming the RobotVacuumCleanerDevice
func (device *RobotVacuumCleanerDevice) Stop() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "stop")
}

// Dock sends a command to return the RobotVacuumCleanerDevice to its charging dock.
func (device *RobotVacuumCleanerDevice) Dock() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "dock")
}

// RobotVacuumCleanerPowerLevel represents the power level of the RobotVacuumCleanerDevice.
type RobotVacuumCleanerPowerLevel int

const (
	RobotVacuumCleanerPowerLevelQuiet    = RobotVacuumCleanerPowerLevel(0)
	RobotVacuumCleanerPowerLevelStandard = RobotVacuumCleanerPowerLevel(1)
	RobotVacuumCleanerPowerLevelStrong   = RobotVacuumCleanerPowerLevel(2)
	RobotVacuumCleanerPowerLevelMax      = RobotVacuumCleanerPowerLevel(3)
)

// SetPowerLevel sends a command to set the suction power level of the RobotVacuumCleanerDevice.
func (device *RobotVacuumCleanerDevice) SetPowerLevel(powerLevel RobotVacuumCleanerPowerLevel) (*CommonResponse, error) {
	if powerLevel < 0 || powerLevel > 3 {
		return nil, fmt.Errorf("invalid powerLevel: %d", powerLevel)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "PowLevel",
		Parameter:   fmt.Sprintf("%d", powerLevel),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// FloorCleaningAction represents the action to be performed floor cleaning mode.
type FloorCleaningAction string

const (
	FloorCleaningActionSweep    FloorCleaningAction = "sweep"
	FloorCleaningActionSweepMop FloorCleaningAction = "sweep_mop"
)

// FloorCleaningParam represents the parameters for floor cleaning mode.
type FloorCleaningParam struct {
	FanLevel   int `json:"fanLevel"`
	WaterLevel int `json:"waterLevel"`
	Times      int `json:"times"`
}

// StartFloorCleaningParam represents the parameters for starting floor cleaning mode.
type StartFloorCleaningParam struct {
	Action FloorCleaningAction `json:"action"`
	Param  FloorCleaningParam  `json:"param"`
}

// NewStartFloorCleaningParam creates a new StartFloorCleaningParam instance with validation.
func NewStartFloorCleaningParam(action FloorCleaningAction, fanLevel int, waterLevel int, times int) (*StartFloorCleaningParam, error) {
	floorCleaningParam, err := NewFloorCleaningParam(fanLevel, waterLevel, times)
	if err != nil {
		return nil, err
	}
	return &StartFloorCleaningParam{
		Action: action,
		Param:  *floorCleaningParam,
	}, nil
}

// NewFloorCleaningParam creates a new FloorCleaningParam instance with validation.
func NewFloorCleaningParam(fanLevel int, waterLevel int, times int) (*FloorCleaningParam, error) {
	if fanLevel < 1 || fanLevel > 4 {
		return nil, fmt.Errorf("invalid fanLevel: %d", fanLevel)
	}
	if waterLevel < 1 || waterLevel > 2 {
		return nil, fmt.Errorf("invalid waterLevel: %d", waterLevel)
	}
	if times < 1 || times > 2639999 {
		return nil, fmt.Errorf("invalid times: %d", times)
	}
	return &FloorCleaningParam{
		FanLevel:   fanLevel,
		WaterLevel: waterLevel,
		Times:      times,
	}, nil
}

// StartClean sends a command to start cleaning the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) StartClean(startFloorCleaningParam *StartFloorCleaningParam) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "startClean",
		Parameter:   startFloorCleaningParam,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// AddWaterForHumi sends a command to refill the mind-blowing Evaporative Humidifier (Auto-refill) in the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) AddWaterForHumi() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "addWaterForHumi")
}

// Pause sends a command to pause the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) Pause() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "pause")
}

// Dock sends a command to return the RobotVacuumCleanerS10Device to its charging dock.
func (device *RobotVacuumCleanerS10Device) Dock() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "dock")
}

// SetVolume sends a command to set the volume of the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) SetVolume(volume int) (*CommonResponse, error) {
	if volume < 0 || volume > 100 {
		return nil, fmt.Errorf("invalid volume: %d", volume)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setVolume",
		Parameter:   fmt.Sprintf("%d", volume),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

type SelfCleaningMode int

const (
	WashMopSelfCleaningMode   = SelfCleaningMode(1)
	DrySelfCleaningMode       = SelfCleaningMode(2)
	TerminateSelfCleaningMode = SelfCleaningMode(3)
)

// SelfClean sends a command to start self-cleaning the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) SelfClean(mode SelfCleaningMode) (*CommonResponse, error) {
	if mode < 1 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "selfClean",
		Parameter:   fmt.Sprintf("%d", mode),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// ChangeParam sends a command to change the cleaning parameters of the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) ChangeParam(floorCleaningParam *FloorCleaningParam) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "changeParam",
		Parameter:   floorCleaningParam,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the HumidifierDevice
func (device *HumidifierDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the HumidifierDevice
func (device *HumidifierDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

type HumidifierMode int

const (
	HumidifierModeAuto   = HumidifierMode(0)
	HumidifierModeLow    = HumidifierMode(101)
	HumidifierModeMedium = HumidifierMode(102)
	HumidifierModeHigh   = HumidifierMode(103)
)

// SetMode sends a command to set the mode of the HumidifierDevice
func (device *HumidifierDevice) SetMode(mode HumidifierMode) (*CommonResponse, error) {
	if (mode < 101 || mode > 103) && (mode != 0) {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetTargetHumidity sends a command to set the target humidity of the HumidifierDevice
func (device *HumidifierDevice) SetTargetHumidity(targetHumidity int) (*CommonResponse, error) {
	if targetHumidity < 0 || targetHumidity > 100 {
		return nil, fmt.Errorf("invalid mode: %d", targetHumidity)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", targetHumidity),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the EvaporativeHumidifierDevice
func (device *EvaporativeHumidifierDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the EvaporativeHumidifierDevice
func (device *EvaporativeHumidifierDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

type EvaporativeHumidifierMode int

const (
	EvaporativeHumidifierModeLevel4   = EvaporativeHumidifierMode(1)
	EvaporativeHumidifierModeLevel3   = EvaporativeHumidifierMode(2)
	EvaporativeHumidifierModeLevel2   = EvaporativeHumidifierMode(3)
	EvaporativeHumidifierModeLevel1   = EvaporativeHumidifierMode(4)
	EvaporativeHumidifierModeHumidity = EvaporativeHumidifierMode(5)
	EvaporativeHumidifierModeSleep    = EvaporativeHumidifierMode(6)
	EvaporativeHumidifierModeAuto     = EvaporativeHumidifierMode(7)
	EvaporativeHumidifierModeDry      = EvaporativeHumidifierMode(8)
)

// EvaporativeHumidifierModeParam represents the parameters for setting the mode of the EvaporativeHumidifierDevice.
type EvaporativeHumidifierModeParam struct {
	Mode           EvaporativeHumidifierMode `json:"mode"`
	TargetHumidity int                       `json:"targetHumidity"`
}

// SetMode sends a command to set the mode of the EvaporativeHumidifierDevice
func (device *EvaporativeHumidifierDevice) SetMode(mode EvaporativeHumidifierMode, targetHumidity int) (*CommonResponse, error) {
	if mode < 1 || mode > 8 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	if targetHumidity < 0 || targetHumidity > 100 {
		return nil, fmt.Errorf("invalid targetHumidity: %d", targetHumidity)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter: EvaporativeHumidifierModeParam{
			Mode:           mode,
			TargetHumidity: targetHumidity,
		},
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetChildLock sends a command to set the child lock of the EvaporativeHumidifierDevice
func (device *EvaporativeHumidifierDevice) SetChildLock(flag bool) (*CommonResponse, error) {
	request := ControlRequest{
		CommandType: "command",
		Command:     "setChildLock",
		Parameter:   fmt.Sprintf("%t", flag),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the AirPurifierDevice
func (device *AirPurifierDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the AirPurifierDevice
func (device *AirPurifierDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

type AirPurifierMode int

const (
	AirPurifierModeNormal = AirPurifierMode(1)
	AirPurifierModeAuto   = AirPurifierMode(2)
	AirPurifierModeSleep  = AirPurifierMode(3)
	AirPurifierModePet    = AirPurifierMode(4)
)

type AirPurifierModeParameter struct {
	Mode    AirPurifierMode `json:"mode"`
	FanGear int             `json:"fanGear,omitempty"`
}

// SetMode sends a command to set the mode of the AirPurifierDevice
func (device *AirPurifierDevice) SetMode(mode AirPurifierMode, fanLevel int) (*CommonResponse, error) {
	if mode < 1 || mode > 4 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	if mode == AirPurifierModeNormal {
		if fanLevel < 1 || fanLevel > 3 {
			return nil, fmt.Errorf("invalid fanLevel: %d", fanLevel)
		}
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter: AirPurifierModeParameter{
			Mode:    mode,
			FanGear: fanLevel,
		},
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetChildLock sends a command to set the child lock of the AirPurifierDevice
func (device *AirPurifierDevice) SetChildLock(flag bool) (*CommonResponse, error) {
	flagInt := 0
	if flag {
		flagInt = 1
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setChildLock",
		Parameter:   flagInt,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetPosition sends a command to set the position of the BlindTiltDevice
func (device *BlindTiltDevice) SetPosition(direction string, position int) (*CommonResponse, error) {
	if direction != "up" && direction != "down" {
		return nil, fmt.Errorf("invalid direction: %s", direction)
	}
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}
	if position%2 != 0 {
		return nil, fmt.Errorf("position must be even: %d", position)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		Parameter:   fmt.Sprintf("%s;%d", direction, position),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// FullyOpen sends a command to fully open the BlindTiltDevice
func (device *BlindTiltDevice) FullyOpen() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "fullyOpen")
}

// CloseUp sends a command to close up the BlindTiltDevice
func (device *BlindTiltDevice) CloseUp() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "closeUp")
}

// CloseDown sends a command to close down the BlindTiltDevice
func (device *BlindTiltDevice) CloseDown() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "closeDown")
}

// TurnOn sends a command to turn on the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

type CirculatorNightLightMode string

const (
	CirculatorNightLightModeTurnOff    = CirculatorNightLightMode("off")
	CirculatorNightLightModeTurnBright = CirculatorNightLightMode("1")
	CirculatorNightLightModeTurnDim    = CirculatorNightLightMode("2")
)

// SetNightLightMode sends a command to set the night light mode of the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) SetNightLightMode(mode CirculatorNightLightMode) (*CommonResponse, error) {
	if mode != CirculatorNightLightModeTurnOff && mode != CirculatorNightLightModeTurnBright && mode != CirculatorNightLightModeTurnDim {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setNightLightMode",
		Parameter:   mode,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

type CirculatorWindMode string

const (
	CirculatorWindModeDirect  = CirculatorWindMode("direct")
	CirculatorWindModeNatural = CirculatorWindMode("natural")
	CirculatorWindModeSleep   = CirculatorWindMode("sleep")
	CirculatorWindModeBaby    = CirculatorWindMode("baby")
)

// SetWindMode sends a command to set the wind mode of the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) SetWindMode(mode CirculatorWindMode) (*CommonResponse, error) {
	if mode != CirculatorWindModeDirect && mode != CirculatorWindModeNatural && mode != CirculatorWindModeSleep && mode != CirculatorWindModeBaby {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setWindMode",
		Parameter:   mode,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetWindSpeed sends a command to set the wind speed of the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) SetWindSpeed(speed int) (*CommonResponse, error) {
	if speed < 1 || speed > 100 {
		return nil, fmt.Errorf("invalid speed: %d", speed)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setWindSpeed",
		Parameter:   fmt.Sprintf("%d", speed),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the CirculatorFanDevice
func (device *CirculatorFanDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// SetNightLightMode sends a command to set the night light mode of the CirculatorFanDevice
func (device *CirculatorFanDevice) SetNightLightMode(mode CirculatorNightLightMode) (*CommonResponse, error) {
	if mode != CirculatorNightLightModeTurnOff && mode != CirculatorNightLightModeTurnBright && mode != CirculatorNightLightModeTurnDim {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setNightLightMode",
		Parameter:   mode,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOff sends a command to turn off the CirculatorFanDevice
func (device *CirculatorFanDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// SetWindMode sends a command to set the wind mode of the CirculatorFanDevice
func (device *CirculatorFanDevice) SetWindMode(mode CirculatorWindMode) (*CommonResponse, error) {
	if mode != CirculatorWindModeDirect && mode != CirculatorWindModeNatural && mode != CirculatorWindModeSleep && mode != CirculatorWindModeBaby {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setWindMode",
		Parameter:   mode,
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetWindSpeed sends a command to set the wind speed of the CirculatorFanDevice
func (device *CirculatorFanDevice) SetWindSpeed(speed int) (*CommonResponse, error) {
	if speed < 1 || speed > 100 {
		return nil, fmt.Errorf("invalid speed: %d", speed)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setWindSpeed",
		Parameter:   fmt.Sprintf("%d", speed),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// SetPosition sends a command to set the position of the RollerShadeDevice
func (device *RollerShadeDevice) SetPosition(position int) (*CommonResponse, error) {
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		Parameter:   fmt.Sprintf("%d", position),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the RelaySwitch1PMDevice
func (device *RelaySwitch1PMDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the RelaySwitch1PMDevice
func (device *RelaySwitch1PMDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the RelaySwitch1PMDevice
func (device *RelaySwitch1PMDevice) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

type RelaySwitchMode int

const (
	RelaySwitchModeToggle    = RelaySwitchMode(0)
	RelaySwitchModeEdge      = RelaySwitchMode(1)
	RelaySwitchModeDetached  = RelaySwitchMode(2)
	RelaySwitchModeMomentary = RelaySwitchMode(3)
)

// SetMode sends a command to set the mode of the RelaySwitch1PMDevice
func (device *RelaySwitch1PMDevice) SetMode(mode RelaySwitchMode) (*CommonResponse, error) {
	if mode < 0 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the RelaySwitch1Device
func (device *RelaySwitch1Device) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the RelaySwitch1Device
func (device *RelaySwitch1Device) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
}

// Toggle sends a command to toggle the RelaySwitch1Device
func (device *RelaySwitch1Device) Toggle() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "toggle")
}

// SetMode sends a command to set the mode of the RelaySwitch1Device
func (device *RelaySwitch1Device) SetMode(mode RelaySwitchMode) (*CommonResponse, error) {
	if mode < 0 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	request := ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}
	return device.Client.SendCommand(device.DeviceID, request)
}

// TurnOn sends a command to turn on the InfraredRemoteDevice
func (device *InfraredRemoteDevice) TurnOn() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOn")
}

// TurnOff sends a command to turn off the InfraredRemoteDevice
func (device *InfraredRemoteDevice) TurnOff() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "turnOff")
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
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "volumeAdd")
}

// VolumeSub sends a command to decrease the volume of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) VolumeSub() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "volumeSub")
}

// ChannelAdd sends a command to increase the channel of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice
func (device *InfraredRemoteTVDevice) ChannelAdd() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "channelAdd")
}

// SetMute sends a command to mute/unmute the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) SetMute() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "setMute")
}

// FastForward sends a command to fast-forward the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) FastForward() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "fastForward")
}

// Rewind sends a command to rewind the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Rewind() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Rewind")
}

// Next sends a command to play the next track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Next() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Next")
}

// Previous sends a command to play the previous track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Previous() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Previous")
}

// Pause sends a command to pause the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Pause() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Pause")
}

// Play sends a command to play/resume the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Play() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Play")
}

// Stop sends a command to stop the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) Stop() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "Stop")
}

// VolumeAdd sends a command to increase the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) VolumeAdd() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "volumeAdd")
}

// VolumeSub sends a command to decrease the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) VolumeSub() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "volumeSub")
}

// Swing sends a command to enable/disable the swing feature of the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) Swing() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "swing")
}

// Timer sends a command to set the timer of the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) Timer() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "timer")
}

// LowSpeed sends a command to set the fan speed to low on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) LowSpeed() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "lowSpeed")
}

// MiddleSpeed sends a command to set the fan speed to middle on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) MiddleSpeed() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "middleSpeed")
}

// HighSpeed sends a command to set the fan speed to high on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) HighSpeed() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "highSpeed")
}

// BrightnessUp sends a command to increase the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) BrightnessUp() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "brightnessUp")
}

// BrightnessDown sends a command to decrease the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) BrightnessDown() (*CommonResponse, error) {
	return sendDefaultParameterCommand(device.Client, device.DeviceID, "brightnessDown")
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
