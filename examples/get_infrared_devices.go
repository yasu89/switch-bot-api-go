package main

import (
	"github.com/yasu89/switch-bot-api-go"
	"log"
	"os"
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

	for _, infraredRemoteDevice := range response.Body.InfraredRemoteList {
		log.Printf("Infrared Remote. DeviceID:%s, DeviceName:%s, RemoteType:%s", infraredRemoteDevice.DeviceID, infraredRemoteDevice.DeviceName, infraredRemoteDevice.RemoteType)
	}
}
