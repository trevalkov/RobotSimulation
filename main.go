package main

import (
	"fmt"
	"time"
)

const worldSize = worldLateral * worldLateral

func tick(robot Robot, world *[worldSize]string) {
	var mark int = placeMark()
	robot.turningWheels = robot.resetTurningWheels()

	for 1 > 0 {
		time.Sleep(1 * time.Second)
		fmt.Println(getCleanScreen())

		robot.mainPos, robot.moveSpeed, robot.moving, robot.turningWheels = robot.move(mark)

		world = drawWorld(world, mark, robot)
		fmt.Println(getWorld(world))

		if robot.mainPos == mark {
			fmt.Println("[*] Resetting...")
			time.Sleep(3 * time.Second)

			tick(robot, world)
		}
	}
}

func main() {
	world := new([worldSize]string)
	world = initWorld(world)

	var robot = initRobot()
	var connectRobot string = robot.connect()
	fmt.Println(connectRobot)

	tick(robot, world)
}
