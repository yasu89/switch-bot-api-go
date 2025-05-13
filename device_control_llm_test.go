package switchbot_test

import (
	"github.com/yasu89/switch-bot-api-go/helpers"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yasu89/switch-bot-api-go"
)

func Test_BotDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.BotDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_BotDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Press",
			expectedBody: `{"commandType": "command","command": "press","parameter": "default"}`,
			parameter:    `{"command":"Press"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.BotDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_CurtainDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.CurtainDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_CurtainDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "SetPosition",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "0,ff,75"}`,
			parameter:    `{"command":"SetPosition","mode":"ff", "position":75}`,
		},
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Pause",
			expectedBody: `{"commandType": "command","command": "pause","parameter": "default"}`,
			parameter:    `{"command":"Pause"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.CurtainDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_CurtainDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "SetPosition(invalid mode)",
			parameter:    `{"command":"SetPosition","mode":"00", "position":75}`,
			errorContain: "Value 00 should be one of the allowed values: 0, 1, ff",
		},
		{
			name:         "SetPosition(invalid position)",
			parameter:    `{"command":"SetPosition","mode":"ff", "position":101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetPosition(missing mode)",
			parameter:    `{"command":"SetPosition", "position":80}`,
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SetPosition(missing position)",
			parameter:    `{"command":"SetPosition","mode":"ff"}`,
			errorContain: "Required property 'position' is missing",
		},
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, Pause, SetPosition",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.CurtainDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_LockDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.LockDevice{}
	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_LockDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "Lock",
			expectedBody: `{"commandType": "command","command": "lock","parameter": "default"}`,
			parameter:    `{"command":"Lock"}`,
		},
		{
			name:         "Unlock",
			expectedBody: `{"commandType": "command","command": "unlock","parameter": "default"}`,
			parameter:    `{"command":"Unlock"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.LockDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_KeypadDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.KeypadDevice{}
	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_KeypadDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "CreateKey(permanent)",
			expectedBody: `{"commandType":"command","command":"createKey","parameter":{"name":"testKey","type":"permanent","password":"123456"}}`,
			parameter:    `{"command":"CreateKey","name":"testKey", "type":"permanent","password":"123456"}`,
		},
		{
			name:         "CreateKey(timeLimit)",
			expectedBody: `{"commandType":"command","command":"createKey","parameter":{"name":"testKey","type":"permanent","password":"123456","startTime":1745080854,"endTime":1745167254}}`,
			parameter:    `{"command":"CreateKey","name":"testKey", "type":"permanent","password":"123456","startTime":1745080854,"endTime":1745167254}`,
		},
		{
			name:         "DeleteKey",
			expectedBody: `{"commandType":"command","command":"deleteKey","parameter":{"id":"testKey"}}`,
			parameter:    `{"command":"DeleteKey","id":"testKey"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.KeypadDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_KeypadDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Create normal key(missing name parameter)",
			parameter:    `{"command":"CreateKey","type":"permanent","password":"123456"}`,
			errorContain: "Required property 'name' is missing",
		},
		{
			name:         "Create normal key(missing password parameter)",
			parameter:    `{"command":"CreateKey","type":"urgent","name":"testKey"}`,
			errorContain: "Required property 'password' is missing",
		},
		{
			name:         "Create normal key(invalid password)",
			parameter:    `{"command":"CreateKey","type":"permanent","name":"testKey","password":"12345"}`,
			errorContain: "Value does not match the required pattern ^[0-9]{6,12}$",
		},
		{
			name:         "Create timeLimit key(missing startTime parameter)",
			parameter:    `{"command":"CreateKey","type":"timeLimit","name":"testKey","password":"123456","endTime":1745167254}`,
			errorContain: "Required property 'startTime' is missing",
		},
		{
			name:         "Create timeLimit key(missing endTime parameter)",
			parameter:    `{"command":"CreateKey","type":"disposable","name":"testKey","password":"123456","startTime":1745167254}`,
			errorContain: "Required property 'endTime' is missing",
		},
		{
			name:         "Create timeLimit key(invalid startTime)",
			parameter:    `{"command":"CreateKey","type":"timeLimit","name":"testKey","password":"123456","startTime":0,"endTime":1745167254}`,
			errorContain: "0 should be at least 100000000",
		},
		{
			name:         "Delete key(missing id parameter)",
			parameter:    `{"command":"DeleteKey"}`,
			errorContain: "Required property 'id' is missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.KeypadDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_CeilingLightDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.CeilingLightDevice{}
	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_CeilingLightDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "50"}`,
			parameter:    `{"command":"SetBrightness","brightness":50}`,
		},
		{
			name:         "SetColorTemperature",
			expectedBody: `{"commandType": "command","command": "setColorTemperature","parameter": "3500"}`,
			parameter:    `{"command":"SetColorTemperature","colorTemperature":3500}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.CeilingLightDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_CeilingLightDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "SetBrightness(invalid brightness)",
			parameter:    `{"command":"SetBrightness","brightness":0}`,
			errorContain: "0 should be at least 1",
		},
		{
			name:         "SetBrightness(brightness too high)",
			parameter:    `{"command":"SetBrightness","brightness":101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetBrightness(missing brightness)",
			parameter:    `{"command":"SetBrightness"}`,
			errorContain: "Required property 'brightness' is missing",
		},
		{
			name:         "SetColorTemperature(invalid colorTemperature too low)",
			parameter:    `{"command":"SetColorTemperature","colorTemperature":2000}`,
			errorContain: "2000 should be at least 2700",
		},
		{
			name:         "SetColorTemperature(invalid colorTemperature too high)",
			parameter:    `{"command":"SetColorTemperature","colorTemperature":7000}`,
			errorContain: "7000 should be at most 6500",
		},
		{
			name:         "SetColorTemperature(missing colorTemperature)",
			parameter:    `{"command":"SetColorTemperature"}`,
			errorContain: "Required property 'colorTemperature' is missing",
		},
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, Toggle, SetBrightness, SetColorTemperature",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.CeilingLightDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_PlugMiniDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.PlugMiniDevice{}
	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_PlugMiniDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.PlugMiniDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_PlugMiniDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, Toggle",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.PlugMiniDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_PlugDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.PlugDevice{}
	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_PlugDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.PlugDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_PlugDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"Toggle"}`,
			errorContain: "Value Toggle should be one of the allowed values: TurnOn, TurnOff",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.PlugDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_StripLightDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.StripLightDevice{}
	schema, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, schema)
}

func Test_StripLightDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "50"}`,
			parameter:    `{"command":"SetBrightness", "brightness": 50}`,
		},
		{
			name:         "SetColor",
			expectedBody: `{"commandType": "command","command": "setColor","parameter": "255:0:0"}`,
			parameter:    `{"command":"SetColor", "red": 255, "green": 0, "blue": 0}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.StripLightDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_StripLightDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: TurnOn, TurnOff, Toggle, SetBrightness, SetColor",
		},
		{
			name:         "SetBrightness(without brightness)",
			parameter:    `{"command": "SetBrightness"}`,
			errorContain: "Required property 'brightness' is missing",
		},
		{
			name:         "SetBrightness(invalid brightness)",
			parameter:    `{"command": "SetBrightness", "brightness": 101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetColor(without red)",
			parameter:    `{"command": "SetColor", "green": 0, "blue": 0}`,
			errorContain: "Required property 'red' is missing",
		},
		{
			name:         "SetColor(without green)",
			parameter:    `{"command": "SetColor", "red": 255, "blue": 0}`,
			errorContain: "Required property 'green' is missing",
		},
		{
			name:         "SetColor(without blue)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 0}`,
			errorContain: "Required property 'blue' is missing",
		},
		{
			name:         "SetColor(invalid red)",
			parameter:    `{"command": "SetColor", "red": 256, "green": 0, "blue": 0}`,
			errorContain: "256 should be at most 255",
		},
		{
			name:         "SetColor(invalid green)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 256, "blue": 0}`,
			errorContain: "256 should be at most 255",
		},
		{
			name:         "SetColor(invalid blue)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 0, "blue": 256}`,
			errorContain: "256 should be at most 255",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.StripLightDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_ColorBulbDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.ColorBulbDevice{}
	schema, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, schema)
}

func Test_ColorBulbDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "50"}`,
			parameter:    `{"command":"SetBrightness", "brightness": 50}`,
		},
		{
			name:         "SetColor",
			expectedBody: `{"commandType": "command","command": "setColor","parameter": "255:0:0"}`,
			parameter:    `{"command":"SetColor", "red": 255, "green": 0, "blue": 0}`,
		},
		{
			name:         "SetColorTemperature",
			expectedBody: `{"commandType": "command","command": "setColorTemperature","parameter": "3000"}`,
			parameter:    `{"command":"SetColorTemperature", "colorTemperature": 3000}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.ColorBulbDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_ColorBulbDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: TurnOn, TurnOff, Toggle, SetBrightness, SetColor, SetColorTemperature",
		},
		{
			name:         "SetBrightness(without brightness)",
			parameter:    `{"command": "SetBrightness"}`,
			errorContain: "Required property 'brightness' is missing",
		},
		{
			name:         "SetBrightness(invalid brightness)",
			parameter:    `{"command": "SetBrightness", "brightness": 101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetBrightness(too low brightness)",
			parameter:    `{"command": "SetBrightness", "brightness": 0}`,
			errorContain: "0 should be at least 1",
		},
		{
			name:         "SetColor(without red)",
			parameter:    `{"command": "SetColor", "green": 0, "blue": 0}`,
			errorContain: "Required property 'red' is missing",
		},
		{
			name:         "SetColor(without green)",
			parameter:    `{"command": "SetColor", "red": 255, "blue": 0}`,
			errorContain: "Required property 'green' is missing",
		},
		{
			name:         "SetColor(without blue)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 0}`,
			errorContain: "Required property 'blue' is missing",
		},
		{
			name:         "SetColor(invalid red)",
			parameter:    `{"command": "SetColor", "red": 256, "green": 0, "blue": 0}`,
			errorContain: "256 should be at most 255",
		},
		{
			name:         "SetColor(invalid green)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 256, "blue": 0}`,
			errorContain: "256 should be at most 255",
		},
		{
			name:         "SetColor(invalid blue)",
			parameter:    `{"command": "SetColor", "red": 255, "green": 0, "blue": 256}`,
			errorContain: "256 should be at most 255",
		},
		{
			name:         "SetColorTemperature(without colorTemperature)",
			parameter:    `{"command": "SetColorTemperature"}`,
			errorContain: "Required property 'colorTemperature' is missing",
		},
		{
			name:         "SetColorTemperature(too low colorTemperature)",
			parameter:    `{"command": "SetColorTemperature", "colorTemperature": 2699}`,
			errorContain: "2699 should be at least 2700",
		},
		{
			name:         "SetColorTemperature(too high colorTemperature)",
			parameter:    `{"command": "SetColorTemperature", "colorTemperature": 6501}`,
			errorContain: "6501 should be at most 6500",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.ColorBulbDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			_, err := device.ExecCommand(testData.parameter)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), testData.errorContain)
		})
	}
}

func Test_RobotVacuumCleanerDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.RobotVacuumCleanerDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_RobotVacuumCleanerDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "Start",
			expectedBody: `{"commandType": "command","command": "start","parameter": "default"}`,
			parameter:    `{"command":"Start"}`,
		},
		{
			name:         "Stop",
			expectedBody: `{"commandType": "command","command": "stop","parameter": "default"}`,
			parameter:    `{"command":"Stop"}`,
		},
		{
			name:         "Dock",
			expectedBody: `{"commandType": "command","command": "dock","parameter": "default"}`,
			parameter:    `{"command":"Dock"}`,
		},
		{
			name:         "SetPowerLevel",
			expectedBody: `{"commandType": "command","command": "PowLevel","parameter": "2"}`,
			parameter:    `{"command":"SetPowerLevel","powerLevel":2}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RobotVacuumCleanerDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_RobotVacuumCleanerDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: Start, Stop, Dock, SetPowerLevel",
		},
		{
			name:         "SetPowerLevel(without powerLevel)",
			parameter:    `{"command":"SetPowerLevel"}`,
			errorContain: "Required property 'powerLevel' is missing",
		},
		{
			name:         "SetPowerLevel(invalid powerLevel)",
			parameter:    `{"command":"SetPowerLevel","powerLevel":5}`,
			errorContain: "5 should be at most 3",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RobotVacuumCleanerDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			_, err := device.ExecCommand(testData.parameter)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), testData.errorContain)
		})
	}
}

func Test_RobotVacuumCleanerS10DeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.RobotVacuumCleanerS10Device{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_RobotVacuumCleanerS10DeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "StartClean",
			expectedBody: `{"commandType": "command","command": "startClean","parameter": {"action":"sweep_mop","param":{"fanLevel":2,"waterLevel":1,"times":1}}}`,
			parameter:    `{"command":"StartClean","action":"sweep_mop","fanLevel":2,"waterLevel":1,"times":1}`,
		},
		{
			name:         "AddWaterForHumi",
			expectedBody: `{"commandType": "command","command": "addWaterForHumi","parameter": "default"}`,
			parameter:    `{"command":"AddWaterForHumi"}`,
		},
		{
			name:         "Pause",
			expectedBody: `{"commandType": "command","command": "pause","parameter": "default"}`,
			parameter:    `{"command":"Pause"}`,
		},
		{
			name:         "Dock",
			expectedBody: `{"commandType": "command","command": "dock","parameter": "default"}`,
			parameter:    `{"command":"Dock"}`,
		},
		{
			name:         "SetVolume",
			expectedBody: `{"commandType": "command","command": "setVolume","parameter": "50"}`,
			parameter:    `{"command":"SetVolume","volume":50}`,
		},
		{
			name:         "SelfClean",
			expectedBody: `{"commandType": "command","command": "selfClean","parameter": "1"}`,
			parameter:    `{"command":"SelfClean","mode":1}`,
		},
		{
			name:         "ChangeParam",
			expectedBody: `{"commandType": "command","command": "changeParam","parameter": {"fanLevel":3,"waterLevel":2,"times":100}}}`,
			parameter:    `{"command":"ChangeParam","fanLevel":3,"waterLevel":2,"times":100}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RobotVacuumCleanerS10Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_RobotVacuumCleanerS10DeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: StartClean, AddWaterForHumi, Pause, Dock, SetVolume, SelfClean, ChangeParam",
		},
		{
			name:         "StartClean(without action)",
			parameter:    `{"command":"StartClean","fanLevel":2,"waterLevel":1,"times":1}`,
			errorContain: "Required property 'action' is missing",
		},
		{
			name:         "StartClean(without fanLevel)",
			parameter:    `{"command":"StartClean","action":"sweep","waterLevel":1,"times":1}`,
			errorContain: "Required property 'fanLevel' is missing",
		},
		{
			name:         "StartClean(without waterLevel)",
			parameter:    `{"command":"StartClean","action":"sweep","fanLevel":2,"times":1}`,
			errorContain: "Required property 'waterLevel' is missing",
		},
		{
			name:         "StartClean(without times)",
			parameter:    `{"command":"StartClean","action":"sweep","fanLevel":2,"waterLevel":1}`,
			errorContain: "Required property 'times' is missing",
		},
		{
			name:         "StartClean(invalid action)",
			parameter:    `{"command":"StartClean","action":"invalid","fanLevel":2,"waterLevel":1,"times":1}`,
			errorContain: "Value invalid should be one of the allowed values: sweep, sweep_mop",
		},
		{
			name:         "StartClean(invalid fanLevel)",
			parameter:    `{"command":"StartClean","action":"sweep","fanLevel":5,"waterLevel":1,"times":1}`,
			errorContain: "5 should be at most 4",
		},
		{
			name:         "StartClean(invalid waterLevel)",
			parameter:    `{"command":"StartClean","action":"sweep","fanLevel":2,"waterLevel":3,"times":1}`,
			errorContain: "3 should be at most 2",
		},
		{
			name:         "StartClean(invalid times)",
			parameter:    `{"command":"StartClean","action":"sweep","fanLevel":2,"waterLevel":1,"times":2640000}`,
			errorContain: "2640000 should be at most 2639999",
		},
		{
			name:         "SetVolume(without volume)",
			parameter:    `{"command":"SetVolume"}`,
			errorContain: "Required property 'volume' is missing",
		},
		{
			name:         "SetVolume(invalid volume)",
			parameter:    `{"command":"SetVolume","volume":101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SelfClean(without mode)",
			parameter:    `{"command":"SelfClean"}`,
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SelfClean(invalid mode)",
			parameter:    `{"command":"SelfClean","mode":4}`,
			errorContain: "4 should be at most 3",
		},
		{
			name:         "ChangeParam(without fanLevel)",
			parameter:    `{"command":"ChangeParam","waterLevel":1,"times":1}`,
			errorContain: "Required property 'fanLevel' is missing",
		},
		{
			name:         "ChangeParam(without waterLevel)",
			parameter:    `{"command":"ChangeParam","fanLevel":2,"times":1}`,
			errorContain: "Required property 'waterLevel' is missing",
		},
		{
			name:         "ChangeParam(without times)",
			parameter:    `{"command":"ChangeParam","fanLevel":2,"waterLevel":1}`,
			errorContain: "Required property 'times' is missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RobotVacuumCleanerS10Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}

			_, err := device.ExecCommand(testData.parameter)
			assert.Error(t, err)
			assert.Contains(t, err.Error(), testData.errorContain)
		})
	}
}

func Test_HumidifierDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.HumidifierDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_HumidifierDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetMode(Auto)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "0"}`,
			parameter:    `{"command":"SetMode","mode":"Auto"}`,
		},
		{
			name:         "SetMode(Low)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "101"}`,
			parameter:    `{"command":"SetMode","mode":"Low"}`,
		},
		{
			name:         "SetMode(Medium)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "102"}`,
			parameter:    `{"command":"SetMode","mode":"Medium"}`,
		},
		{
			name:         "SetMode(High)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "103"}`,
			parameter:    `{"command":"SetMode","mode":"High"}`,
		},
		{
			name:         "SetTargetHumidity",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "50"}`,
			parameter:    `{"command":"SetTargetHumidity","targetHumidity":50}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.HumidifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_HumidifierDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, SetMode, SetTargetHumidity",
		},
		{
			name:         "SetMode(missing mode)",
			parameter:    `{"command":"SetMode"}`,
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SetMode(invalid mode)",
			parameter:    `{"command":"SetMode","mode":"Invalid"}`,
			errorContain: "Invalid should be one of the allowed values: Auto, Low, Medium, High",
		},
		{
			name:         "SetTargetHumidity(missing targetHumidity)",
			parameter:    `{"command":"SetTargetHumidity"}`,
			errorContain: "Required property 'targetHumidity' is missing",
		},
		{
			name:         "SetTargetHumidity(targetHumidity too low)",
			parameter:    `{"command":"SetTargetHumidity","targetHumidity":-1}`,
			errorContain: "-1 should be at least 0",
		},
		{
			name:         "SetTargetHumidity(targetHumidity too high)",
			parameter:    `{"command":"SetTargetHumidity","targetHumidity":101}`,
			errorContain: "101 should be at most 100",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.HumidifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_EvaporativeHumidifierDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.EvaporativeHumidifierDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_EvaporativeHumidifierDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetMode(Level4)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":1,"targetHumidity":50}}`,
			parameter:    `{"command":"SetMode","mode":1,"targetHumidity":50}`,
		},
		{
			name:         "SetMode(Level3)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":2,"targetHumidity":60}}`,
			parameter:    `{"command":"SetMode","mode":2,"targetHumidity":60}`,
		},
		{
			name:         "SetMode(Level2)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":3,"targetHumidity":70}}`,
			parameter:    `{"command":"SetMode","mode":3,"targetHumidity":70}`,
		},
		{
			name:         "SetMode(Level1)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":4,"targetHumidity":80}}`,
			parameter:    `{"command":"SetMode","mode":4,"targetHumidity":80}`,
		},
		{
			name:         "SetMode(Humidity mode)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":5,"targetHumidity":50}}`,
			parameter:    `{"command":"SetMode","mode":5,"targetHumidity":50}`,
		},
		{
			name:         "SetMode(Sleep mode)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":6,"targetHumidity":50}}`,
			parameter:    `{"command":"SetMode","mode":6,"targetHumidity":50}`,
		},
		{
			name:         "SetMode(Auto mode)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":7,"targetHumidity":50}}`,
			parameter:    `{"command":"SetMode","mode":7,"targetHumidity":50}`,
		},
		{
			name:         "SetMode(Drying mode)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":8,"targetHumidity":50}}`,
			parameter:    `{"command":"SetMode","mode":8,"targetHumidity":50}`,
		},
		{
			name:         "SetChildLock(true)",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": "true"}`,
			parameter:    `{"command":"SetChildLock","childLock":true}`,
		},
		{
			name:         "SetChildLock(false)",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": "false"}`,
			parameter:    `{"command":"SetChildLock","childLock":false}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.EvaporativeHumidifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_EvaporativeHumidifierDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, SetMode, SetChildLock",
		},
		{
			name:         "SetMode(missing mode)",
			parameter:    `{"command":"SetMode", "targetHumidity":50}`,
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SetMode(missing targetHumidity)",
			parameter:    `{"command":"SetMode","mode":1}`,
			errorContain: "Required property 'targetHumidity' is missing",
		},
		{
			name:         "SetMode(invalid mode)",
			parameter:    `{"command":"SetMode","mode":999,"targetHumidity":50}`,
			errorContain: "999 should be at most 8",
		},
		{
			name:         "SetMode(targetHumidity too low)",
			parameter:    `{"command":"SetMode","mode":1,"targetHumidity":-1}`,
			errorContain: "-1 should be at least 0",
		},
		{
			name:         "SetMode(targetHumidity too high)",
			parameter:    `{"command":"SetMode","mode":1,"targetHumidity":101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetChildLock(missing childLock)",
			parameter:    `{"command":"SetChildLock"}`,
			errorContain: "Required property 'childLock' is missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.EvaporativeHumidifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_AirPurifierDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.AirPurifierDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_AirPurifierDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetMode(Normal)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":1,"fanGear":1}}`,
			parameter:    `{"command":"SetMode","mode":1,"fanLevel":1}`,
		},
		{
			name:         "SetMode(Normal with fanLevel 2)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":1,"fanGear":2}}`,
			parameter:    `{"command":"SetMode","mode":1,"fanLevel":2}`,
		},
		{
			name:         "SetMode(Normal with fanLevel 3)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":1,"fanGear":3}}`,
			parameter:    `{"command":"SetMode","mode":1,"fanLevel":3}`,
		},
		{
			name:         "SetMode(Auto)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":2}}`,
			parameter:    `{"command":"SetMode","mode":2}`,
		},
		{
			name:         "SetMode(Sleep)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":3}}`,
			parameter:    `{"command":"SetMode","mode":3}`,
		},
		{
			name:         "SetMode(Manual)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":4}}`,
			parameter:    `{"command":"SetMode","mode":4}`,
		},
		{
			name:         "SetChildLock(true)",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": 1}`,
			parameter:    `{"command":"SetChildLock","childLock":true}`,
		},
		{
			name:         "SetChildLock(false)",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": 0}`,
			parameter:    `{"command":"SetChildLock","childLock":false}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.AirPurifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_AirPurifierDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"Invalid"}`,
			errorContain: "Value Invalid should be one of the allowed values: TurnOn, TurnOff, SetMode, SetChildLock",
		},
		{
			name:         "SetMode(missing mode)",
			parameter:    `{"command":"SetMode"}`,
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SetMode(invalid mode too low)",
			parameter:    `{"command":"SetMode","mode":0}`,
			errorContain: "0 should be at least 1",
		},
		{
			name:         "SetMode(invalid mode too high)",
			parameter:    `{"command":"SetMode","mode":5}`,
			errorContain: "5 should be at most 4",
		},
		{
			name:         "SetMode(Normal missing fanLevel)",
			parameter:    `{"command":"SetMode","mode":1}`,
			errorContain: "Required property 'fanLevel' is missing",
		},
		{
			name:         "SetMode(Normal invalid fanLevel too low)",
			parameter:    `{"command":"SetMode","mode":1,"fanLevel":0}`,
			errorContain: "0 should be at least 1",
		},
		{
			name:         "SetMode(Normal invalid fanLevel too high)",
			parameter:    `{"command":"SetMode","mode":1,"fanLevel":4}`,
			errorContain: "4 should be at most 3",
		},
		{
			name:         "SetChildLock(missing childLock)",
			parameter:    `{"command":"SetChildLock"}`,
			errorContain: "Required property 'childLock' is missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.AirPurifierDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_BlindTiltDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.BlindTiltDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_BlindTiltDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "SetPosition(up, 50)",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "up;50"}`,
			parameter:    `{"command":"SetPosition","direction":"up","position":50}`,
		},
		{
			name:         "SetPosition(down, 80)",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "down;80"}`,
			parameter:    `{"command":"SetPosition","direction":"down","position":80}`,
		},
		{
			name:         "FullyOpen",
			expectedBody: `{"commandType": "command","command": "fullyOpen","parameter": "default"}`,
			parameter:    `{"command":"FullyOpen"}`,
		},
		{
			name:         "CloseUp",
			expectedBody: `{"commandType": "command","command": "closeUp","parameter": "default"}`,
			parameter:    `{"command":"CloseUp"}`,
		},
		{
			name:         "CloseDown",
			expectedBody: `{"commandType": "command","command": "closeDown","parameter": "default"}`,
			parameter:    `{"command":"CloseDown"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.BlindTiltDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_BlindTiltDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: SetPosition, FullyOpen, CloseUp, CloseDown",
		},
		{
			name:         "Invalid direction",
			parameter:    `{"command":"SetPosition","direction":"invalid","position":50}`,
			errorContain: "Value invalid should be one of the allowed values: up, down",
		},
		{
			name:         "Position too low",
			parameter:    `{"command":"SetPosition","direction":"up","position":-10}`,
			errorContain: "-10 should be at least 0",
		},
		{
			name:         "Position too high",
			parameter:    `{"command":"SetPosition","direction":"up","position":110}`,
			errorContain: "110 should be at most 100",
		},
		{
			name:         "Position not even",
			parameter:    `{"command":"SetPosition","direction":"up","position":51}`,
			errorContain: "position must be even: 51",
		},
		{
			name:         "Missing required parameter",
			parameter:    `{"command":"SetPosition"}`,
			errorContain: "Required properties 'direction', 'position' are missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.BlindTiltDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_BatteryCirculatorFanDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.BatteryCirculatorFanDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_BatteryCirculatorFanDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetNightLightMode(off)",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "off"}`,
			parameter:    `{"command":"SetNightLightMode","nightLight":"off"}`,
		},
		{
			name:         "SetNightLightMode(1)",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "1"}`,
			parameter:    `{"command":"SetNightLightMode","nightLight":"1"}`,
		},
		{
			name:         "SetNightLightMode(2)",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "2"}`,
			parameter:    `{"command":"SetNightLightMode","nightLight":"2"}`,
		},
		{
			name:         "SetWindMode(direct)",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "direct"}`,
			parameter:    `{"command":"SetWindMode","windMode":"direct"}`,
		},
		{
			name:         "SetWindMode(natural)",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "natural"}`,
			parameter:    `{"command":"SetWindMode","windMode":"natural"}`,
		},
		{
			name:         "SetWindMode(sleep)",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "sleep"}`,
			parameter:    `{"command":"SetWindMode","windMode":"sleep"}`,
		},
		{
			name:         "SetWindMode(baby)",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "baby"}`,
			parameter:    `{"command":"SetWindMode","windMode":"baby"}`,
		},
		{
			name:         "SetWindSpeed(50)",
			expectedBody: `{"commandType": "command","command": "setWindSpeed","parameter": "50"}`,
			parameter:    `{"command":"SetWindSpeed","windSpeed":50}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.BatteryCirculatorFanDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_BatteryCirculatorFanDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid command",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: TurnOn, TurnOff, SetNightLightMode, SetWindMode, SetWindSpeed",
		},
		{
			name:         "Invalid night light mode",
			parameter:    `{"command":"SetNightLightMode","nightLight":"invalid"}`,
			errorContain: "Value invalid should be one of the allowed values: off, 1, 2",
		},
		{
			name:         "Invalid wind mode",
			parameter:    `{"command":"SetWindMode","windMode":"invalid"}`,
			errorContain: "Value invalid should be one of the allowed values: direct, natural, sleep, baby",
		},
		{
			name:         "Wind speed too low",
			parameter:    `{"command":"SetWindSpeed","windSpeed":0}`,
			errorContain: "0 should be at least 1",
		},
		{
			name:         "Wind speed too high",
			parameter:    `{"command":"SetWindSpeed","windSpeed":101}`,
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetNightLightMode(without nightLight)",
			parameter:    `{"command":"SetNightLightMode"}`,
			errorContain: "Required property 'nightLight' is missing",
		},
		{
			name:         "SetWindMode(without windMode)",
			parameter:    `{"command":"SetWindMode"}`,
			errorContain: "Required property 'windMode' is missing",
		},
		{
			name:         "SetWindSpeed(without windSpeed)",
			parameter:    `{"command":"SetWindSpeed"}`,
			errorContain: "Required property 'windSpeed' is missing",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.BatteryCirculatorFanDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_CirculatorFanDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.CirculatorFanDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_CirculatorFanDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetNightLightMode(off)",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "off"}`,
			parameter:    `{"command":"SetNightLightMode","nightLight":"off"}`,
		},
		{
			name:         "SetWindMode(direct)",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "direct"}`,
			parameter:    `{"command":"SetWindMode","windMode":"direct"}`,
		},
		{
			name:         "SetWindSpeed(50)",
			expectedBody: `{"commandType": "command","command": "setWindSpeed","parameter": "50"}`,
			parameter:    `{"command":"SetWindSpeed","windSpeed":50}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.CirculatorFanDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_RollerShadeDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.RollerShadeDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_RollerShadeDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "SetPosition",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "50"}`,
			parameter:    `{"command":"SetPosition","position":50}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RollerShadeDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_RollerShadeDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "InvalidCommand",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: `Value InvalidCommand should be one of the allowed values: SetPosition"`,
		},
		{
			name:         "SetPositionWithoutPosition",
			parameter:    `{"command":"SetPosition"}`,
			errorContain: `Required property 'position' is missing`,
		},
		{
			name:         "SetPositionWithInvalidPosition",
			parameter:    `{"command":"SetPosition","position":-1}`,
			errorContain: "-1 should be at least 0",
		},
		{
			name:         "SetPositionWithInvalidPositionTooLarge",
			parameter:    `{"command":"SetPosition","position":101}`,
			errorContain: "101 should be at most 100",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RollerShadeDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_RelaySwitch1DeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.RelaySwitch1Device{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_RelaySwitch1DeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
		{
			name:         "SetMode(Toggle)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "0"}`,
			parameter:    `{"command":"SetMode","mode":0}`,
		},
		{
			name:         "SetMode(Edge)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "1"}`,
			parameter:    `{"command":"SetMode","mode":1}`,
		},
		{
			name:         "SetMode(Detached)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "2"}`,
			parameter:    `{"command":"SetMode","mode":2}`,
		},
		{
			name:         "SetMode(Momentary)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "3"}`,
			parameter:    `{"command":"SetMode","mode":3}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RelaySwitch1Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_RelaySwitch1DeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "InvalidCommand",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: `Value InvalidCommand should be one of the allowed values`,
		},
		{
			name:         "SetModeWithoutMode",
			parameter:    `{"command":"SetMode"}`,
			errorContain: `Required property 'mode' is missing`,
		},
		{
			name:         "SetModeWithInvalidModeTooSmall",
			parameter:    `{"command":"SetMode","mode":-1}`,
			errorContain: "-1 should be at least 0",
		},
		{
			name:         "SetModeWithInvalidModeTooLarge",
			parameter:    `{"command":"SetMode","mode":4}`,
			errorContain: "4 should be at most 3",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RelaySwitch1Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_RelaySwitch1PMDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.RelaySwitch1PMDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_RelaySwitch1PMDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			parameter:    `{"command":"Toggle"}`,
		},
		{
			name:         "SetMode(Toggle)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "0"}`,
			parameter:    `{"command":"SetMode","mode":0}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RelaySwitch1PMDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteAirConditionerDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.InfraredRemoteAirConditionerDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_InfraredRemoteAirConditionerDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetAll(Auto,Auto,On)",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "25,1,1,on"}`,
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"fan":1,"powerState":"on"}`,
		},
		{
			name:         "SetAll(Cool,Low,Off)",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "20,2,2,off"}`,
			parameter:    `{"command":"SetAll","temperatureCelsius":20,"mode":2,"fan":2,"powerState":"off"}`,
		},
		{
			name:         "SetAll(Dry,Medium,On)",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "22,3,3,on"}`,
			parameter:    `{"command":"SetAll","temperatureCelsius":22,"mode":3,"fan":3,"powerState":"on"}`,
		},
		{
			name:         "SetAll(Fan,High,Off)",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "24,4,4,off"}`,
			parameter:    `{"command":"SetAll","temperatureCelsius":24,"mode":4,"fan":4,"powerState":"off"}`,
		},
		{
			name:         "SetAll(Heat,Auto,On)",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "28,5,1,on"}`,
			parameter:    `{"command":"SetAll","temperatureCelsius":28,"mode":5,"fan":1,"powerState":"on"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteAirConditionerDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteAirConditionerDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "InvalidCommand",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: `Value InvalidCommand should be one of the allowed values`,
		},
		{
			name:         "SetAllWithoutTemperature",
			parameter:    `{"command":"SetAll","mode":1,"fan":1,"powerState":"on"}`,
			errorContain: `Required property 'temperatureCelsius' is missing`,
		},
		{
			name:         "SetAllWithoutMode",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"fan":1,"powerState":"on"}`,
			errorContain: `Required property 'mode' is missing`,
		},
		{
			name:         "SetAllWithoutFan",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"powerState":"on"}`,
			errorContain: `Required property 'fan' is missing`,
		},
		{
			name:         "SetAllWithoutPowerState",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"fan":1}`,
			errorContain: `Required property 'powerState' is missing`,
		},
		{
			name:         "SetAllWithInvalidTemperatureTooLow",
			parameter:    `{"command":"SetAll","temperatureCelsius":-11,"mode":1,"fan":1,"powerState":"on"}`,
			errorContain: `-11 should be at least -10`,
		},
		{
			name:         "SetAllWithInvalidTemperatureTooHigh",
			parameter:    `{"command":"SetAll","temperatureCelsius":41,"mode":1,"fan":1,"powerState":"on"}`,
			errorContain: `41 should be at most 40`,
		},
		{
			name:         "SetAllWithInvalidModeTooLow",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":0,"fan":1,"powerState":"on"}`,
			errorContain: `0 should be at least 1`,
		},
		{
			name:         "SetAllWithInvalidModeTooHigh",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":6,"fan":1,"powerState":"on"}`,
			errorContain: `6 should be at most 5`,
		},
		{
			name:         "SetAllWithInvalidFanTooLow",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"fan":0,"powerState":"on"}`,
			errorContain: `0 should be at least 1`,
		},
		{
			name:         "SetAllWithInvalidFanTooHigh",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"fan":5,"powerState":"on"}`,
			errorContain: `5 should be at most 4`,
		},
		{
			name:         "SetAllWithInvalidPowerState",
			parameter:    `{"command":"SetAll","temperatureCelsius":25,"mode":1,"fan":1,"powerState":"invalid"}`,
			errorContain: `Value invalid should be one of the allowed values`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteAirConditionerDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_InfraredRemoteTVDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.InfraredRemoteTVDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_InfraredRemoteTVDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "VolumeAdd",
			expectedBody: `{"commandType": "command","command": "volumeAdd","parameter": "default"}`,
			parameter:    `{"command":"VolumeAdd"}`,
		},
		{
			name:         "VolumeSub",
			expectedBody: `{"commandType": "command","command": "volumeSub","parameter": "default"}`,
			parameter:    `{"command":"VolumeSub"}`,
		},
		{
			name:         "ChannelAdd",
			expectedBody: `{"commandType": "command","command": "channelAdd","parameter": "default"}`,
			parameter:    `{"command":"ChannelAdd"}`,
		},
		{
			name:         "ChannelSub",
			expectedBody: `{"commandType": "command","command": "channelSub","parameter": "default"}`,
			parameter:    `{"command":"ChannelSub"}`,
		},
		{
			name:         "SetChannel",
			expectedBody: `{"commandType": "command","command": "SetChannel","parameter": "10"}`,
			parameter:    `{"command":"SetChannel","channel":10}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteTVDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteTVDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "InvalidCommand",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: `Value InvalidCommand should be one of the allowed values`,
		},
		{
			name:         "SetChannelWithoutChannel",
			parameter:    `{"command":"SetChannel"}`,
			errorContain: `Required property 'channel' is missing`,
		},
		{
			name:         "SetChannelWithInvalidChannelTooSmall",
			parameter:    `{"command":"SetChannel","channel":0}`,
			errorContain: `0 should be at least 1`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteTVDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_InfraredRemoteDvdPlayerDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.InfraredRemoteDvdPlayerDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_InfraredRemoteDvdPlayerDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "SetMute",
			expectedBody: `{"commandType": "command","command": "setMute","parameter": "default"}`,
			parameter:    `{"command":"SetMute"}`,
		},
		{
			name:         "FastForward",
			expectedBody: `{"commandType": "command","command": "fastForward","parameter": "default"}`,
			parameter:    `{"command":"FastForward"}`,
		},
		{
			name:         "Rewind",
			expectedBody: `{"commandType": "command","command": "Rewind","parameter": "default"}`,
			parameter:    `{"command":"Rewind"}`,
		},
		{
			name:         "Next",
			expectedBody: `{"commandType": "command","command": "Next","parameter": "default"}`,
			parameter:    `{"command":"Next"}`,
		},
		{
			name:         "Previous",
			expectedBody: `{"commandType": "command","command": "Previous","parameter": "default"}`,
			parameter:    `{"command":"Previous"}`,
		},
		{
			name:         "Pause",
			expectedBody: `{"commandType": "command","command": "Pause","parameter": "default"}`,
			parameter:    `{"command":"Pause"}`,
		},
		{
			name:         "Play",
			expectedBody: `{"commandType": "command","command": "Play","parameter": "default"}`,
			parameter:    `{"command":"Play"}`,
		},
		{
			name:         "Stop",
			expectedBody: `{"commandType": "command","command": "Stop","parameter": "default"}`,
			parameter:    `{"command":"Stop"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteDvdPlayerDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteDvdPlayerDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "InvalidCommand",
			parameter:    `{"command":"InvalidCommand"}`,
			errorContain: `Value InvalidCommand should be one of the allowed values`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteDvdPlayerDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_InfraredRemoteSpeakerDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.InfraredRemoteSpeakerDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_InfraredRemoteSpeakerDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "VolumeAdd",
			expectedBody: `{"commandType": "command","command": "volumeAdd","parameter": "default"}`,
			parameter:    `{"command":"VolumeAdd"}`,
		},
		{
			name:         "VolumeSub",
			expectedBody: `{"commandType": "command","command": "volumeSub","parameter": "default"}`,
			parameter:    `{"command":"VolumeSub"}`,
		},
		{
			name:         "SetMute",
			expectedBody: `{"commandType": "command","command": "setMute","parameter": "default"}`,
			parameter:    `{"command":"SetMute"}`,
		},
		{
			name:         "FastForward",
			expectedBody: `{"commandType": "command","command": "fastForward","parameter": "default"}`,
			parameter:    `{"command":"FastForward"}`,
		},
		{
			name:         "Rewind",
			expectedBody: `{"commandType": "command","command": "Rewind","parameter": "default"}`,
			parameter:    `{"command":"Rewind"}`,
		},
		{
			name:         "Next",
			expectedBody: `{"commandType": "command","command": "Next","parameter": "default"}`,
			parameter:    `{"command":"Next"}`,
		},
		{
			name:         "Previous",
			expectedBody: `{"commandType": "command","command": "Previous","parameter": "default"}`,
			parameter:    `{"command":"Previous"}`,
		},
		{
			name:         "Pause",
			expectedBody: `{"commandType": "command","command": "Pause","parameter": "default"}`,
			parameter:    `{"command":"Pause"}`,
		},
		{
			name:         "Play",
			expectedBody: `{"commandType": "command","command": "Play","parameter": "default"}`,
			parameter:    `{"command":"Play"}`,
		},
		{
			name:         "Stop",
			expectedBody: `{"commandType": "command","command": "Stop","parameter": "default"}`,
			parameter:    `{"command":"Stop"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteSpeakerDevice{
				InfraredRemoteDvdPlayerDevice: switchbot.InfraredRemoteDvdPlayerDevice{
					InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
						Client:   client,
						DeviceID: "ABCDEF123456",
					},
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteSpeakerDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: TurnOn, TurnOff, VolumeAdd, VolumeSub, SetMute, FastForward, Rewind, Next, Previous, Pause, Play, Stop",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteSpeakerDevice{
				InfraredRemoteDvdPlayerDevice: switchbot.InfraredRemoteDvdPlayerDevice{
					InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
						Client:   client,
						DeviceID: "ABCDEF123456",
					},
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}

func Test_InfraredRemoteFanDeviceGetCommandParameterJSONSchema(t *testing.T) {
	device := &switchbot.InfraredRemoteFanDevice{}

	description, err := device.GetCommandParameterJSONSchema()
	assert.NoError(t, err)
	assert.NotEmpty(t, description)
}

func Test_InfraredRemoteFanDeviceExecCommand(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		parameter    string
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			parameter:    `{"command":"TurnOn"}`,
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			parameter:    `{"command":"TurnOff"}`,
		},
		{
			name:         "Swing",
			expectedBody: `{"commandType": "command","command": "swing","parameter": "default"}`,
			parameter:    `{"command":"Swing"}`,
		},
		{
			name:         "Timer",
			expectedBody: `{"commandType": "command","command": "timer","parameter": "default"}`,
			parameter:    `{"command":"Timer"}`,
		},
		{
			name:         "LowSpeed",
			expectedBody: `{"commandType": "command","command": "lowSpeed","parameter": "default"}`,
			parameter:    `{"command":"LowSpeed"}`,
		},
		{
			name:         "MiddleSpeed",
			expectedBody: `{"commandType": "command","command": "middleSpeed","parameter": "default"}`,
			parameter:    `{"command":"MiddleSpeed"}`,
		},
		{
			name:         "HighSpeed",
			expectedBody: `{"commandType": "command","command": "highSpeed","parameter": "default"}`,
			parameter:    `{"command":"HighSpeed"}`,
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteFanDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			response, err := device.ExecCommand(testData.parameter)
			assert.NoError(t, err)
			assert.NotNil(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func Test_InfraredRemoteFanDeviceExecCommandInvalid(t *testing.T) {
	testDataList := []struct {
		name         string
		parameter    string
		errorContain string
	}{
		{
			name:         "Invalid Command",
			parameter:    `{"command": "InvalidCommand"}`,
			errorContain: "Value InvalidCommand should be one of the allowed values: TurnOn, TurnOff, Swing, Timer, LowSpeed, MiddleSpeed, HighSpeed",
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteFanDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:   client,
					DeviceID: "ABCDEF123456",
				},
			}
			_, err := device.ExecCommand(testData.parameter)
			assert.ErrorContains(t, err, testData.errorContain)
		})
	}
}
