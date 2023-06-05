package main

import "fmt"

type Car struct {
}

var driver *CarDriver

func init() {
	///init car driver
	driver = &CarDriver{}
	fmt.Println("Start run")
	driver.Up = make(chan int)
	driver.Left = make(chan int)
	driver.Right = make(chan int)
	driver.Down = make(chan int)
	go driver.Run()
	fmt.Println("Start run end")
}

func (car *Car) Action(t string, action string) {
	switch t {
	case "press":
		car.press(action)
	case "blur":
		car.blur(action)
	}
}

func (car *Car) press(action string) {
	switch action {
	case "up":
		driver.Up <- 1
	case "left":
		driver.Left <- 1
	case "right":
		driver.Right <- 1
	case "down":
		fmt.Println("in down")
		driver.Down <- 1
		fmt.Println("in down done")
	}
}

func (car *Car) blur(action string) {
	switch action {
	case "up":
		driver.Up <- 0
	case "left":
		driver.Left <- 0
	case "right":
		driver.Right <- 0
	case "down":
		driver.Down <- 0
	}
}
