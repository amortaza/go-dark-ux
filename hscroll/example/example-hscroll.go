package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/hscroll"
	"github.com/amortaza/go-dark-ux/label"
	"strconv"
	"github.com/amortaza/go-dark-ux/border"
)

func initialize() {
	go_dark_ux.Init()
}

var i int

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(800, 700)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(780,680)

			hscroll.Id("wow2").Left(10).Top(300).Width(750)
			hscroll.On(func(v float32) {
				i = int(v * 101)
			})
			hscroll.End()

			label.Id("label").Label(strconv.Itoa(i)+"%").Left(250).Top(100).Width(350).Height(200).FontSize(200)
			border.Wire(255, 255, 0)
			label.End()
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1280, 1024, initialize, tick, uninit )

	fmt.Println("bye!")
}


