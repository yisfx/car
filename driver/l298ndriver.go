package driver

import "github.com/stianeikeland/go-rpio/v4"

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
	in3.Mode(rpio.Pwm)
	in3.Freq(64000)

	in4 = rpio.Pin(18)
	in4.Mode(rpio.Pwm)
	in4.Freq(64000)

	in1 = rpio.Pin(27)
	in1.Mode(rpio.Pwm)
	in1.Freq(64000)

	in2 = rpio.Pin(22)
	in2.Mode(rpio.Pwm)
	in2.Freq(64000)

	in3.DutyCycle(0, 32)
	in4.DutyCycle(0, 32)
	in1.DutyCycle(0, 32)
	in2.DutyCycle(0, 32)

}

func setPin(pin rpio.Pin, p uint32) {
	pin.DutyCycle(p, 32)
}

func (driver *L298NDriver) Close() {
	rpio.Close()
}

func (driver *L298NDriver) IN1(p uint32) {
	setPin(in1, p)
}

func (driver *L298NDriver) IN2(p uint32) {
	setPin(in2, p)
}

func (driver *L298NDriver) IN3(p uint32) {
	setPin(in3, p)
}
func (driver *L298NDriver) IN4(p uint32) {
	setPin(in4, p)
}
