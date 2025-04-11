# SwitchBot API in Go

[日本語はこちら](README_ja.md)

A Golang library for the [SwitchBot API v1.1](https://github.com/OpenWonderLabs/SwitchBotAPI).

Each device has a dedicated struct that exposes only the APIs available for that specific device.<br>
This design promotes both safety and simplicity for users.

## Support Feature

- Devices
  - ✅️ Get device list
  - 🚧 Get device status
    - Support is currently limited to certain devices
  - Send device control command
    - 🚧 Physical devices
      - Support is currently limited to certain devices
    - ✅ Virtual infrared remote devices
- Scenes
  - ❌ Get scene list
  - ❌ Execute manual scenes
- Webhooks
  - ❌ Configure webhook
  - ❌ Get webhook configuration
  - ❌ Update webhook configuration
  - ❌ Delete webhook
  - ❌ Receive events from webhook

## Installing

```shell
$ go get github.com/yasu89/switch-bot-api-go
```

## Current Support Status (2025/04/12)

- A ✅ in the “Verification” column indicates that the feature has been tested and verified using an actual device.

### Physical Devices

| Device                               | Struct Definition | Get Status | Send Command | Verification |
|:-------------------------------------|:-----------------:|:----------:|:------------:|:------------:|
| Bot                                  |         ✅         |     ✅      |      ✅       |      ✅       |
| Curtain                              |         ✅         |     ❌      |      ✅       |              |
| Curtain 3                            |         ✅         |     ❌      |      ✅       |              |
| Hub                                  |        ✅️         |     ❌      |      -       |              |
| Hub Plus                             |         ✅         |     ❌      |      -       |              |
| Hub Mini                             |         ✅         |     ❌      |      -       |      ✅       |
| Hub 2                                |         ✅         |     ✅      |      -       |      ✅       |
| Meter                                |         ✅         |     ✅      |      -       |      ✅       |
| Meter Plus                           |         ✅         |     ✅      |      -       |              |
| Outdoor Meter                        |         ✅         |     ✅      |      -       |              |
| Meter Pro                            |         ✅         |     ✅      |      -       |              |
| Meter Pro CO2                        |         ✅         |     ✅      |      -       |              |
| Lock                                 |         ✅         |     ✅      |      ✅       |              |
| Lock Pro                             |         ✅         |     ✅      |      ✅       |              |
| Keypad                               |         ❌         |     ❌      |      ❌       |              |
| Keypad Touch                         |         ❌         |     ❌      |      ❌       |              |
| Remote                               |         ✅         |     ❌      |      -       |              |
| Motion Sensor                        |         ✅         |     ❌      |      -       |              |
| Contact Sensor                       |         ❌         |     ❌      |      -       |              |
| Water Leak Detector                  |         ❌         |     ❌      |      -       |              |
| Ceiling Light                        |         ❌         |     ❌      |      ❌       |              |
| Ceiling Light Pro                    |         ❌         |     ❌      |      ❌       |              |
| Plug Mini (US)                       |         ❌         |     ❌      |      ❌       |              |
| Plug Mini (JP)                       |         ❌         |     ❌      |      ❌       |              |
| Plug                                 |         ❌         |     ❌      |      ❌       |              |
| Strip Light                          |         ❌         |     ❌      |      ❌       |              |
| Color Bulb                           |         ❌         |     ❌      |      ❌       |              |
| Robot Vacuum Cleaner S1              |         ❌         |     ❌      |      ❌       |              |
| Robot Vacuum Cleaner S1 Plus         |         ❌         |     ❌      |      ❌       |              |
| Mini Robot Vacuum K10+               |         ❌         |     ❌      |      ❌       |              |
| Mini Robot Vacuum K10+ Pro           |         ❌         |     ❌      |      ❌       |              |
| K10+ Pro Combo                       |         ❌         |     ❌      |      ❌       |              |
| Floor Cleaning Robot S10             |         ❌         |     ❌      |      ❌       |              |
| Humidifier                           |         ❌         |     ❌      |      ❌       |              |
| Evaporative Humidifier               |         ❌         |     ❌      |      ❌       |              |
| Evaporative Humidifier (Auto-refill) |         ❌         |     ❌      |      ❌       |              |
| Air Purifier VOC                     |         ❌         |     ❌      |      ❌       |              |
| Air Purifier Table VOC               |         ❌         |     ❌      |      ❌       |              |
| Air Purifier PM2.5                   |         ❌         |     ❌      |      ❌       |              |
| Air Purifier Table PM2.5             |         ❌         |     ❌      |      ❌       |              |
| Indoor Cam                           |         ❌         |     -      |      -       |              |
| Pan/Tilt Cam                         |         ❌         |     -      |      -       |              |
| Pan/Tilt Cam 2K                      |         ❌         |     -      |      -       |              |
| Blind Tilt                           |         ❌         |     ❌      |      ❌       |              |
| Battery Circulator Fan               |         ❌         |     ❌      |      ❌       |              |
| Circulator Fan                       |         ❌         |     ❌      |      ❌       |              |
| Roller Shade                         |         ❌         |     ❌      |      ❌       |              |
| Relay Switch 1PM                     |         ❌         |     ❌      |      ❌       |              |
| Relay Switch 1                       |         ❌         |     ❌      |      ❌       |              |

### Virtual Infrared Remote Devices

| Virtual Infrared Remote Device | Struct Definition | Send Command | Verification |
|:-------------------------------|:-----------------:|:------------:|:------------:|
| Air Conditioner                |         ✅         |      ✅       |      ✅       |
| TV                             |         ✅         |      ✅       |              |
| Light                          |         ✅         |      ✅       |      ✅       |
| Streamer                       |         ✅         |      ✅       |              |
| Set Top Box                    |         ✅         |      ✅       |              |
| DVD Player                     |         ✅         |      ✅       |              |
| Fan                            |         ✅         |      ✅       |              |
| Projector                      |         ✅         |      ✅       |              |
| Camera                         |         ✅         |      ✅       |              |
| Air Purifier                   |         ✅         |      ✅       |              |
| Speaker                        |         ✅         |      ✅       |              |
| Water Heater                   |         ✅         |      ✅       |              |
| Robot Vacuum Cleaner           |         ✅         |      ✅       |              |
| Others                         |         ✅         |      ✅       |              |
