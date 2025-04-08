package switchbot

import (
	"encoding/json"
	"fmt"
)

type GetDevicesResponse struct {
	CommonResponse
	Body GetDevicesResponseBody `json:"body"`
}

type GetDevicesResponseBody struct {
	DeviceList []interface{} `json:"deviceList"`
}

type CommonDevice struct {
	DeviceID    string `json:"deviceId"`
	DeviceType  string `json:"deviceType"`
	HubDeviceId string `json:"hubDeviceId"`
}

type CommonDeviceListItem struct {
	CommonDevice
	Client             *Client
	DeviceName         string `json:"deviceName"`
	EnableCloudService bool   `json:"enableCloudService"`
}

type BotDevice struct {
	CommonDeviceListItem
}

type CurtainDevice struct {
	CommonDeviceListItem
	CurtainDevicesIds []string `json:"curtainDevicesIds"`
	Calibrate         bool     `json:"calibrate"`
	Group             bool     `json:"group"`
	Master            bool     `json:"master"`
	OpenDirection     string   `json:"openDirection"`
}

type HubDevice struct {
	CommonDeviceListItem
}

type Hub2Device struct {
	CommonDeviceListItem
}

type MeterDevice struct {
	CommonDeviceListItem
}

type MeterProCo2Device struct {
	CommonDeviceListItem
}

type LockDevice struct {
	CommonDeviceListItem
	Group          bool     `json:"group"`
	Master         bool     `json:"master"`
	GroupName      string   `json:"groupName"`
	LockDevicesIds []string `json:"lockDevicesIds"`
}

type MotionSensorDevice struct {
	CommonDeviceListItem
}

type RemoteDevice struct {
	CommonDeviceListItem
}

func GetDevicesResponseParser(response *GetDevicesResponse) ResponseParser {
	return func(client *Client, bodyBytes []byte) error {
		err := json.Unmarshal(bodyBytes, response)
		if err != nil {
			return err
		}

		var parsedDevices []interface{}
		for _, deviceInterface := range response.Body.DeviceList {
			device, ok := deviceInterface.(map[string]interface{})
			if !ok {
				return fmt.Errorf("failed to cast device to map[string]interface{}")
			}
			jsonString, err := json.Marshal(device)
			if err != nil {
				return err
			}

			deviceType, ok := device["deviceType"].(string)
			if !ok {
				return fmt.Errorf("failed to cast deviceType to string")
			}

			var parsed interface{}
			switch deviceType {
			case "Bot":
				parsed = &BotDevice{}
				parsed.(*BotDevice).Client = client
			case "Curtain", "Curtain3":
				parsed = &CurtainDevice{}
				parsed.(*CurtainDevice).Client = client
			case "Hub", "Hub Plus", "Hub Mini":
				parsed = &HubDevice{}
				parsed.(*HubDevice).Client = client
			case "Hub 2":
				parsed = &Hub2Device{}
				parsed.(*Hub2Device).Client = client
			case "Meter", "MeterPlus", "WoIOSensor", "MeterPro":
				parsed = &MeterDevice{}
				parsed.(*MeterDevice).Client = client
			case "MeterPro(CO2)":
				parsed = &MeterProCo2Device{}
				parsed.(*MeterProCo2Device).Client = client
			case "Smart Lock", "Smart Lock Pro":
				parsed = &LockDevice{}
				parsed.(*LockDevice).Client = client
			case "Motion Sensor":
				parsed = &MotionSensorDevice{}
				parsed.(*MotionSensorDevice).Client = client
			case "Remote":
				parsed = &RemoteDevice{}
				parsed.(*RemoteDevice).Client = client
			default:
				parsed = &CommonDeviceListItem{}
				parsed.(*CommonDeviceListItem).Client = client
			}

			err = json.Unmarshal(jsonString, parsed)
			if err != nil {
				return err
			}
			parsedDevices = append(parsedDevices, parsed)
		}
		response.Body.DeviceList = parsedDevices

		return nil
	}
}

func (client *Client) GetDevices() (*GetDevicesResponse, error) {
	response := &GetDevicesResponse{}
	err := client.GetRequest("/devices", GetDevicesResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}
