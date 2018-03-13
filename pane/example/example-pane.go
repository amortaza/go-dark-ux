package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/pane"
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

			pane.Id("mypane").Label("Cool").End()
			//
			//checkbox.Id("mycheckbox").Label("Hi!").Left(30).Top(30).Width(160).Height(60)
			//checkbox.OnClick(func() {
			//	fmt.Println("checkbox clicked")
			//})
			//checkbox.End()
			//
			//button.Id("mybutton").Label("Bye!").Left(30).Top(90).Width(160).Height(60)
			//button.OnClick(func() {
			//	fmt.Println("button clicked")
			//})
			//button.End()
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

	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1024, 768, initialize, tick, uninit )

	fmt.Println("bye!")
}


