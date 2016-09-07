ctx.BeginFrame(w, h, 1.0)
x++
y++
w -= 2
h -= 2

if (inHasBack) {
	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, 1.0)
	ctx.SetFillColor(vgoRGBA(inBackRed,inBackGreen,inBackBlue, inBackAlpha))
	ctx.Fill()
}

ctx.SetFontSize(FontSize)
ctx.SetFontFace("sans")

ctx.SetTextAlign(vgoAlignLeft | vgoAlignMiddle)

var btnx = 10.0 // w / 120 - sysGetTextWidth(text) / 120
var btny = h / 2 - 1

ctx.SetFontBlur(1.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, 255))
ctx.Text(btnx, btny, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(inTextRed,inTextGreen,inTextBlue, inTextAlpha))
ctx.Text(btnx, btny, text)

// test border
if (0) {
    ctx.BeginPath()
    ctx.RoundedRect(x, y, w, h, 1.0)

    ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 180))
    ctx.SetStrokeWidth(2.0)
    ctx.Stroke()
}

ctx.EndFrame()