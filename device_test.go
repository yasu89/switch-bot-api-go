package switchbot_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/yasu89/switch-bot-api-go"
	"github.com/yasu89/switch-bot-api-go/helpers"
)

func TestGetPhysicalDevices(t *testing.T) {
	// Tests do not cover devices that only embed CommonDeviceListItem/InfraredRemoteDevice with the same structure.

	t.Run("BotDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ABCDEF123456",
			"deviceType":         "Bot",
			"hubDeviceId":        "123456789",
			"deviceName":         "BotDevice",
			"enableCloudService": true,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.BotDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ABCDEF123456",
						DeviceType:  "Bot",
						HubDeviceId: "123456789",
					},
					Client:             client,
					DeviceName:         "BotDevice",
					EnableCloudService: true,
				},
			},
		})
	})

	t.Run("CurtainDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ABCDEF123456",
			"deviceType":         "Curtain3",
			"hubDeviceId":        "123456789",
			"deviceName":         "CurtainDevice",
			"enableCloudService": false,
			"curtainDevicesIds":  []string{"BBBBBBBB", "CCCCCCCC", "DDDDDDDD"},
			"calibrate":          true,
			"group":              true,
			"master":             false,
			"openDirection":      "left",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.CurtainDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ABCDEF123456",
						DeviceType:  "Curtain3",
						HubDeviceId: "123456789",
					},
					Client:             client,
					DeviceName:         "CurtainDevice",
					EnableCloudService: false,
				},
				CurtainDevicesIds: []string{"BBBBBBBB", "CCCCCCCC", "DDDDDDDD"},
				Calibrate:         true,
				Group:             true,
				Master:            false,
				OpenDirection:     "left",
			},
		})
	})

	t.Run("LockDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ABCDEF123456",
			"deviceType":         "Smart Lock",
			"hubDeviceId":        "123456789",
			"deviceName":         "LockDevice",
			"enableCloudService": true,
			"group":              true,
			"master":             false,
			"groupName":          "LockGroup",
			"lockDevicesIds":     []string{"AAAAAAAA"},
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.LockDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ABCDEF123456",
						DeviceType:  "Smart Lock",
						HubDeviceId: "123456789",
					},
					Client:             client,
					DeviceName:         "LockDevice",
					EnableCloudService: true,
				},
				Group:          true,
				Master:         false,
				GroupName:      "LockGroup",
				LockDevicesIds: []string{"AAAAAAAA"},
			},
		})
	})

	t.Run("KeypadDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ABCDEF123456",
			"deviceType":         "Keypad Touch",
			"hubDeviceId":        "123456789",
			"deviceName":         "KeypadDevice",
			"enableCloudService": true,
			"lockDevicesIds":     []string{},
			"keyList": []map[string]interface{}{
				{
					"id":         1,
					"name":       "firstKey",
					"type":       "permanent",
					"password":   "z41UoV93PIS0OYElzUd7nwA9TO6XxSDlf9N+P4nFuJw=",
					"iv":         "71fbf00383b6e214dc08b8b94183cf30",
					"status":     "normal",
					"createTime": 1744814218,
				},
				{
					"id":         2,
					"name":       "secondKey",
					"type":       "timeLimit",
					"password":   "z6aSgCwa4+0a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5afa=",
					"iv":         "af0b1a2c3d4e5f6g7h8i9j0k1l2m3n4o5",
					"status":     "expired",
					"createTime": 1746025200,
				},
			},
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.KeypadDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ABCDEF123456",
						DeviceType:  "Keypad Touch",
						HubDeviceId: "123456789",
					},
					Client:             client,
					DeviceName:         "KeypadDevice",
					EnableCloudService: true,
				},
				LockDevicesIds: []string{},
				KeyList: []switchbot.KeyListItem{
					{
						Id:         1,
						Name:       "firstKey",
						Type:       "permanent",
						Password:   "z41UoV93PIS0OYElzUd7nwA9TO6XxSDlf9N+P4nFuJw=",
						Iv:         "71fbf00383b6e214dc08b8b94183cf30",
						Status:     "normal",
						CreateTime: 1744814218,
					},
					{
						Id:         2,
						Name:       "secondKey",
						Type:       "timeLimit",
						Password:   "z6aSgCwa4+0a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5afa=",
						Iv:         "af0b1a2c3d4e5f6g7h8i9j0k1l2m3n4o5",
						Status:     "expired",
						CreateTime: 1746025200,
					},
				},
			},
		})
	})

	t.Run("BlindTiltDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":            "XYZ123456789",
			"deviceType":          "Blind Tilt",
			"hubDeviceId":         "987654321",
			"deviceName":          "BlindTiltDevice",
			"enableCloudService":  true,
			"version":             1,
			"blindTiltDevicesIds": []string{"BBBBBBBB"},
			"calibrate":           true,
			"group":               true,
			"master":              false,
			"direction":           "up",
			"slidePosition":       50,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.BlindTiltDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "XYZ123456789",
						DeviceType:  "Blind Tilt",
						HubDeviceId: "987654321",
					},
					Client:             client,
					DeviceName:         "BlindTiltDevice",
					EnableCloudService: true,
				},
				Version:             1,
				BlindTiltDevicesIds: []string{"BBBBBBBB"},
				Calibrate:           true,
				Group:               true,
				Master:              false,
				Direction:           "up",
				SlidePosition:       50,
			},
		})
	})

	t.Run("RollerShadeDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ROLLER123456",
			"deviceType":         "Roller Shade",
			"hubDeviceId":        "HUB123456",
			"deviceName":         "RollerShadeDevice",
			"enableCloudService": true,
			"bleVersion":         "1.0",
			"groupingDevicesIds": []string{"GROUP123"},
			"group":              true,
			"master":             true,
			"groupName":          "RollerGroup",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.RollerShadeDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ROLLER123456",
						DeviceType:  "Roller Shade",
						HubDeviceId: "HUB123456",
					},
					Client:             client,
					DeviceName:         "RollerShadeDevice",
					EnableCloudService: true,
				},
				BleVersion:         "1.0",
				GroupingDevicesIds: []string{"GROUP123"},
				Group:              true,
				Master:             true,
				GroupName:          "RollerGroup",
			},
		})
	})

	t.Run("TwoDevices", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "ABCDEF123456",
			"deviceType":         "Hub 2",
			"hubDeviceId":        "123456789",
			"deviceName":         "Hub2Device",
			"enableCloudService": true,
		})
		switchBotMock.AddDevice(map[string]interface{}{
			"deviceId":           "GHIJKL987654",
			"deviceType":         "MeterPro(CO2)",
			"hubDeviceId":        "987654321",
			"deviceName":         "MeterProCO2Device",
			"enableCloudService": true,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()
		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		response, err := client.GetDevices()
		if err != nil {
			t.Fatal(err)
		}

		assertDevices(t, response, []interface{}{
			&switchbot.Hub2Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "ABCDEF123456",
						DeviceType:  "Hub 2",
						HubDeviceId: "123456789",
					},
					Client:             client,
					DeviceName:         "Hub2Device",
					EnableCloudService: true,
				},
			},
			&switchbot.MeterProCo2Device{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID:    "GHIJKL987654",
						DeviceType:  "MeterPro(CO2)",
						HubDeviceId: "987654321",
					},
					Client:             client,
					DeviceName:         "MeterProCO2Device",
					EnableCloudService: true,
				},
			},
		})
	})

}

func TestGetInfraredRemoteDevices(t *testing.T) {
	switchBotMock := helpers.NewSwitchBotMock(t)
	switchBotMock.AddInfraredDevice(map[string]interface{}{
		"deviceId":    "AIRCONDITIONER123456",
		"deviceName":  "AirConditionerDevice",
		"remoteType":  "Air Conditioner",
		"hubDeviceId": "HUB123456",
	})
	switchBotMock.AddInfraredDevice(map[string]interface{}{
		"deviceId":    "SETTOPBOX123456",
		"deviceName":  "SetTopBoxDevice",
		"remoteType":  "Set Top Box",
		"hubDeviceId": "HUB123456",
	})
	switchBotMock.AddInfraredDevice(map[string]interface{}{
		"deviceId":    "OTHERDEVICE123456",
		"deviceName":  "OthersDevice",
		"remoteType":  "Others",
		"hubDeviceId": "HUB123456",
	})
	testServer := switchBotMock.NewTestServer()

	client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

	response, err := client.GetDevices()
	if err != nil {
		t.Fatal(err)
	}

	assertInfraredDevices(t, response, []interface{}{
		&switchbot.InfraredRemoteAirConditionerDevice{
			InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
				Client:      client,
				DeviceID:    "AIRCONDITIONER123456",
				DeviceName:  "AirConditionerDevice",
				RemoteType:  "Air Conditioner",
				HubDeviceId: "HUB123456",
			},
		},
		&switchbot.InfraredRemoteSetTopBoxDevice{
			InfraredRemoteTVDevice: switchbot.InfraredRemoteTVDevice{
				InfraredRemoteDevice: switchbot.InfraredRemoteDevice{
					Client:      client,
					DeviceID:    "SETTOPBOX123456",
					DeviceName:  "SetTopBoxDevice",
					RemoteType:  "Set Top Box",
					HubDeviceId: "HUB123456",
				},
			},
		},
		&switchbot.InfraredRemoteOthersDevice{
			Client:      client,
			DeviceID:    "OTHERDEVICE123456",
			DeviceName:  "OthersDevice",
			RemoteType:  "Others",
			HubDeviceId: "HUB123456",
		},
	})
}

func assertDevices(t *testing.T, response *switchbot.GetDevicesResponse, expectedList []interface{}) {
	t.Helper()

	if response.StatusCode != 100 {
		t.Fatalf("Expected status code 100, got %d", response.StatusCode)
	}

	if len(response.Body.DeviceList) != len(expectedList) {
		t.Fatalf("Expected %d device, got %d", len(expectedList), len(response.Body.DeviceList))
	}

	if len(response.Body.InfraredRemoteList) != 0 {
		t.Fatalf("Expected 0 infrared remote, got %d", len(response.Body.InfraredRemoteList))
	}

	for i, device := range response.Body.DeviceList {
		if reflect.TypeOf(device) != reflect.TypeOf(expectedList[i]) {
			t.Fatalf("Expected type %T, got %T", expectedList[i], device)
		}

		if !reflect.DeepEqual(device, expectedList[i]) {
			t.Fatalf("expected %s, actual %s", jsonDump(t, expectedList[i]), jsonDump(t, device))
		}
	}
}

func assertInfraredDevices(t *testing.T, response *switchbot.GetDevicesResponse, expectedList []interface{}) {
	t.Helper()

	if response.StatusCode != 100 {
		t.Fatalf("Expected status code 100, got %d", response.StatusCode)
	}

	if len(response.Body.DeviceList) != 0 {
		t.Fatalf("Expected 0 device, got %d", len(response.Body.DeviceList))
	}

	if len(response.Body.InfraredRemoteList) != len(expectedList) {
		t.Fatalf("Expected %d infrared remote, got %d", len(expectedList), len(response.Body.InfraredRemoteList))
	}

	for i, infrared := range response.Body.InfraredRemoteList {
		if reflect.TypeOf(infrared) != reflect.TypeOf(expectedList[i]) {
			t.Fatalf("Expected type %T, got %T", expectedList[i], infrared)
		}

		if !reflect.DeepEqual(infrared, expectedList[i]) {
			t.Fatalf("expected %s, actual %s", jsonDump(t, expectedList[i]), jsonDump(t, infrared))
		}
	}
}

func jsonDump(t *testing.T, data interface{}) string {
	t.Helper()

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}
	return string(b)
}
