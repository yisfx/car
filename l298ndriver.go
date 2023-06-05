package main

import (
	"github.com/stianeikeland/go-rpio/v4"
)

type L298NDriver struct {
}

var in1 rpio.Pin
var in2 rpio.Pin

var in3 rpio.Pin
var in4 rpio.Pin

func init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	in3 = rpio.Pin(17)
	in3.Output()
	in4 = rpio.Pin(18)
	in4.Output()
	in1 = rpio.Pin(27)
	in1.Output()
	in2 = rpio.Pin(22)
	in2.Output()

	in1.Low()
	in2.Low()
	in3.Low()
	in4.Low()
}

func setPin(pin rpio.Pin, f bool) {
	if f {
		pin.High()
	} else {
		pin.Low()
	}
}

func (driver *L298NDriver) Close() {
	rpio.Close()
}

func (driver *L298NDriver) IN1(f bool) {
	setPin(in1, f)
}

func (driver *L298NDriver) IN2(f bool) {
	setPin(in2, f)
}

func (driver *L298NDriver) IN3(f bool) {
	setPin(in3, f)
}
func (driver *L298NDriver) IN4(f bool) {
	setPin(in4, f)
}
