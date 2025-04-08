package switchbot

import (
	"encoding/json"
)

func GetDeviceStatusResponseParser(response interface{}) ResponseParser {
	return func(client *Client, bodyBytes []byte) error {
		err := json.Unmarshal(bodyBytes, response)
		if err != nil {
			return err
		}
		return nil
	}
}

type BotDeviceStatusBody struct {
	CommonDevice
	Power      string `json:"power"`
	Battery    int    `json:"battery"`
	Version    string `json:"version"`
	DeviceMode string `json:"deviceMode"`
}

type BotDeviceStatusResponse struct {
	CommonResponse
	Body *BotDeviceStatusBody `json:"body"`
}

func (device *BotDevice) GetStatus() (*BotDeviceStatusResponse, error) {
	response := &BotDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type Hub2DeviceStatusBody struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	LightLevel  int     `json:"lightLevel"`
	Version     string  `json:"version"`
	Humidity    int     `json:"humidity"`
}

type Hub2DeviceStatusResponse struct {
	CommonResponse
	Body *Hub2DeviceStatusBody `json:"body"`
}

func (device *Hub2Device) GetStatus() (*Hub2DeviceStatusResponse, error) {
	response := &Hub2DeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type MeterDeviceStatusBody struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
}

type MeterDeviceStatusResponse struct {
	CommonResponse
	Body *MeterDeviceStatusBody `json:"body"`
}

func (device *MeterDevice) GetStatus() (*MeterDeviceStatusResponse, error) {
	response := &MeterDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type MeterProCo2DeviceStatusBody struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
	CO2         int     `json:"CO2"`
}

type MeterProCo2DeviceStatusResponse struct {
	CommonResponse
	Body *MeterProCo2DeviceStatusBody `json:"body"`
}

func (device *MeterProCo2Device) GetStatus() (*MeterProCo2DeviceStatusResponse, error) {
	response := &MeterProCo2DeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type LockDeviceStatusBody struct {
	CommonDevice
	Battery   int    `json:"battery"`
	Version   string `json:"version"`
	LockState string `json:"lockState"`
	DoorState string `json:"doorState"`
	Calibrate bool   `json:"calibrate"`
}

type LockDeviceStatusResponse struct {
	CommonResponse
	Body *LockDeviceStatusBody `json:"body"`
}

func (device *LockDevice) GetStatus() (*LockDeviceStatusResponse, error) {
	response := &LockDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type MotionSensorDeviceStatusBody struct {
	CommonDevice
	Battery      int    `json:"battery"`
	Version      string `json:"version"`
	MoveDetected bool   `json:"moveDetected"`
	OpenState    string `json:"openState"`
	Brightness   string `json:"brightness"`
}

type MotionSensorDeviceStatusResponse struct {
	CommonResponse
	Body *MotionSensorDeviceStatusBody `json:"body"`
}

func (device *MotionSensorDevice) GetStatus() (*MotionSensorDeviceStatusResponse, error) {
	response := &MotionSensorDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}
