// Package enums defines the GPIOMode type, which is used to set the direction of a GPIO pin.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The GPIOMode type is a string that can be either "in" or "out". These values represent the direction of a GPIO pin.
// "in" means the GPIO pin is set to input mode, and "out" means the GPIO pin is set to output mode.
//
// The package provides constants for these two values, IN and OUT.
package enums

type GPIOMode string

const (
	// IN represents the input mode of a GPIO pin.
	IN = "in"
	// OUT represents the output mode of a GPIO pin.
	OUT = "out"
)
