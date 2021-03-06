package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-dark-ux/button"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(640, 480)

		label.Id("first").
			HasBack(true).
				BackColor4i(255, 0, 0, 255).
					Left(10).Top(10).
						Label("Hello World!").
							End()

		button.Id("two").Left(10).Top(120).Width(200).Height(90).Label("Push")
		button.OnClick(func() {
			fmt.Println("Pushed")
		})
		button.End()
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


