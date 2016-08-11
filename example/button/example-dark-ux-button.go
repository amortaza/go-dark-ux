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
var x int
func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(800, 600)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(640,480)
			bl.Div()
			{
				bl.Id("green")
				bl.Pos(10,10)
				bl.Dim(300,200)

				button.Id("1").Label("Wusup").Left(10).Top(30).Width(110).Height(30)
				button.OnClick(func() {
					fmt.Println("clicked")
				})
				button.End()

				//bl.Div()
				//{
				//	bl.Id("blue")
				//	bl.Pos(x,10)
				//	bl.Dim(100,100)
				//
				//
				//	x += 1
				//}
				//bl.End()
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


