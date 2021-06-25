package main

import (
	"flag"
	"gopkg.in/RockKeeper/go-rpi-rgb-led-matrix.v1"
	"image/color"
	"time"
)

var (
	rows                     = flag.Int("led-rows", 64, "number of rows supported")
	cols                     = flag.Int("led-cols", 64, "number of columns supported")
	parallel                 = flag.Int("led-parallel", 1, "number of daisy-chained panels")
	chain                    = flag.Int("led-chain", 2, "number of displays daisy-chained")
	brightness               = flag.Int("brightness", 100, "brightness (0-100)")
	hardware_mapping         = flag.String("led-gpio-mapping", "regular", "Name of GPIO mapping used.")
	show_refresh             = flag.Bool("led-show-refresh", false, "Show refresh rate.")
	inverse_colors           = flag.Bool("led-inverse", false, "Switch if your matrix has inverse colors on.")
	disable_hardware_pulsing = flag.Bool("led-no-hardware-pulse", false, "Don't use hardware pin-pulse generation.")
)

func main() {
	config := &rgbmatrix.DefaultConfig
	config.Rows = *rows
	config.Cols = *cols
	config.Parallel = *parallel
	config.ChainLength = *chain
	config.Brightness = *brightness
	config.HardwareMapping = *hardware_mapping
	config.ShowRefreshRate = *show_refresh
	config.InverseColors = *inverse_colors
	config.DisableHardwarePulsing = *disable_hardware_pulsing

	m, err := rgbmatrix.NewRGBLedMatrix(config)
	fatal(err)

	c := rgbmatrix.NewCanvas(m)

	start := time.Now()
	defer func() {
		time.Sleep((time.Second * 30) - time.Since(start))
		c.Close()
	}()

	col := color.RGBA{
		R: 150,
		G: 1,
		B: 255,
		A: 255,
	}
	c.DrawText(1, 1, "Hello", col)
	c.DrawText(1, 10, "Hello", col)
	c.DrawText(1, 20, "Hello", col)
	c.DrawText(1, 30, "Hello", col)
	c.DrawText(1, 40, "Hello", col)
	c.DrawText(1, 50, "Hello", col)

	c.Render()

}

func init() {
	flag.Parse()
}

func fatal(err error) {
	if err != nil {
		panic(err)
	}
}
