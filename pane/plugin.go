package pane

import (
	"github.com/amortaza/go-bellina"
	"fmt"
)

var plugin *Plugin

type Plugin struct {
}

func (c *Plugin) Name() string {
	return "pane"
}

func (c *Plugin) Init() {
	fmt.Println("pane plugin Init is being called")
	gStateByNode = make(map[string] *State)
}

func (c *Plugin) Tick() {
}

func (c *Plugin) Reset() {
}

func (c *Plugin) OnNodeAdded(node *bl.Node) {
}

func (c *Plugin) OnNodeRemoved(node *bl.Node) {
}

func (c *Plugin) Uninit() {
}


