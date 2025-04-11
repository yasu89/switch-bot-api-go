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
	DeviceList         []interface{} `json:"deviceList"`
	InfraredRemoteList []interface{} `json:"infraredRemoteList"`
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

type InfraredRemoteDevice struct {
	Client      *Client
	DeviceID    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	RemoteType  string `json:"remoteType"`
	HubDeviceId string `json:"hubDeviceId"`
}

// InfraredRemoteAirConditionerDevice represents an infrared remote-controlled air conditioner device.
type InfraredRemoteAirConditionerDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteTVDevice represents an infrared remote-controlled TV device.
type InfraredRemoteTVDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteLightDevice represents an infrared remote-controlled light device.
type InfraredRemoteLightDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteStreamerDevice represents an infrared remote-controlled streamer device.
type InfraredRemoteStreamerDevice struct {
	InfraredRemoteTVDevice
}

// InfraredRemoteSetTopBoxDevice represents an infrared remote-controlled set-top box device.
type InfraredRemoteSetTopBoxDevice struct {
	InfraredRemoteTVDevice
}

// InfraredRemoteDvdPlayerDevice represents an infrared remote-controlled DVD player device.
type InfraredRemoteDvdPlayerDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteFanDevice represents an infrared remote-controlled fan device.
type InfraredRemoteFanDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteProjectorDevice represents an infrared remote-controlled projector device.
type InfraredRemoteProjectorDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteCameraDevice represents an infrared remote-controlled camera device.
type InfraredRemoteCameraDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteAirPurifierDevice represents an infrared remote-controlled air purifier device.
type InfraredRemoteAirPurifierDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteSpeakerDevice represents an infrared remote-controlled speaker device.
type InfraredRemoteSpeakerDevice struct {
	InfraredRemoteDvdPlayerDevice
}

// InfraredRemoteWaterHeaterDevice represents an infrared remote-controlled water heater device.
type InfraredRemoteWaterHeaterDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteRobotVacuumCleanerDevice represents an infrared remote-controlled robot vacuum cleaner device.
type InfraredRemoteRobotVacuumCleanerDevice struct {
	InfraredRemoteDevice
}

// InfraredRemoteOthersDevice represents an infrared remote-controlled device of other types.
type InfraredRemoteOthersDevice struct {
	Client      *Client
	DeviceID    string `json:"deviceId"`
	DeviceName  string `json:"deviceName"`
	RemoteType  string `json:"remoteType"`
	HubDeviceId string `json:"hubDeviceId"`
}

func GetDevicesResponseParser(response *GetDevicesResponse) ResponseParser {
	return func(client *Client, bodyBytes []byte) error {
		err := json.Unmarshal(bodyBytes, response)
		if err != nil {
			return err
		}

		// Parse the device list
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

		// Set the Client for each InfraredRemoteDevice
		var parsedInfraredRemoteDevices []interface{}
		for _, infraredRemoteDeviceInterface := range response.Body.InfraredRemoteList {
			infraredRemoteDevice, ok := infraredRemoteDeviceInterface.(map[string]interface{})
			if !ok {
				return fmt.Errorf("failed to cast infraredRemoteDevice to map[string]interface{}")
			}
			jsonString, err := json.Marshal(infraredRemoteDevice)
			if err != nil {
				return err
			}

			remoteType, ok := infraredRemoteDevice["remoteType"].(string)
			if !ok {
				return fmt.Errorf("failed to cast remoteType to string")
			}

			var parsedInfrared interface{}
			switch remoteType {
			case "Air Conditioner":
				parsedInfrared = &InfraredRemoteAirConditionerDevice{}
				parsedInfrared.(*InfraredRemoteAirConditionerDevice).Client = client
			case "TV":
				parsedInfrared = &InfraredRemoteTVDevice{}
				parsedInfrared.(*InfraredRemoteTVDevice).Client = client
			case "Light":
				parsedInfrared = &InfraredRemoteLightDevice{}
				parsedInfrared.(*InfraredRemoteLightDevice).Client = client
			case "Streamer":
				parsedInfrared = &InfraredRemoteStreamerDevice{}
				parsedInfrared.(*InfraredRemoteStreamerDevice).Client = client
			case "Set Top Box":
				parsedInfrared = &InfraredRemoteSetTopBoxDevice{}
				parsedInfrared.(*InfraredRemoteSetTopBoxDevice).Client = client
			case "DVD Player":
				parsedInfrared = &InfraredRemoteDvdPlayerDevice{}
				parsedInfrared.(*InfraredRemoteDvdPlayerDevice).Client = client
			case "Fan":
				parsedInfrared = &InfraredRemoteFanDevice{}
				parsedInfrared.(*InfraredRemoteFanDevice).Client = client
			case "Projector":
				parsedInfrared = &InfraredRemoteProjectorDevice{}
				parsedInfrared.(*InfraredRemoteProjectorDevice).Client = client
			case "Camera":
				parsedInfrared = &InfraredRemoteCameraDevice{}
				parsedInfrared.(*InfraredRemoteCameraDevice).Client = client
			case "Air Purifier":
				parsedInfrared = &InfraredRemoteAirPurifierDevice{}
				parsedInfrared.(*InfraredRemoteAirPurifierDevice).Client = client
			case "Speaker":
				parsedInfrared = &InfraredRemoteSpeakerDevice{}
				parsedInfrared.(*InfraredRemoteSpeakerDevice).Client = client
			case "Water Heater":
				parsedInfrared = &InfraredRemoteWaterHeaterDevice{}
				parsedInfrared.(*InfraredRemoteWaterHeaterDevice).Client = client
			case "Robot Vacuum Cleaner":
				parsedInfrared = &InfraredRemoteRobotVacuumCleanerDevice{}
				parsedInfrared.(*InfraredRemoteRobotVacuumCleanerDevice).Client = client
			case "Others":
				parsedInfrared = &InfraredRemoteOthersDevice{}
				parsedInfrared.(*InfraredRemoteOthersDevice).Client = client
			default:
				parsedInfrared = &InfraredRemoteDevice{}
				parsedInfrared.(*InfraredRemoteDevice).Client = client
			}

			err = json.Unmarshal(jsonString, parsedInfrared)
			if err != nil {
				return err
			}
			parsedInfraredRemoteDevices = append(parsedInfraredRemoteDevices, parsedInfrared)
		}
		response.Body.InfraredRemoteList = parsedInfraredRemoteDevices

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
