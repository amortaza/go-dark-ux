ctx.BeginFrame(w, h, 1.0)

var ix = x + 1
var iy = y + 1
var iw = w - 2
var ih = h - 2

var col = vgoRGBA(30, 30, 30, 60.0)

ctx.SetFontSize(36.0)
ctx.SetFontFace("sans")
ctx.SetTextAlign(vgoAlignLeft | vgoAlignTop)

var tx = x + 12, ty = y + 6

ctx.SetFontBlur(2.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, 125))
ctx.Text(tx, ty, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(255, 255, 255, 160))
ctx.Text(tx, ty, text)

// border
ctx.BeginPath()
ctx.RoundedRect(ix, iy, iw, ih, 1.0)

ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 80))
ctx.SetStrokeWidth(3.0)
ctx.Stroke()

// test border
if (0) {
    ctx.BeginPath()
    ctx.RoundedRect(x, y, w, h, 1.0)

    ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 180))
    ctx.SetStrokeWidth(2.0)
    ctx.Stroke()
}

ctx.EndFrame()