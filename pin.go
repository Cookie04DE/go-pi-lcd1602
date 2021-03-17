package lcd1602

import (
	"fmt"
	"os"

	"github.com/warthog618/gpiod"
)

type Pin interface {
	Input()
	Output()
	High()
	Low()
}

type GpiodPin struct {
	*gpiod.Line
}

func (gp GpiodPin) Input() {
	gp.Reconfigure(gpiod.AsInput)
}

func (gp GpiodPin) Output() {
	gp.Reconfigure(gpiod.AsOutput())
}

func (gp GpiodPin) High() {
	if err := gp.SetValue(1); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (gp GpiodPin) Low() {
	if err := gp.SetValue(0); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
