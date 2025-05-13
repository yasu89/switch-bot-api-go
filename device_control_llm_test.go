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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOn\"}",
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOff\"}",
		},
		{
			name:         "Press",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"press\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Press\"}",
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setPosition\",\"parameter\": \"0,ff,75\"}",
			parameter:    "{\"command\":\"SetPosition\",\"mode\":\"ff\", \"position\":75}",
		},
		{
			name:         "TurnOn",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOn\"}",
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOff\"}",
		},
		{
			name:         "Pause",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"pause\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Pause\"}",
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
			parameter:    "{\"command\":\"SetPosition\",\"mode\":\"00\", \"position\":75}",
			errorContain: "Value 00 should be one of the allowed values: 0, 1, ff",
		},
		{
			name:         "SetPosition(invalid position)",
			parameter:    "{\"command\":\"SetPosition\",\"mode\":\"ff\", \"position\":101}",
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetPosition(missing mode)",
			parameter:    "{\"command\":\"SetPosition\", \"position\":80}",
			errorContain: "Required property 'mode' is missing",
		},
		{
			name:         "SetPosition(missing position)",
			parameter:    "{\"command\":\"SetPosition\",\"mode\":\"ff\"}",
			errorContain: "Required property 'position' is missing",
		},
		{
			name:         "Invalid command",
			parameter:    "{\"command\":\"Invalid\"}",
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"lock\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Lock\"}",
		},
		{
			name:         "Unlock",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"unlock\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Unlock\"}",
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
			expectedBody: "{\"commandType\":\"command\",\"command\":\"createKey\",\"parameter\":{\"name\":\"testKey\",\"type\":\"permanent\",\"password\":\"123456\"}}",
			parameter:    "{\"command\":\"CreateKey\",\"name\":\"testKey\", \"type\":\"permanent\",\"password\":\"123456\"}",
		},
		{
			name:         "CreateKey(timeLimit)",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"createKey\",\"parameter\":{\"name\":\"testKey\",\"type\":\"permanent\",\"password\":\"123456\",\"startTime\":1745080854,\"endTime\":1745167254}}",
			parameter:    "{\"command\":\"CreateKey\",\"name\":\"testKey\", \"type\":\"permanent\",\"password\":\"123456\",\"startTime\":1745080854,\"endTime\":1745167254}",
		},
		{
			name:         "DeleteKey",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"deleteKey\",\"parameter\":{\"id\":\"testKey\"}}",
			parameter:    "{\"command\":\"DeleteKey\",\"id\":\"testKey\"}",
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
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"permanent\",\"password\":\"123456\"}",
			errorContain: "Required property 'name' is missing",
		},
		{
			name:         "Create normal key(missing password parameter)",
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"urgent\",\"name\":\"testKey\"}",
			errorContain: "Required property 'password' is missing",
		},
		{
			name:         "Create normal key(invalid password)",
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"permanent\",\"name\":\"testKey\",\"password\":\"12345\"}",
			errorContain: "Value does not match the required pattern ^[0-9]{6,12}$",
		},
		{
			name:         "Create timeLimit key(missing startTime parameter)",
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"timeLimit\",\"name\":\"testKey\",\"password\":\"123456\",\"endTime\":1745167254}",
			errorContain: "Required property 'startTime' is missing",
		},
		{
			name:         "Create timeLimit key(missing endTime parameter)",
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"disposable\",\"name\":\"testKey\",\"password\":\"123456\",\"startTime\":1745167254}",
			errorContain: "Required property 'endTime' is missing",
		},
		{
			name:         "Create timeLimit key(invalid startTime)",
			parameter:    "{\"command\":\"CreateKey\",\"type\":\"timeLimit\",\"name\":\"testKey\",\"password\":\"123456\",\"startTime\":0,\"endTime\":1745167254}",
			errorContain: "0 should be at least 100000000",
		},
		{
			name:         "Delete key(missing id parameter)",
			parameter:    "{\"command\":\"DeleteKey\"}",
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOn\"}",
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOff\"}",
		},
		{
			name:         "Toggle",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"toggle\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Toggle\"}",
		},
		{
			name:         "SetBrightness",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setBrightness\",\"parameter\": \"50\"}",
			parameter:    "{\"command\":\"SetBrightness\",\"brightness\":50}",
		},
		{
			name:         "SetColorTemperature",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setColorTemperature\",\"parameter\": \"3500\"}",
			parameter:    "{\"command\":\"SetColorTemperature\",\"colorTemperature\":3500}",
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
			parameter:    "{\"command\":\"SetBrightness\",\"brightness\":0}",
			errorContain: "0 should be at least 1",
		},
		{
			name:         "SetBrightness(brightness too high)",
			parameter:    "{\"command\":\"SetBrightness\",\"brightness\":101}",
			errorContain: "101 should be at most 100",
		},
		{
			name:         "SetBrightness(missing brightness)",
			parameter:    "{\"command\":\"SetBrightness\"}",
			errorContain: "Required property 'brightness' is missing",
		},
		{
			name:         "SetColorTemperature(invalid colorTemperature too low)",
			parameter:    "{\"command\":\"SetColorTemperature\",\"colorTemperature\":2000}",
			errorContain: "2000 should be at least 2700",
		},
		{
			name:         "SetColorTemperature(invalid colorTemperature too high)",
			parameter:    "{\"command\":\"SetColorTemperature\",\"colorTemperature\":7000}",
			errorContain: "7000 should be at most 6500",
		},
		{
			name:         "SetColorTemperature(missing colorTemperature)",
			parameter:    "{\"command\":\"SetColorTemperature\"}",
			errorContain: "Required property 'colorTemperature' is missing",
		},
		{
			name:         "Invalid command",
			parameter:    "{\"command\":\"Invalid\"}",
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOn\"}",
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOff\"}",
		},
		{
			name:         "Toggle",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"toggle\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"Toggle\"}",
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
			parameter:    "{\"command\":\"Invalid\"}",
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOn\"}",
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			parameter:    "{\"command\":\"TurnOff\"}",
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
			parameter:    "{\"command\":\"Toggle\"}",
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
