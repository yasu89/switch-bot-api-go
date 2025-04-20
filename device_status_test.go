package switchbot_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	switchbot "github.com/yasu89/switch-bot-api-go"
)

func TestGetStatusAndGetAnyStatusBody(t *testing.T) {
	t.Run("BotDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123456/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123456/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123456",
		"deviceType": "Bot",
		"hubDeviceId": "123456789",
		"power": "ON",
		"battery": 100,
		"version": "1.0",
		"deviceMode": "pressMode"
    },
    "message": "success"
}`))
			}),
		)
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CurtainDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123457/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123457/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123457",
		"deviceType": "Curtain",
		"hubDeviceId": "123456789",
		"calibrate": true,
		"group": false,
		"moving": false,
		"battery": 80,
		"version": "1.1",
		"slidePosition": "50"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.CurtainDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123457",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.CurtainDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123457",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("Hub2Device", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123458/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123458/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123458",
		"deviceType": "Hub 2",
		"hubDeviceId": "123456789",
		"temperature": 22.5,
		"lightLevel": 300,
		"version": "1.2",
		"humidity": 60
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.Hub2Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123458",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.Hub2DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123458",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MeterDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123459/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123459/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123459",
		"deviceType": "MeterPlus",
		"hubDeviceId": "123456789",
		"temperature": 25.5,
		"version": "1.3",
		"battery": 90,
		"humidity": 50
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MeterDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123459",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.MeterDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123459",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MeterProCo2Device", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123460/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123460/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123460",
		"deviceType": "MeterPro(CO2)",
		"hubDeviceId": "123456789",
		"temperature": 23.0,
		"version": "1.4",
		"battery": 85,
		"humidity": 55,
		"CO2": 400
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MeterProCo2Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123460",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.MeterProCo2DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123460",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("LockDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123461/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123461/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123461",
		"deviceType": "Smart Lock",
		"hubDeviceId": "123456789",
		"battery": 70,
		"version": "1.5",
		"lockState": "locked",
		"doorState": "closed",
		"calibrate": true
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.LockDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123461",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.LockDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123461",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("KeypadDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123462/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123462/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123462",
		"deviceType": "Keypad Touch",
		"hubDeviceId": "123456789"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.KeypadDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123462",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.KeypadDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123462",
				DeviceType:  "Keypad Touch",
				HubDeviceId: "123456789",
			},
		}

		assertResponse(t, &status.CommonResponse)
		assertBody(t, status.Body, expectedBody)

		// Test GetAnyStatusBody() method
		anyStatus, err := device.GetAnyStatusBody()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("MotionSensorDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123463/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123463/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123463",
		"deviceType": "Motion Sensor",
		"hubDeviceId": "123456789",
		"battery": 80,
		"version": "1.6",
		"moveDetected": true,
		"openState": "closed",
		"brightness": "bright"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.MotionSensorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123463",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.MotionSensorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123463",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("ContactSensorDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123464/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123464/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123464",
		"deviceType": "Contact Sensor",
		"hubDeviceId": "123456789",
		"battery": 75,
		"version": "1.7",
		"moveDetected": false,
		"openState": "open",
		"brightness": "dim"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.ContactSensorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123464",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.ContactSensorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123464",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("WaterLeakDetectorDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123465/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123465/status', got '%s'", r.URL.Path)
				}
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
    "body": {
        "deviceId": "ABCDEF123465",
        "deviceType": "Water Detector",
        "hubDeviceId": "123456789",
        "battery": 60,
        "version": "1.8",
		"status": true
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.WaterLeakDetectorDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123465",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.WaterLeakDetectorDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123465",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CeilingLightDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123466/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123466/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123466",
		"deviceType": "Ceiling Light",
		"hubDeviceId": "123456789",
		"power": "ON",
		"version": "1.9",
		"brightness": 80,
		"colorTemperature": 4000
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.CeilingLightDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123466",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.CeilingLightDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123466",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("PlugMiniDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123467/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123467/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123467",
		"deviceType": "Plug Mini (JP)",
		"hubDeviceId": "123456789",
		"voltage": 220.5,
		"version": "2.0",
		"weight": 1.2,
		"electricityOfDay": 15,
		"electricCurrent": 0.5
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.PlugMiniDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123467",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.PlugMiniDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123467",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("PlugDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123468/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123468/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123468",
		"deviceType": "Plug",
		"hubDeviceId": "123456789",
		"power": "ON",
		"version": "2.1"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.PlugDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123468",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.PlugDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123468",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("StripLightDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123469/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123469/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123469",
		"deviceType": "Strip Light",
		"hubDeviceId": "123456789",
		"power": "OFF",
		"version": "2.2",
		"brightness": 70,
		"color": "#FF5733"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.StripLightDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123469",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.StripLightDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123469",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("ColorBulbDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123470/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123470/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123470",
		"deviceType": "Color Bulb",
		"hubDeviceId": "123456789",
		"power": "ON",
		"brightness": 90,
		"version": "2.3",
		"color": "#00FF00",
		"colorTemperature": 5000
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.ColorBulbDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123470",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.ColorBulbDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123470",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RobotVacuumCleanerDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123471/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123471/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123471",
		"deviceType": "Robot Vacuum Cleaner S1",
		"hubDeviceId": "123456789",
		"workingStatus": "cleaning",
		"onlineStatus": "online",
		"battery": 75
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RobotVacuumCleanerDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123471",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.RobotVacuumCleanerDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123471",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RobotVacuumCleanerS10Device", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123472/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123472/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123472",
		"deviceType": "Robot Vacuum Cleaner S10",
		"hubDeviceId": "123456789",
		"workingStatus": "mopping",
		"onlineStatus": "offline",
		"battery": 50,
		"waterBaseBattery": 80,
		"taskType": "mop"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RobotVacuumCleanerS10Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123472",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.RobotVacuumCleanerS10DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123472",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("HumidifierDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123473/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123473/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123473",
		"deviceType": "Humidifier",
		"hubDeviceId": "123456789",
		"power": "ON",
		"humidity": 45,
		"temperature": 22,
		"nebulizationEfficiency": 3,
		"auto": true,
		"childLock": false,
		"sound": true,
		"lackWater": false
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.HumidifierDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123473",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.HumidifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123473",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("EvaporativeHumidifierDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123474/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123474/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123474",
		"deviceType": "Humidifier2",
		"hubDeviceId": "123456789",
		"power": "OFF",
		"humidity": 50,
		"mode": 2,
		"drying": false,
		"childLock": true,
		"filterElement": {
			"effectiveUsageHours": 100,
			"usedHours": 20
		},
		"version": 1
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.EvaporativeHumidifierDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123474",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.EvaporativeHumidifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123474",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("AirPurifierDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123475/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123475/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123475",
		"deviceType": "Air Purifier VOC",
		"hubDeviceId": "123456789",
		"power": "ON",
		"version": "1.0",
		"mode": 3,
		"childLock": false
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.AirPurifierDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123475",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.AirPurifierDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123475",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("BlindTiltDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123476/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123476/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123476",
		"deviceType": "Blind Tilt",
		"hubDeviceId": "123456789",
		"version": 1,
		"calibrate": true,
		"group": false,
		"moving": false,
		"direction": "up",
		"slidePosition": 50
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.BlindTiltDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123476",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.BlindTiltDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123476",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("BatteryCirculatorFanDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123477/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123477/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123477",
		"deviceType": "Battery Circulator Fan",
		"hubDeviceId": "123456789",
		"mode": "normal",
		"version": "1.0",
		"battery": 85,
		"power": "ON",
		"nightStatus": "OFF",
		"oscillation": "ON",
		"verticalOscillation": "OFF",
		"chargingStatus": "charging",
		"fanSpeed": 3
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.BatteryCirculatorFanDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123477",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.BatteryCirculatorFanDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123477",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("CirculatorFanDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123478/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123478/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123478",
		"deviceType": "Circulator Fan",
		"hubDeviceId": "123456789",
		"mode": "eco",
		"version": "1.1",
		"power": "OFF",
		"nightStatus": "ON",
		"oscillation": "OFF",
		"verticalOscillation": "ON",
		"fanSpeed": 2
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.CirculatorFanDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123478",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.CirculatorFanDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123478",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RollerShadeDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123479/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123479/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123479",
		"deviceType": "Roller Shade",
		"hubDeviceId": "123456789",
		"version": "1.0",
		"calibrate": true,
		"battery": 90,
		"moving": false,
		"slidePosition": 75
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RollerShadeDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123479",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.RollerShadeDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123479",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RelaySwitch1PMDevice", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123480/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123480/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123480",
		"deviceType": "Relay Switch 1PM",
		"hubDeviceId": "123456789",
		"switchStatus": 1,
		"voltage": 220,
		"version": "1.1",
		"power": 50,
		"usedElectricity": 100,
		"electricCurrent": 10
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RelaySwitch1PMDevice{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123480",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.RelaySwitch1PMDeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123480",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		assertBody(t, anyStatus, expectedBody)
	})

	t.Run("RelaySwitch1Device", func(t *testing.T) {
		testServer := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.URL.Path != "/devices/ABCDEF123481/status" {
					t.Fatalf("Expected path '/devices/ABCDEF123481/status', got '%s'", r.URL.Path)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
    "statusCode": 100,
	"body": {
		"deviceId": "ABCDEF123481",
		"deviceType": "Relay Switch 1",
		"hubDeviceId": "123456789",
		"switchStatus": 0,
		"version": "1.2"
    },
    "message": "success"
}`))
			}),
		)
		defer testServer.Close()

		client := switchbot.NewClient("secret", "token", switchbot.OptionBaseApiURL(testServer.URL))

		device := &switchbot.RelaySwitch1Device{
			CommonDeviceListItem: switchbot.CommonDeviceListItem{
				CommonDevice: switchbot.CommonDevice{
					DeviceID: "ABCDEF123481",
				},
				Client: client,
			},
		}
		status, err := device.GetStatus()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expectedBody := &switchbot.RelaySwitch1DeviceStatusBody{
			CommonDevice: switchbot.CommonDevice{
				DeviceID:    "ABCDEF123481",
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
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
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
