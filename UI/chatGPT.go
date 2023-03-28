package UI

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

// chatGPTScreen loads chatGPT model & do something
func chatGPTScreen(_ fyne.Window) fyne.CanvasObject {
	return container.NewGridWrap(fyne.NewSize(90, 90),
		//canvas.NewImageFromResource(theme.FyneLogo()),
		canvas.NewText("Text", color.NRGBA{0, 0x80, 0, 0xff}),
	)
}
