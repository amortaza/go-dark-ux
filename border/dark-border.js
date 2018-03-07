ctx.BeginFrame(w, h, 1.0)

    var ox = x
    var oy = y
    var ow = w
    var oh = h

    ctx.BeginPath()
    ctx.RoundedRect(ox, oy, ow, oh, 1.0)

	if (inFill == 1) {
		ctx.SetFillColor(vgoRGBA(inRed, inGreen, inBlue, 255))
		ctx.Fill()
	}

    ctx.SetStrokeColor(vgoRGBA(inRed, inGreen, inBlue, 255))
    ctx.SetStrokeWidth(3.0)
    ctx.Stroke()

ctx.EndFrame()