ctx.BeginFrame(w, h, 1.0)

// 234,134,60, 250
// 214,134,60, 250
// 214; 234; 160; // green

console.log(inState)

if (inState == 0) {

    // default
    hasBack = 0

    backRed = backGreen = backBlue = 255
    backAlpha = 0

    textRed = 214; textGreen = 134; textBlue = 60;
    blurrAlpha = textAlpha = 255
}
else if (inState == 1) {

    // hover
    hasBack = 1

    backRed = backGreen = backBlue = 255
    backAlpha = 50

    textRed = 214; textGreen = 134; textBlue = 60;
    blurrAlpha = textAlpha = 255
}
else if (inState == 2) {

    // pick (click down)
    hasBack = 1

    backRed = backGreen = backBlue = 255
    backAlpha = 90

    textRed = 214; textGreen = 234; textBlue = 160;
    blurrAlpha = textAlpha = 255
}

if (hasBack) {
	ctx.BeginPath()
	ctx.RoundedRect(x, y, w, h, 1.0)

	ctx.SetFillColor(vgoRGBA(backRed,backGreen,backBlue, backAlpha))
	ctx.Fill()
}

ctx.SetFontSize(36.0)
ctx.SetFontFace("sans")

ctx.SetTextAlign(vgoAlignLeft | vgoAlignMiddle)

var left = 20.0
var top = h / 2 - 1

ctx.SetFontBlur(1.0)
ctx.SetFillColor(vgoRGBA(0, 0, 0, blurrAlpha))
ctx.Text(left, top, text)

ctx.SetFontBlur(0.0)
ctx.SetFillColor(vgoRGBA(textRed,textGreen,textBlue, textAlpha))
ctx.Text(left, top, text)

ctx.EndFrame()