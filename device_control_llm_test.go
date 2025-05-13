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
			assert.ErrorContains(t, err, testData.errorContain)
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
			expectedBody: `{"commandType": "command","command": "changeParam","parameter": {"fanLevel":3,"waterLevel":2,"times":100}}`,
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
