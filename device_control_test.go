package switchbot_test

import (
	"encoding/json"
	"github.com/yasu89/switch-bot-api-go"
	"image/color"
	"net/http"
	"net/http/httptest"
	"reflect"
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Press",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"press\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.BotDevice) (*switchbot.CommonResponse, error) { return device.Press() },
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setPosition\",\"parameter\": \"0,ff,75\"}",
			method: func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) {
				return device.SetPosition(switchbot.CurtainPositionModeDefault, 75)
			},
		},
		{
			name:         "TurnOn",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Pause",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"pause\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CurtainDevice) (*switchbot.CommonResponse, error) { return device.Pause() },
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\":\"command\",\"command\":\"createKey\",\"parameter\":{\"name\":\"testKey\",\"type\":\"permanent\",\"password\":\"123456\"}}",
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
			expectedBody: "{\"commandType\":\"command\",\"command\":\"createKey\",\"parameter\":{\"name\":\"testKey\",\"type\":\"permanent\",\"password\":\"123456\",\"startTime\":1745080854,\"endTime\":1745167254}}",
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
			expectedBody: "{\"commandType\":\"command\",\"command\":\"deleteKey\",\"parameter\":{\"id\":\"testKey\"}}",
			method: func(device *switchbot.KeypadDevice) (*switchbot.CommonResponse, error) {
				return device.DeleteKey("testKey")
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"lock\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.LockDevice) (*switchbot.CommonResponse, error) { return device.Lock() },
		},
		{
			name:         "Unlock",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"unlock\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.LockDevice) (*switchbot.CommonResponse, error) { return device.Unlock() },
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOn\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.TurnOn() },
		},
		{
			name:         "TurnOff",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"turnOff\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.TurnOff() },
		},
		{
			name:         "Toggle",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"toggle\",\"parameter\": \"default\"}",
			method:       func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) { return device.Toggle() },
		},
		{
			name:         "SetBrightness",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setBrightness\",\"parameter\": \"50\"}",
			method: func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(50)
			},
		},
		{
			name:         "SetColorTemperature",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setColorTemperature\",\"parameter\": \"3500\"}",
			method: func(device *switchbot.CeilingLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetColorTemperature(3500)
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setBrightness\",\"parameter\": \"55\"}",
			method: func(device *switchbot.StripLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(55)
			},
		},
		{
			name:         "SetColor",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setColor\",\"parameter\": \"255:100:0\"}",
			method: func(device *switchbot.StripLightDevice) (*switchbot.CommonResponse, error) {
				return device.SetColor(color.RGBA{R: 255, G: 100, B: 0, A: 0})
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setBrightness\",\"parameter\": \"55\"}",
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetBrightness(55)
			},
		},
		{
			name:         "SetColor",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setColor\",\"parameter\": \"255:100:0\"}",
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetColor(color.RGBA{R: 255, G: 100, B: 0, A: 0})
			},
		},
		{
			name:         "SetColorTemperature",
			expectedBody: "{\"commandType\": \"command\",\"command\": \"setColorTemperature\",\"parameter\": \"5000\"}",
			method: func(device *switchbot.ColorBulbDevice) (*switchbot.CommonResponse, error) {
				return device.SetColorTemperature(5000)
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
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
			expectedBody: "{\"commandType\": \"command\",\"command\": \"PowLevel\",\"parameter\": \"3\"}",
			method: func(device *switchbot.RobotVacuumCleanerDevice) (*switchbot.CommonResponse, error) {
				return device.SetPowerLevel(switchbot.RobotVacuumCleanerPowerLevelMax)
			},
		},
	}

	for _, testData := range testDataList {
		t.Run(testData.name, func(t *testing.T) {
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
		})
	}
}

func TestRobotVacuumCleanerS10Device(t *testing.T) {
	testDataList := []struct {
		name         string
		expectedBody string
		method       func(*switchbot.RobotVacuumCleanerS10Device) (*switchbot.CommonResponse, error)
	}{
		{
			name:         "StartClean",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"startClean\",\"parameter\":{\"action\":\"sweep\",\"param\":{\"fanLevel\":1,\"waterLevel\":2,\"times\":100}}}",
			method: func(device *switchbot.RobotVacuumCleanerS10Device) (*switchbot.CommonResponse, error) {
				startFloorCleaningParam, err := switchbot.NewStartFloorCleaningParam(switchbot.FloorCleaningActionSweep, 1, 2, 100)
				if err != nil {
					return nil, err
				}
				return device.StartClean(startFloorCleaningParam)
			},
		},
		{
			name:         "SetVolume",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"setVolume\",\"parameter\":\"30\"}",
			method: func(device *switchbot.RobotVacuumCleanerS10Device) (*switchbot.CommonResponse, error) {
				return device.SetVolume(30)
			},
		},
		{
			name:         "SelfClean",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"selfClean\",\"parameter\":\"1\"}",
			method: func(device *switchbot.RobotVacuumCleanerS10Device) (*switchbot.CommonResponse, error) {
				return device.SelfClean(switchbot.WashMopSelfCleaningMode)
			},
		},
		{
			name:         "ChangeParam",
			expectedBody: "{\"commandType\":\"command\",\"command\":\"changeParam\",\"parameter\":{\"fanLevel\":2,\"waterLevel\":1,\"times\":20000}}",
			method: func(device *switchbot.RobotVacuumCleanerS10Device) (*switchbot.CommonResponse, error) {
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
			testServer := newTestCommandServer(t, testData.expectedBody)
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
			response, err := testData.method(device)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			assertResponse(t, response)
		})
	}
}

func newTestCommandServer(t *testing.T, expectedBody string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path != "/devices/ABCDEF123456/commands" {
				t.Fatalf("Expected path '/devices/ABCDEF123456/commands', got '%s'", r.URL.Path)
			}

			var expectedObject map[string]interface{}
			if err := json.Unmarshal([]byte(expectedBody), &expectedObject); err != nil {
				t.Fatalf("Failed to unmarshal expected body: %v", err)
			}

			var actualObject map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&actualObject); err != nil {
				t.Fatalf("Failed to decode actual body: %v", err)
			}

			if !reflect.DeepEqual(expectedObject, actualObject) {
				t.Fatalf("Expected body %v, got %v", expectedObject, actualObject)
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{
    "statusCode": 100,
	"body": {},
    "message": "success"
}`))
		}),
	)
}
