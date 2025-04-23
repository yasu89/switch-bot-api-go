package helpers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/yasu89/switch-bot-api-go"
)

// HttpMockHandler is a struct that represents a mock HTTP handler.
type HttpMockHandler struct {
	Method  string
	Path    string
	Count   int
	Handler http.HandlerFunc
}

// IsMatch checks if the method and path match the handler's method and path.
func (h *HttpMockHandler) IsMatch(method string, path string) bool {
	if method != h.Method {
		return false
	}

	if path != h.Path {
		return false
	}

	return true
}

// SwitchBotMock is a mock for the SwitchBot API.
type SwitchBotMock struct {
	t        *testing.T
	handlers []*HttpMockHandler
}

// NewSwitchBotMock creates a new instance of SwitchBotMock.
func NewSwitchBotMock(t *testing.T) *SwitchBotMock {
	return &SwitchBotMock{
		t: t,
	}
}

// RegisterDevicesMock registers a mock response for the device's endpoint.
func (s *SwitchBotMock) RegisterDevicesMock(devices []interface{}, infraredDevices []interface{}) {
	s.handlers = append(s.handlers, &HttpMockHandler{
		Method: http.MethodGet,
		Path:   "/devices",
		Count:  0,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			response := switchbot.GetDevicesResponse{
				CommonResponse: switchbot.CommonResponse{
					StatusCode: 100,
					Message:    "success",
				},
				Body: switchbot.GetDevicesResponseBody{
					DeviceList:         devices,
					InfraredRemoteList: infraredDevices,
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
		},
	})
}

// RegisterStatusMock registers a mock response for a specific device's status.
func (s *SwitchBotMock) RegisterStatusMock(deviceId string, mockBody interface{}) {
	s.handlers = append(s.handlers, &HttpMockHandler{
		Method: http.MethodGet,
		Path:   "/devices/" + deviceId + "/status",
		Count:  0,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			response := struct {
				switchbot.CommonResponse
				Body interface{} `json:"body"`
			}{
				CommonResponse: switchbot.CommonResponse{
					StatusCode: 100,
					Message:    "success",
				},
				Body: mockBody,
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
		},
	})
}

// RegisterCommandMock registers a mock response for a specific device's command.
func (s *SwitchBotMock) RegisterCommandMock(deviceId string, expectedBody string) {
	s.handlers = append(s.handlers, &HttpMockHandler{
		Method: http.MethodPost,
		Path:   "/devices/" + deviceId + "/commands",
		Count:  0,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			var expectedObject map[string]interface{}
			if err := json.Unmarshal([]byte(expectedBody), &expectedObject); err != nil {
				s.t.Fatalf("Failed to unmarshal expected body: %v", err)
			}

			var actualObject map[string]interface{}
			if err := json.NewDecoder(r.Body).Decode(&actualObject); err != nil {
				s.t.Fatalf("Failed to decode actual body: %v", err)
			}
			if !reflect.DeepEqual(expectedObject, actualObject) {
				s.t.Fatalf("Expected body %v, got %v", expectedObject, actualObject)
			}

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			_, err := w.Write([]byte(`{
    "statusCode": 100,
	"body": {},
    "message": "success"
}`))
			if err != nil {
				s.t.Fatalf("Failed to write response: %v", err)
			}
		},
	})
}

// AssertCallCount checks the number of times a specific method and path were called.
func (s *SwitchBotMock) AssertCallCount(method string, path string, expected int) {
	for _, handler := range s.handlers {
		if handler.IsMatch(method, path) {
			if handler.Count != expected {
				s.t.Fatalf("Expected %d calls to %s %s, got %d", expected, method, path, handler.Count)
			}
			return
		}
	}
	s.t.Fatalf("No handler found for %s %s", method, path)
}

// NewTestServer creates a new test server with the mock handlers.
func (s *SwitchBotMock) NewTestServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, handler := range s.handlers {
				if handler.IsMatch(r.Method, r.URL.Path) {
					handler.Count++
					handler.Handler(w, r)
					return
				}
			}
			s.t.Fatalf("No handler found for %s %s", r.Method, r.URL.Path)
		}),
	)
}
