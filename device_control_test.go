package switchbot_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/yasu89/switch-bot-api-go"
	"github.com/yasu89/switch-bot-api-go/helpers"
	"image/color"
	"net/http"
	"testing"
)

func TestBotDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.BotDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Press",
			expectedBody: `{"commandType": "command","command": "press","parameter": "default"}`,
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.Press() },
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestCurtainDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.CurtainDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetPosition",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "0,ff,75"}`,
			method: func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) {
				return device.SetPosition(switchbot.CurtainPositionModeDefault, 75)
			},
		},
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Pause",
			expectedBody: `{"commandType": "command","command": "pause","parameter": "default"}`,
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.Pause() },
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestKeypadDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.KeypadDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "CreateKey(permanent)",
			expectedBody: `{"commandType":"command","command":"createKey","parameter":{"name":"testKey","type":"permanent","password":"123456"}}`,
			method: func(device *switchbot.KeypadDevice) (*switchbot.CommonResponse, error) {
				keypadKey, err := switchbot.NewKeypadKey("testKey", "permanent", "123456", 0, 0)
				if err != nil {
					return nil, err
				}
				return device.CreateKey(keypadKey)
			},
		},
		{
			name:         "CreateKey(timeLimit)",
			expectedBody: `{"commandType":"command","command":"createKey","parameter":{"name":"testKey","type":"permanent","password":"123456","startTime":1745080854,"endTime":1745167254}}`,
			method: func(device *switchbot.KeypadDevice) (*switchbot.CommonResponse, error) {
				keypadKey, err := switchbot.NewKeypadKey("testKey", "permanent", "123456", 1745080854, 1745167254)
				if err != nil {
					return nil, err
				}
				return device.CreateKey(keypadKey)
			},
		},
		{
			name:         "DeleteKey",
			expectedBody: `{"commandType":"command","command":"deleteKey","parameter":{"id":"testKey"}}`,
			method: func(device *switchbot.KeypadDevice) (*switchbot.CommonResponse, error) {
				return device.DeleteKey("testKey")
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestLockDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.LockDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "Lock",
			expectedBody: `{"commandType": "command","command": "lock","parameter": "default"}`,
			method:       func(device *switchbot.LockDevice) (*switchbot.CommonResponse, error) { return device.Lock() },
		},
		{
			name:         "Unlock",
			expectedBody: `{"commandType": "command","command": "unlock","parameter": "default"}`,
			method:       func(device *switchbot.LockDevice) (*switchbot.CommonResponse, error) { return device.Unlock() },
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestCeilingLightDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Toggle",
			expectedBody: `{"commandType": "command","command": "toggle","parameter": "default"}`,
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.Toggle() },
		},
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "50"}`,
			method: func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(50)
			},
		},
		{
			name:         "SetColorTemperature",
			expectedBody: `{"commandType": "command","command": "setColorTemperature","parameter": "3500"}`,
			method: func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetColorTemperature(3500)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestStripLightDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.StripLightDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "55"}`,
			method: func(device *switchbot.StripLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(55)
			},
		},
		{
			name:         "SetColor",
			expectedBody: `{"commandType": "command","command": "setColor","parameter": "255:100:0"}`,
			method: func(device *switchbot.StripLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetColor(color.RGBA{R: 255, G: 100, B: 0, A: 0})
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestColorBulbDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetBrightness",
			expectedBody: `{"commandType": "command","command": "setBrightness","parameter": "55"}`,
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(55)
			},
		},
		{
			name:         "SetColor",
			expectedBody: `{"commandType": "command","command": "setColor","parameter": "255:100:0"}`,
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetColor(color.RGBA{R: 255, G: 100, B: 0, A: 0})
			},
		},
		{
			name:         "SetColorTemperature",
			expectedBody: `{"commandType": "command","command": "setColorTemperature","parameter": "5000"}`,
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetColorTemperature(5000)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestRobotVacuumCleanerDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RobotVacuumCleanerDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetPowerLevel",
			expectedBody: `{"commandType": "command","command": "PowLevel","parameter": "3"}`,
			method: func(device *switchbot.RobotVacuumCleanerDevice) (*switchbot.CommonResponse, error) {
				return device.SetPowerLevel(switchbot.RobotVacuumCleanerPowerLevelMax)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestRobotVacuumCleanerSDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RobotVacuumCleanerSDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "StartClean",
			expectedBody: `{"commandType":"command","command":"startClean","parameter":{"action":"sweep","param":{"fanLevel":1,"waterLevel":2,"times":100}}}`,
			method: func(device *switchbot.RobotVacuumCleanerSDevice) (*switchbot.CommonResponse, error) {
				startFloorCleaningParam, err := switchbot.NewStartFloorCleaningParam(switchbot.FloorCleaningActionSweep, 1, 2, 100)
				if err != nil {
					return nil, err
				}
				return device.StartClean(startFloorCleaningParam)
			},
		},
		{
			name:         "SetVolume",
			expectedBody: `{"commandType":"command","command":"setVolume","parameter":"30"}`,
			method: func(device *switchbot.RobotVacuumCleanerSDevice) (*switchbot.CommonResponse, error) {
				return device.SetVolume(30)
			},
		},
		{
			name:         "SelfClean",
			expectedBody: `{"commandType":"command","command":"selfClean","parameter":"1"}`,
			method: func(device *switchbot.RobotVacuumCleanerSDevice) (*switchbot.CommonResponse, error) {
				return device.SelfClean(switchbot.WashMopSelfCleaningMode)
			},
		},
		{
			name:         "ChangeParam",
			expectedBody: `{"commandType":"command","command":"changeParam","parameter":{"fanLevel":2,"waterLevel":1,"times":20000}}`,
			method: func(device *switchbot.RobotVacuumCleanerSDevice) (*switchbot.CommonResponse, error) {
				floorCleaningParam, err := switchbot.NewFloorCleaningParam(2, 1, 20000)
				if err != nil {
					return nil, err
				}
				return device.ChangeParam(floorCleaningParam)
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.RobotVacuumCleanerSDevice{
				CommonDeviceListItem: switchbot.CommonDeviceListItem{
					CommonDevice: switchbot.CommonDevice{
						DeviceID: "ABCDEF123456",
					},
					Client: client,
				},
			}
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestHumidifierDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.HumidifierDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetMode",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "102"}`,
			method: func(device *switchbot.HumidifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.HumidifierModeMedium)
			},
		},
		{
			name:         "SetTargetHumidity",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "60"}`,
			method: func(device *switchbot.HumidifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetTargetHumidity(60)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestEvaporativeHumidifierDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.EvaporativeHumidifierDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetMode",
			expectedBody: `{"commandType":"command","command":"setMode","parameter":{"mode":7,"targetHumidity":60}}`,
			method: func(device *switchbot.EvaporativeHumidifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.EvaporativeHumidifierModeAuto, 60)
			},
		},
		{
			name:         "SetChildLock",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": "true"}`,
			method: func(device *switchbot.EvaporativeHumidifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetChildLock(true)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestAirPurifierDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.AirPurifierDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetMode(Normal)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":1,"fanGear":3}}`,
			method: func(device *switchbot.AirPurifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.AirPurifierModeNormal, 3)
			},
		},
		{
			name:         "SetMode(Auto)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": {"mode":2}}`,
			method: func(device *switchbot.AirPurifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.AirPurifierModeAuto, 0)
			},
		},
		{
			name:         "SetChildLock",
			expectedBody: `{"commandType": "command","command": "setChildLock","parameter": 1}`,
			method: func(device *switchbot.AirPurifierDevice) (*switchbot.CommonResponse, error) {
				return device.SetChildLock(true)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestBlindTiltDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.BlindTiltDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetPosition(up)",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "up;50"}`,
			method: func(device *switchbot.BlindTiltDevice) (*switchbot.CommonResponse, error) {
				return device.SetPosition("up", 50)
			},
		},
		{
			name:         "SetPosition(down)",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "down;80"}`,
			method: func(device *switchbot.BlindTiltDevice) (*switchbot.CommonResponse, error) {
				return device.SetPosition("down", 80)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestBatteryCirculatorFanDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.BatteryCirculatorFanDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetNightLightMode",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "1"}`,
			method: func(device *switchbot.BatteryCirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetNightLightMode(switchbot.CirculatorNightLightModeTurnBright)
			},
		},
		{
			name:         "SetWindMode",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "natural"}`,
			method: func(device *switchbot.BatteryCirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetWindMode(switchbot.CirculatorWindModeNatural)
			},
		},
		{
			name:         "SetWindSpeed",
			expectedBody: `{"commandType": "command","command": "setWindSpeed","parameter": "50"}`,
			method: func(device *switchbot.BatteryCirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetWindSpeed(50)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestCirculatorFanDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.CirculatorFanDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetNightLightMode",
			expectedBody: `{"commandType": "command","command": "setNightLightMode","parameter": "2"}`,
			method: func(device *switchbot.CirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetNightLightMode(switchbot.CirculatorNightLightModeTurnDim)
			},
		},
		{
			name:         "SetWindMode",
			expectedBody: `{"commandType": "command","command": "setWindMode","parameter": "sleep"}`,
			method: func(device *switchbot.CirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetWindMode(switchbot.CirculatorWindModeSleep)
			},
		},
		{
			name:         "SetWindSpeed",
			expectedBody: `{"commandType": "command","command": "setWindSpeed","parameter": "75"}`,
			method: func(device *switchbot.CirculatorFanDevice) (*switchbot.CommonResponse, error) {
				return device.SetWindSpeed(75)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestRollerShadeDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RollerShadeDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetPosition",
			expectedBody: `{"commandType": "command","command": "setPosition","parameter": "50"}`,
			method: func(device *switchbot.RollerShadeDevice) (*switchbot.CommonResponse, error) {
				return device.SetPosition(50)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestRelaySwitch1PMDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RelaySwitch1PMDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetMode(Toggle)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "0"}`,
			method: func(device *switchbot.RelaySwitch1PMDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.RelaySwitchModeToggle)
			},
		},
		{
			name:         "SetMode(Momentary)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "3"}`,
			method: func(device *switchbot.RelaySwitch1PMDevice) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.RelaySwitchModeMomentary)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestRelaySwitch1Device(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RelaySwitch1Device) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetMode(Toggle)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "0"}`,
			method: func(device *switchbot.RelaySwitch1Device) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.RelaySwitchModeToggle)
			},
		},
		{
			name:         "SetMode(Momentary)",
			expectedBody: `{"commandType": "command","command": "setMode","parameter": "3"}`,
			method: func(device *switchbot.RelaySwitch1Device) (*switchbot.CommonResponse, error) {
				return device.SetMode(switchbot.RelaySwitchModeMomentary)
			},
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
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestInfraredRemoteAirConditionerDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.InfraredRemoteAirConditionerDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "TurnOn",
			expectedBody: `{"commandType": "command","command": "turnOn","parameter": "default"}`,
			method: func(device *switchbot.InfraredRemoteAirConditionerDevice) (*switchbot.CommonResponse, error) {
				return device.TurnOn()
			},
		},
		{
			name:         "TurnOff",
			expectedBody: `{"commandType": "command","command": "turnOff","parameter": "default"}`,
			method: func(device *switchbot.InfraredRemoteAirConditionerDevice) (*switchbot.CommonResponse, error) {
				return device.TurnOff()
			},
		},
		{
			name:         "SetAll",
			expectedBody: `{"commandType": "command","command": "setAll","parameter": "25,2,4,on"}`,
			method: func(device *switchbot.InfraredRemoteAirConditionerDevice) (*switchbot.CommonResponse, error) {
				return device.SetAll(25, switchbot.AirConditionerModeCool, switchbot.AirConditionerFanModeHigh, switchbot.AirConditionerPowerStateOn)
			},
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
					Client:     client,
					DeviceID:   "ABCDEF123456",
					DeviceName: "Test AC",
				},
			}
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestInfraredRemoteTVDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.InfraredRemoteTVDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "SetChannel",
			expectedBody: `{"commandType": "command","command": "SetChannel","parameter": "5"}`,
			method: func(device *switchbot.InfraredRemoteTVDevice) (*switchbot.CommonResponse, error) {
				return device.SetChannel(5)
			},
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
					Client:     client,
					DeviceID:   "ABCDEF123456",
					DeviceName: "Test TV",
				},
			}
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}

func TestInfraredRemoteOthersDevice(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.InfraredRemoteOthersDevice) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "CustomCommand",
			expectedBody: `{"commandType": "command","command": "customize","parameter": "testButton"}`,
			method: func(device *switchbot.InfraredRemoteOthersDevice) (*switchbot.CommonResponse, error) {
				return device.CustomCommand("testButton")
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			switchBotMock := helpers.NewSwitchBotMock(t)
			switchBotMock.RegisterCommandMock("ABCDEF123456", testData.expectedBody)
			testServer := switchBotMock.NewTestServer()
			defer testServer.Close()

			client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))
			device := &switchbot.InfraredRemoteOthersDevice{
				Client:     client,
				DeviceID:   "ABCDEF123456",
				DeviceName: "Test Others",
			}
			response, err := testData.method(device)
			assert.NoError(t, err)
			assertResponse(t, response)
			switchBotMock.AssertCallCount(http.MethodPost, "/devices/ABCDEF123456/commands", 1)
		})
	}
}
