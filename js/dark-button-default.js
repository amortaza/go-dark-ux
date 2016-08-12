ctx.BeginFrame(w, h, 1.0)

x++
y++
w -= 2
h -= 2

var cornerRadius = 4.0

var bb = 100.0

ctx.BeginPath()
ctx.RoundedRect(x, y, w, h, cornerRadius)
ctx.SetFillColor(vgoRGBA(bb, bb, bb, 255))
ctx.Fill()

var c1 = 255.0
var c2 = 0.0
var bg = vgoLinearGradient(x, y, x, y+h, vgoRGBA(c1, c1, c1, 16), vgoRGBA(c2, c2, c2, 20))

ctx.BeginPath()
ctx.RoundedRect(x, y, w, h, cornerRadius)

ctx.SetFillPaint(bg)
ctx.Fill()

ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 80))
ctx.SetStrokeWidth(2.0)
ctx.Stroke()

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

ctx.EndFrame()