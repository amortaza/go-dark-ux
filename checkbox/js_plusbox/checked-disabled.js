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
ctx.SetFillColor(vgoRGBA(195, 195, 195, 150))
ctx.Text(tx, ty, text)


// var boxy = y + (h - boxw) * 0.5
//
// var bg = vgoBoxGradient(x, boxy, boxw, boxw, 3.0, 3.0, vgoRGBA(0, 0, 0, 32), vgoRGBA(0, 0, 0, 92))
// ctx.BeginPath()
// ctx.RoundedRect(x, boxy, boxw, boxw, 3.0)
// ctx.SetFillPaint(bg)
// ctx.Fill()





ctx.SetFontSize(80.0)
ctx.SetFontFace("icons")

var btnx = x + 19.0
var btny = y + h*0.5

ctx.SetFillColor(vgoRGBA(155, 155, 155, 155))
ctx.SetTextAlign(vgoAlignCenter | vgoAlignMiddle)
ctx.Text(btnx, btny, GuiIconCarrotDown)


// test border
/*
ctx.BeginPath()
ctx.RoundedRect(x, y, w, h, 1.0)

ctx.SetStrokeColor(vgoRGBA(0, 0, 0, 80))
ctx.SetStrokeWidth(2.0)
ctx.Stroke()
*/

ctx.EndFrame()