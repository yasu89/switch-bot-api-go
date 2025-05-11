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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}
