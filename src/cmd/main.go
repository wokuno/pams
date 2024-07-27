package main

import (
	"fmt"
	"os"

	"go.bug.st/serial"
)

func main() {

	// Open the serial port
	port, err := serial.Open("/dev/ttyACM0", &serial.Mode{BaudRate: 9600})
	if err != nil {
		fmt.Println("Failed to open serial port:", err)
		os.Exit(1)
	}
	defer port.Close()

	// get current temperature
	highestTemp := getHighestTemp()
	fmt.Println("Highest temperature is:", highestTemp)

	fanSpeed := "10\n"

	if highestTemp > 40 {
		fanSpeed = "20\n"
	}

	if highestTemp > 50 {
		fanSpeed = "50\n"
	}

	if highestTemp > 70 {
		fanSpeed = "75\n"
	}

	if highestTemp > 80 {
		fanSpeed = "100\n"
	}

	fmt.Println("Setting fan speed to:", fanSpeed)
	// Send the fan speed signal to the serial port
	_, err = port.Write([]byte(fanSpeed))
	if err != nil {
		fmt.Println("Failed to send fan speed signal:", err)
		os.Exit(1)
	}
}
