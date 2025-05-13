package switchbot

import (
	"encoding/json"
	"fmt"

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
		KeypadDeviceCommandCreateNormalKeyIfExposer{},
		KeypadDeviceCommandCreateTimeLimitKeyIfExposer{},
		KeypadDeviceCommandDeleteKeyIfExposer{},
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
	Brightness       int      `json:"brightness" title:"Brightness" minimum:"1" maximum:"100" description:"Brightness level (0-100)"`
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
		CeilingLightDeviceCommandSetBrightnessIfExposer{},
		CeilingLightDeviceCommandSetColorTemperatureIfExposer{},
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
