// Package gpio provides functionality to interact with GPIO (General Purpose Input/Output) pins.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The GPIO struct represents a GPIO pin with its properties.
// It includes methods to export and unexport the GPIO, set the mode to active low,
// write and read the GPIO value, and set the GPIO Direction.
//
// Fields:
// Number: The GPIO Number.
// ActiveLow: A boolean indicating whether the GPIO is active low.
// Value: The current value of the GPIO.
// Direction: The Direction of the GPIO, can be either "in" or "out".
//
// Methods:
// ExportGpio: Exports the GPIO to the sysfs. If the GPIO is already exported, it does nothing.
// UnexportGpio: Unexports the GPIO from the sysfs.
// SetModeActiveLow: Sets the GPIO mode to active low. If ActiveLow is true, the GPIO is set to active low mode, otherwise it is set to normal mode.
// WriteGpioValue: Writes a value to the GPIO. If Value is true, it writes "1" to the GPIO value file, otherwise it writes "0".
// ReadGpioValue: Reads the current value of the GPIO. It reads the GPIO value file and returns true if the value is "1", false if the value is "0".
// SetDirectionGpio: Sets the Direction of the GPIO. It writes the Direction field value to the GPIO Direction file.
package linuxgpio

import (
	"fmt"
	"os"
	"strings"
	"time"

	enums "github.com/daniel38192/go-gpio/gpio/enums"
	generalconstants "github.com/daniel38192/go-gpio/gpio/utils/constants/general"
)

type GPIO struct {
	number    int
	activeLow bool
	direction enums.GPIOMode
}

func NewGpio(gpioNumber int, activeLow bool, direction enums.GPIOMode) GPIO {
	gpio := GPIO{number: gpioNumber, activeLow: activeLow, direction: direction}
	gpio.Init()
	return gpio
}

// Init initializes the GPIO pin, exporting it, then setting the direction, and finaly set the mode.

func (gpio GPIO) Init() {
	gpio.exportGpio()
	time.Sleep(time.Millisecond * 189)
	gpio.setDirectionGpio()
	time.Sleep(time.Millisecond * 189)
	gpio.setModeActiveLow()
}

// ExportGpio exports the GPIO to the sysfs. If the GPIO is already exported, it does nothing.
func (gpio GPIO) exportGpio() {

	var sysGpio = fmt.Sprint(gpio.number)

	if _, err := os.Stat(generalconstants.PathToGpioBase + "gpio" + sysGpio); err != nil {
		if os.IsNotExist(err) {
			err2 := os.WriteFile(generalconstants.PathToGpioBase+"export", []byte(sysGpio), 0666)
			if err2 != nil {
				panic(err2)
			}
		}
	}
}

// UnexportGpio unexports the GPIO from the sysfs.
func (gpio GPIO) UnexportGpio() {
	var sysGpio = fmt.Sprint(gpio.number)
	err := os.WriteFile(generalconstants.PathToGpioBase+"unexport", []byte(sysGpio), 0666)
	if err != nil {
		panic(err)
	}

}

// SetModeActiveLow sets the GPIO mode to active low. If activeLow is true, the GPIO is set to active low mode, otherwise it is set to normal mode.
func (gpio GPIO) setModeActiveLow() {
	var sysGpio = fmt.Sprint(gpio.number)
	var err error
	if gpio.activeLow {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/active_low", []byte("1"), 0666)
	} else {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/active_low", []byte("0"), 0666)
	}

	if err != nil {
		panic(err)
	}
}

// WriteGpioValue writes a value to the GPIO. If Value is true, it writes "1" to the GPIO value file, otherwise it writes "0".
func (gpio GPIO) WriteGpioValue(gpioValue bool) {

	if gpio.direction == enums.IN {
		fmt.Println("cannot write gpio value because gpio" + fmt.Sprint(gpio.number) + "is actually configured as an input")
		os.Exit(1)
	}

	var sysGpio = fmt.Sprint(gpio.number)

	var err error

	if gpioValue {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/value", []byte("1"), 0666)
	} else {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/value", []byte("0"), 0666)
	}

	if err != nil {
		panic(err)
	}
}

// ReadGpioValue reads the current value of the GPIO. It reads the GPIO value file and returns true if the value is "1", false if the value is "0".
func (gpio GPIO) ReadGpioValue() bool {
	var value bool
	var sysGpio = fmt.Sprint(gpio.number)
	e, err := os.ReadFile(generalconstants.PathToGpioBase + "gpio" + sysGpio + "/value")
	gpioStat := strings.TrimSuffix(string(e), "\n")
	if err != nil {
		panic(err)
	}

	if gpioStat == "1" {
		value = true
	} else if gpioStat == "0" {
		value = false
	} else {
		fmt.Println("unsupported gpio value for reading")
	}

	return value

}

// SetDirectionGpio sets the direction of the GPIO. It writes the direction field value to the GPIO direction file.
func (gpio GPIO) setDirectionGpio() {
	var sysGpio = fmt.Sprint(gpio.number)
	err := os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/direction", []byte(gpio.direction), 0666)

	if err != nil {
		panic(err)
	}
}
