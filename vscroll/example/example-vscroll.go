package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/vscroll"
	"github.com/amortaza/go-dark-ux/label"
	"strconv"
	"github.com/amortaza/go-dark-ux/hscroll"
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

			vscroll.Id("wow1").Left(10).Top(30).Height(450)
			vscroll.On(func(v float32) {
				i = int(v * 101)
			})
			vscroll.Settle()
			vscroll.End()

			vscroll.Id("wow2").Left(480).Top(30).Height(450)
			vscroll.On(func(v float32) {
				i = int(v * 101)
			})
			vscroll.Settle()
			vscroll.End()

			hscroll.Id("wow3").Left(30).Top(20).Width(450)
			hscroll.On(func(v float32) {
				i = int(v * 101)
			})
			hscroll.End()

			hscroll.Id("wow4").Left(30).Top(470).Width(450)
			hscroll.On(func(v float32) {
				i = int(v * 101)
			})
			hscroll.End()

			label.Id("label").Label(strconv.Itoa(i)+"%").Left(100).Top(100).Width(350).Height(200).FontSize(200)
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


