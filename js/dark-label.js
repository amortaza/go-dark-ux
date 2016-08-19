ctx.BeginFrame(w, h, 1.0)

x++
y++
w -= 2
h -= 2

ctx.SetFontSize(36.0)
ctx.SetFontFace("sans")

ctx.SetTextAlign(vgoAlignLeft | vgoAlignMiddle)

var btnx = w / 2 - sysGetTextWidth(text) / 2
var btny = h / 2 - 1

ctx.SetFontBlur(1.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, 255))
ctx.Text(btnx, btny, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(255, 255, 255, 170))
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