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

			checkbox.Id("1").Label("Enabled Checked").Left(10).Top(10).Width(160).Height(60).Checked(false)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("2").Label("Enabled").Left(10).Top(100).Width(160).Height(60).Enabled(true).Checked(true)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("3").Label("Disabled Checked").Left(10).Top(190).Width(160).Height(60).Enabled(false).Checked(true)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			checkbox.Id("4").Label("Disabled Unchecked").Left(10).Top(280).Width(160).Height(60).Enabled(false).Checked(false)
			checkbox.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			checkbox.End()

			button.Id("mybutton").Label("Enabled").Left(10).Top(370).Width(160).Height(60)
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
	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1280, 1024, initialize, tick, uninit )

	fmt.Println("bye!")
}


