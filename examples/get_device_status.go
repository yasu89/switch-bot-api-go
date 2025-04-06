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
	deviceId, ok := os.LookupEnv("SWITCH_BOT_DEVICE_ID")
	if !ok {
		log.Fatal("SWITCH_BOT_DEVICE_ID environment variable is required")
	}

	client := switchbot.NewClient(secret, token)

	response, err := client.GetDeviceStatus(deviceId)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	// NOTE: MeterDevice Example
	meterDeviceStatus := response.Body.(*switchbot.MeterDeviceStatus)
	log.Printf("Battery: %d%%, Temperature:%0.1f°C, Humidity:%d%%", meterDeviceStatus.Battery, meterDeviceStatus.Temperature, meterDeviceStatus.Humidity)
}
