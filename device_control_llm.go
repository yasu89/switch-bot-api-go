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
