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
- **Shared Device Support**: Some devices with identical functionality share the same struct (e.g., Lock Ultra uses LockDevice, S10/S20 robots use RobotVacuumCleanerSDevice)

### Implementation Guidelines

**API Documentation Compliance**:
- **ALWAYS** refer to the official SwitchBot API documentation at https://raw.githubusercontent.com/OpenWonderLabs/SwitchBotAPI/main/README.md when implementing new device support
- Device status fields, command parameters, and data types must match the official API specification exactly
- When adding new devices, verify the device type strings and supported operations from the official documentation
- Status response structures should reflect the exact JSON schema provided in the API docs

**Device Implementation Process**:
1. Check the official API documentation for device specifications
2. Verify device type strings and available commands
3. Implement status structures with correct field names and data types
4. Add device to the parser with the exact `deviceType` string from the API
5. Update README.md and README_ja.md to reflect support status

### Authentication
The library uses SwitchBot API v1.1 authentication requiring:
- `SWITCH_BOT_TOKEN` environment variable
- `SWITCH_BOT_SECRET` environment variable for HMAC signature generation

Examples in the `examples/` directory demonstrate proper usage patterns for different operations.
