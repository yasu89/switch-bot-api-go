package switchbot

import (
	"encoding/json"
	"fmt"
	"image/color"

	jsonschemaValidation "github.com/kaptinlin/jsonschema"
	"github.com/swaggest/jsonschema-go"
)

type ExecutableCommandDevice interface {
	GetCommandParameterJSONSchema() (string, error)
	ExecCommand(jsonString string) (*CommonResponse, error)
}

// validateAndUnmarshalJSON validates the JSON string against the schema and unmarshal it into the target
func validateAndUnmarshalJSON(device ExecutableCommandDevice, jsonString string, target interface{}) error {
	schemaJSON, err := device.GetCommandParameterJSONSchema()
	if err != nil {
		return err
	}

	compiler := jsonschemaValidation.NewCompiler()
	schema, err := compiler.Compile([]byte(schemaJSON))
	if err != nil {
		return err
	}

	var instance map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &instance)
	if err != nil {
		return err
	}

	result := schema.Validate(instance)
	if !result.IsValid() {
		errorDetails, _ := json.Marshal(result.ToList())
		return fmt.Errorf("invalid command parameter: %s", string(errorDetails))
	}

	return json.Unmarshal([]byte(jsonString), target)
}

// reflectJSONSchema returns the JSON schema for the given parameter
func reflectJSONSchema(parameter interface{}) (string, error) {
	reflector := jsonschema.Reflector{}
	schema, err := reflector.Reflect(parameter, jsonschema.InlineRefs)
	if err != nil {
		return "", err
	}
	jsonString, err := json.Marshal(schema)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

// BotDeviceCommandParameter is a struct that represents the command parameter for the BotDevice
type BotDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Press" description:"TurnOn:set to OFF state, TurnOff:set to ON state, Press:trigger press" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the BotDevice
func (device *BotDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter BotDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Press":
		return device.Press()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the BotDevice command parameter
func (device *BotDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(BotDeviceCommandParameter{})
}

// CurtainDeviceCommandParameter is a struct that represents the command parameter for the CurtainDevice
type CurtainDeviceCommandParameter struct {
	Command  string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Pause,SetPosition" description:"TurnOn:equivalent to set position to 100, TurnOff:equivalent to set position to 0, Pause:set to PAUSE state, SetPosition:set position" required:"true"`
	Mode     string   `json:"mode" title:"Mode" enum:"0,1,ff" description:"0:performance mode, 1:silent mode, ff:default mode"`
	Position int      `json:"position" title:"Position" minimum:"0" maximum:"100"`
	_        struct{} `additionalProperties:"false"`
}

// JSONSchemaIf returns the JSON schema if block for the CurtainDevice command parameter
func (parameter *CurtainDeviceCommandParameter) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetPosition" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the CurtainDevice command parameter
func (parameter *CurtainDeviceCommandParameter) JSONSchemaThen() interface{} {
	return struct {
		Mode     string `json:"mode" title:"Mode" enum:"0,1,ff" description:"0:performance mode, 1:silent mode, ff:default mode" required:"true"`
		Position int    `json:"position" title:"Position" minimum:"0" maximum:"100" required:"true"`
	}{}
}

// ExecCommand sends a command to the CurtainDevice
func (device *CurtainDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter CurtainDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Pause":
		return device.Pause()
	case "SetPosition":
		return device.SetPosition(CurtainPositionMode(parameter.Mode), parameter.Position)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the CurtainDevice command parameter
func (device *CurtainDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(CurtainDeviceCommandParameter{})
}

// LockDeviceCommandParameter is a struct that represents the command parameter for the LockDevice
type LockDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"Lock,Unlock" description:"Lock:rotate to locked position, Unlock:rotate to unlocked position" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the LockDevice
func (device *LockDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter LockDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "Lock":
		return device.Lock()
	case "Unlock":
		return device.Unlock()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the LockDevice command parameter
func (device *LockDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(LockDeviceCommandParameter{})
}

// KeypadDeviceCommandParameter is a struct that represents the command parameter for the KeypadDevice
type KeypadDeviceCommandParameter struct {
	Command   string   `json:"command" title:"Command" enum:"CreateKey,DeleteKey" required:"true" description:"CreateKey:create a new passcode, DeleteKey:delete an existing passcode"`
	Id        string   `json:"id" title:"ID" description:"the id of the passcode"`
	Name      string   `json:"name" title:"Name" description:"a unique name for the passcode"`
	Type      string   `json:"type" title:"Type" enum:"permanent,timeLimit,disposable,urgent" description:"type of the passcode. permanent, a permanent passcode. timeLimit, a temporary passcode. disposable, a one-time passcode. urgent, an emergency passcode."`
	Password  string   `json:"password" title:"Password" pattern:"^[0-9]{6,12}$" description:"a 6 to 12-digit passcode in plain text"`
	StartTime int64    `json:"startTime" title:"StartTime" minimum:"100000000" maximum:"9999999999" description:"set the time the passcode becomes valid from, mandatory for one-time passcode and temporary passcode. a 10-digit timestamp(Unix timestamp)."`
	EndTime   int64    `json:"endTime" title:"EndTime" minimum:"100000000" maximum:"9999999999" description:"set the time the passcode becomes expired, mandatory for one-time passcode and temporary passcode. a 10-digit timestamp(Unix timestamp)."`
	_         struct{} `additionalProperties:"false"`
}

// KeypadDeviceCommandCreateNormalKeyIfExposer represents the CreateKey command parameters for types: permanent and urgent.
type KeypadDeviceCommandCreateNormalKeyIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the KeypadDeviceCommandCreateNormalKeyIfExposer parameter
func (parameter *KeypadDeviceCommandCreateNormalKeyIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"CreateKey" required:"true"`
		Type    string `json:"type" enum:"permanent,urgent" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the KeypadDeviceCommandCreateNormalKeyIfExposer parameter
func (parameter *KeypadDeviceCommandCreateNormalKeyIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Name     string `json:"name" required:"true"`
		Password string `json:"password" required:"true"`
	}{}
}

// KeypadDeviceCommandCreateTimeLimitKeyIfExposer represents the CreateKey command parameters for type: timeLimit and disposable
type KeypadDeviceCommandCreateTimeLimitKeyIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the KeypadDeviceCommandCreateTimeLimitKeyIfExposer parameter
func (parameter *KeypadDeviceCommandCreateTimeLimitKeyIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"CreateKey" required:"true"`
		Type    string `json:"type" enum:"timeLimit,disposable" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the KeypadDeviceCommandCreateTimeLimitKeyIfExposer parameter
func (parameter *KeypadDeviceCommandCreateTimeLimitKeyIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Name      string `json:"name" required:"true"`
		Password  string `json:"password" required:"true"`
		StartTime int64  `json:"startTime" required:"true"`
		EndTime   int64  `json:"endTime" required:"true"`
	}{}
}

// KeypadDeviceCommandDeleteKeyIfExposer represents the DeleteKey command parameters
type KeypadDeviceCommandDeleteKeyIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the KeypadDeviceCommandCreateNormalKeyIfExposer parameter
func (parameter *KeypadDeviceCommandDeleteKeyIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"DeleteKey" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the KeypadDeviceCommandCreateNormalKeyIfExposer parameter
func (parameter *KeypadDeviceCommandDeleteKeyIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Id string `json:"id" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the KeypadDevice command parameter
func (parameter *KeypadDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&KeypadDeviceCommandCreateNormalKeyIfExposer{},
		&KeypadDeviceCommandCreateTimeLimitKeyIfExposer{},
		&KeypadDeviceCommandDeleteKeyIfExposer{},
	}
}

// ExecCommand sends a command to the KeypadDevice
func (device *KeypadDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter KeypadDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "CreateKey":
		key, err := NewKeypadKey(parameter.Name, parameter.Type, parameter.Password, parameter.StartTime, parameter.EndTime)
		if err != nil {
			return nil, err
		}
		return device.CreateKey(key)
	case "DeleteKey":
		return device.DeleteKey(parameter.Id)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the KeypadDevice command parameter
func (device *KeypadDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(KeypadDeviceCommandParameter{})
}

// CeilingLightDeviceCommandParameter is a struct that represents the command parameter for the CeilingLightDevice
type CeilingLightDeviceCommandParameter struct {
	Command          string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Toggle,SetBrightness,SetColorTemperature" description:"TurnOn:turn on the ceiling light, TurnOff:turn off the ceiling light, Toggle:toggle the ceiling light, SetBrightness:set brightness, SetColorTemperature:set color temperature" required:"true"`
	Brightness       int      `json:"brightness" title:"Brightness" minimum:"1" maximum:"100" description:"Brightness level (1-100)"`
	ColorTemperature int      `json:"colorTemperature" title:"ColorTemperature" minimum:"2700" maximum:"6500" description:"Color temperature in Kelvin (2700-6500)"`
	_                struct{} `additionalProperties:"false"`
}

// CeilingLightDeviceCommandSetBrightnessIfExposer represents the SetBrightness command parameters
type CeilingLightDeviceCommandSetBrightnessIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the CeilingLightDevice command parameter for SetBrightness
func (parameter *CeilingLightDeviceCommandSetBrightnessIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetBrightness" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the CeilingLightDevice command parameter for SetBrightness
func (parameter *CeilingLightDeviceCommandSetBrightnessIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Brightness int `json:"brightness" required:"true"`
	}{}
}

// CeilingLightDeviceCommandSetColorTemperatureIfExposer represents the SetColorTemperature command parameters
type CeilingLightDeviceCommandSetColorTemperatureIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the CeilingLightDevice command parameter for SetColorTemperature
func (parameter *CeilingLightDeviceCommandSetColorTemperatureIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetColorTemperature" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the CeilingLightDevice command parameter for SetColorTemperature
func (parameter *CeilingLightDeviceCommandSetColorTemperatureIfExposer) JSONSchemaThen() interface{} {
	return struct {
		ColorTemperature int `json:"colorTemperature" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the CeilingLightDevice command parameter
func (parameter *CeilingLightDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&CeilingLightDeviceCommandSetBrightnessIfExposer{},
		&CeilingLightDeviceCommandSetColorTemperatureIfExposer{},
	}
}

// ExecCommand sends a command to the CeilingLightDevice
func (device *CeilingLightDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter CeilingLightDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	case "SetBrightness":
		return device.SetBrightness(parameter.Brightness)
	case "SetColorTemperature":
		return device.SetColorTemperature(parameter.ColorTemperature)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the CeilingLightDevice command parameter
func (device *CeilingLightDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(CeilingLightDeviceCommandParameter{})
}

// PlugMiniDeviceCommandParameter is a struct that represents the command parameter for the PlugMiniDevice
type PlugMiniDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Toggle" description:"TurnOn:turn on the plug, TurnOff:turn off the plug, Toggle:toggle the plug state" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the PlugMiniDevice
func (device *PlugMiniDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter PlugMiniDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the PlugMiniDevice command parameter
func (device *PlugMiniDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(PlugMiniDeviceCommandParameter{})
}

// PlugDeviceCommandParameter is a struct that represents the command parameter for the PlugDevice
type PlugDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff" description:"TurnOn:turn on the plug, TurnOff:turn off the plug" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the PlugDevice
func (device *PlugDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter PlugDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the PlugDevice command parameter
func (device *PlugDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(PlugDeviceCommandParameter{})
}

// StripLightDeviceCommandParameter is a struct that represents the command parameter for the StripLightDevice
type StripLightDeviceCommandParameter struct {
	Command    string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Toggle,SetBrightness,SetColor" description:"TurnOn:turn on the strip light, TurnOff:turn off the strip light, Toggle:toggle the strip light, SetBrightness:set brightness, SetColor:set color" required:"true"`
	Brightness int      `json:"brightness" title:"Brightness" minimum:"1" maximum:"100" description:"Brightness level (1-100)"`
	Red        int      `json:"red" title:"Red" minimum:"0" maximum:"255" description:"Red color value (0-255)"`
	Green      int      `json:"green" title:"Green" minimum:"0" maximum:"255" description:"Green color value (0-255)"`
	Blue       int      `json:"blue" title:"Blue" minimum:"0" maximum:"255" description:"Blue color value (0-255)"`
	_          struct{} `additionalProperties:"false"`
}

// StripLightDeviceCommandSetBrightnessIfExposer represents the SetBrightness command parameters
type StripLightDeviceCommandSetBrightnessIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the StripLightDevice command parameter for SetBrightness
func (parameter *StripLightDeviceCommandSetBrightnessIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetBrightness" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the StripLightDevice command parameter for SetBrightness
func (parameter *StripLightDeviceCommandSetBrightnessIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Brightness int `json:"brightness" required:"true"`
	}{}
}

// StripLightDeviceCommandSetColorIfExposer represents the SetColor command parameters
type StripLightDeviceCommandSetColorIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the StripLightDevice command parameter for SetColor
func (parameter *StripLightDeviceCommandSetColorIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetColor" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the StripLightDevice command parameter for SetColor
func (parameter *StripLightDeviceCommandSetColorIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Red   int `json:"red" required:"true"`
		Green int `json:"green" required:"true"`
		Blue  int `json:"blue" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the StripLightDevice command parameter
func (parameter *StripLightDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&StripLightDeviceCommandSetBrightnessIfExposer{},
		&StripLightDeviceCommandSetColorIfExposer{},
	}
}

// ExecCommand sends a command to the StripLightDevice
func (device *StripLightDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter StripLightDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	case "SetBrightness":
		return device.SetBrightness(parameter.Brightness)
	case "SetColor":
		return device.SetColor(color.RGBA{R: uint8(parameter.Red), G: uint8(parameter.Green), B: uint8(parameter.Blue), A: 255})
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the StripLightDevice command parameter
func (device *StripLightDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(StripLightDeviceCommandParameter{})
}

// ColorBulbDeviceCommandParameter is a struct that represents the command parameter for the ColorBulbDevice
type ColorBulbDeviceCommandParameter struct {
	Command          string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Toggle,SetBrightness,SetColor,SetColorTemperature" description:"TurnOn:turn on the bulb, TurnOff:turn off the bulb, Toggle:toggle the bulb, SetBrightness:set brightness, SetColor:set color, SetColorTemperature:set color temperature" required:"true"`
	Brightness       int      `json:"brightness" title:"Brightness" minimum:"1" maximum:"100" description:"Brightness level (1-100)"`
	Red              int      `json:"red" title:"Red" minimum:"0" maximum:"255" description:"Red color value (0-255)"`
	Green            int      `json:"green" title:"Green" minimum:"0" maximum:"255" description:"Green color value (0-255)"`
	Blue             int      `json:"blue" title:"Blue" minimum:"0" maximum:"255" description:"Blue color value (0-255)"`
	ColorTemperature int      `json:"colorTemperature" title:"ColorTemperature" minimum:"2700" maximum:"6500" description:"Color temperature in Kelvin (2700-6500)"`
	_                struct{} `additionalProperties:"false"`
}

// ColorBulbDeviceCommandSetBrightnessIfExposer represents the SetBrightness command parameters
type ColorBulbDeviceCommandSetBrightnessIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the ColorBulbDevice command parameter for SetBrightness
func (parameter *ColorBulbDeviceCommandSetBrightnessIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetBrightness" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the ColorBulbDevice command parameter for SetBrightness
func (parameter *ColorBulbDeviceCommandSetBrightnessIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Brightness int `json:"brightness" required:"true"`
	}{}
}

// ColorBulbDeviceCommandSetColorIfExposer represents the SetColor command parameters
type ColorBulbDeviceCommandSetColorIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the ColorBulbDevice command parameter for SetColor
func (parameter *ColorBulbDeviceCommandSetColorIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetColor" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the ColorBulbDevice command parameter for SetColor
func (parameter *ColorBulbDeviceCommandSetColorIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Red   int `json:"red" required:"true"`
		Green int `json:"green" required:"true"`
		Blue  int `json:"blue" required:"true"`
	}{}
}

// ColorBulbDeviceCommandSetColorTemperatureIfExposer represents the SetColorTemperature command parameters
type ColorBulbDeviceCommandSetColorTemperatureIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the ColorBulbDevice command parameter for SetColorTemperature
func (parameter *ColorBulbDeviceCommandSetColorTemperatureIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetColorTemperature" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the ColorBulbDevice command parameter for SetColorTemperature
func (parameter *ColorBulbDeviceCommandSetColorTemperatureIfExposer) JSONSchemaThen() interface{} {
	return struct {
		ColorTemperature int `json:"colorTemperature" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the ColorBulbDevice command parameter
func (parameter *ColorBulbDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&ColorBulbDeviceCommandSetBrightnessIfExposer{},
		&ColorBulbDeviceCommandSetColorIfExposer{},
		&ColorBulbDeviceCommandSetColorTemperatureIfExposer{},
	}
}

// ExecCommand sends a command to the ColorBulbDevice
func (device *ColorBulbDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter ColorBulbDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	case "SetBrightness":
		return device.SetBrightness(parameter.Brightness)
	case "SetColor":
		return device.SetColor(color.RGBA{R: uint8(parameter.Red), G: uint8(parameter.Green), B: uint8(parameter.Blue), A: 255})
	case "SetColorTemperature":
		return device.SetColorTemperature(parameter.ColorTemperature)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the ColorBulbDevice command parameter
func (device *ColorBulbDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(ColorBulbDeviceCommandParameter{})
}

// RobotVacuumCleanerDeviceCommandParameter is a struct that represents the command parameter for the RobotVacuumCleanerDevice
type RobotVacuumCleanerDeviceCommandParameter struct {
	Command    string   `json:"command" title:"Command" enum:"Start,Stop,Dock,SetPowerLevel" description:"Start:start vacuuming, Stop:stop vacuuming, Dock:return to charging dock, SetPowerLevel:set the suction power level" required:"true"`
	PowerLevel int      `json:"powerLevel" title:"PowerLevel" minimum:"0" maximum:"3" description:"Power level: 0:Quiet, 1:Standard, 2:Strong, 3:Max"`
	_          struct{} `additionalProperties:"false"`
}

// RobotVacuumCleanerDeviceCommandSetPowerLevelIfExposer represents the SetPowerLevel command parameters
type RobotVacuumCleanerDeviceCommandSetPowerLevelIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RobotVacuumCleanerDevice command parameter for SetPowerLevel
func (parameter *RobotVacuumCleanerDeviceCommandSetPowerLevelIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetPowerLevel" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RobotVacuumCleanerDevice command parameter for SetPowerLevel
func (parameter *RobotVacuumCleanerDeviceCommandSetPowerLevelIfExposer) JSONSchemaThen() interface{} {
	return struct {
		PowerLevel int `json:"powerLevel" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the RobotVacuumCleanerDevice command parameter
func (parameter *RobotVacuumCleanerDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&RobotVacuumCleanerDeviceCommandSetPowerLevelIfExposer{},
	}
}

// ExecCommand sends a command to the RobotVacuumCleanerDevice
func (device *RobotVacuumCleanerDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter RobotVacuumCleanerDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "Start":
		return device.Start()
	case "Stop":
		return device.Stop()
	case "Dock":
		return device.Dock()
	case "SetPowerLevel":
		return device.SetPowerLevel(RobotVacuumCleanerPowerLevel(parameter.PowerLevel))
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the RobotVacuumCleanerDevice command parameter
func (device *RobotVacuumCleanerDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(RobotVacuumCleanerDeviceCommandParameter{})
}

// RobotVacuumCleanerS10DeviceCommandParameter is a struct that represents the command parameter for the RobotVacuumCleanerS10Device
type RobotVacuumCleanerS10DeviceCommandParameter struct {
	Command    string   `json:"command" title:"Command" enum:"StartClean,AddWaterForHumi,Pause,Dock,SetVolume,SelfClean,ChangeParam" description:"StartClean:start cleaning, AddWaterForHumi:refill the humidifier, Pause:pause cleaning, Dock:return to charging dock, SetVolume:set volume level, SelfClean:start self-cleaning, ChangeParam:change cleaning parameters" required:"true"`
	Action     string   `json:"action" title:"Action" enum:"sweep,sweep_mop" description:"sweep:sweep only, sweep_mop:sweep and mop"`
	FanLevel   int      `json:"fanLevel" title:"FanLevel" minimum:"1" maximum:"4" description:"Fan level (1-4)"`
	WaterLevel int      `json:"waterLevel" title:"WaterLevel" minimum:"1" maximum:"2" description:"Water level (1-2)"`
	Times      int      `json:"times" title:"Times" minimum:"1" maximum:"2639999" description:"the number of cycles"`
	Volume     int      `json:"volume" title:"Volume" minimum:"0" maximum:"100" description:"Volume level (0-100)"`
	Mode       int      `json:"mode" title:"Mode" minimum:"1" maximum:"3" description:"Self-cleaning mode: 1:wash mop, 2:dry, 3:terminate"`
	_          struct{} `additionalProperties:"false"`
}

// RobotVacuumCleanerS10DeviceCommandStartCleanIfExposer represents the StartClean command parameters
type RobotVacuumCleanerS10DeviceCommandStartCleanIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RobotVacuumCleanerS10Device command parameter for StartClean
func (parameter *RobotVacuumCleanerS10DeviceCommandStartCleanIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"StartClean" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RobotVacuumCleanerS10Device command parameter for StartClean
func (parameter *RobotVacuumCleanerS10DeviceCommandStartCleanIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Action     string `json:"action" required:"true"`
		FanLevel   int    `json:"fanLevel" required:"true"`
		WaterLevel int    `json:"waterLevel" required:"true"`
		Times      int    `json:"times" required:"true"`
	}{}
}

// RobotVacuumCleanerS10DeviceCommandSetVolumeIfExposer represents the SetVolume command parameters
type RobotVacuumCleanerS10DeviceCommandSetVolumeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RobotVacuumCleanerS10Device command parameter for SetVolume
func (parameter *RobotVacuumCleanerS10DeviceCommandSetVolumeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetVolume" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RobotVacuumCleanerS10Device command parameter for SetVolume
func (parameter *RobotVacuumCleanerS10DeviceCommandSetVolumeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Volume int `json:"volume" required:"true"`
	}{}
}

// RobotVacuumCleanerS10DeviceCommandSelfCleanIfExposer represents the SelfClean command parameters
type RobotVacuumCleanerS10DeviceCommandSelfCleanIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RobotVacuumCleanerS10Device command parameter for SelfClean
func (parameter *RobotVacuumCleanerS10DeviceCommandSelfCleanIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SelfClean" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RobotVacuumCleanerS10Device command parameter for SelfClean
func (parameter *RobotVacuumCleanerS10DeviceCommandSelfCleanIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode int `json:"mode" required:"true"`
	}{}
}

// RobotVacuumCleanerS10DeviceCommandChangeParamIfExposer represents the ChangeParam command parameters
type RobotVacuumCleanerS10DeviceCommandChangeParamIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RobotVacuumCleanerS10Device command parameter for ChangeParam
func (parameter *RobotVacuumCleanerS10DeviceCommandChangeParamIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"ChangeParam" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RobotVacuumCleanerS10Device command parameter for ChangeParam
func (parameter *RobotVacuumCleanerS10DeviceCommandChangeParamIfExposer) JSONSchemaThen() interface{} {
	return struct {
		FanLevel   int `json:"fanLevel" required:"true"`
		WaterLevel int `json:"waterLevel" required:"true"`
		Times      int `json:"times" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the RobotVacuumCleanerS10Device command parameter
func (parameter *RobotVacuumCleanerS10DeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&RobotVacuumCleanerS10DeviceCommandStartCleanIfExposer{},
		&RobotVacuumCleanerS10DeviceCommandSetVolumeIfExposer{},
		&RobotVacuumCleanerS10DeviceCommandSelfCleanIfExposer{},
		&RobotVacuumCleanerS10DeviceCommandChangeParamIfExposer{},
	}
}

// ExecCommand sends a command to the RobotVacuumCleanerS10Device
func (device *RobotVacuumCleanerS10Device) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter RobotVacuumCleanerS10DeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "StartClean":
		startParam, err := NewStartFloorCleaningParam(FloorCleaningAction(parameter.Action), parameter.FanLevel, parameter.WaterLevel, parameter.Times)
		if err != nil {
			return nil, err
		}
		return device.StartClean(startParam)
	case "AddWaterForHumi":
		return device.AddWaterForHumi()
	case "Pause":
		return device.Pause()
	case "Dock":
		return device.Dock()
	case "SetVolume":
		return device.SetVolume(parameter.Volume)
	case "SelfClean":
		return device.SelfClean(SelfCleaningMode(parameter.Mode))
	case "ChangeParam":
		floorParam, err := NewFloorCleaningParam(parameter.FanLevel, parameter.WaterLevel, parameter.Times)
		if err != nil {
			return nil, err
		}
		return device.ChangeParam(floorParam)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the RobotVacuumCleanerS10Device command parameter
func (device *RobotVacuumCleanerS10Device) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(RobotVacuumCleanerS10DeviceCommandParameter{})
}

// HumidifierDeviceCommandParameter is a struct that represents the command parameter for the HumidifierDevice
type HumidifierDeviceCommandParameter struct {
	Command        string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetMode,SetTargetHumidity" description:"TurnOn:turn on the humidifier, TurnOff:turn off the humidifier, SetMode:set the mode of the humidifier, SetTargetHumidity:set the target humidity" required:"true"`
	Mode           string   `json:"mode" title:"Mode" enum:"Auto,Low,Medium,High" description:"Auto: auto mode, Low:34%, Medium:67%, High:100%"`
	TargetHumidity int      `json:"targetHumidity" title:"TargetHumidity" minimum:"0" maximum:"100" description:"Target humidity level (0-100%)"`
	_              struct{} `additionalProperties:"false"`
}

// HumidifierDeviceCommandSetModeIfExposer represents the SetMode command parameters
type HumidifierDeviceCommandSetModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the HumidifierDevice command parameter for SetMode
func (parameter *HumidifierDeviceCommandSetModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the HumidifierDevice command parameter for SetMode
func (parameter *HumidifierDeviceCommandSetModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode string `json:"mode" required:"true"`
	}{}
}

// HumidifierDeviceCommandSetTargetHumidityIfExposer represents the SetTargetHumidity command parameters
type HumidifierDeviceCommandSetTargetHumidityIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the HumidifierDevice command parameter for SetTargetHumidity
func (parameter *HumidifierDeviceCommandSetTargetHumidityIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetTargetHumidity" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the HumidifierDevice command parameter for SetTargetHumidity
func (parameter *HumidifierDeviceCommandSetTargetHumidityIfExposer) JSONSchemaThen() interface{} {
	return struct {
		TargetHumidity int `json:"targetHumidity" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the HumidifierDevice command parameter
func (parameter *HumidifierDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&HumidifierDeviceCommandSetModeIfExposer{},
		&HumidifierDeviceCommandSetTargetHumidityIfExposer{},
	}
}

// parseHumidifierMode parses a string into a HumidifierMode
func parseHumidifierMode(s string) (HumidifierMode, error) {
	switch s {
	case "Auto":
		return HumidifierModeAuto, nil
	case "Low":
		return HumidifierModeLow, nil
	case "Medium":
		return HumidifierModeMedium, nil
	case "High":
		return HumidifierModeHigh, nil
	default:
		return 0, fmt.Errorf("invalid humidifier mode: %s", s)
	}
}

// ExecCommand sends a command to the HumidifierDevice
func (device *HumidifierDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter HumidifierDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetMode":
		mode, err := parseHumidifierMode(parameter.Mode)
		if err != nil {
			return nil, err
		}
		return device.SetMode(mode)
	case "SetTargetHumidity":
		return device.SetTargetHumidity(parameter.TargetHumidity)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the HumidifierDevice command parameter
func (device *HumidifierDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(HumidifierDeviceCommandParameter{})
}

// EvaporativeHumidifierDeviceCommandParameter is a struct that represents the command parameter for the EvaporativeHumidifierDevice
type EvaporativeHumidifierDeviceCommandParameter struct {
	Command        string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetMode,SetChildLock" description:"TurnOn:turn on device, TurnOff:turn off device, SetMode:set the mode, SetChildLock:set the child lock" required:"true"`
	Mode           int      `json:"mode" title:"Mode" minimum:"1" maximum:"8" description:"1:Level 4, 2:Level 3, 3:Level 2, 4:Level 1, 5:humidity mode, 6:sleep mode, 7:auto mode, 8:drying mode"`
	TargetHumidity int      `json:"targetHumidity" title:"TargetHumidity" minimum:"0" maximum:"100" description:"Target humidity level (0-100%)"`
	ChildLock      bool     `json:"childLock" title:"ChildLock" description:"true:lock, false:unlock"`
	_              struct{} `additionalProperties:"false"`
}

// EvaporativeHumidifierDeviceCommandSetModeIfExposer represents the SetMode command parameters
type EvaporativeHumidifierDeviceCommandSetModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the EvaporativeHumidifierDevice command parameter for SetMode
func (parameter *EvaporativeHumidifierDeviceCommandSetModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the EvaporativeHumidifierDevice command parameter for SetMode
func (parameter *EvaporativeHumidifierDeviceCommandSetModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode           int `json:"mode" required:"true"`
		TargetHumidity int `json:"targetHumidity" required:"true"`
	}{}
}

// EvaporativeHumidifierDeviceCommandSetChildLockIfExposer represents the SetChildLock command parameters
type EvaporativeHumidifierDeviceCommandSetChildLockIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the EvaporativeHumidifierDevice command parameter for SetChildLock
func (parameter *EvaporativeHumidifierDeviceCommandSetChildLockIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetChildLock" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the EvaporativeHumidifierDevice command parameter for SetChildLock
func (parameter *EvaporativeHumidifierDeviceCommandSetChildLockIfExposer) JSONSchemaThen() interface{} {
	return struct {
		ChildLock bool `json:"childLock" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the EvaporativeHumidifierDevice command parameter
func (parameter *EvaporativeHumidifierDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&EvaporativeHumidifierDeviceCommandSetModeIfExposer{},
		&EvaporativeHumidifierDeviceCommandSetChildLockIfExposer{},
	}
}

// ExecCommand sends a command to the EvaporativeHumidifierDevice
func (device *EvaporativeHumidifierDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter EvaporativeHumidifierDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetMode":
		return device.SetMode(EvaporativeHumidifierMode(parameter.Mode), parameter.TargetHumidity)
	case "SetChildLock":
		return device.SetChildLock(parameter.ChildLock)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the EvaporativeHumidifierDevice command parameter
func (device *EvaporativeHumidifierDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(EvaporativeHumidifierDeviceCommandParameter{})
}

// AirPurifierDeviceCommandParameter is a struct that represents the command parameter for the AirPurifierDevice
type AirPurifierDeviceCommandParameter struct {
	Command   string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetMode,SetChildLock" description:"TurnOn:turn on device, TurnOff:turn off device, SetMode:set the mode, SetChildLock:set the child lock" required:"true"`
	Mode      int      `json:"mode" title:"Mode" minimum:"1" maximum:"4" description:"1:Normal mode, 2:Auto mode, 3:Sleep mode, 4:Manual mode"`
	FanLevel  int      `json:"fanLevel" title:"FanLevel" minimum:"1" maximum:"3" description:"Fan speed level (1-3) for Normal mode"`
	ChildLock bool     `json:"childLock" title:"ChildLock" description:"true:lock, false:unlock"`
	_         struct{} `additionalProperties:"false"`
}

// AirPurifierDeviceCommandSetModeIfExposer represents the SetMode command parameters
type AirPurifierDeviceCommandSetModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the AirPurifierDevice command parameter for SetMode
func (parameter *AirPurifierDeviceCommandSetModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the AirPurifierDevice command parameter for SetMode
func (parameter *AirPurifierDeviceCommandSetModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode int `json:"mode" required:"true"`
	}{}
}

// AirPurifierDeviceCommandSetNormalModeIfExposer represents the SetMode off NormalMode command parameters
type AirPurifierDeviceCommandSetNormalModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the AirPurifierDevice command parameter for SetMode
func (parameter *AirPurifierDeviceCommandSetNormalModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetMode" required:"true"`
		Mode    int    `json:"mode" const:"1" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the AirPurifierDevice command parameter for SetMode
func (parameter *AirPurifierDeviceCommandSetNormalModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode     int `json:"mode" required:"true"`
		FanLevel int `json:"fanLevel" required:"true"`
	}{}
}

// AirPurifierDeviceCommandSetChildLockIfExposer represents the SetChildLock command parameters
type AirPurifierDeviceCommandSetChildLockIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the AirPurifierDevice command parameter for SetChildLock
func (parameter *AirPurifierDeviceCommandSetChildLockIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetChildLock" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the AirPurifierDevice command parameter for SetChildLock
func (parameter *AirPurifierDeviceCommandSetChildLockIfExposer) JSONSchemaThen() interface{} {
	return struct {
		ChildLock bool `json:"childLock" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the AirPurifierDevice command parameter
func (parameter *AirPurifierDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&AirPurifierDeviceCommandSetModeIfExposer{},
		&AirPurifierDeviceCommandSetNormalModeIfExposer{},
		&AirPurifierDeviceCommandSetChildLockIfExposer{},
	}
}

// ExecCommand sends a command to the AirPurifierDevice
func (device *AirPurifierDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter AirPurifierDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetMode":
		return device.SetMode(AirPurifierMode(parameter.Mode), parameter.FanLevel)
	case "SetChildLock":
		return device.SetChildLock(parameter.ChildLock)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the AirPurifierDevice command parameter
func (device *AirPurifierDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(AirPurifierDeviceCommandParameter{})
}

// BlindTiltDeviceCommandParameter is a struct that represents the command parameter for the BlindTiltDevice
type BlindTiltDeviceCommandParameter struct {
	Command   string   `json:"command" title:"Command" enum:"SetPosition,FullyOpen,CloseUp,CloseDown" description:"SetPosition:set the position of the blind, FullyOpen:fully open the blind, CloseUp:close up the blind, CloseDown:close down the blind" required:"true"`
	Direction string   `json:"direction" title:"Direction" enum:"up,down" description:"Direction of the blind (up or down)"`
	Position  int      `json:"position" title:"Position" minimum:"0" maximum:"100" description:"Position value (0-100, must be even number)"`
	_         struct{} `additionalProperties:"false"`
}

// BlindTiltDeviceCommandSetPositionIfExposer represents the SetPosition command parameters
type BlindTiltDeviceCommandSetPositionIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the BlindTiltDevice command parameter for SetPosition
func (parameter *BlindTiltDeviceCommandSetPositionIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetPosition" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the BlindTiltDevice command parameter for SetPosition
func (parameter *BlindTiltDeviceCommandSetPositionIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Direction string `json:"direction" required:"true"`
		Position  int    `json:"position" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the BlindTiltDevice command parameter
func (parameter *BlindTiltDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&BlindTiltDeviceCommandSetPositionIfExposer{},
	}
}

// ExecCommand sends a command to the BlindTiltDevice
func (device *BlindTiltDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter BlindTiltDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "SetPosition":
		if parameter.Position%2 != 0 {
			return nil, fmt.Errorf("position must be even: %d", parameter.Position)
		}
		return device.SetPosition(parameter.Direction, parameter.Position)
	case "FullyOpen":
		return device.FullyOpen()
	case "CloseUp":
		return device.CloseUp()
	case "CloseDown":
		return device.CloseDown()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the BlindTiltDevice command parameter
func (device *BlindTiltDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(BlindTiltDeviceCommandParameter{})
}

// CirculatorFanDeviceCommandParameter is a struct that represents the command parameter for the BatteryCirculatorFanDevice and CirculatorFanDevice
type CirculatorFanDeviceCommandParameter struct {
	Command    string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetNightLightMode,SetWindMode,SetWindSpeed" description:"TurnOn:turn on device, TurnOff:turn off device, SetNightLightMode:set night light mode, SetWindMode:set wind mode, SetWindSpeed:set wind speed" required:"true"`
	NightLight string   `json:"nightLight" title:"NightLight" enum:"off,1,2" description:"Night light mode: off:turn off, 1:bright, 2:dim"`
	WindMode   string   `json:"windMode" title:"WindMode" enum:"direct,natural,sleep,baby" description:"Wind mode: direct:direct wind, natural:natural wind, sleep:sleep wind, baby:ultra quiet mode"`
	WindSpeed  int      `json:"windSpeed" title:"WindSpeed" minimum:"1" maximum:"100" description:"Wind speed (1-100)"`
	_          struct{} `additionalProperties:"false"`
}

// CirculatorFanDeviceCommandSetNightLightModeIfExposer represents the SetNightLightMode command parameters
type CirculatorFanDeviceCommandSetNightLightModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetNightLightMode
func (parameter *CirculatorFanDeviceCommandSetNightLightModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetNightLightMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetNightLightMode
func (parameter *CirculatorFanDeviceCommandSetNightLightModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		NightLight string `json:"nightLight" required:"true"`
	}{}
}

// CirculatorFanDeviceCommandSetWindModeIfExposer represents the SetWindMode command parameters
type CirculatorFanDeviceCommandSetWindModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetWindMode
func (parameter *CirculatorFanDeviceCommandSetWindModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetWindMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetWindMode
func (parameter *CirculatorFanDeviceCommandSetWindModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		WindMode string `json:"windMode" required:"true"`
	}{}
}

// CirculatorFanDeviceCommandSetWindSpeedIfExposer represents the SetWindSpeed command parameters
type CirculatorFanDeviceCommandSetWindSpeedIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetWindSpeed
func (parameter *CirculatorFanDeviceCommandSetWindSpeedIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetWindSpeed" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter for SetWindSpeed
func (parameter *CirculatorFanDeviceCommandSetWindSpeedIfExposer) JSONSchemaThen() interface{} {
	return struct {
		WindSpeed int `json:"windSpeed" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the BatteryCirculatorFanDevice and CirculatorFanDevice command parameter
func (parameter *CirculatorFanDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&CirculatorFanDeviceCommandSetNightLightModeIfExposer{},
		&CirculatorFanDeviceCommandSetWindModeIfExposer{},
		&CirculatorFanDeviceCommandSetWindSpeedIfExposer{},
	}
}

// ExecCommand sends a command to the BatteryCirculatorFanDevice
func (device *BatteryCirculatorFanDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter CirculatorFanDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetNightLightMode":
		return device.SetNightLightMode(CirculatorNightLightMode(parameter.NightLight))
	case "SetWindMode":
		return device.SetWindMode(CirculatorWindMode(parameter.WindMode))
	case "SetWindSpeed":
		return device.SetWindSpeed(parameter.WindSpeed)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the BatteryCirculatorFanDevice command parameter
func (device *BatteryCirculatorFanDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(CirculatorFanDeviceCommandParameter{})
}

// ExecCommand sends a command to the CirculatorFanDevice
func (device *CirculatorFanDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter CirculatorFanDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetNightLightMode":
		return device.SetNightLightMode(CirculatorNightLightMode(parameter.NightLight))
	case "SetWindMode":
		return device.SetWindMode(CirculatorWindMode(parameter.WindMode))
	case "SetWindSpeed":
		return device.SetWindSpeed(parameter.WindSpeed)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the CirculatorFanDevice command parameter
func (device *CirculatorFanDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(CirculatorFanDeviceCommandParameter{})
}

// RollerShadeDeviceCommandParameter is a struct that represents the command parameter for the RollerShadeDevice
type RollerShadeDeviceCommandParameter struct {
	Command  string   `json:"command" title:"Command" enum:"SetPosition" description:"SetPosition:set position" required:"true"`
	Position int      `json:"position" title:"Position" minimum:"0" maximum:"100" description:"Position (0-100)" required:"true"`
	_        struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the RollerShadeDevice
func (device *RollerShadeDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter RollerShadeDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "SetPosition":
		return device.SetPosition(parameter.Position)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the RollerShadeDevice command parameter
func (device *RollerShadeDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(RollerShadeDeviceCommandParameter{})
}

// RelaySwitch1DeviceCommandParameter is a struct that represents the command parameter for the RelaySwitch1Device
type RelaySwitch1DeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Toggle,SetMode" description:"TurnOn:turn on the relay switch, TurnOff:turn off the relay switch, Toggle:toggle the relay switch state, SetMode:set the mode of the relay switch" required:"true"`
	Mode    int      `json:"mode" title:"Mode" minimum:"0" maximum:"3" description:"Mode (0:toggle mode, 1:edge switch mode, 2:detached switch mode, 3:momentary switch mode)"`
	_       struct{} `additionalProperties:"false"`
}

// RelaySwitch1DeviceCommandSetModeIfExposer represents the SetMode command parameters
type RelaySwitch1DeviceCommandSetModeIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the RelaySwitch1Device command parameter for SetMode
func (parameter *RelaySwitch1DeviceCommandSetModeIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetMode" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the RelaySwitch1Device command parameter for SetMode
func (parameter *RelaySwitch1DeviceCommandSetModeIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Mode int `json:"mode" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the RelaySwitch1Device command parameter
func (parameter *RelaySwitch1DeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&RelaySwitch1DeviceCommandSetModeIfExposer{},
	}
}

// ExecCommand sends a command to the RelaySwitch1Device
func (device *RelaySwitch1Device) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter RelaySwitch1DeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	case "SetMode":
		return device.SetMode(RelaySwitchMode(parameter.Mode))
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the RelaySwitch1Device command parameter
func (device *RelaySwitch1Device) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(RelaySwitch1DeviceCommandParameter{})
}

// ExecCommand sends a command to the RelaySwitch1PMDevice
func (device *RelaySwitch1PMDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter RelaySwitch1DeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Toggle":
		return device.Toggle()
	case "SetMode":
		return device.SetMode(RelaySwitchMode(parameter.Mode))
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the RelaySwitch1PMDevice command parameter
func (device *RelaySwitch1PMDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(RelaySwitch1DeviceCommandParameter{})
}

// InfraredRemoteAirConditionerDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteAirConditionerDevice
type InfraredRemoteAirConditionerDeviceCommandParameter struct {
	Command            string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetAll" description:"TurnOn:turn on the air conditioner, TurnOff:turn off the air conditioner, SetAll:configure all parameters of the air conditioner" required:"true"`
	TemperatureCelsius int      `json:"temperatureCelsius" title:"Temperature (Celsius)" minimum:"-10" maximum:"40" description:"Temperature in Celsius (-10 to 40)"`
	Mode               int      `json:"mode" title:"Mode" minimum:"1" maximum:"5" description:"Mode (1:auto, 2:cool, 3:dry, 4:fan, 5:heat)"`
	Fan                int      `json:"fan" title:"Fan" minimum:"1" maximum:"4" description:"Fan mode (1:auto, 2:low, 3:medium, 4:high)"`
	PowerState         string   `json:"powerState" title:"Power State" enum:"on,off" description:"Power state (on/off)"`
	_                  struct{} `additionalProperties:"false"`
}

// InfraredRemoteAirConditionerDeviceCommandSetAllIfExposer represents the SetAll command parameters
type InfraredRemoteAirConditionerDeviceCommandSetAllIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the InfraredRemoteAirConditionerDevice command parameter for SetAll
func (parameter *InfraredRemoteAirConditionerDeviceCommandSetAllIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetAll" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the InfraredRemoteAirConditionerDevice command parameter for SetAll
func (parameter *InfraredRemoteAirConditionerDeviceCommandSetAllIfExposer) JSONSchemaThen() interface{} {
	return struct {
		TemperatureCelsius int    `json:"temperatureCelsius" required:"true"`
		Mode               int    `json:"mode" required:"true"`
		Fan                int    `json:"fan" required:"true"`
		PowerState         string `json:"powerState" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the InfraredRemoteAirConditionerDevice command parameter
func (parameter *InfraredRemoteAirConditionerDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&InfraredRemoteAirConditionerDeviceCommandSetAllIfExposer{},
	}
}

// ExecCommand sends a command to the InfraredRemoteAirConditionerDevice
func (device *InfraredRemoteAirConditionerDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteAirConditionerDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetAll":
		return device.SetAll(
			parameter.TemperatureCelsius,
			AirConditionerMode(parameter.Mode),
			AirConditionerFanMode(parameter.Fan),
			AirConditionerPowerState(parameter.PowerState),
		)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteAirConditionerDevice command parameter
func (device *InfraredRemoteAirConditionerDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteAirConditionerDeviceCommandParameter{})
}

// InfraredRemoteTVDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteTVDevice
type InfraredRemoteTVDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,VolumeAdd,VolumeSub,ChannelAdd,ChannelSub,SetChannel" description:"TurnOn:turn on the TV, TurnOff:turn off the TV, VolumeAdd:increase volume, VolumeSub:decrease volume, ChannelAdd:increase channel, ChannelSub: decrease channel, SetChannel:set specific channel" required:"true"`
	Channel int      `json:"channel" title:"Channel" minimum:"1" description:"Channel number for SetChannel command"`
	_       struct{} `additionalProperties:"false"`
}

// InfraredRemoteTVDeviceCommandSetChannelIfExposer represents the SetChannel command parameters
type InfraredRemoteTVDeviceCommandSetChannelIfExposer struct{}

// JSONSchemaIf returns the JSON schema if block for the InfraredRemoteTVDevice command parameter for SetChannel
func (parameter *InfraredRemoteTVDeviceCommandSetChannelIfExposer) JSONSchemaIf() interface{} {
	return struct {
		Command string `json:"command" const:"SetChannel" required:"true"`
	}{}
}

// JSONSchemaThen returns the JSON schema then block for the InfraredRemoteTVDevice command parameter for SetChannel
func (parameter *InfraredRemoteTVDeviceCommandSetChannelIfExposer) JSONSchemaThen() interface{} {
	return struct {
		Channel int `json:"channel" required:"true"`
	}{}
}

// JSONSchemaAllOf returns the JSON schema allOf block for the InfraredRemoteTVDevice command parameter
func (parameter *InfraredRemoteTVDeviceCommandParameter) JSONSchemaAllOf() []interface{} {
	return []interface{}{
		&InfraredRemoteTVDeviceCommandSetChannelIfExposer{},
	}
}

// ExecCommand sends a command to the InfraredRemoteTVDevice
func (device *InfraredRemoteTVDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteTVDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "VolumeAdd":
		return device.VolumeAdd()
	case "VolumeSub":
		return device.VolumeSub()
	case "ChannelAdd":
		return device.ChannelAdd()
	case "ChannelSub":
		return device.ChannelSub()
	case "SetChannel":
		return device.SetChannel(parameter.Channel)
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteTVDevice command parameter
func (device *InfraredRemoteTVDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteTVDeviceCommandParameter{})
}

// InfraredRemoteDvdPlayerDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteDvdPlayerDevice
type InfraredRemoteDvdPlayerDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,SetMute,FastForward,Rewind,Next,Previous,Pause,Play,Stop" description:"TurnOn:turn on the DVD player, TurnOff:turn off the DVD player, SetMute:mute/unmute, FastForward:fast forward, Rewind:rewind, Next:next track, Previous:previous track, Pause:pause, Play:start, Stop:stop" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the InfraredRemoteDvdPlayerDevice
func (device *InfraredRemoteDvdPlayerDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteDvdPlayerDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "SetMute":
		return device.SetMute()
	case "FastForward":
		return device.FastForward()
	case "Rewind":
		return device.Rewind()
	case "Next":
		return device.Next()
	case "Previous":
		return device.Previous()
	case "Pause":
		return device.Pause()
	case "Play":
		return device.Play()
	case "Stop":
		return device.Stop()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteDvdPlayerDevice command parameter
func (device *InfraredRemoteDvdPlayerDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteDvdPlayerDeviceCommandParameter{})
}

// InfraredRemoteSpeakerDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteSpeakerDevice
type InfraredRemoteSpeakerDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,VolumeAdd,VolumeSub,SetMute,FastForward,Rewind,Next,Previous,Pause,Play,Stop" description:"TurnOn:turn on the speaker, TurnOff:turn off the speaker, VolumeAdd:increase volume, VolumeSub:decrease volume, SetMute:mute/unmute, FastForward:fast forward, Rewind:rewind, Next:next track, Previous:previous track, Pause:pause, Play:start, Stop:stop" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the InfraredRemoteSpeakerDevice
func (device *InfraredRemoteSpeakerDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteSpeakerDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "VolumeAdd":
		return device.VolumeAdd()
	case "VolumeSub":
		return device.VolumeSub()
	case "SetMute":
		return device.SetMute()
	case "FastForward":
		return device.FastForward()
	case "Rewind":
		return device.Rewind()
	case "Next":
		return device.Next()
	case "Previous":
		return device.Previous()
	case "Pause":
		return device.Pause()
	case "Play":
		return device.Play()
	case "Stop":
		return device.Stop()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteSpeakerDevice command parameter
func (device *InfraredRemoteSpeakerDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteSpeakerDeviceCommandParameter{})
}

// InfraredRemoteFanDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteFanDevice
type InfraredRemoteFanDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,Swing,Timer,LowSpeed,MiddleSpeed,HighSpeed" description:"TurnOn:turn on the fan, TurnOff:turn off the fan, Swing:enable/disable swing feature, Timer:set timer, LowSpeed:set fan speed to low, MiddleSpeed:set fan speed to middle, HighSpeed:set fan speed to high" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the InfraredRemoteFanDevice
func (device *InfraredRemoteFanDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteFanDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "Swing":
		return device.Swing()
	case "Timer":
		return device.Timer()
	case "LowSpeed":
		return device.LowSpeed()
	case "MiddleSpeed":
		return device.MiddleSpeed()
	case "HighSpeed":
		return device.HighSpeed()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteFanDevice command parameter
func (device *InfraredRemoteFanDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteFanDeviceCommandParameter{})
}

// InfraredRemoteLightDeviceCommandParameter is a struct that represents the command parameter for the InfraredRemoteLightDevice
type InfraredRemoteLightDeviceCommandParameter struct {
	Command string   `json:"command" title:"Command" enum:"TurnOn,TurnOff,BrightnessUp,BrightnessDown" description:"TurnOn:turn on the light, TurnOff:turn off the light, BrightnessUp:increase brightness, BrightnessDown:decrease brightness" required:"true"`
	_       struct{} `additionalProperties:"false"`
}

// ExecCommand sends a command to the InfraredRemoteLightDevice
func (device *InfraredRemoteLightDevice) ExecCommand(jsonString string) (*CommonResponse, error) {
	var parameter InfraredRemoteLightDeviceCommandParameter
	if err := validateAndUnmarshalJSON(device, jsonString, &parameter); err != nil {
		return nil, err
	}

	switch parameter.Command {
	case "TurnOn":
		return device.TurnOn()
	case "TurnOff":
		return device.TurnOff()
	case "BrightnessUp":
		return device.BrightnessUp()
	case "BrightnessDown":
		return device.BrightnessDown()
	default:
		return nil, fmt.Errorf("invalid Command: %s", parameter.Command)
	}
}

// GetCommandParameterJSONSchema returns the JSON schema for the InfraredRemoteLightDevice command parameter
func (device *InfraredRemoteLightDevice) GetCommandParameterJSONSchema() (string, error) {
	return reflectJSONSchema(InfraredRemoteLightDeviceCommandParameter{})
}
