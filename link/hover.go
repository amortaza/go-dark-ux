package link

import (
	"github.com/amortaza/go-bellina-plugins/hover"
	"github.com/amortaza/go-bellina-plugins/click"
)

type HoverEvent struct {

	linkId string
	payload string
}

func setup_hover(state *State) {

	hover.On(func (v interface{}) {

		state.dirty = true

		e := v.(*hover.Event)

		if e.IsInEvent {

			state.state = 1

		} else {
			state.state = 0
		}
	})

	click.On_WithLifeCycle(

		func(e interface{}) {
			state.state = 0
			state.dirty = true

			if state.onClick != nil {

				event := &HoverEvent{linkId: state.linkId, payload: state.payload}

				state.onClick(event)
			}
		},

		func(e interface{}) {
			state.state = 2
			state.dirty = true
		},

		func(e interface{}) {
			state.state = 0
			state.dirty = true
		})
}
