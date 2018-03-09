package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(1024, 768)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(100,100)
			bl.Dim(800,600)

			bl.Div()
			{
				bl.Id("green")
				bl.Pos(100,100)
				bl.Dim(400,300)

				border.Wire(0, 255, 0)
			}
			bl.End()

			border.Fill(255, 0, 0)
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
	go_dark_ux.Uninit()
}

func init() {
	runtime.LockOSThread()
}

func main() {
	bl.Start( hal_g5.NewHal(), "Dark UX", 1280, 1024, initialize, tick, uninit )

	fmt.Println("bye!")
}


