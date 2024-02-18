// Package gpio provides functionality to interact with GPIO (General Purpose Input/Output) pins.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The GPIO struct represents a GPIO pin with its properties.
// It includes methods to export and unexport the GPIO, set the mode to active low,
// write and read the GPIO value, and set the GPIO direction.
//
// Fields:
// Number: The GPIO number.
// ActiveLow: A boolean indicating whether the GPIO is active low.
// Value: The current value of the GPIO.
// Direction: The direction of the GPIO, can be either "in" or "out".
//
// Methods:
// ExportGpio: Exports the GPIO to the sysfs. If the GPIO is already exported, it does nothing.
// UnexportGpio: Unexports the GPIO from the sysfs.
// SetModeActiveLow: Sets the GPIO mode to active low. If ActiveLow is true, the GPIO is set to active low mode, otherwise it is set to normal mode.
// WriteGpioValue: Writes a value to the GPIO. If Value is true, it writes "1" to the GPIO value file, otherwise it writes "0".
// ReadGpioValue: Reads the current value of the GPIO. It reads the GPIO value file and returns true if the value is "1", false if the value is "0".
// SetDirectionGpio: Sets the direction of the GPIO. It writes the Direction field value to the GPIO direction file.
package linuxgpio

import (
	"fmt"
	"os"
	"strings"

	enums "github.com/daniel38192/go-gpio/enums"
	generalconstants "github.com/daniel38192/go-gpio/utils/constants/general"
	kernelutils "github.com/daniel38192/go-gpio/utils/kernelutils"
)

type GPIO struct {
	Number    int
	ActiveLow bool
	Value     bool
	Direction enums.GPIOMode
}

// ExportGpio exports the GPIO to the sysfs. If the GPIO is already exported, it does nothing.
func (gpio GPIO) ExportGpio() {

	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)

	if _, err := os.Stat(generalconstants.PathToGpioBase + "gpio" + sysGpio); err != nil {
		if os.IsNotExist(err) {
			err2 := os.WriteFile(generalconstants.PathToGpioBase+"export", []byte(sysGpio), 0666)
			if err2 != nil {
				fmt.Println("failed to open gpio export file for writing")
				os.Exit(1)
			}
		}
	}
}

// UnexportGpio unexports the GPIO from the sysfs.
func (gpio GPIO) UnexportGpio() {
	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)

	err2 := os.WriteFile(generalconstants.PathToGpioBase+"unexport", []byte(sysGpio), 0666)
	if err2 != nil {
		fmt.Println("failed to open gpio unexport file for writing")
		os.Exit(1)
	}

}

// SetModeActiveLow sets the GPIO mode to active low. If ActiveLow is true, the GPIO is set to active low mode, otherwise it is set to normal mode.
func (gpio GPIO) SetModeActiveLow() {
	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)
	var err error
	if gpio.ActiveLow {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/active_low", []byte("1"), 0666)
	} else {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/active_low", []byte("0"), 0666)
	}

	if err != nil {
		fmt.Println("failed to open gpio active_low file for writing")
		os.Exit(1)

	}
}

// WriteGpioValue writes a value to the GPIO. If Value is true, it writes "1" to the GPIO value file, otherwise it writes "0".
func (gpio GPIO) WriteGpioValue() {
	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)

	var err error

	if gpio.Value {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/value", []byte("1"), 0666)
	} else {
		err = os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/value", []byte("0"), 0666)
	}

	if err != nil {
		fmt.Println("failed to open gpio value file for writing")
		os.Exit(1)
	}
}

// ReadGpioValue reads the current value of the GPIO. It reads the GPIO value file and returns true if the value is "1", false if the value is "0".
func (gpio GPIO) ReadGpioValue() bool {
	var value bool
	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)
	e, err := os.ReadFile(generalconstants.PathToGpioBase + "gpio" + sysGpio + "/value")
	gpioStat := strings.TrimSuffix(string(e), "\n")
	if err != nil {
		fmt.Println("failed to open gpio value file for reading")
		os.Exit(1)
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

// SetDirectionGpio sets the direction of the GPIO. It writes the Direction field value to the GPIO direction file.
func (gpio GPIO) SetDirectionGpio() {
	var sysGpio = fmt.Sprint(kernelutils.GetGpioBase() + gpio.Number)
	err := os.WriteFile(generalconstants.PathToGpioBase+"gpio"+sysGpio+"/direction", []byte(gpio.Direction), 0666)

	if err != nil {
		fmt.Println("failed to open gpio direction file for writing")
		os.Exit(1)

	}
}
