# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Testing
- `go test ./...` - Run all tests
- `go test -v ./...` - Run tests with verbose output
- `go test -v ./... -coverprofile=coverage.out` - Run tests with coverage report

### Build and Development
- `go build` - Build the package
- `go mod download` - Download dependencies
- `go mod tidy` - Clean up module dependencies
- `go fmt ./...` - Format all Go files
- `go vet ./...` - Run Go vet for static analysis

### Examples
- `go run examples/get_devices/get_devices.go` - Run device listing example
- `go run examples/get_device_status/get_device_status.go` - Run device status example
- `go run examples/command_device/command_device.go` - Run device control example

## Architecture

This is a Go library for the SwitchBot API v1.1 that provides type-safe access to SwitchBot devices and infrared remote devices.

### Core Components

- **Client (`switch_bot.go`)**: Main API client with authentication, request handling, and debug capabilities. Uses HMAC-SHA256 signature authentication with the SwitchBot API.

- **Device Management (`device.go`)**: Contains comprehensive device type definitions for both physical devices (Bot, Curtain, Hub, Meter, etc.) and virtual infrared remote devices (Air Conditioner, TV, Light, etc.). Each device type has dedicated structs that expose only relevant APIs for that device.

- **Device Status (`device_status.go`)**: Handles retrieving status information from devices with device-specific response parsing.

- **Device Control (`device_control.go`, `device_control_json.go`)**: Manages sending control commands to devices with JSON schema validation for command parameters.

### Key Design Patterns

- **Type Safety**: Each device has a dedicated struct that exposes only the APIs available for that specific device
- **Dynamic Type Resolution**: The `GetDevicesResponseParser` dynamically creates appropriate device structs based on `deviceType` and `remoteType` fields
- **Client Injection**: All device structs contain a `*Client` field for API operations
- **Response Parsing**: Custom response parsers handle the complex JSON unmarshaling for different device types

### Authentication
The library uses SwitchBot API v1.1 authentication requiring:
- `SWITCH_BOT_TOKEN` environment variable
- `SWITCH_BOT_SECRET` environment variable for HMAC signature generation

Examples in the `examples/` directory demonstrate proper usage patterns for different operations.