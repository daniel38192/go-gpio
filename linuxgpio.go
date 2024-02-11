package linuxgpio

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ExportGpio(gpioNumber int) {

	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)

	if _, err := os.Stat("/sys/class/gpio/gpio" + sysGpio); err != nil {
		if os.IsNotExist(err) {
			err2 := os.WriteFile("/sys/class/gpio/export", []byte(sysGpio), 0666)
			if err2 != nil {
				fmt.Println("failed to open gpio export file for writing")
				os.Exit(1)
			}
		}
	}

}

func UnexportGpio(gpioNumber int) {
	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)

	err2 := os.WriteFile("/sys/class/gpio/unexport", []byte(sysGpio), 0666)
	if err2 != nil {
		fmt.Println("failed to open gpio unexport file for writing")
		os.Exit(1)
	}

}

func SetModeActiveLow(gpioNumber int, activeLow bool) {
	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)
	var err error
	if activeLow == true {
		err = os.WriteFile("/sys/class/gpio/gpio"+sysGpio+"/active_low", []byte("1"), 0666)
	} else if activeLow == false {
		err = os.WriteFile("/sys/class/gpio/gpio"+sysGpio+"/active_low", []byte("0"), 0666)
	}

	if err != nil {
		fmt.Println("failed to open gpio active_low file for writing")
		os.Exit(1)

	}
}

func WriteGpioValue(gpioNumber int, value bool) {
	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)

	var err error

	if value == true {
		err = os.WriteFile("/sys/class/gpio/gpio"+sysGpio+"/value", []byte("1"), 0666)
	} else if value == false {
		err = os.WriteFile("/sys/class/gpio/gpio"+sysGpio+"/value", []byte("0"), 0666)
	}

	if err != nil {
		fmt.Println("failed to open gpio value file for writing")
		os.Exit(1)

	}
}

func ReadGpioValue(gpioNumber int) bool {
	var value bool
	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)
	e, err := os.ReadFile("/sys/class/gpio/gpio" + sysGpio + "/value")
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

func SetDirectionGpio(gpioNumber int, direction string) {
	var sysGpio = fmt.Sprint(GetGpioBase() + gpioNumber)

	var err error

	err = os.WriteFile("/sys/class/gpio/gpio"+sysGpio+"/direction", []byte(direction), 0666)

	if err != nil {
		fmt.Println("failed to open gpio direction file for writing")
		os.Exit(1)

	}
}

func GetGpioBase() int {

	var gpiobase string

	sysgpiochipdir := GpioChipSYSList()

	for e := sysgpiochipdir.Front(); e != nil; e = e.Next() {

		name := fmt.Sprint(e.Value)

		path := "/sys/class/gpio/" + name + "/base"

		i, err := os.ReadFile(path)

		gpiobase = string(i)
		if err != nil {
			fmt.Println("error at reading base file")
			os.Exit(1)

		}

	}

	gpiobaseint, err1 := strconv.Atoi(strings.TrimSuffix(gpiobase, "\n"))

	if err1 != nil {
		fmt.Println("unsupported gpiobase number")
		os.Exit(1)
	}

	return gpiobaseint

}

func GpioChipSYSList() list.List {

	sysgpiochipdir := list.New()

	sysdir, err := os.ReadDir("/sys/class/gpio/")

	if err != nil {
		fmt.Println("error at opening sys dir")
		os.Exit(1)

	}

	for _, e := range sysdir {

		if strings.Contains(e.Name(), "gpiochip") == true {
			sysgpiochipdir.PushFront(e.Name())
		}

	}

	return *sysgpiochipdir

}

func GpioDirectionIN() string { return "in" }

func GpioDirectionOUT() string { return "out" }
