package main

import (
	"fmt"
	"github.com/amortaza/go-xel2"
	"github.com/amortaza/go-g5"
	gl3 "github.com/chsc/gogl/gl33"
	"time"
	"github.com/amortaza/go-ux"
	"github.com/amortaza/go-dark-ux"
)

var w, h int = 0, 0

func afterGL() {

	e := gl3.Init()

	if e != nil {
		panic("ok")
	}

	g5.Init()

	ux.Init()

	go_dark_ux.Init()
}

func onDelete() {
	g5.Uninit()
	ux.Uninit()
}

func onLoop() {
	g5.PushView(w/2,h/2)

	g5.Clear(.33,.33,.33, 1.0)

	go_dark_ux.DrawButton_Default(10, 10, 95, 30, "OK")
	//go_dark_ux.DrawCheckbox_Unchecked_Enabled(10, 10, 95, 30, "OK")

	g5.PopView()

	time.Sleep(500 * time.Millisecond)
}

func onResize(a,b int) {
	w,h=a,b
	fmt.Println(w, " ", h)
}

func main() {
	xel.Init("Welcome, home!", 1280, 960)

	xel.SetCallbacks(afterGL, onLoop, onDelete, onResize, nil, nil, nil )
	xel.Loop()
	xel.Uninit()
}

