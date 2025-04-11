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

	for _, infraredRemoteDevice := range response.Body.InfraredRemoteList {
		switch infraredRemoteDevice.(type) {
		case *switchbot.InfraredRemoteLightDevice:
			lightDevice := infraredRemoteDevice.(*switchbot.InfraredRemoteLightDevice)
			log.Printf("Light. DeviceID:%s, DeviceName:%s, RemoteType:%s", lightDevice.DeviceID, lightDevice.DeviceName, lightDevice.RemoteType)

			commandResponse, err := lightDevice.TurnOn()
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			log.Printf("StatusCode: %d, Message: %s", commandResponse.StatusCode, commandResponse.Message)
		}
	}
}
