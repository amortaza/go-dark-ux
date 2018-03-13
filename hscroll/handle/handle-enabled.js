ctx.BeginFrame(w, h, 1.0)

y--
h+=3

var ix = x + 1
var iy = y + 1
var iw = w - 2
var ih = h - 2

var c = 240.0
var bg = vgoBoxGradient(x-4, y+1, w+8, h - 4.0, 15.0, 5.0, vgoRGBA(c,c,c, 132.0), vgoRGBA(32.0, 32.0, 32.0, 32.0))
var radius = 8.0

ctx.BeginPath()
ctx.RoundedRect(x, y, w, h, radius)
ctx.SetFillPaint(bg)
ctx.Fill()

ctx.BeginPath()
ctx.RoundedRect(ix,iy,iw,ih, radius)
ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 99))
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