package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/radiobox/group"
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

			group.Id("wow2").Label("Best way to code Bellina?").Left(150).Top(330).Width(350)
			group.AddChoice("At Night")
			group.AddChoice("During the Day")
			group.AddChoice("Any time")
			group.OnClick(func() {
				fmt.Println("checkbox clicked")
			})
			group.End()
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


