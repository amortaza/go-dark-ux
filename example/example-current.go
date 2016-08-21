package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-oob"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/border"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
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
			bl.Id("zoo")
			bl.Pos(50,50)
			bl.Dim(800, 600)
			border.Draw()

			bl.Div()
			{
				bl.Id("one")
				bl.Pos(100,100)
				bl.Dim(100,100)

				border.Draw()
			}
			bl.End()

			bl.Div()
			{
				bl.Id("two")
				bl.Pos(100,100)
				bl.Dim(100,100)

				border.Draw()
			}
			bl.End()

			vert.Id("one").Spacing(10).End()
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


