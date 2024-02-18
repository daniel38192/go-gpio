// Module: gpio
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The gpio module provides functionality to interact with GPIO (General Purpose Input/Output) pins on a Linux system.
// It includes a GPIO struct and methods to export and unexport the GPIO, set the mode to active low,
// write and read the GPIO value, and set the GPIO direction.
//
// The module is organized into several files:
//
// linuxgpio.go: This file defines the GPIO struct and its methods. The GPIO struct represents a GPIO pin with its properties.
//
// enums/GPIOMode.go: This file defines the GPIOMode type, which is used to set the direction of a GPIO pin.
//
// utils/KernelUtils.go: This file provides utility functions to interact with the Linux kernel's sysfs interface for GPIOs.
//
// utils/constants/general/Constants.go: This file defines constants used in the gpio module.
//
// go.mod: This file defines the module's dependencies.
//
// The gpio module uses the sysfs interface provided by the Linux kernel to interact with GPIO pins.
// The sysfs interface is a virtual file system provided by the Linux kernel that allows user space programs to interact with kernel objects.
// In the case of GPIOs, each GPIO pin is represented as a directory in the sysfs file system, and the properties of the GPIO pin can be read and written by reading and writing files in that directory.
//
// The gpio module provides a high-level interface to this sysfs interface, allowing the user to interact with GPIO pins using simple methods and properties, without needing to know the details of the sysfs interface.
module linuxgpio

go 1.21.6
