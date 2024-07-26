# paco Autonomous Management System
### Current Features:
- Fan Control

### Installation:
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

arduino-cli compile -b arduino:avr:uno pams.ino
arduino-cli upload -v -b arduino:avr:uno -p /dev/ttyACM0
```
