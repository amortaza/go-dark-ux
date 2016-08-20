package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-oob"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/button"
)

func initialize() {
	go_dark_ux.Init()
}
var x int = 10
func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(640, 480)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(x,10)
			bl.Dim(320,240)

			button.Id("1").Label("Hi!").Left(x).Top(10).Width(160).Height(60)
			button.OnClick(func() {
				fmt.Println("clicked")
			})
			button.End()

			button.Id("2").Label("Bye!").Left(x).Top(90).Width(160).Height(60)
			button.OnClick(func() {
				fmt.Println("clicked")
			})
			button.End()

			//x += 3
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
	bl.Start( haloob.New(), 1280, 1024, "Bellina v0.2", initialize, tick, uninit )

	fmt.Println("bye!")
}


