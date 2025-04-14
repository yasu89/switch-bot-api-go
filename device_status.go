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

type KeypadDeviceStatusBody struct {
	CommonDevice
}

type KeypadStatusResponse struct {
	CommonResponse
	Body *KeypadDeviceStatusBody `json:"body"`
}

func (device *KeypadDevice) GetStatus() (*KeypadStatusResponse, error) {
	response := &KeypadStatusResponse{}
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

type ContactSensorDeviceStatusBody struct {
	CommonDevice
	Battery      int    `json:"battery"`
	Version      string `json:"version"`
	MoveDetected bool   `json:"moveDetected"`
	OpenState    string `json:"openState"`
	Brightness   string `json:"brightness"`
}

type ContactSensorDeviceStatusResponse struct {
	CommonResponse
	Body *ContactSensorDeviceStatusBody `json:"body"`
}

func (device *ContactSensorDevice) GetStatus() (*ContactSensorDeviceStatusResponse, error) {
	response := &ContactSensorDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type WaterLeakDetectorDeviceStatusBody struct {
	CommonDevice
	Battery int    `json:"battery"`
	Version string `json:"version"`
	Status  bool   `json:"status"`
}

type WaterLeakDetectorDeviceStatusResponse struct {
	CommonResponse
	Body *WaterLeakDetectorDeviceStatusBody `json:"body"`
}

func (device *WaterLeakDetectorDevice) GetStatus() (*WaterLeakDetectorDeviceStatusResponse, error) {
	response := &WaterLeakDetectorDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type CeilingLightDeviceStatusBody struct {
	CommonDevice
	Power            string `json:"power"`
	Version          string `json:"version"`
	Brightness       int    `json:"brightness"`
	ColorTemperature int    `json:"colorTemperature"`
}

type CeilingLightDeviceStatusResponse struct {
	CommonResponse
	Body *CeilingLightDeviceStatusBody `json:"body"`
}

func (device *CeilingLightDevice) GetStatus() (*CeilingLightDeviceStatusResponse, error) {
	response := &CeilingLightDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type PlugMiniDeviceStatusBody struct {
	CommonDevice
	Voltage          float64 `json:"voltage"`
	Version          string  `json:"version"`
	Weight           float64 `json:"weight"`
	ElectricityOfDay int     `json:"electricityOfDay"`
	ElectricCurrent  float64 `json:"electricCurrent"`
}

type PlugMiniDeviceStatusResponse struct {
	CommonResponse
	Body *PlugMiniDeviceStatusBody `json:"body"`
}

func (device *PlugMiniDevice) GetStatus() (*PlugMiniDeviceStatusResponse, error) {
	response := &PlugMiniDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type PlugDeviceStatusBody struct {
	CommonDevice
	Power   string `json:"power"`
	Version string `json:"version"`
}

type PlugDeviceStatusResponse struct {
	CommonResponse
	Body *PlugDeviceStatusBody `json:"body"`
}

func (device *PlugDevice) GetStatus() (*PlugDeviceStatusResponse, error) {
	response := &PlugDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type StripLightDeviceStatusBody struct {
	CommonDevice
	Power      string `json:"power"`
	Version    string `json:"version"`
	Brightness int    `json:"brightness"`
	Color      string `json:"color"`
}

type StripLightDeviceStatusResponse struct {
	CommonResponse
	Body *StripLightDeviceStatusBody `json:"body"`
}

func (device *StripLightDevice) GetStatus() (*StripLightDeviceStatusResponse, error) {
	response := &StripLightDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type ColorBulbDeviceStatusBody struct {
	CommonDevice
	Power            string `json:"power"`
	Brightness       int    `json:"brightness"`
	Version          string `json:"version"`
	Color            string `json:"color"`
	ColorTemperature int    `json:"colorTemperature"`
}

type ColorBulbDeviceStatusResponse struct {
	CommonResponse
	Body *ColorBulbDeviceStatusBody `json:"body"`
}

func (device *ColorBulbDevice) GetStatus() (*ColorBulbDeviceStatusResponse, error) {
	response := &ColorBulbDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type RobotVacuumCleanerDeviceStatusBody struct {
	CommonDevice
	WorkingStatus string `json:"workingStatus"`
	OnlineStatus  string `json:"onlineStatus"`
	Battery       int    `json:"battery"`
}

type RobotVacuumCleanerDeviceStatusResponse struct {
	CommonResponse
	Body *RobotVacuumCleanerDeviceStatusBody `json:"body"`
}

func (device *RobotVacuumCleanerDevice) GetStatus() (*RobotVacuumCleanerDeviceStatusResponse, error) {
	response := &RobotVacuumCleanerDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type RobotVacuumCleanerS10DeviceStatusBody struct {
	CommonDevice
	WorkingStatus    string `json:"workingStatus"`
	OnlineStatus     string `json:"onlineStatus"`
	Battery          int    `json:"battery"`
	WaterBaseBattery int    `json:"waterBaseBattery"`
	TaskType         string `json:"taskType"`
}

type RobotVacuumCleanerS10DeviceStatusResponse struct {
	CommonResponse
	Body *RobotVacuumCleanerS10DeviceStatusBody `json:"body"`
}

func (device *RobotVacuumCleanerS10Device) GetStatus() (*RobotVacuumCleanerS10DeviceStatusResponse, error) {
	response := &RobotVacuumCleanerS10DeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type HumidifierDeviceStatusBody struct {
	CommonDevice
	Power                  string `json:"power"`
	Humidity               int    `json:"humidity"`
	Temperature            int    `json:"temperature"`
	NebulizationEfficiency int    `json:"nebulizationEfficiency"`
	Auto                   bool   `json:"auto"`
	ChildLock              bool   `json:"childLock"`
	Sound                  bool   `json:"sound"`
	LackWater              bool   `json:"lackWater"`
}

type HumidifierDeviceStatusResponse struct {
	CommonResponse
	Body *HumidifierDeviceStatusBody `json:"body"`
}

func (device *HumidifierDevice) GetStatus() (*HumidifierDeviceStatusResponse, error) {
	response := &HumidifierDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type EvaporativeHumidifierDeviceFilterElement struct {
	EffectiveUsageHours int `json:"effectiveUsageHours"`
	UsedHours           int `json:"usedHours"`
}

type EvaporativeHumidifierDeviceStatusBody struct {
	CommonDevice
	Power         string                                   `json:"power"`
	Humidity      int                                      `json:"humidity"`
	Mode          int                                      `json:"mode"`
	Drying        bool                                     `json:"drying"`
	ChildLock     bool                                     `json:"childLock"`
	FilterElement EvaporativeHumidifierDeviceFilterElement `json:"filterElement"`
	Version       int                                      `json:"version"`
}

type EvaporativeHumidifierDeviceStatusResponse struct {
	CommonResponse
	Body *EvaporativeHumidifierDeviceStatusBody `json:"body"`
}

func (device *EvaporativeHumidifierDevice) GetStatus() (*EvaporativeHumidifierDeviceStatusResponse, error) {
	response := &EvaporativeHumidifierDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type AirPurifierDeviceStatusBody struct {
	CommonDevice
	Power     string `json:"power"`
	Version   string `json:"version"`
	Mode      int    `json:"mode"`
	ChildLock bool   `json:"childLock"`
}

type AirPurifierDeviceStatusResponse struct {
	CommonResponse
	Body *AirPurifierDeviceStatusBody `json:"body"`
}

func (device *AirPurifierDevice) GetStatus() (*AirPurifierDeviceStatusResponse, error) {
	response := &AirPurifierDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type BlindTiltDeviceStatusBody struct {
	CommonDevice
	Version       int    `json:"version"`
	Calibrate     bool   `json:"calibrate"`
	Group         bool   `json:"group"`
	Moving        bool   `json:"moving"`
	Direction     string `json:"direction"`
	SlidePosition int    `json:"slidePosition"`
}

type BlindTiltDeviceStatusResponse struct {
	CommonResponse
	Body *BlindTiltDeviceStatusBody `json:"body"`
}

func (device *BlindTiltDevice) GetStatus() (*BlindTiltDeviceStatusResponse, error) {
	response := &BlindTiltDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type BatteryCirculatorFanDeviceStatusBody struct {
	CommonDevice
	Mode                string `json:"mode"`
	Version             string `json:"version"`
	Battery             int    `json:"battery"`
	Power               string `json:"power"`
	NightStatus         string `json:"nightStatus"`
	Oscillation         string `json:"oscillation"`
	VerticalOscillation string `json:"verticalOscillation"`
	ChargingStatus      string `json:"chargingStatus"`
	FanSpeed            int    `json:"fanSpeed"`
}

type BatteryCirculatorFanDeviceStatusResponse struct {
	CommonResponse
	Body *BatteryCirculatorFanDeviceStatusBody `json:"body"`
}

func (device *BatteryCirculatorFanDevice) GetStatus() (*BatteryCirculatorFanDeviceStatusResponse, error) {
	response := &BatteryCirculatorFanDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type CirculatorFanDeviceStatusBody struct {
	CommonDevice
	Mode                string `json:"mode"`
	Version             string `json:"version"`
	Power               string `json:"power"`
	NightStatus         string `json:"nightStatus"`
	Oscillation         string `json:"oscillation"`
	VerticalOscillation string `json:"verticalOscillation"`
	FanSpeed            int    `json:"fanSpeed"`
}

type CirculatorFanDeviceStatusResponse struct {
	CommonResponse
	Body *CirculatorFanDeviceStatusBody `json:"body"`
}

func (device *CirculatorFanDevice) GetStatus() (*CirculatorFanDeviceStatusResponse, error) {
	response := &CirculatorFanDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}

type RollerShadeDeviceStatusBody struct {
	CommonDevice
	Version       string `json:"version"`
	Calibrate     bool   `json:"calibrate"`
	Battery       int    `json:"battery"`
	Moving        bool   `json:"moving"`
	SlidePosition int    `json:"slidePosition"`
}

type RollerShadeDeviceStatusResponse struct {
	CommonResponse
	Body *RollerShadeDeviceStatusBody `json:"body"`
}

func (device *RollerShadeDevice) GetStatus() (*RollerShadeDeviceStatusResponse, error) {
	response := &RollerShadeDeviceStatusResponse{}
	err := device.Client.GetRequest("/devices/"+device.DeviceID+"/status", GetDeviceStatusResponseParser(response))
	if err != nil {
		return nil, err
	}
	return response, nil
}
