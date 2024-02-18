// Package general defines constants used in the gpio module.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The general package includes one main constant:
//
// PathToGpioBase: This constant represents the path to the GPIO base in the sysfs interface.
//
// This constant is used by the functions in the kernelutils package to interact with the sysfs interface for GPIOs.
package general

const (
	// PathToGpioBase represents the path to the GPIO base in the sysfs interface.
	PathToGpioBase = "/sys/class/gpio/"
)
