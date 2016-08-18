package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-oob"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/pane"
	"github.com/amortaza/go-dark-ux/checkbox"
	"github.com/amortaza/go-dark-ux/radiobox/group"
)

func initialize() {
	go_dark_ux.Init()
}

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

			pane.Id("mypane").Label("Cool").Left(10).Top(10).Width(500).Height(400).End()

			//radiobox.AddChoice("One")
			//radiobox.AddChoice("Two")
			//radiobox.AddChoice("Three")
			radio.Id("wow").Label("Hi!").Left(30).Top(30).Width(160)
			radio.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			radio.End()

			checkbox.Id("mycheckbox").Label("Hi!").Left(30).Top(230).Width(160).Height(60)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()
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


