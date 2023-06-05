package main

import "fmt"

type CarDriver struct {
	Up    chan int
	Left  chan int
	Right chan int
	Down  chan int
}

var up, left, right, down int = 0, 0, 0, 0
var l298nDriver *L298NDriver

func init() {
	l298nDriver = &L298NDriver{}

}
func (driver *CarDriver) Close() {
	l298nDriver.Close()
}
func (driver *CarDriver) Run() {
	for {
		select {
		case v := <-driver.Up:
			up = v
			driver.up()
		case v := <-driver.Left:
			left = v
			driver.left()
		case v := <-driver.Right:
			right = v
			driver.right()
		case v := <-driver.Down:
			down = v
			driver.down()
		}
	}
}

func ToLeft() {
	l298nDriver.IN1(false)
	l298nDriver.IN2(false)
	l298nDriver.IN3(true)
	l298nDriver.IN4(false)

	///右正传 in3=高，in4=低
	//左不转 in1=低，in2=低
}
func ToRight() {

	l298nDriver.IN3(false)
	l298nDriver.IN4(false)
	l298nDriver.IN1(true)
	l298nDriver.IN2(false)
	///右不转  in3=低，in4=低
	//左正转  in1=高，in2=低
}

func Stop() {
	l298nDriver.IN3(false)
	l298nDriver.IN4(false)
	l298nDriver.IN1(false)
	l298nDriver.IN2(false)
	///右不转  in3=低，in4=低
	//左不转 in1=低，in2=低
}
func ToGO() {
	l298nDriver.IN3(true)
	l298nDriver.IN4(false)
	l298nDriver.IN1(true)
	l298nDriver.IN2(false)
	///左正传 in1=高，in2=低
	//右正传  in3=高，in4=低
}

func ToLeftBack() {
	l298nDriver.IN1(false)
	l298nDriver.IN2(false)
	l298nDriver.IN3(false)
	l298nDriver.IN4(true)
	//左不转 in1=低，in2=低
	///右倒转  in3=低，in4=高
}

func ToRightBack() {

	l298nDriver.IN1(false)
	l298nDriver.IN2(true)
	l298nDriver.IN3(false)
	l298nDriver.IN4(false)

	///右不转  in3=低，in4=低
	///左倒转  in1=低，in2=高
}

func ToBack() {

	fmt.Print("back:")
	l298nDriver.IN1(false)
	l298nDriver.IN2(true)
	l298nDriver.IN3(false)
	l298nDriver.IN4(true)
	///右倒转  in3=低，in4=高
	///左倒转  in1=低，in2=高
}

func (driver *CarDriver) up() {
	if up == 1 {
		if left == 1 {
			ToLeft()
			///右正传 in3=高，in4=低
			//左不转 in1=低，in2=低

		}
		if right == 1 {
			ToRight()
			///右不转  in3=低，in4=低
			//左正转  in1=高，in2=低
		}
		if left == 0 && right == 0 {
			ToGO()
			///左正传 in1=高，in2=低
			//右正传  in3=高，in4=低
		}
	}

	if up == 0 {
		Stop()
		///右不转  in3=低，in4=低
		//左不转 in1=低，in2=低
	}

}
func (driver *CarDriver) left() {
	if up == 1 {
		if left == 1 {
			ToLeft()
		}
		if left == 0 {
			ToGO()
		}

	}
	if down == 1 {
		if left == 1 {
			ToLeftBack()
		}
		if left == 0 {
			ToBack()
		}
	}
}
func (driver *CarDriver) right() {
	if up == 1 {
		if right == 1 {
			ToRight()
		}
		if right == 0 {
			ToGO()
		}

	}
	if down == 1 {
		if right == 1 {
			ToRightBack()
		}
		if right == 0 {
			ToBack()
		}
	}
}

func (driver *CarDriver) down() {
	fmt.Println("down:", down, left)
	if down == 1 {
		if left == 1 {
			ToLeft()
		}
		if left == 0 {
			ToBack()
		}
	}
	if down == 0 {
		Stop()
	}
}
