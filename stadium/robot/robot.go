package robot

import (
	"RobotRace-Go/stadium"
	"RobotRace-Go/stadium/robot/utils"
	"log"
	"time"
)

type Robot interface {
	SetCommandString(command string)
	Start()
	RegisterHost(stadium *stadium.Stadium)
}

type Racer struct {
	racerName     string
	commandString string
	delay         time.Duration
	position      utils.Point2d
	direction     *utils.Direction

	rank     int
	raceHost *stadium.Stadium
}

func (robot *Racer) RegisterHost(stadium *stadium.Stadium) {
	robot.raceHost = stadium
}

func (robot *Racer) move() {
	robot.position = utils.Point2d{
		X: robot.position.X + robot.direction.GetHorizontalModifier(),
		Y: robot.position.Y + robot.direction.GetVerticalModifier(),
	}
}

func (robot *Racer) turnLeft() {
	direction, err := robot.direction.GetLeft()
	if err != nil {
		panic(err)
	}

	robot.direction = direction
}

func (robot *Racer) turnRight() {
	direction, err := robot.direction.GetRight()
	if err != nil {
		panic(err)
	}

	robot.direction = direction
}

func (robot *Racer) SetCommandString(command string) {
	robot.commandString = command
}

func (robot *Racer) Start() {
	log.Printf("[%s] Command Strings is: %s\n", robot.racerName, robot.commandString)
	for _, command := range robot.commandString {
		switch command {
		case 'F':
			robot.move()
			log.Printf("[%s] Moved Forward to: %s\n", robot.racerName, robot.position)
		case 'L':
			robot.turnLeft()
			log.Printf("[%s] Turned Left towards: %s\n", robot.racerName, robot.direction.GetName())
		case 'R':
			robot.turnRight()
			log.Printf("[%s] Turned Right towards: %s\n", robot.racerName, robot.direction.GetName())
		default:
			log.Printf("[%s] Unknown Command: %c\n", robot.racerName, command)
		}
		time.Sleep(robot.delay)
	}
	log.Printf("[%s] Completed Run!\n", robot.racerName)
	robot.rank = robot.raceHost.GetRank()
}

func NewRobot(name string, delay time.Duration) Racer {
	robot := Racer{
		racerName: name,
		delay:     delay,
		position: utils.Point2d{
			X: 0,
			Y: 0,
		},
		direction:     utils.GetNorth(),
		commandString: "",
		rank:          0,
		raceHost:      nil,
	}
	return robot
}
