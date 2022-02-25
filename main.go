package main

import (
	"RobotRace-Go/stadium"
	"RobotRace-Go/stadium/robot"
	"time"
)

type racerData struct {
	name          string
	delay         time.Duration
	commandString string
}

var host = stadium.NewStadium()

func initRace() {
	arr := [...]racerData{
		{"R1", 250 * time.Millisecond, "FFFFF"},
		{"R2", 120 * time.Millisecond, "FLRFF"},
		{"R3", 120 * time.Millisecond, "FLFFFRF"},
		{"R4", 130 * time.Millisecond, "LFFFFF"},
		{"R5", 150 * time.Millisecond, "RFFFFFF"},
	}

	for _, data := range arr {
		robotTemp := robot.NewRobot(data.name, data.delay)
		robotTemp.SetCommandString(data.commandString)

		host.AddRacer(&robotTemp)
	}
}

func startRace() {
	host.StartRace()
}

func main() {
	initRace()
	startRace()
}
