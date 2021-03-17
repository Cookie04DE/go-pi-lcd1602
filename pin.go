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
	if err := gp.Reconfigure(gpiod.AsInput); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (gp GpiodPin) Output() {
	if err := gp.Reconfigure(gpiod.AsOutput()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
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
