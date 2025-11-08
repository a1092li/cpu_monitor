# CPU Monitor - CPU Core Load Monitoring for Linux

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![Platform](https://img.shields.io/badge/Platform-Linux-lightgrey.svg)](https://linux.org)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](CONTRIBUTING.md)

A command-line utility for real-time CPU core load monitoring on Linux systems.

## ✨ Features

- 📊 Real-time monitoring of individual CPU cores
- 🔄 Live updates every second
- 🎯 Flexible core selection via command-line arguments
- 📈 Proper formatting for cores 0-999
- 🐧 Works on any Linux distribution
- 🚀 Lightweight and fast

## 🚀 Quick Start

### Installation via Go
```bash
go install github.com/a1092li/cpu-monitor/src@latest
```
## Build from Source
```
git clone https://github.com/a1092li/cpu-monitor.git
cd cpu-monitor
make build
```
## Download Binary

Download the latest release from Releases

## Usage
```
# Monitor default cores (0,1,2,3)
./cpu_monitor

# Monitor specific cores
./cpu_monitor 0 2 4 6

# Monitor single core
./cpu_monitor 1

# Monitor cores with high numbers
./cpu_monitor 0 16 32 64

# Show help
./cpu_monitor -h
```
## Example Output
```
CPU Core Load Monitor: [0 2 4 6]
Press Ctrl+C to exit
------------------------------------------------------------
Time: 14:30:25 | Core 0:  15.3% | Core 2:  22.1% | Core 4:   8.7% | Core 6:  95.2% 
```
## Development

### Requirements

Go 1.21 or higher
Linux system

### Build
```
make build          # Build for current platform
make build-all      # Build for all platforms
make test           # Run tests
make install        # Install to GOPATH
```
## Contributing
We welcome contributions! Please see our Contributing Guide for details.

## 📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

## 🔧 How It Works
The program reads data from /proc/stat and calculates CPU load using the formula:

Load = 100% × (Working Time - Idle Time) / Total Time
Data is updated every second for smooth real-time display.

## Supported Platforms
- ✅ Rocky Linux 8+
- ✅ CentOS 7+
- ✅ Ubuntu 16.04+
- ✅ Debian 9+
- ✅ Alpine Linux
- ✅ Any other Linux distribution

## 🐛 Reporting Issues
Found a bug? Please open an issue with:

- Linux distribution and version
- Go version
- Steps to reproduce
- Expected vs actual behavior

## 📚 Documentation
- Installation Guide
https://docs/INSTALL.md
- Usage Examples
https://docs/USAGE.md
- Contributing Guide
https://contributing.md/
