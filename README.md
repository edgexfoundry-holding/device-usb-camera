# USB Camera Device Service
[![Build Status](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-usb-camera/job/main/badge/icon)](https://jenkins.edgexfoundry.org/view/EdgeX%20Foundry%20Project/job/edgexfoundry/job/device-usb-camera/job/main/) [![Code Coverage](https://codecov.io/gh/edgexfoundry/device-usb-camera/branch/main/graph/badge.svg?token=K4V4LAJYYW)](https://codecov.io/gh/edgexfoundry/device-usb-camera) [![Go Report Card](https://goreportcard.com/badge/github.com/edgexfoundry/device-usb-camera)](https://goreportcard.com/report/github.com/edgexfoundry/device-usb-camera) [![GitHub Latest Dev Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-usb-camera?include_prereleases&sort=semver&label=latest-dev)](https://github.com/edgexfoundry/device-usb-camera/tags) ![GitHub Latest Stable Tag)](https://img.shields.io/github/v/tag/edgexfoundry/device-usb-camera?sort=semver&label=latest-stable) [![GitHub License](https://img.shields.io/github/license/edgexfoundry/device-usb-camera)](https://choosealicense.com/licenses/apache-2.0/) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edgexfoundry/device-usb-camera) [![GitHub Pull Requests](https://img.shields.io/github/issues-pr-raw/edgexfoundry/device-usb-camera)](https://github.com/edgexfoundry/device-usb-camera/pulls) [![GitHub Contributors](https://img.shields.io/github/contributors/edgexfoundry/device-usb-camera)](https://github.com/edgexfoundry/device-usb-camera/contributors) [![GitHub Committers](https://img.shields.io/badge/team-committers-green)](https://github.com/orgs/edgexfoundry/teams/device-usb-camera-committers/members) [![GitHub Commit Activity](https://img.shields.io/github/commit-activity/m/edgexfoundry/device-usb-camera)](https://github.com/edgexfoundry/device-usb-camera/commits)


## Overview
The USB Device Service is a microservice created to address the lack of standardization and automation of camera discovery and onboarding. EdgeX Foundry is a flexible microservice-based architecture created to promote the interoperability of multiple device interface combinations at the edge. In an EdgeX deployment, the USB Device Service controls and communicates with USB cameras, while EdgeX Foundry presents a standard interface to application developers. With normalized connectivity protocols and a vendor-neutral architecture, EdgeX paired with USB Camera Device Service, simplifies deployment of edge camera devices.

Specifically, the device service uses V4L2 API to get camera metadata, FFmpeg framework to capture video frames and stream them to an [RTSP server](https://github.com/aler9/rtsp-simple-server), which is embedded in the dockerized device service. This allows the video stream to be integrated into the [larger architecture](#how-it-works).

Use the USB Device Service to streamline and scale your edge camera device deployment. 

## Build with NATS Messaging
Currently, the NATS Messaging capability (NATS MessageBus) is opt-in at build time.
This means that the published Docker image and Snaps do not include the NATS messaging capability.

The following make commands will build the local binary or local Docker image with NATS messaging
capability included.
```makefile
make build-nats
make docker-nats
```

The locally built Docker image can then be used in place of the published Docker image in your compose file.
See [Compose Builder](https://github.com/edgexfoundry/edgex-compose/tree/main/compose-builder#gen) `nat-bus` option to generate compose file for NATS and local dev images.

## How It Works

The figure below illustrates the software flow through the architecture components.

![high-level-arch](./docs/images/USBDeviceServiceArch.png)
<p align="left">
      <i>Figure 1: Software Flow</i>
</p>

1. **EdgeX Device Discovery:** Camera device microservices probe network and platform for video devices at a configurable interval. Devices that do not currently exist and that satisfy Provision Watcher filter criteria are added to `Core Metadata`.
2. **Application Device Discovery:** The microservices then query `Core Metadata` for devices and associated configuration.
3. **Application Device Configuration:** The configuration and triggering of device actions are done through a REST API representing the resources of the video device.
4. **Pipeline Control:** The application initiates the `Video Analytics Pipeline` through HTTP Post Request.
5. **Publish Inference Events/Data:** Analytics inferences are formatted and passed to the destination message bus specified in the request.
6. **Export Data:** Publish prepared (transformed, enriched, filtered, etc.) and groomed (formatted, compressed, encrypted, etc.) data to external systems (be it analytics package, enterprise or on-premises application, cloud systems like Azure IoT, AWS IoT, or Google IoT Core, etc.

# Getting Started

To set up your system, follow [this guide.](./docs/setup.md)  
For a full walkthrough on how to use this service and RTSP streaming, follow [this guide.](./docs/general-usage.md)  


### CameraStatus Command
Use the following query to determine the status of the camera.
URL parameter:
- **DeviceName**: The name of the camera
- **InputIndex**: indicates the current index of the video input (if a camera only has one source for video, the index needs to be set to '0')
```
curl -X GET http://localhost:59882/api/v2/device/name/<DeviceName>/CameraStatus?InputIndex=0 | jq -r '"CameraStatus: " + (.event.readings[].value|tostring)'
```
   Example Output: 
   ```
    CameraStatus: 0
   ```
   **Response meanings**:
| Response   | Description |
| ---------- | ----------- |
| 0 | Ready |
| 1 | No Power |
| 2 | No Signal |
| 3 | No Color |  
## License
[Apache-2.0](LICENSE)
