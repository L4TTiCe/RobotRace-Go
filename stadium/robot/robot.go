package robot

import (
	"RobotRace-Go/stadium"
	"RobotRace-Go/stadium/robot/utils"
	"fmt"
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

func (robot *Racer) String() string {
	return fmt.Sprintf("%s @ %s ^%s", robot.racerName, robot.position, robot.direction)
}

func (robot *Racer) RegisterHost(stadium *stadium.Stadium) {
	robot.raceHost = stadium
}

func (robot *Racer) Announce() {
	log.Printf("[%s]\tFinished with Rank %d\n", robot, robot.rank)
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

func (robot *Racer) GetCommandString() string {
	return robot.commandString
}

func (robot *Racer) GetDelay() time.Duration {
	return robot.delay
}

func (robot *Racer) Start() {
	timestamp := time.Now()

	for _, command := range robot.commandString {
		switch command {
		case 'F':
			robot.move()
			log.Printf("[%s]\tMoved Forward\t(%dms Since last action)\n",
				robot,
				time.Since(timestamp).Milliseconds())
		case 'L':
			tempDirection := robot.direction.GetName()
			robot.turnLeft()
			log.Printf("[%s]\tTurned Left (%s -> %s)\t(%dms Since last action)\n",
				robot,
				tempDirection,
				robot.direction,
				time.Since(timestamp).Milliseconds())
		case 'R':
			tempDirection := robot.direction.GetName()
			robot.turnRight()
			log.Printf("[%s]\tTurned Right (%s -> %s)\t(%dms Since last action)\n",
				robot,
				tempDirection,
				robot.direction,
				time.Since(timestamp).Milliseconds())
		default:
			log.Printf("[%s]\tUnknown Command: %c\t(%dms Since last action)\n",
				robot,
				command,
				time.Since(timestamp).Milliseconds())
		}
		timestamp = time.Now()
		time.Sleep(robot.delay)
	}
	log.Printf("[%s]\tCompleted Run!\n", robot)
	robot.rank = robot.raceHost.GetRank(robot)
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
