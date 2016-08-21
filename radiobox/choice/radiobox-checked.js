ctx.BeginFrame(w, h, 1.0)

var boxw = 36.0
var alpha = 60.0
var col = vgoRGBA(30, 30, 30, alpha)

ctx.SetFontSize(36.0)
ctx.SetFontFace("sans")
ctx.SetTextAlign(vgoAlignLeft | vgoAlignMiddle)

var tx = x + boxw + 16.0, ty = y + h * 0.5 - 3.0

ctx.SetFontBlur(2.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, 125))
ctx.Text(tx, ty, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(255, 255, 255, 160))
ctx.Text(tx, ty, text)

var r = 17.0
var innerR = 10.0

var cx = x + r + 1;
var cy = y + h * 0.5

ctx.BeginPath()
ctx.Circle(cx,cy, r)
ctx.SetFillColor(vgoRGBA(63,60,63,105))
ctx.Fill()

ctx.BeginPath()

alpha = 70
// Outer Circle
ctx.Circle(cx,cy, r)
ctx.SetStrokeWidth(2.0)
ctx.SetStrokeColor(vgoRGBA(0, 0, 0, alpha))
ctx.Stroke()

// Inner Circle
ctx.BeginPath()
ctx.Circle(cx,cy, innerR)
ctx.SetFillColor(vgoRGBA(215,215,213,165))
ctx.Fill()

ctx.SetStrokeColor(vgoRGBA(0, 0, 0, alpha + 20))
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