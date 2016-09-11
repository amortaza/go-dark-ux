package listpane

func ScrollToPercent(listpaneId string, value float32) {
	state := ensureState(listpaneId)

	state.offset = -int(100 * value)
}
