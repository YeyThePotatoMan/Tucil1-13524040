package main

import (
	"image/color"
)

var (
	n      int
	g      []string
	ans    []int
	found  bool
	cnt    int
	stop   bool
	delay  int = 50
	uiChan chan bool
)

var Colors = []color.Color{
	color.RGBA{255, 105, 97, 255},
	color.RGBA{255, 180, 71, 255},
	color.RGBA{253, 253, 150, 255},
	color.RGBA{119, 221, 119, 255},
	color.RGBA{119, 190, 255, 255},
	color.RGBA{189, 178, 255, 255},
	color.RGBA{255, 150, 255, 255},
	color.RGBA{128, 128, 128, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{0, 255, 255, 255},
	color.RGBA{255, 0, 255, 255},
	color.RGBA{192, 192, 192, 255},
	color.RGBA{128, 0, 0, 255},
	color.RGBA{128, 128, 0, 255},
	color.RGBA{0, 128, 0, 255},
	color.RGBA{128, 0, 128, 255},
	color.RGBA{0, 128, 128, 255},
	color.RGBA{0, 0, 128, 255},
	color.RGBA{255, 140, 0, 255},
	color.RGBA{153, 50, 204, 255},
	color.RGBA{46, 139, 87, 255},
	color.RGBA{70, 130, 180, 255},
	color.RGBA{210, 105, 30, 255},
}
