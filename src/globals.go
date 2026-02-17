package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var (
	n     int
	g     []string
	ans   []int
	found bool
	cnt   int
	stop  bool
	delay int = 50

	grid   *fyne.Container
	rects  [][]*canvas.Rectangle
	texts  [][]*canvas.Image
	uiChan chan bool
)

var Colors = []color.Color{
	color.RGBA{220, 20, 60, 255},
	color.RGBA{65, 105, 225, 255},
	color.RGBA{34, 139, 34, 255},
	color.RGBA{255, 215, 0, 255},
	color.RGBA{138, 43, 226, 255},
	color.RGBA{255, 140, 0, 255},
	color.RGBA{0, 128, 128, 255},
	color.RGBA{255, 105, 180, 255},
	color.RGBA{112, 128, 144, 255},
	color.RGBA{139, 69, 19, 255},
	color.RGBA{0, 191, 255, 255},
	color.RGBA{154, 205, 50, 255},
	color.RGBA{210, 105, 30, 255},
	color.RGBA{75, 0, 130, 255},
	color.RGBA{255, 127, 80, 255},
	color.RGBA{70, 130, 180, 255},
	color.RGBA{128, 128, 0, 255},
	color.RGBA{218, 112, 214, 255},
	color.RGBA{178, 34, 34, 255},
	color.RGBA{95, 158, 160, 255},
	color.RGBA{189, 183, 107, 255},
	color.RGBA{147, 112, 219, 255},
	color.RGBA{46, 139, 87, 255},
	color.RGBA{240, 128, 128, 255},
	color.RGBA{100, 149, 237, 255},
	color.RGBA{255, 165, 0, 255},
} // credits to : https://www.w3schools.com/cssref/css_colors.php
