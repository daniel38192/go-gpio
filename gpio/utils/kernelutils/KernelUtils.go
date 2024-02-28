// Package kernelutils provides utility functions to interact with the Linux kernel's sysfs interface for GPIOs.
//
// Author: Luis D. Nuñez V. (daniel38192)
//
// The kernelutils package includes two main functions:
//
// GetGpioBase: This function returns the base address of the GPIOs in the system. It reads the base address from the sysfs interface.
//
// GpioChipSYSList: This function returns a list of all GPIO chips in the system. It reads the list of GPIO chips from the sysfs interface.
//
// These functions provide a low-level interface to the sysfs interface for GPIOs. They are used by the higher-level functions in the gpio package to interact with the GPIOs.
package kernelutils

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"

	generalconstants "github.com/daniel38192/go-gpio/gpio/utils/constants/general"
)

// GetGpioBase returns the base address of the GPIOs in the system. It reads the base address from the sysfs interface.
func GetGpioBase() int {

	var gpiobase string

	sysgpiochipdir := GpioChipSYSList()

	for e := sysgpiochipdir.Front(); e != nil; e = e.Next() {

		name := fmt.Sprint(e.Value)

		path := generalconstants.PathToGpioBase + name + "/base"

		i, err := os.ReadFile(path)

		gpiobase = string(i)
		if err != nil {
			panic(err)

		}

	}

	gpiobaseint, err1 := strconv.Atoi(strings.TrimSuffix(gpiobase, "\n"))

	if err1 != nil {
		panic(err1)
	}

	return gpiobaseint

}

// GpioChipSYSList returns a list of all GPIO chips in the system. It reads the list of GPIO chips from the sysfs interface.
func GpioChipSYSList() list.List {

	sysgpiochipdir := list.New()

	sysdir, err := os.ReadDir(generalconstants.PathToGpioBase)

	if err != nil {
		panic(err)
	}

	for _, e := range sysdir {

		if strings.Contains(e.Name(), "gpiochip") {
			sysgpiochipdir.PushFront(e.Name())
		}

	}

	return *sysgpiochipdir

}
