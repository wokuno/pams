
# paco Autonomous Management System

## Current Features

- Fan Control
- Fan Speed Daemon

## Installation

### Requirements

- Arduino-cli
- Go 1.22
- Systemd
- Currently supporting openSUSE

### Arduino Installation

``` bash
# Make sure your user can use Serial Ports
sudo usermod -aG dialout $USER

# Install arduino-cli
curl -fsSL https://raw.githubusercontent.com/arduino/arduino-cli/master/install.sh | sh

# Initalize Arduino enviroment
arduino-cli config init
arduino-cli core update-index

# Find the board, should be an Arduino Uno on /dev/ttyACM0
arduino-cli board list

# Currently running Arduino Uno which is an AVR chip
arduino-cli core install arduino:avr

arduino-cli compile -b arduino:avr:uno src/arduino/pams.ino
arduino-cli upload -v -b arduino:avr:uno -p /dev/ttyACM0
```

### pAMS Installation

```bash
./src/scripts/build_pams.sh
sudo ./src/scripts/install_pams.sh
```

### Uninstall

```bash
sudo ./src/scripts/uninstall_pams.sh
```