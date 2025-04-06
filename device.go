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
	DeviceName         string `json:"deviceName"`
	EnableCloudService bool   `json:"enableCloudService"`
}

type BotDevice struct {
	CommonDeviceListItem
}

type HubDevice struct {
	CommonDeviceListItem
}

type MeterDevice struct {
	CommonDeviceListItem
}

type MeterPlusDevice struct {
	CommonDeviceListItem
}

type OutdoorMeterDevice struct {
	CommonDeviceListItem
}

type MeterProDevice struct {
	CommonDeviceListItem
}

type MeterProCo2Device struct {
	CommonDeviceListItem
}

type MotionSensorDevice struct {
	CommonDeviceListItem
}

type RemoteDevice struct {
	CommonDeviceListItem
}

func GetDevicesResponseParser(response *GetDevicesResponse) ResponseParser {
	return func(bodyBytes []byte) error {
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
			case "Hub", "Hub Plus", "Hub Mini", "Hub 2":
				parsed = &HubDevice{}
			case "Meter":
				parsed = &MeterDevice{}
			case "MeterPlus":
				parsed = &MeterPlusDevice{}
			case "WoIOSensor":
				parsed = &OutdoorMeterDevice{}
			case "MeterPro":
				parsed = &MeterProDevice{}
			case "MeterPro(CO2)":
				parsed = &MeterProCo2Device{}
			case "Motion Sensor":
				parsed = &MotionSensorDevice{}
			case "Remote":
				parsed = &RemoteDevice{}
			default:
				parsed = &CommonDeviceListItem{}
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
