// Package linuxpwm provides a Go interface for controlling Pulse Width Modulation (PWM) signals on Linux systems using the sysfs interface.
//
// This package offers functionalities to manage PWM controllers and channels, allowing users to configure parameters such as period, duty cycle, polarity, and enable/disable PWM signals.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// Usage:
// 	pwm := linuxpwm.NewPWM("pwmchip0", "pwm0", 1000000, polarity.Normal)
// 	pwm.SetPolarity() // Set polarity
// 	pwm.SetDutyCycle(500000) // Set duty cycle
// 	pwm.Enable(true) // Enable PWM signal
//
// For proper functionality, ensure that the PWM kernel module is enabled and configured on your Linux system.

package linuxpwm

import (
	"os"
	"strconv"

	"github.com/daniel38192/go-gpio/pwm/polarity"
	"github.com/daniel38192/go-gpio/pwm/sysfspath"
)

// PWM represents a Pulse Width Modulation controller and channel on a Linux system.
type PWM struct {
	controller string
	channel    string
	period     int
	polarity   polarity.PWMpolarity
}

// NewPWM creates a new PWM instance with the specified parameters and initializes it
func NewPWM(controller string, channel string, period int, polarity polarity.PWMpolarity) PWM {

	pwm := PWM{controller: controller, channel: channel, period: period, polarity: polarity}
	pwm.Init()
	return pwm

}

// Init initializes the PWM by exporting the PWM channel and setting the period.
func (pwm PWM) Init() {
	pwm.exportPWM()
	pwm.setPeriod()
}

// exportPWM exports the PWM channel if it is not already exported.
func (pwm PWM) exportPWM() {
	sysfsChannel := pwm.channel
	channelNumber := sysfsChannel[len(sysfsChannel)-1]
	if _, err := os.Stat(sysfspath.SysfsPWMPath + pwm.controller + "/" + pwm.channel); err != nil {
		if os.IsNotExist(err) {
			err2 := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+"export", []byte(string(channelNumber)), 0666)
			if err2 != nil {
				panic(err2)
			}
		}
	}
}

// Enable enables or disables the PWM signal.
func (pwm PWM) Enable(enable bool) {

	if enable {
		err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/enable", []byte("1"), 0666)
		if err != nil {
			panic(err)
		}
	} else {
		err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/enable", []byte("0"), 0666)
		if err != nil {
			panic(err)
		}
	}

}

// SetDutyCycle sets the duty cycle of the PWM signal, for example put "500000" to set duty cycle to 50% with a period of "1000000".
func (pwm PWM) SetDutyCycle(dutyCycle int) {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/duty_cycle", []byte(strconv.Itoa(dutyCycle)), 0666)
	if err != nil {
		panic(err)
	}
}

// setPeriod sets the period of the PWM signal, for example put "1000000" to set the period to 1 second.
func (pwm PWM) setPeriod() {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/period", []byte(strconv.Itoa(pwm.period)), 0666)
	if err != nil {
		panic(err)
	}
}

// SetPolarity sets the polarity of the PWM signal.
func (pwm PWM) SetPolarity() {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/polarity", []byte(pwm.polarity), 0666)
	if err != nil {
		panic(err)
	}
}

// UnexportPWM unexports the PWM channel.
func (pwm PWM) UnexportPWM() {
	sysfsChannel := pwm.channel
	channelNumber := sysfsChannel[len(sysfsChannel)-1]
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+"unexport", []byte(string(channelNumber)), 0666)
	if err != nil {
		panic(err)
	}
}
