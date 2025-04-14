# SwitchBot API in Go

[English version is here.](README.md)

[SwitchBot API v1.1](https://github.com/OpenWonderLabs/SwitchBotAPI)を使用するためのGo言語用のライブラリです。

各デバイスで構造体を定義し、そのデバイスで利用可能なAPIのみを実装することで安全でシンプルに使用できるようにしています。

## 対応機能

- Devices
  - ✅️ デバイス一覧の取得
  - 🚧 デバイスのステータス取得
    - 現在は一部デバイスのみ対応しています
  - コマンドの送信
    - 🚧 物理デバイス
      - 現在は一部デバイスのみ対応しています
    - ✅ 赤外線リモコン
- Scenes
  - ❌ シーン一覧の取得
  - ❌ シーンの手動実行
- Webhooks
  - ❌ Webhookの設定
  - ❌ Webhookの設定取得
  - ❌ Webhookの設定更新
  - ❌ Webhookの削除
  - ❌ Webhookからのイベント受信

# インストール方法

```shell
$ go get github.com/yasu89/switch-bot-api-go
```

## 現在のサポート状況 (2025/04/12)

- 検証済み列に✅がある場合は、実際のデバイスを使用してテストおよび検証されたことを示します。

### 物理デバイス

| デバイス                                 | 構造体定義 | ステータス取得 | コマンド送信 | 検証済み |
|:-------------------------------------|:-----:|:-------:|:------:|:----:|
| Bot                                  |   ✅   |    ✅    |   ✅    |  ✅   |
| Curtain                              |   ✅   |    ❌    |   ✅    |      |
| Curtain 3                            |   ✅   |    ❌    |   ✅    |      |
| Hub                                  |  ✅️   |    ❌    |   -    |      |
| Hub Plus                             |   ✅   |    ❌    |   -    |      |
| Hub Mini                             |   ✅   |    ❌    |   -    |  ✅   |
| Hub 2                                |   ✅   |    ✅    |   -    |  ✅   |
| Meter                                |   ✅   |    ✅    |   -    |  ✅   |
| Meter Plus                           |   ✅   |    ✅    |   -    |      |
| Outdoor Meter                        |   ✅   |    ✅    |   -    |      |
| Meter Pro                            |   ✅   |    ✅    |   -    |      |
| Meter Pro CO2                        |   ✅   |    ✅    |   -    |      |
| Lock                                 |   ✅   |    ✅    |   ✅    |      |
| Lock Pro                             |   ✅   |    ✅    |   ✅    |      |
| Keypad                               |   ✅   |    ✅    |   ✅    |      |
| Keypad Touch                         |   ✅   |    ✅    |   ✅    |      |
| Remote                               |   ✅   |    ❌    |   -    |      |
| Motion Sensor                        |   ✅   |    ❌    |   -    |      |
| Contact Sensor                       |   ✅   |    ✅    |   -    |      |
| Water Leak Detector                  |   ✅   |    ✅    |   -    |      |
| Ceiling Light                        |   ✅   |    ✅    |   ✅    |      |
| Ceiling Light Pro                    |   ✅   |    ✅    |   ✅    |      |
| Plug Mini (US)                       |   ✅   |    ✅    |   ✅    |      |
| Plug Mini (JP)                       |   ✅   |    ✅    |   ✅    |      |
| Plug                                 |   ✅   |    ✅    |   ✅    |      |
| Strip Light                          |   ✅   |    ✅    |   ✅    |      |
| Color Bulb                           |   ✅   |    ✅    |   ✅    |      |
| Robot Vacuum Cleaner S1              |   ❌   |    ❌    |   ❌    |      |
| Robot Vacuum Cleaner S1 Plus         |   ❌   |    ❌    |   ❌    |      |
| Mini Robot Vacuum K10+               |   ❌   |    ❌    |   ❌    |      |
| Mini Robot Vacuum K10+ Pro           |   ❌   |    ❌    |   ❌    |      |
| K10+ Pro Combo                       |   ❌   |    ❌    |   ❌    |      |
| Floor Cleaning Robot S10             |   ❌   |    ❌    |   ❌    |      |
| Humidifier                           |   ❌   |    ❌    |   ❌    |      |
| Evaporative Humidifier               |   ❌   |    ❌    |   ❌    |      |
| Evaporative Humidifier (Auto-refill) |   ❌   |    ❌    |   ❌    |      |
| Air Purifier VOC                     |   ❌   |    ❌    |   ❌    |      |
| Air Purifier Table VOC               |   ❌   |    ❌    |   ❌    |      |
| Air Purifier PM2.5                   |   ❌   |    ❌    |   ❌    |      |
| Air Purifier Table PM2.5             |   ❌   |    ❌    |   ❌    |      |
| Indoor Cam                           |   ❌   |    -    |   -    |      |
| Pan/Tilt Cam                         |   ❌   |    -    |   -    |      |
| Pan/Tilt Cam 2K                      |   ❌   |    -    |   -    |      |
| Blind Tilt                           |   ❌   |    ❌    |   ❌    |      |
| Battery Circulator Fan               |   ❌   |    ❌    |   ❌    |      |
| Circulator Fan                       |   ❌   |    ❌    |   ❌    |      |
| Roller Shade                         |   ❌   |    ❌    |   ❌    |      |
| Relay Switch 1PM                     |   ❌   |    ❌    |   ❌    |      |
| Relay Switch 1                       |   ❌   |    ❌    |   ❌    |      |

### 赤外線リモコン

| 赤外線リモコン              | 構造体定義 | コマンド送信 | 検証済み |
|:---------------------|:-----:|:------:|:----:|
| Air Conditioner      |   ✅   |   ✅    |  ✅   |
| TV                   |   ✅   |   ✅    |      |
| Light                |   ✅   |   ✅    |  ✅   |
| Streamer             |   ✅   |   ✅    |      |
| Set Top Box          |   ✅   |   ✅    |      |
| DVD Player           |   ✅   |   ✅    |      |
| Fan                  |   ✅   |   ✅    |      |
| Projector            |   ✅   |   ✅    |      |
| Camera               |   ✅   |   ✅    |      |
| Air Purifier         |   ✅   |   ✅    |      |
| Speaker              |   ✅   |   ✅    |      |
| Water Heater         |   ✅   |   ✅    |      |
| Robot Vacuum Cleaner |   ✅   |   ✅    |      |
| Others               |   ✅   |   ✅    |      |
