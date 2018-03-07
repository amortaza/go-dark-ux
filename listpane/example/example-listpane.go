package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/listpane"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		bl.Pos(0,0)
		bl.Dim(800, 700)

		bl.Div()
		{
			bl.Id("red")
			bl.Pos(10,10)
			bl.Dim(780,680)

			listpane.Id("wow2").Left(30).Top(30).Width(250).Height(250)
			listpane.On(func(v float32) {
				//i = int(v * 101)
			})
			listpane.Item("One")
			listpane.Item("Two")
			listpane.Item("Three")
			listpane.Item("Four")
			listpane.Item("Five")
			listpane.End()

			//label.Id("label").Label(strconv.Itoa(i)+"%").Left(250).Top(100).Width(350).Height(200).FontSize(200)
			//border.Draw()
			//label.End()
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


