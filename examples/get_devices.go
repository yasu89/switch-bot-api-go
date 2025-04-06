package main

import (
	"log"
	"os"

	"github.com/yasu89/switch-bot-api-go"
)

func main() {
	token, ok := os.LookupEnv("SWITCH_BOT_TOKEN")
	if !ok {
		log.Fatal("SWITCH_BOT_TOKEN environment variable is required")
	}
	secret, ok := os.LookupEnv("SWITCH_BOT_SECRET")
	if !ok {
		log.Fatal("SWITCH_BOT_SECRET environment variable is required")
	}

	client := switchbot.NewClient(secret, token)

	response, err := client.GetDevices()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, device := range response.Body.DeviceList {
		switch device.(type) {
		case *switchbot.HubDevice:
			device := device.(*switchbot.HubDevice)
			log.Printf("Hub Device. DeviceID:%s, DeviceName:%s", device.DeviceID, device.DeviceName)
		case *switchbot.MeterDevice:
			device := device.(*switchbot.MeterDevice)
			log.Printf("Meter Device. DeviceID:%s, DeviceName:%s", device.DeviceID, device.DeviceName)
		}
	}
}
