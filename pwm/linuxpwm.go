package linuxpwm

import (
	"os"
	"strconv"

	"github.com/daniel38192/go-gpio/pwm/polarity"
	"github.com/daniel38192/go-gpio/pwm/sysfspath"
)

type PWM struct {
	controller string
	channel    string
	period     int
	polarity   polarity.PWMpolarity
}

func NewPWM(controller string, channel string, period int, polarity polarity.PWMpolarity) PWM {

	pwm := PWM{controller: controller, channel: channel, period: period, polarity: polarity}
	pwm.Init()
	return pwm

}

func (pwm PWM) Init() {
	pwm.exportPWM()
	pwm.setPeriod()
}

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

func (pwm PWM) SetDutyCycle(dutyCycle int) {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/duty_cycle", []byte(strconv.Itoa(dutyCycle)), 0666)
	if err != nil {
		panic(err)
	}
}

func (pwm PWM) setPeriod() {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/period", []byte(strconv.Itoa(pwm.period)), 0666)
	if err != nil {
		panic(err)
	}
}

func (pwm PWM) SetPolarity() {
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+pwm.channel+"/polarity", []byte(pwm.polarity), 0666)
	if err != nil {
		panic(err)
	}
}

func (pwm PWM) UnexportPWM() {
	sysfsChannel := pwm.channel
	channelNumber := sysfsChannel[len(sysfsChannel)-1]
	err := os.WriteFile(sysfspath.SysfsPWMPath+pwm.controller+"/"+"unexport", []byte(string(channelNumber)), 0666)
	if err != nil {
		panic(err)
	}
}
