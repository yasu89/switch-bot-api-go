package switchbot_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/yasu89/switch-bot-api-go"
)

func TestBotDevice(t *testing.T) {
	t.Run("TurnOn", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "turnOn",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.TurnOn()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})

	t.Run("TurnOff", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "turnOff",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.TurnOff()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})

	t.Run("Press", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "press",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.Press()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})
}

func TestCurtainDevice(t *testing.T) {
	t.Run("SetPosition", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "setPosition",
	"parameter": "0,ff,75"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.SetPosition(switchbot.CurtainPositionModeDefault, 75)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})

	t.Run("TurnOn", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "turnOn",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.TurnOn()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})

	t.Run("TurnOff", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "turnOff",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.TurnOff()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})

	t.Run("Pause", func(t *testing.T) {
		expectedBody := `
{
	"commandType": "command",
	"command": "pause",
	"parameter": "default"
}`

		testServer := newTestCommandServer(t, expectedBody)
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
		response, err := device.Pause()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		assertResponse(t, response)
	})
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
