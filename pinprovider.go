package lcd1602

import (
	"fmt"
	"os"

	"github.com/stianeikeland/go-rpio"
	"github.com/warthog618/gpiod"
)

type PinProvider interface {
	Open()
	Close()
	GetPin(line int) Pin
}

type RpioProvider struct {
}

//function should be called before executing any other code!
func (rp RpioProvider) Open() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (rp RpioProvider) Close() {
	rpio.Close()
}

func (rp RpioProvider) GetPin(line int) Pin {
	return rpio.Pin(line)
}

type GpiodProvider struct {
	chip     *gpiod.Chip
	chipName string
}

func (gpp *GpiodProvider) Open() {
	chip, err := gpiod.NewChip(gpp.chipName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	gpp.chip = chip
}

func (gpp *GpiodProvider) Close() {
	err := gpp.chip.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (gpp *GpiodProvider) GetPin(line int) Pin {
	l, err := gpp.chip.RequestLine(line)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return GpiodPin{Line: l}
}
