package switchbot_test

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	switchbot "github.com/yasu89/switch-bot-api-go"
	"github.com/yasu89/switch-bot-api-go/helpers"
)

func TestGetStatusAndGetAnyStatusBody(t *testing.T) {
	t.Run("BotDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Bot",
			"hubDeviceId": "123456789",
			"power":       "ON",
			"battery":     100,
			"version":     "1.0",
			"deviceMode":  "pressMode",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.BotDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Bot",
				HubDeviceId: "123456789",
			},
			Power:      "ON",
			Battery:    100,
			Version:    "1.0",
			DeviceMode: "pressMode",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CurtainDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":      "ABCDEF123456",
			"deviceType":    "Curtain",
			"hubDeviceId":   "123456789",
			"calibrate":     true,
			"group":         false,
			"moving":        false,
			"battery":       80,
			"version":       "1.1",
			"slidePosition": "50",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.CurtainDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Curtain",
				HubDeviceId: "123456789",
			},
			Calibrate:     true,
			Group:         false,
			Moving:        false,
			Battery:       80,
			Version:       "1.1",
			SlidePosition: "50",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("Hub2Device", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Hub 2",
			"hubDeviceId": "123456789",
			"temperature": 22.5,
			"lightLevel":  300,
			"version":     "1.2",
			"humidity":    60,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.Hub2Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.Hub2DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Hub 2",
				HubDeviceId: "123456789",
			},
			Temperature: 22.5,
			LightLevel:  300,
			Version:     "1.2",
			Humidity:    60,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("Hub3Device", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":     "ABCDEF123456",
			"deviceType":   "Hub 3",
			"hubDeviceId":  "123456789",
			"temperature":  23.5,
			"lightLevel":   400,
			"version":      "1.3",
			"humidity":     65,
			"moveDetected": true,
			"online":       "online",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.Hub3Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.Hub3DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Hub 3",
				HubDeviceId: "123456789",
			},
			Temperature:  23.5,
			LightLevel:   400,
			Version:      "1.3",
			Humidity:     65,
			MoveDetected: true,
			Online:       "online",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MeterDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "MeterPlus",
			"hubDeviceId": "123456789",
			"temperature": 25.5,
			"version":     "1.3",
			"battery":     90,
			"humidity":    50,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MeterDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.MeterDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "MeterPlus",
				HubDeviceId: "123456789",
			},
			Temperature: 25.5,
			Version:     "1.3",
			Battery:     90,
			Humidity:    50,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MeterProCo2Device", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "MeterPro(CO2)",
			"hubDeviceId": "123456789",
			"temperature": 23.0,
			"version":     "1.4",
			"battery":     85,
			"humidity":    55,
			"CO2":         400,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MeterProCo2Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.MeterProCo2DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "MeterPro(CO2)",
				HubDeviceId: "123456789",
			},
			Temperature: 23.0,
			Version:     "1.4",
			Battery:     85,
			Humidity:    55,
			CO2:         400,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("LockDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Smart Lock",
			"hubDeviceId": "123456789",
			"battery":     70,
			"version":     "1.5",
			"lockState":   "locked",
			"doorState":   "closed",
			"calibrate":   true,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.LockDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Smart Lock",
				HubDeviceId: "123456789",
			},
			Battery:   70,
			Version:   "1.5",
			LockState: "locked",
			DoorState: "closed",
			Calibrate: true,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("KeypadDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Keypad Touch",
			"hubDeviceId": "123456789",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.KeypadDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Keypad Touch",
				HubDeviceId: "123456789",
			},
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MotionSensorDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":     "ABCDEF123456",
			"deviceType":   "Motion Sensor",
			"hubDeviceId":  "123456789",
			"battery":      80,
			"version":      "1.6",
			"moveDetected": true,
			"openState":    "closed",
			"brightness":   "bright",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MotionSensorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.MotionSensorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Motion Sensor",
				HubDeviceId: "123456789",
			},
			Battery:      80,
			Version:      "1.6",
			MoveDetected: true,
			OpenState:    "closed",
			Brightness:   "bright",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("ContactSensorDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":     "ABCDEF123456",
			"deviceType":   "Contact Sensor",
			"hubDeviceId":  "123456789",
			"battery":      75,
			"version":      "1.7",
			"moveDetected": false,
			"openState":    "open",
			"brightness":   "dim",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.ContactSensorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.ContactSensorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Contact Sensor",
				HubDeviceId: "123456789",
			},
			Battery:      75,
			Version:      "1.7",
			MoveDetected: false,
			OpenState:    "open",
			Brightness:   "dim",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("WaterLeakDetectorDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Water Detector",
			"hubDeviceId": "123456789",
			"battery":     60,
			"version":     "1.8",
			"status":      true,
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.WaterLeakDetectorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.WaterLeakDetectorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Water Detector",
				HubDeviceId: "123456789",
			},
			Battery: 60,
			Version: "1.8",
			Status:  true,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CeilingLightDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":         "ABCDEF123456",
			"deviceType":       "Ceiling Light",
			"hubDeviceId":      "123456789",
			"power":            "ON",
			"version":          "1.9",
			"brightness":       80,
			"colorTemperature": 4000,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.CeilingLightDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Ceiling Light",
				HubDeviceId: "123456789",
			},
			Power:            "ON",
			Version:          "1.9",
			Brightness:       80,
			ColorTemperature: 4000,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("PlugMiniDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":         "ABCDEF123456",
			"deviceType":       "Plug Mini (JP)",
			"hubDeviceId":      "123456789",
			"voltage":          220.5,
			"version":          "2.0",
			"weight":           1.2,
			"electricityOfDay": 15,
			"electricCurrent":  0.5,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.PlugMiniDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Plug Mini (JP)",
				HubDeviceId: "123456789",
			},
			Voltage:          220.5,
			Version:          "2.0",
			Weight:           1.2,
			ElectricityOfDay: 15,
			ElectricCurrent:  0.5,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("PlugDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Plug",
			"hubDeviceId": "123456789",
			"power":       "ON",
			"version":     "2.1",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.PlugDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Plug",
				HubDeviceId: "123456789",
			},
			Power:   "ON",
			Version: "2.1",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("StripLightDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Strip Light",
			"hubDeviceId": "123456789",
			"power":       "OFF",
			"version":     "2.2",
			"brightness":  70,
			"color":       "#FF5733",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.StripLightDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Strip Light",
				HubDeviceId: "123456789",
			},
			Power:      "OFF",
			Version:    "2.2",
			Brightness: 70,
			Color:      "#FF5733",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("ColorBulbDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":         "ABCDEF123456",
			"deviceType":       "Color Bulb",
			"hubDeviceId":      "123456789",
			"power":            "ON",
			"brightness":       90,
			"version":          "2.3",
			"color":            "#00FF00",
			"colorTemperature": 5000,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.ColorBulbDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Color Bulb",
				HubDeviceId: "123456789",
			},
			Power:            "ON",
			Brightness:       90,
			Version:          "2.3",
			Color:            "#00FF00",
			ColorTemperature: 5000,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RobotVacuumCleanerDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":      "ABCDEF123456",
			"deviceType":    "Robot Vacuum Cleaner S1",
			"hubDeviceId":   "123456789",
			"workingStatus": "cleaning",
			"onlineStatus":  "online",
			"battery":       75,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RobotVacuumCleanerDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Robot Vacuum Cleaner S1",
				HubDeviceId: "123456789",
			},
			WorkingStatus: "cleaning",
			OnlineStatus:  "online",
			Battery:       75,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RobotVacuumCleanerSDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":         "ABCDEF123456",
			"deviceType":       "Robot Vacuum Cleaner S10",
			"hubDeviceId":      "123456789",
			"workingStatus":    "mopping",
			"onlineStatus":     "offline",
			"battery":          50,
			"waterBaseBattery": 80,
			"taskType":         "mop",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RobotVacuumCleanerSDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Robot Vacuum Cleaner S10",
				HubDeviceId: "123456789",
			},
			WorkingStatus:    "mopping",
			OnlineStatus:     "offline",
			Battery:          50,
			WaterBaseBattery: 80,
			TaskType:         "mop",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RobotVacuumCleanerComboDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":      "ABCDEF123456",
			"deviceType":    "Robot Vacuum Cleaner K10+ Pro Combo",
			"hubDeviceId":   "123456789",
			"workingStatus": "Clearing",
			"onlineStatus":  "online",
			"battery":       80,
			"taskType":      "backToCharge",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RobotVacuumCleanerComboDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}

		status, err := device.GetStatus()
		assert.NoError(t, err)

		expectedBody := &switchbot.RobotVacuumCleanerComboDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Robot Vacuum Cleaner K10+ Pro Combo",
				HubDeviceId: "123456789",
			},
			WorkingStatus: "Clearing",
			OnlineStatus:  "online",
			Battery:       80,
			TaskType:      "backToCharge",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("HumidifierDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":               "ABCDEF123456",
			"deviceType":             "Humidifier",
			"hubDeviceId":            "123456789",
			"power":                  "ON",
			"humidity":               45,
			"temperature":            22,
			"nebulizationEfficiency": 3,
			"auto":                   true,
			"childLock":              false,
			"sound":                  true,
			"lackWater":              false,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.HumidifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Humidifier",
				HubDeviceId: "123456789",
			},
			Power:                  "ON",
			Humidity:               45,
			Temperature:            22,
			NebulizationEfficiency: 3,
			Auto:                   true,
			ChildLock:              false,
			Sound:                  true,
			LackWater:              false,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("EvaporativeHumidifierDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Humidifier2",
			"hubDeviceId": "123456789",
			"power":       "OFF",
			"humidity":    50,
			"mode":        2,
			"drying":      false,
			"childLock":   true,
			"filterElement": map[string]interface{}{
				"effectiveUsageHours": 100,
				"usedHours":           20,
			},
			"version": 1,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.EvaporativeHumidifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Humidifier2",
				HubDeviceId: "123456789",
			},
			Power:     "OFF",
			Humidity:  50,
			Mode:      2,
			Drying:    false,
			ChildLock: true,
			FilterElement: switchbot.EvaporativeHumidifierDeviceFilterElement{
				EffectiveUsageHours: 100,
				UsedHours:           20,
			},
			Version: 1,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("AirPurifierDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":    "ABCDEF123456",
			"deviceType":  "Air Purifier VOC",
			"hubDeviceId": "123456789",
			"power":       "ON",
			"version":     "1.0",
			"mode":        3,
			"childLock":   false,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.AirPurifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Air Purifier VOC",
				HubDeviceId: "123456789",
			},
			Power:     "ON",
			Version:   "1.0",
			Mode:      3,
			ChildLock: false,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("BlindTiltDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":      "ABCDEF123456",
			"deviceType":    "Blind Tilt",
			"hubDeviceId":   "123456789",
			"version":       1,
			"calibrate":     true,
			"group":         false,
			"moving":        false,
			"direction":     "up",
			"slidePosition": 50,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.BlindTiltDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Blind Tilt",
				HubDeviceId: "123456789",
			},
			Version:       1,
			Calibrate:     true,
			Group:         false,
			Moving:        false,
			Direction:     "up",
			SlidePosition: 50,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("BatteryCirculatorFanDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":            "ABCDEF123456",
			"deviceType":          "Battery Circulator Fan",
			"hubDeviceId":         "123456789",
			"mode":                "normal",
			"version":             "1.0",
			"battery":             85,
			"power":               "ON",
			"nightStatus":         "OFF",
			"oscillation":         "ON",
			"verticalOscillation": "OFF",
			"chargingStatus":      "charging",
			"fanSpeed":            3,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.BatteryCirculatorFanDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Battery Circulator Fan",
				HubDeviceId: "123456789",
			},
			Mode:                "normal",
			Version:             "1.0",
			Battery:             85,
			Power:               "ON",
			NightStatus:         "OFF",
			Oscillation:         "ON",
			VerticalOscillation: "OFF",
			ChargingStatus:      "charging",
			FanSpeed:            3,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CirculatorFanDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":            "ABCDEF123456",
			"deviceType":          "Circulator Fan",
			"hubDeviceId":         "123456789",
			"mode":                "eco",
			"version":             "1.1",
			"power":               "OFF",
			"nightStatus":         "ON",
			"oscillation":         "OFF",
			"verticalOscillation": "ON",
			"fanSpeed":            2,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.CirculatorFanDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Circulator Fan",
				HubDeviceId: "123456789",
			},
			Mode:                "eco",
			Version:             "1.1",
			Power:               "OFF",
			NightStatus:         "ON",
			Oscillation:         "OFF",
			VerticalOscillation: "ON",
			FanSpeed:            2,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RollerShadeDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":      "ABCDEF123456",
			"deviceType":    "Roller Shade",
			"hubDeviceId":   "123456789",
			"version":       "1.0",
			"calibrate":     true,
			"battery":       90,
			"moving":        false,
			"slidePosition": 75,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RollerShadeDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Roller Shade",
				HubDeviceId: "123456789",
			},
			Version:       "1.0",
			Calibrate:     true,
			Battery:       90,
			Moving:        false,
			SlidePosition: 75,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RelaySwitch1PMDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":        "ABCDEF123456",
			"deviceType":      "Relay Switch 1PM",
			"hubDeviceId":     "123456789",
			"switchStatus":    1,
			"voltage":         220,
			"version":         "1.1",
			"power":           50,
			"usedElectricity": 100,
			"electricCurrent": 10,
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RelaySwitch1PMDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Relay Switch 1PM",
				HubDeviceId: "123456789",
			},
			SwitchStatus:    1,
			Voltage:         220,
			Version:         "1.1",
			Power:           50,
			UsedElectricity: 100,
			ElectricCurrent: 10,
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RelaySwitch1Device", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":     "ABCDEF123456",
			"deviceType":   "Relay Switch 1",
			"hubDeviceId":  "123456789",
			"switchStatus": 0,
			"version":      "1.2",
		})
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
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RelaySwitch1DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Relay Switch 1",
				HubDeviceId: "123456789",
			},
			SwitchStatus: 0,
			Version:      "1.2",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RelaySwitch2PMDevice", func(t *testing.T) {
		switchBotMock := helpers.NewSwitchBotMock(t)
		switchBotMock.RegisterStatusMock("ABCDEF123456", map[string]interface{}{
			"deviceId":               "ABCDEF123456",
			"deviceType":             "Relay Switch 2PM",
			"hubDeviceId":            "123456789",
			"online":                 true,
			"switch1Status":          1,
			"switch2Status":          0,
			"switch1voltage":         120,
			"switch2voltage":         0,
			"version":                "V3.1-6.3",
			"switch1power":           50,
			"switch2power":           0,
			"switch1usedElectricity": 1000,
			"switch2usedElectricity": 0,
			"switch1electricCurrent": 500,
			"switch2electricCurrent": 0,
			"calibrate":              true,
			"position":               50,
			"isStuck":                "false",
		})
		testServer := switchBotMock.NewTestServer()
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RelaySwitch2PMDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123456",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		assert.NoError(t, err)

		switchBotMock.AssertCallCount(http.MethodGet, "/devices/ABCDEF123456/status", 1)

		expectedBody := &switchbot.RelaySwitch2PMDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123456",
				DeviceType:  "Relay Switch 2PM",
				HubDeviceId: "123456789",
			},
			Online:                 true,
			Switch1Status:          1,
			Switch2Status:          0,
			Switch1Voltage:         120,
			Switch2Voltage:         0,
			Version:                "V3.1-6.3",
			Switch1Power:           50,
			Switch2Power:           0,
			Switch1UsedElectricity: 1000,
			Switch2UsedElectricity: 0,
			Switch1ElectricCurrent: 500,
			Switch2ElectricCurrent: 0,
			Calibrate:              true,
			Position:               50,
			IsStuck:                "false",
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		assert.NoError(t, err)
		assertBody(t, anyStatus, expectedBody)
	})
}

func assertResponse(t *testing.T, response *switchbot.CommonResponse) {
	t.Helper()

	if response.StatusCode != 100 {
		t.Fatalf("Expected status code 100, got %d", response.StatusCode)
	}

	if response.Message != "success" {
		t.Fatalf("Expected message 'success', got '%s'", response.Message)
	}
}

func assertBody(t *testing.T, response interface{}, expected interface{}) {
	t.Helper()

	if !reflect.DeepEqual(response, expected) {
		t.Fatalf("Expected body %s, got %s", jsonDump(t, expected), jsonDump(t, response))
	}
}
