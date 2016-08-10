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

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(640, 512)

		bl.Div()
		{
			bl.Id("1")
			bl.Pos(10,10)
			bl.Dim(300,240)
			bl.Div()
			{
				bl.Id("1/1")
				bl.Pos(10,10)
				bl.Dim(300,240)

				bl.Div()
				{
					bl.Id("1/1/1")
					bl.Pos(10,10)
					bl.Dim(300,240)

					button.Id("1").Label("Wusup").Width(10).Height(10).Left(30).Top(30)
					button.OnClick(func() {
						fmt.Println("clicked")
					})
					button.End()
				}
				bl.End()
			}
			bl.End()
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


