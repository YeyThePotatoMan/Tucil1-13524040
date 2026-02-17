package main

import (
	"bufio"
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func read_file(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}

	sc := bufio.NewScanner(f)
	g = []string{}
	for sc.Scan() {
		cnt += 1
		line := sc.Text()
		if len(line) > 0 {
			g = append(g, line)
		}
	}
	// TODO : add validation

	n = len(g)
	ans = make([]int, n)
	for i := range ans {
		ans[i] = -1
	}
	return true
}

func refresh_grid() {
	if n == 0 {
		return
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			txt := ""
			if ans[r] == c {
				txt = "#"
			}
			texts[r][c].Text = txt
			texts[r][c].Refresh()
		}
	}
}

func build_grid() {
	rects = make([][]*canvas.Rectangle, n)
	texts = make([][]*canvas.Text, n)
	objects := []fyne.CanvasObject{}

	for r := 0; r < n; r++ {
		rects[r] = make([]*canvas.Rectangle, n)
		texts[r] = make([]*canvas.Text, n)
		for c := 0; c < n; c++ {
			rect := canvas.NewRectangle(Colors[g[r][c]-'A'])
			rect.StrokeColor = Colors[g[r][c]-'A']

			text := canvas.NewText("", color.Black)
			text.TextStyle.Bold = true
			text.Alignment = fyne.TextAlignCenter
			text.TextSize = 20

			rects[r][c] = rect
			texts[r][c] = text

			stack := container.NewStack(rect, text)
			objects = append(objects, stack)
		}
	}
	grid.Layout = layout.NewGridLayout(n)
	grid.Objects = objects
	grid.Refresh()
}

func main() {
	fmt.Println("working")
	a := app.New()
	w := a.NewWindow("Tucil 1")
	w.Resize(fyne.NewSize(800, 600))

	grid = container.New(layout.NewGridLayout(1))
	lblTime := widget.NewLabel("Time: 0ms")
	lblIter := widget.NewLabel("Iter: 0")

	btn1 := widget.NewButton("Solver 1 (Live without pruning)", func() {})
	btn2 := widget.NewButton("Solver 2 (Live with Pruning)", func() {})
	btn3 := widget.NewButton("Solver 3 (No live updates)", func() {})

	slider := widget.NewSlider(1, 500)
	slider.SetValue(50)
	slider.OnChanged = func(v float64) { delay = int(v) }

	uiChan = make(chan bool)
	go func() {
		for range uiChan {
			refresh_grid()
		}
	}()

	sidepanel := container.NewVBox(
		btn1, btn2, btn3,
		slider,
		lblTime, lblIter,
	)

	split := container.NewHSplit(sidepanel, container.NewPadded(grid))
	w.SetContent(split)
	w.ShowAndRun()
}
