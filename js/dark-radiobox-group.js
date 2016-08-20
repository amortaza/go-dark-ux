ctx.BeginFrame(w, h, 1.0)

var boxw = 36.0
var alpha = 60.0
var col = vgoRGBA(30, 30, 30, alpha)

ctx.SetFontSize(36.0)
ctx.SetFontFace("sans")
ctx.SetTextAlign(vgoAlignLeft | vgoAlignTop)

var tx = x + 6, ty = y

ctx.SetFontBlur(2.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, 125))
ctx.Text(tx, ty, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(255, 255, 255, 160))
ctx.Text(tx, ty, text)

// test border
if (0) {
    ctx.BeginPath()
    ctx.RoundedRect(x, y, w, h, 1.0)

    ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 180))
    ctx.SetStrokeWidth(2.0)
    ctx.Stroke()
}

ctx.EndFrame()