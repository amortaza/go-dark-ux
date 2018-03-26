package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/accordian-menu/tile"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		accordian_menu_tile.Use().Title("one")
		accordian_menu_tile.End()

	}
	bl.End()
}

func init() {
	runtime.LockOSThread()
}

func main() {

	bl.Start( hal_g5.NewHal(), "Bellina v0.2", 1280, 1024, initialize, tick, nil )

	fmt.Println("bye!")
}


