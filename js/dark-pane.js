ctx.BeginFrame(w, h, 1.0)

var ox = x
var oy = y
var ow = w
var oh = h

var c = 85.0

var cornerRadius = 4.0

ctx.BeginPath()
ctx.RoundedRect(ox, oy, ow, oh, cornerRadius)
ctx.SetFillColor(vgoRGBA(c,c,c,255))
ctx.Fill()

// Header
if (0) {
    var headerPaint = vgoLinearGradient(x, y, x, y+15.0, vgoRGBA(255, 255, 255, 8), vgoRGBA(0, 0, 0, 10))

    ctx.BeginPath()
    ctx.RoundedRect(x+1.0, y+1.0, w-2.0, 30.0, cornerRadius-1.0)
    ctx.SetFillPaint(headerPaint)
    ctx.Fill()
    ctx.BeginPath()
    ctx.MoveTo(x+0.5, y+0.5+30)
    ctx.LineTo(x+0.5+w-1, y+0.5+30)
    ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 32))
    ctx.Stroke()

    ctx.SetFontSize(18.0)
    ctx.SetFontFace("sans")
    ctx.SetTextAlign(vgoAlignCenter | vgoAlignMiddle)

    ctx.SetFontBlur(1.0)
    ctx.SetFillColor(vgoRGBA(0, 0, 0, 238))
    ctx.Text(x+30.0, y+16.0, text)

    ctx.SetFontBlur(0.0)
    ctx.SetFillColor(vgoRGBA(255, 255, 255, 160))
    ctx.Text(x+30.0, y+16.0, text)
 }

    ctx.BeginPath()
    ctx.RoundedRect(ox, oy, ow, oh, 1.0)

    var c = 15
    ctx.SetStrokeColor(vgoRGBA(c, c, c, 180))
    ctx.SetStrokeWidth(3.0)
    ctx.Stroke()

ctx.EndFrame()