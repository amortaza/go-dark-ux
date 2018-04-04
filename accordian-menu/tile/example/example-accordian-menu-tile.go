package main

import (
	"runtime"
	"fmt"
	"github.com/amortaza/go-bellina"
	"github.com/amortaza/go-hal-g5"
	"github.com/amortaza/go-dark-ux"
	"github.com/amortaza/go-dark-ux/accordian-menu/tile"
	"github.com/amortaza/go-bellina-plugins/click"
	"github.com/amortaza/go-bellina-plugins/animation"
	"github.com/amortaza/go-bellina-plugins/layout/vert"
	"github.com/amortaza/go-dark-ux/label"
	"github.com/amortaza/go-bellina-plugins/layout/pack"
)

func initialize() {
	go_dark_ux.Init()
}

func tick() {

	bl.Root()
	{
		accordian_menu_tile.Id("acc-1").Title("one")
		{
			items();
		}
		accordian_menu_tile.End()

		accordian_menu_tile.Id("acc-2").Title("two")
		{
			items2();
		}
		accordian_menu_tile.End()

		vert.Use().Spacing(10).End()
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

func items() {

	bl.Div()
	{
		bl.Id("i-1")
		bl.Pos(10,100)

		label.Id("aace").Label("Incidents").Width(200).Height(50).End()

		pack.Use().End()

		//border.Wire(150,120,20)
	}
	bl.End()

	bl.Div()
	{
		bl.Id("i-2")
		bl.Pos(10,100)

		label.Id("dict").Label("Dictionary").Width(200).Height(50).End()
		//border.Wire(50,200,20)

		pack.Use().End()

		click.On(func(i interface{}) {
			animation.StartPath("i-2", "hi2", float32(100), float32(300), 50, func(shadow *bl.ShadowNode, value float32) {
				shadow.Height = int(value)
			})
		})
	}
	bl.End()

	bl.Div()
	{
		bl.Id("i-3")
		bl.Pos(10,100)

		label.Id("Business Rules").Label("Business Rules").Width(400).Height(50).End()

		pack.Use().End()
		//border.Wire(10,20,200)
	}
	bl.End()

}

func items2() {

	bl.Div()
	{
		bl.Id("j-1")
		bl.Dim(100,50)
		bl.Pos(10,100)

		label.Id("aBusiness Rules").Label("Scheduled Jobs").Width(400).Height(50).End()
		pack.Use().End()

		//border.Wire_TopsKids(150,120,20)
	}
	bl.End()

	bl.Div()
	{
		bl.Id("j-2")
		bl.Dim(100,50)
		bl.Pos(10,100)

		label.Id("bBusiness Rules").Label("Script Includes").Width(400).Height(50).End()
		pack.Use().End()
		//border.Wire(50,200,20)

		click.On(func(i interface{}) {
			animation.StartPath("j-2", "hi3", float32(100), float32(300), 50, func(shadow *bl.ShadowNode, value float32) {
				shadow.Height = int(value)
			})
		})
	}
	bl.End()

	bl.Div()
	{
		bl.Id("j-3")
		bl.Dim(100,50)
		bl.Pos(10,100)

		label.Id("cBusiness Rules").Label("Access Controls").Width(400).Height(50).End()
		pack.Use().End()
		//border.Wire(10,20,200)
	}
	bl.End()

}

