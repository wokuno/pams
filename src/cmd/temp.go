package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type CPUTemps struct {
	Adapter    string     `json:"Adapter,omitempty"`
	PackageID0 PackageID0 `json:"Package id 0,omitempty"`
	Core0      Core0      `json:"Core 0,omitempty"`
	Core1      Core1      `json:"Core 1,omitempty"`
	Core2      Core2      `json:"Core 2,omitempty"`
	Core3      Core3      `json:"Core 3,omitempty"`
}
type PackageID0 struct {
	Temp1Input     int `json:"temp1_input,omitempty"`
	Temp1Max       int `json:"temp1_max,omitempty"`
	Temp1Crit      int `json:"temp1_crit,omitempty"`
	Temp1CritAlarm int `json:"temp1_crit_alarm,omitempty"`
}
type Core0 struct {
	Temp2Input     int `json:"temp2_input,omitempty"`
	Temp2Max       int `json:"temp2_max,omitempty"`
	Temp2Crit      int `json:"temp2_crit,omitempty"`
	Temp2CritAlarm int `json:"temp2_crit_alarm,omitempty"`
}
type Core1 struct {
	Temp3Input     int `json:"temp3_input,omitempty"`
	Temp3Max       int `json:"temp3_max,omitempty"`
	Temp3Crit      int `json:"temp3_crit,omitempty"`
	Temp3CritAlarm int `json:"temp3_crit_alarm,omitempty"`
}
type Core2 struct {
	Temp4Input     int `json:"temp4_input,omitempty"`
	Temp4Max       int `json:"temp4_max,omitempty"`
	Temp4Crit      int `json:"temp4_crit,omitempty"`
	Temp4CritAlarm int `json:"temp4_crit_alarm,omitempty"`
}
type Core3 struct {
	Temp5Input     int `json:"temp5_input,omitempty"`
	Temp5Max       int `json:"temp5_max,omitempty"`
	Temp5Crit      int `json:"temp5_crit,omitempty"`
	Temp5CritAlarm int `json:"temp5_crit_alarm,omitempty"`
}

func getHighestTemp() int {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Getting CPU temperatures")
	cmd := exec.Command("bash", "-c", "sensors -j | sed 's/coretemp-isa-0000/coretemp/' | jq '.coretemp'")
	// The `Output` method executes the command and
	// collects the output, returning its value
	rawOut, err := cmd.Output()
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not run command: ", err)
	}

	temps := CPUTemps{}
	// Unmarshal the JSON byte slice to a predefined struct
	err = json.Unmarshal(rawOut, &temps)
	if err != nil {
		// if there was any error, print it here
		fmt.Println("could not unmarshal json: ", err)
	}

	// Get the highest temperature from the struct
	highestTemp := int(0)

	if temps.PackageID0.Temp1Input > highestTemp {
		fmt.Println("PackageID0.Temp1Input:", temps.PackageID0.Temp1Input)
		highestTemp = temps.PackageID0.Temp1Input
	} else if temps.Core0.Temp2Input > highestTemp {
		fmt.Println("Core0.Temp2Input:", temps.Core0.Temp2Input)
		highestTemp = temps.Core0.Temp2Input
	} else if temps.Core1.Temp3Input > highestTemp {
		fmt.Println("Core1.Temp3Input:", temps.Core1.Temp3Input)
		highestTemp = temps.Core1.Temp3Input
	} else if temps.Core2.Temp4Input > highestTemp {
		fmt.Println("Core2.Temp4Input:", temps.Core2.Temp4Input)
		highestTemp = temps.Core2.Temp4Input
	} else if temps.Core3.Temp5Input > highestTemp {
		fmt.Println("Core3.Temp5Input:", temps.Core3.Temp5Input)
		highestTemp = temps.Core3.Temp5Input
	}

	return highestTemp
}
