package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/xray"
	"github.com/amortaza/go-dark-ux/border"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(640, 480)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(640,480)
			border.Wire(255, 0, 0)
		}
		bl.End()

		bl.Div()
		{
			bl.Id("blue")
			bl.Pos(120,410)
			bl.Dim(640,480)
			border.Wire(255, 0, 250)
		}
		bl.End()

		xray.Use()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1400, 1100, initialize, tick, uninit )

	fmt.Println("bye!")
}


