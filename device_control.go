package switchbot

import (
	"fmt"
	"image/color"
	"regexp"
)

type ControlRequest struct {
	Command     string      `json:"command"`
	Parameter   interface{} `json:"parameter"`
	CommandType string      `json:"commandType"`
}

// GetTurnOnCommand generates a ControlRequest to turn on the device for the TurnOnOffDevice.
func (device *TurnOnOffDevice) GetTurnOnCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "turnOn",
		Parameter:   "default",
	}
}

// GetTurnOffCommand generates a ControlRequest to turn off the device for the TurnOnOffDevice.
func (device *TurnOnOffDevice) GetTurnOffCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "turnOff",
		Parameter:   "default",
	}
}

// GetToggleCommand generates a ControlRequest to toggle the device for the CeilingLightDevice.
func (device *ToggleCommandDevice) GetToggleCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "toggle",
		Parameter:   "default",
	}
}

// GetPressCommand generates a ControlRequest to press the device for the BotDevice.
func (device *BotDevice) GetPressCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "press",
		Parameter:   "default",
	}
}

type CurtainPositionMode string

const (
	CurtainPositionModePerformance CurtainPositionMode = "0"
	CurtainPositionModeSilent      CurtainPositionMode = "1"
	CurtainPositionModeDefault     CurtainPositionMode = "ff"
)

// GetSetPositionCommand generates a ControlRequest to set the position of the curtain for the CurtainDevice.
func (device *CurtainDevice) GetSetPositionCommand(mode CurtainPositionMode, position int) (*ControlRequest, error) {
	if mode != CurtainPositionModePerformance && mode != CurtainPositionModeSilent && mode != CurtainPositionModeDefault {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}

	return &ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		// MEMO: The "index0" parameter is unclear, so it is fixed to 0 for now.
		Parameter: fmt.Sprintf("0,%s,%d", mode, position),
	}, nil
}

// GetPauseCommand generates a ControlRequest to pause the curtain for the CurtainDevice.
func (device *CurtainDevice) GetPauseCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "pause",
		Parameter:   "default",
	}
}

// GetLockCommand generates a ControlRequest to lock the device for the LockDevice.
func (device *LockDevice) GetLockCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "lock",
		Parameter:   "default",
	}
}

// GetUnlockCommand generates a ControlRequest to unlock the device for the LockDevice.
func (device *LockDevice) GetUnlockCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "unlock",
		Parameter:   "default",
	}
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

// GetCreateKeyCommand generates a ControlRequest to create a key for the KeypadDevice.
// Note: The result of this request is not returned by this method but is asynchronously returned via a webhook.
func (device *KeypadDevice) GetCreateKeyCommand(keypadKey *KeypadKey) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "createKey",
		Parameter:   keypadKey,
	}
}

// GetDeleteKeyCommand create command to delete a key for the KeypadDevice.
// Note: The result of this request is not returned by this method but is asynchronously returned via a webhook.
func (device *KeypadDevice) GetDeleteKeyCommand(id string) *ControlRequest {
	deleteKeyParameter := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "deleteKey",
		Parameter:   deleteKeyParameter,
	}
}

// GetSetBrightnessCommand generates a ControlRequest to set the brightness of the device for the CeilingLightDevice.
func (device *CeilingLightDevice) GetSetBrightnessCommand(brightness int) (*ControlRequest, error) {
	if brightness < 0 || brightness > 100 {
		return nil, fmt.Errorf("invalid brightness: %d", brightness)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setBrightness",
		Parameter:   fmt.Sprintf("%d", brightness),
	}, nil
}

// GetSetColorTemperatureCommand generates a ControlRequest to set the color temperature of the device for the CeilingLightDevice.
func (device *CeilingLightDevice) GetSetColorTemperatureCommand(colorTemperature int) (*ControlRequest, error) {
	if colorTemperature < 2700 || colorTemperature > 6500 {
		return nil, fmt.Errorf("invalid colorTemperature: %d", colorTemperature)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setColorTemperature",
		Parameter:   fmt.Sprintf("%d", colorTemperature),
	}, nil
}

// GetSetBrightnessCommand generates a ControlRequest to set the brightness of the device for the StripLightDevice.
func (device *StripLightDevice) GetSetBrightnessCommand(brightness int) (*ControlRequest, error) {
	if brightness < 0 || brightness > 100 {
		return nil, fmt.Errorf("invalid brightness: %d", brightness)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setBrightness",
		Parameter:   fmt.Sprintf("%d", brightness),
	}, nil
}

// GetSetColorCommand generates a ControlRequest to set the color of the device for the StripLightDevice.
func (device *StripLightDevice) GetSetColorCommand(color color.RGBA) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "setColor",
		Parameter:   fmt.Sprintf("%d:%d:%d", color.R, color.G, color.B),
	}
}

// GetSetBrightnessCommand generates a ControlRequest to set the brightness of the device for the ColorBulbDevice.
func (device *ColorBulbDevice) GetSetBrightnessCommand(brightness int) (*ControlRequest, error) {
	if brightness < 0 || brightness > 100 {
		return nil, fmt.Errorf("invalid brightness: %d", brightness)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setBrightness",
		Parameter:   fmt.Sprintf("%d", brightness),
	}, nil
}

// GetSetColorCommand generates a ControlRequest to set the color of the device for the ColorBulbDevice.
func (device *ColorBulbDevice) GetSetColorCommand(color color.RGBA) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "setColor",
		Parameter:   fmt.Sprintf("%d:%d:%d", color.R, color.G, color.B),
	}
}

// GetSetColorTemperatureCommand generates a ControlRequest to set the color temperature of the device for the ColorBulbDevice.
func (device *ColorBulbDevice) GetSetColorTemperatureCommand(colorTemperature int) (*ControlRequest, error) {
	if colorTemperature < 2700 || colorTemperature > 6500 {
		return nil, fmt.Errorf("invalid colorTemperature: %d", colorTemperature)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setColorTemperature",
		Parameter:   fmt.Sprintf("%d", colorTemperature),
	}, nil
}

// GetStartCommand generates a ControlRequest to start vacuuming the RobotVacuumCleanerDevice.
func (device *RobotVacuumCleanerDevice) GetStartCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "start",
		Parameter:   "default",
	}
}

// GetStopCommand generates a ControlRequest to stop vacuuming the RobotVacuumCleanerDevice.
func (device *RobotVacuumCleanerDevice) GetStopCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "stop",
		Parameter:   "default",
	}
}

// GetDockCommand generates a ControlRequest to dock the RobotVacuumCleanerDevice.
func (device *RobotVacuumCleanerDevice) GetDockCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "dock",
		Parameter:   "default",
	}
}

// RobotVacuumCleanerPowerLevel represents the power level of the RobotVacuumCleanerDevice.
type RobotVacuumCleanerPowerLevel int

const (
	RobotVacuumCleanerPowerLevelQuiet    = RobotVacuumCleanerPowerLevel(0)
	RobotVacuumCleanerPowerLevelStandard = RobotVacuumCleanerPowerLevel(1)
	RobotVacuumCleanerPowerLevelStrong   = RobotVacuumCleanerPowerLevel(2)
	RobotVacuumCleanerPowerLevelMax      = RobotVacuumCleanerPowerLevel(3)
)

// GetSetPowerLevelCommand generates a ControlRequest to set the power level of the RobotVacuumCleanerDevice.
func (device *RobotVacuumCleanerDevice) GetSetPowerLevelCommand(powerLevel RobotVacuumCleanerPowerLevel) (*ControlRequest, error) {
	if powerLevel < 0 || powerLevel > 3 {
		return nil, fmt.Errorf("invalid powerLevel: %d", powerLevel)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "PowLevel",
		Parameter:   fmt.Sprintf("%d", powerLevel),
	}, nil
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

// GetStartCleanCommand generates a ControlRequest to start cleaning the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetStartCleanCommand(startFloorCleaningParam StartFloorCleaningParam) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "startClean",
		Parameter:   startFloorCleaningParam,
	}
}

// GetAddWaterForHumiCommand generates a ControlRequest to refill the mind-blowing Evaporative Humidifier (Auto-refill) in the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetAddWaterForHumiCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "addWaterForHumi",
		Parameter:   "default",
	}
}

// GetPauseCommand generates a ControlRequest to pause the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetPauseCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "pause",
		Parameter:   "default",
	}
}

// GetDockCommand generates a ControlRequest to dock the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetDockCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "dock",
		Parameter:   "default",
	}
}

// GetSetVolumeCommand generates a ControlRequest to set the volume of the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetSetVolumeCommand(volume int) (*ControlRequest, error) {
	if volume < 0 || volume > 100 {
		return nil, fmt.Errorf("invalid volume: %d", volume)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setVolume",
		Parameter:   fmt.Sprintf("%d", volume),
	}, nil
}

// SelfCleaningMode represents the self-cleaning mode of the RobotVacuumCleanerS10Device.
type SelfCleaningMode int

const (
	WashMopSelfCleaningMode   = SelfCleaningMode(1)
	DrySelfCleaningMode       = SelfCleaningMode(2)
	TerminateSelfCleaningMode = SelfCleaningMode(3)
)

// GetSelfCleanCommand generates a ControlRequest to start self-cleaning the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetSelfCleanCommand(mode SelfCleaningMode) (*ControlRequest, error) {
	if mode < 1 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "selfClean",
		Parameter:   fmt.Sprintf("%d", mode),
	}, nil
}

// GetChangeParamCommand generates a ControlRequest to change the parameters of the RobotVacuumCleanerS10Device.
func (device *RobotVacuumCleanerS10Device) GetChangeParamCommand(floorCleaningParam FloorCleaningParam) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "changeParam",
		Parameter:   floorCleaningParam,
	}
}

type HumidifierMode int

const (
	HumidifierModeAuto   = HumidifierMode(0)
	HumidifierModeLow    = HumidifierMode(101)
	HumidifierModeMedium = HumidifierMode(102)
	HumidifierModeHigh   = HumidifierMode(103)
)

// GetSetModeCommand generates a ControlRequest to set the mode of the HumidifierDevice.
func (device *HumidifierDevice) GetSetModeCommand(mode HumidifierMode) (*ControlRequest, error) {
	if (mode < 101 || mode > 103) && (mode != 0) {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}, nil
}

// GetSetTargetHumidityCommand generates a ControlRequest to set the target humidity of the HumidifierDevice.
func (device *HumidifierDevice) GetSetTargetHumidityCommand(targetHumidity int) (*ControlRequest, error) {
	if targetHumidity < 0 || targetHumidity > 100 {
		return nil, fmt.Errorf("invalid mode: %d", targetHumidity)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", targetHumidity),
	}, nil
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

// GetSetModeCommand generates a ControlRequest to set the mode of the EvaporativeHumidifierDevice.
func (device *EvaporativeHumidifierDevice) GetSetModeCommand(mode EvaporativeHumidifierMode, targetHumidity int) (*ControlRequest, error) {
	if mode < 1 || mode > 8 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	if targetHumidity < 0 || targetHumidity > 100 {
		return nil, fmt.Errorf("invalid targetHumidity: %d", targetHumidity)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   "default",
	}, nil
}

// GetSetChildLockCommand generates a ControlRequest to set the child lock of the EvaporativeHumidifierDevice.
func (device *EvaporativeHumidifierDevice) GetSetChildLockCommand(flag bool) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "setChildLock",
		Parameter:   fmt.Sprintf("%t", flag),
	}
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

// GetSetModeCommand generates a ControlRequest to set the mode of the AirPurifierDevice.
func (device *AirPurifierDevice) GetSetModeCommand(mode AirPurifierMode, fanLevel int) (*ControlRequest, error) {
	if mode < 1 || mode > 4 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	if fanLevel < 1 || fanLevel > 3 {
		return nil, fmt.Errorf("invalid fanLevel: %d", fanLevel)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter: AirPurifierModeParameter{
			Mode:    mode,
			FanGear: fanLevel,
		},
	}, nil
}

// GetSetChildLockCommand generates a ControlRequest to set the child lock of the AirPurifierDevice.
func (device *AirPurifierDevice) GetSetChildLockCommand(flag bool) *ControlRequest {
	flagInt := 0
	if flag {
		flagInt = 1
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setChildLock",
		Parameter:   flagInt,
	}
}

// GetSetPositionCommand generates a ControlRequest to set the position of the BlindTiltDevice.
func (device *BlindTiltDevice) GetSetPositionCommand(direction string, position int) (*ControlRequest, error) {
	if direction != "up" && direction != "down" {
		return nil, fmt.Errorf("invalid direction: %s", direction)
	}
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}
	if position%2 != 0 {
		return nil, fmt.Errorf("position must be even: %d", position)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		Parameter:   fmt.Sprintf("%s;%d", direction, position),
	}, nil
}

// GetFullyOpenCommand generates a ControlRequest to fully open the BlindTiltDevice.
func (device *BlindTiltDevice) GetFullyOpenCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "fullyOpen",
		Parameter:   "default",
	}
}

// GetCloseUpCommand generates a ControlRequest to close up the BlindTiltDevice.
func (device *BlindTiltDevice) GetCloseUpCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "closeUp",
		Parameter:   "default",
	}
}

// GetCloseDownCommand generates a ControlRequest to close down the BlindTiltDevice.
func (device *BlindTiltDevice) GetCloseDownCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "closeDown",
		Parameter:   "default",
	}
}

type CirculatorNightLightMode string

const (
	CirculatorNightLightModeTurnOff    = CirculatorNightLightMode("off")
	CirculatorNightLightModeTurnBright = CirculatorNightLightMode("1")
	CirculatorNightLightModeTurnDim    = CirculatorNightLightMode("2")
)

// GetSetNightLightModeCommand generates a ControlRequest to set the night light mode of the BatteryCirculatorFanDevice.
func (device *BatteryCirculatorFanDevice) GetSetNightLightModeCommand(mode CirculatorNightLightMode) (*ControlRequest, error) {
	if mode != CirculatorNightLightModeTurnOff && mode != CirculatorNightLightModeTurnBright && mode != CirculatorNightLightModeTurnDim {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setNightLightMode",
		Parameter:   mode,
	}, nil
}

type CirculatorWindMode string

const (
	CirculatorWindModeDirect  = CirculatorWindMode("direct")
	CirculatorWindModeNatural = CirculatorWindMode("natural")
	CirculatorWindModeSleep   = CirculatorWindMode("sleep")
	CirculatorWindModeBaby    = CirculatorWindMode("baby")
)

// GetSetWindModeCommand generates a ControlRequest to set the wind mode of the BatteryCirculatorFanDevice.
func (device *BatteryCirculatorFanDevice) GetSetWindModeCommand(mode CirculatorWindMode) (*ControlRequest, error) {
	if mode != CirculatorWindModeDirect && mode != CirculatorWindModeNatural && mode != CirculatorWindModeSleep && mode != CirculatorWindModeBaby {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setWindMode",
		Parameter:   mode,
	}, nil
}

// GetSetWindSpeedCommand generates a ControlRequest to set the wind speed of the BatteryCirculatorFanDevice.
func (device *BatteryCirculatorFanDevice) GetSetWindSpeedCommand(speed int) (*ControlRequest, error) {
	if speed < 1 || speed > 100 {
		return nil, fmt.Errorf("invalid speed: %d", speed)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setWindSpeed",
		Parameter:   fmt.Sprintf("%d", speed),
	}, nil
}

// GetSetNightLightModeCommand generates a ControlRequest to set the night light mode of the CirculatorFanDevice.
func (device *CirculatorFanDevice) GetSetNightLightModeCommand(mode CirculatorNightLightMode) (*ControlRequest, error) {
	if mode != CirculatorNightLightModeTurnOff && mode != CirculatorNightLightModeTurnBright && mode != CirculatorNightLightModeTurnDim {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setNightLightMode",
		Parameter:   mode,
	}, nil
}

// GetSetWindModeCommand generates a ControlRequest to set the wind mode of the CirculatorFanDevice.
func (device *CirculatorFanDevice) GetSetWindModeCommand(mode CirculatorWindMode) (*ControlRequest, error) {
	if mode != CirculatorWindModeDirect && mode != CirculatorWindModeNatural && mode != CirculatorWindModeSleep && mode != CirculatorWindModeBaby {
		return nil, fmt.Errorf("invalid mode: %s", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setWindMode",
		Parameter:   mode,
	}, nil
}

// GetSetWindSpeedCommand generates a ControlRequest to set the wind speed of the CirculatorFanDevice.
func (device *CirculatorFanDevice) GetSetWindSpeedCommand(speed int) (*ControlRequest, error) {
	if speed < 1 || speed > 100 {
		return nil, fmt.Errorf("invalid speed: %d", speed)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setWindSpeed",
		Parameter:   fmt.Sprintf("%d", speed),
	}, nil
}

// GetSetPositionCommand generates a ControlRequest to set the position of the RollerShadeDevice.
func (device *RollerShadeDevice) GetSetPositionCommand(position int) (*ControlRequest, error) {
	if position < 0 || position > 100 {
		return nil, fmt.Errorf("invalid position: %d", position)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setPosition",
		Parameter:   fmt.Sprintf("%d", position),
	}, nil
}

type RelaySwitchMode int

const (
	RelaySwitchModeToggle    = RelaySwitchMode(0)
	RelaySwitchModeEdge      = RelaySwitchMode(1)
	RelaySwitchModeDetached  = RelaySwitchMode(2)
	RelaySwitchModeMomentary = RelaySwitchMode(3)
)

// GetSetModeCommand generates a ControlRequest to set the mode of the RelaySwitch1PMDevice.
func (device *RelaySwitch1PMDevice) GetSetModeCommand(mode RelaySwitchMode) (*ControlRequest, error) {
	if mode < 0 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}, nil
}

// GetSetModeCommand generates a ControlRequest to set the mode of the RelaySwitch1Device.
func (device *RelaySwitch1Device) GetSetModeCommand(mode RelaySwitchMode) (*ControlRequest, error) {
	if mode < 0 || mode > 3 {
		return nil, fmt.Errorf("invalid mode: %d", mode)
	}
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMode",
		Parameter:   fmt.Sprintf("%d", mode),
	}, nil
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

// GetSetAllCommand generates a ControlRequest to set the mode, fan speed, and power state of the InfraredRemoteAirConditionerDevice.
func (device *InfraredRemoteAirConditionerDevice) GetSetAllCommand(
	temperatureCelsius int, mode AirConditionerMode, fan AirConditionerFanMode, powerState AirConditionerPowerState,
) (*ControlRequest, error) {
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
	return &ControlRequest{
		CommandType: "command",
		Command:     "setAll",
		Parameter:   parameter,
	}, nil
}

// GetSetChannelCommand generates a ControlRequest to set the channel of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice.
func (device *InfraredRemoteTVDevice) GetSetChannelCommand(channel int) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "SetChannel",
		Parameter:   fmt.Sprintf("%d", channel),
	}
}

// GetVolumeAddCommand generates a ControlRequest to increase the volume of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice.
func (device *InfraredRemoteTVDevice) GetVolumeAddCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "volumeAdd",
		Parameter:   "default",
	}
}

// GetVolumeSubCommand generates a ControlRequest to decrease the volume of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice.
func (device *InfraredRemoteTVDevice) GetVolumeSubCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "volumeSub",
		Parameter:   "default",
	}
}

// GetChannelAddCommand generates a ControlRequest to increase the channel of the InfraredRemoteTVDevice / InfraredRemoteStreamerDevice / InfraredRemoteSetTopBoxDevice.
func (device *InfraredRemoteTVDevice) GetChannelAddCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "channelAdd",
		Parameter:   "default",
	}
}

// GetSetMuteCommand generates a ControlRequest to mute/unmute the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetSetMuteCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "setMute",
		Parameter:   "default",
	}
}

// GetFastForwardCommand generates a ControlRequest to fast-forward the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetFastForwardCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "FastForward",
		Parameter:   "default",
	}
}

// GetRewindCommand generates a ControlRequest to rewind the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetRewindCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Rewind",
		Parameter:   "default",
	}
}

// GetNextCommand generates a ControlRequest to play the next track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetNextCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Next",
		Parameter:   "default",
	}
}

// GetPreviousCommand generates a ControlRequest to play the previous track on the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetPreviousCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Previous",
		Parameter:   "default",
	}
}

// GetPauseCommand generates a ControlRequest to pause the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetPauseCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Pause",
		Parameter:   "default",
	}
}

// GetPlayCommand generates a ControlRequest to play/resume the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetPlayCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Play",
		Parameter:   "default",
	}
}

// GetStopCommand generates a ControlRequest to stop the InfraredRemoteDvdPlayerDevice / InfraredRemoteSpeakerDevice
func (device *InfraredRemoteDvdPlayerDevice) GetStopCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "Stop",
		Parameter:   "default",
	}
}

// GetVolumeAddCommand generates a ControlRequest to increase the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) GetVolumeAddCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "volumeAdd",
		Parameter:   "default",
	}
}

// GetVolumeSubCommand generates a ControlRequest to decrease the volume of the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) GetVolumeSubCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "volumeSub",
		Parameter:   "default",
	}
}

// GetSwingCommand generates a ControlRequest to enable/disable the swing feature of the InfraredRemoteFanDevice.
func (device *InfraredRemoteFanDevice) GetSwingCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "swing",
		Parameter:   "default",
	}
}

// GetTimerCommand generates a ControlRequest to set the timer of the InfraredRemoteFanDevice.
func (device *InfraredRemoteFanDevice) GetTimerCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "timer",
		Parameter:   "default",
	}
}

// GetLowSpeedCommand generates a ControlRequest to set the fan speed to low on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) GetLowSpeedCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "lowSpeed",
		Parameter:   "default",
	}
}

// GetMiddleSpeedCommand generates a ControlRequest to set the fan speed to middle on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) GetMiddleSpeedCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "middleSpeed",
		Parameter:   "default",
	}
}

// GetHighSpeedCommand generates a ControlRequest to set the fan speed to high on the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) GetHighSpeedCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "highSpeed",
		Parameter:   "default",
	}
}

// GetBrightnessUpCommand generates a ControlRequest to increase the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) GetBrightnessUpCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "brightnessUp",
		Parameter:   "default",
	}
}

// GetBrightnessDownCommand generates a ControlRequest to decrease the brightness of the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) GetBrightnessDownCommand() *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "brightnessDown",
		Parameter:   "default",
	}
}

// GetCustomCommand generates a ControlRequest to send a custom command to the InfraredRemoteOthersDevice.
func (device *InfraredRemoteOthersDevice) GetCustomCommand(buttonName string) *ControlRequest {
	return &ControlRequest{
		CommandType: "command",
		Command:     "customize",
		Parameter:   buttonName,
	}
}

// SendCommand sends a command to the device with the specified deviceId.
func (client *Client) SendCommand(deviceId string, request *ControlRequest) (*CommonResponse, error) {
	return client.PostRequest("/devices/"+deviceId+"/commands", request)
}
