ctx.BeginFrame(inViewWidth, inViewHeight, 1.0)

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
    else {
        ctx.SetStrokeColor(vgoRGBA(inRed, inGreen, inBlue, 255))
        ctx.SetStrokeWidth(3.0)
        ctx.Stroke()
    }

    ctx.SetFontSize(30.0)
    ctx.SetFontFace("sans")

    ctx.SetTextAlign(vgoAlignLeft | vgoAlignMiddle)

    ctx.SetFontBlur(1.0)
    ctx.SetFillColor(vgoRGBA(0, 0, 0, 255))
    ctx.Text(ox + 20.0, oy + 20.0, text)

    ctx.SetFontBlur(0.0)
    ctx.SetFillColor(vgoRGBA(255, 255, 255, 255))
    ctx.Text(ox + 20.0, oy + 20.0, text)


ctx.EndFrame()