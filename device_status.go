package switchbot

import (
	"encoding/json"
	"fmt"
)

type GetDeviceStatusResponse struct {
	CommonResponse
	Body interface{} `json:"body"`
}

type BotDeviceStatus struct {
	CommonDevice
	Power      string `json:"power"`
	Battery    int    `json:"battery"`
	Version    string `json:"version"`
	DeviceMode string `json:"deviceMode"`
}

type MeterDeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
}

type MeterPlusDeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
}

type OutdoorMeterDeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
}

type MeterProDeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
}

type MeterProCo2DeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	Version     string  `json:"version"`
	Battery     int     `json:"battery"`
	Humidity    int     `json:"humidity"`
	CO2         int     `json:"CO2"`
}

type Hub2DeviceStatus struct {
	CommonDevice
	Temperature float64 `json:"temperature"`
	LightLevel  int     `json:"lightLevel"`
	Version     string  `json:"version"`
	Humidity    int     `json:"humidity"`
}

type MotionSensorDeviceStatus struct {
	CommonDevice
	Battery      int    `json:"battery"`
	Version      string `json:"version"`
	MoveDetected bool   `json:"moveDetected"`
	OpenState    string `json:"openState"`
	Brightness   string `json:"brightness"`
}

func GetDeviceStatusResponseParser(response *GetDeviceStatusResponse) ResponseParser {
	return func(client *Client, bodyBytes []byte) error {
		err := json.Unmarshal(bodyBytes, response)
		if err != nil {
			return err
		}

		body, ok := response.Body.(map[string]interface{})
		if !ok {
			return fmt.Errorf("failed to cast body to map[string]interface{}")
		}
		bodyString, err := json.Marshal(body)
		if err != nil {
			return err
		}

		deviceType, ok := body["deviceType"].(string)
		if !ok {
			return fmt.Errorf("failed to cast deviceType to string")
		}

		var parsed interface{}
		switch deviceType {
		case "Bot":
			parsed = &BotDeviceStatus{}
		case "Meter":
			parsed = &MeterDeviceStatus{}
		case "MeterPlus":
			parsed = &MeterPlusDeviceStatus{}
		case "WoIOSensor":
			parsed = &OutdoorMeterDeviceStatus{}
		case "MeterPro":
			parsed = &MeterProDeviceStatus{}
		case "MeterPro(CO2)":
			parsed = &MeterProCo2DeviceStatus{}
		case "Hub 2":
			parsed = &Hub2DeviceStatus{}
		case "MotionSensor":
			parsed = &MotionSensorDeviceStatus{}
		default:
			parsed = &CommonDevice{}
		}
		err = json.Unmarshal(bodyString, parsed)
		if err != nil {
			return err
		}
		response.Body = parsed

		return nil
	}
}

func (client *Client) GetDeviceStatus(deviceId string) (*GetDeviceStatusResponse, error) {
	response := &GetDeviceStatusResponse{}
	err := client.GetRequest("/devices/"+deviceId+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}
