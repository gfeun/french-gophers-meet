// Connects to an WS2812 RGB LED strip with 10 LEDS.
//
package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ws2812"
)

var leds [8]color.RGBA

func main() {
	var ledColor color.RGBA

	neo := machine.D2
	neo.Configure(machine.PinConfig{Mode: machine.PinOutput})

	ws := ws2812.New(neo)
	rg := false

	for {
		rg = !rg
		if rg {
			ledColor = color.RGBA{R: 0xaa, G: 0x00, B: 0x00}
		} else {
			ledColor = color.RGBA{R: 0x00, G: 0xaa, B: 0x00}
		}

		for i := range leds {
			leds[i] = ledColor
		}

		ws.WriteColors(leds[:])
		time.Sleep(1000 * time.Millisecond)
	}
}
