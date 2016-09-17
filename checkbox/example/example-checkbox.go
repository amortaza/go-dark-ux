package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/checkbox"
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

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(640,480)

			checkbox.Id("1").Label("Hi!").Left(10).Top(10).Width(160).Height(60)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("2").Label("Hi!").Left(10).Top(100).Width(160).Height(60).Enabled(true).Checked(false)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("3").Label("Hi!").Left(10).Top(190).Width(160).Height(60).Enabled(false).Checked(true)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("4").Label("Hi!").Left(10).Top(280).Width(160).Height(60).Enabled(false).Checked(false)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			button.Id("mybutton").Label("Bye!").Left(10).Top(370).Width(160).Height(60)
			button.OnClick(func() {
				fmt.Println("button clicked")
			})
			button.End()
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


