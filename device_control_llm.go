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
	schema, err := reflector.Reflect(parameter)
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
