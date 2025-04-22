package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/yasu89/switch-bot-api-go"
)

type SwitchBotMock struct {
	t               *testing.T
	devices         []interface{}
	infraredDevices []interface{}
	handlers        map[string]interface{}
}

func NewSwitchBotMock(t *testing.T) *SwitchBotMock {
	return &SwitchBotMock{
		t: t,
	}
}

// AddDevice adds a device to the mock server's response.
func (s *SwitchBotMock) AddDevice(device interface{}) {
	s.devices = append(s.devices, device)
}

// AddInfraredDevice adds an infrared device to the mock server's response.
func (s *SwitchBotMock) AddInfraredDevice(device interface{}) {
	s.infraredDevices = append(s.infraredDevices, device)
}

func (s *SwitchBotMock) NewTestServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/devices" {
				if r.Method != http.MethodGet {
					s.t.Fatalf("/devices API expected GET request, got %s", r.Method)
				}

				response := switchbot.GetDevicesResponse{
					CommonResponse: switchbot.CommonResponse{
						StatusCode: 100,
						Message:    "success",
					},
					Body: switchbot.GetDevicesResponseBody{
						DeviceList:         s.devices,
						InfraredRemoteList: s.infraredDevices,
					},
				}
				responseJsonText, err := json.Marshal(response)
				if err != nil {
					s.t.Fatalf("Failed to marshal response: %v", err)
				}

				w.WriteHeader(http.StatusOK)
				w.Header().Set("Content-Type", "application/json")
				_, err = w.Write(responseJsonText)
				if err != nil {
					s.t.Fatalf("Failed to write response: %v", err)
				}
			}
		}),
	)
}
