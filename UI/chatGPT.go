package UI

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// chatGPTScreen loads chatGPT screen & do something
func chatGPTScreen(w fyne.Window) fyne.CanvasObject {
	// 创建一个垂直布局
	outPutTab := "	"
	testText := "近年来，随着科技的发展和社会的进步，人们的生活方式和工作方式发生了很大的变化。现代人的生活节奏越来越快，人们对于时间的利用也越来越高效。同时，互联网的普及也使得人们的信息获取变得更加便捷和快速。在这样的背景下，人们对于生活质量和工作效率的要求也越来越高。在这个时代，我们需要不断地学习和提升自己的能力，以适应社会的发展和变化。只有不断地学习和更新自己的知识和技能，才能在激烈的竞争中获得更好的发展机会。同时，我们也需要注重身体健康，保持良好的生活习惯和健康的饮食。只有身体健康，才能更好地投入到工作和学习中，提高自己的工作效率和学习成绩。总之，现代社会对于人们的要求越来越高，我们需要不断地提高自己的能力和素质，以适应社会的发展和变化。同时，我们也需要注重身体健康和良好的生活习惯，以保持良好的心态和精神状态。只有这样，我们才能在这个时代中获得更好的发展和生活。"
	output := widget.NewMultiLineEntry()
	output.Wrapping = fyne.TextWrapBreak //自动换行
	output.TextStyle = fyne.TextStyle{
		Bold: true,
	}
	output.Text = outPutTab + testText
	input := widget.NewMultiLineEntry()
	input.Wrapping = fyne.TextWrapBreak //自动换行

	//TODO 通知实现
	//TODO 进度条
	// button := widget.NewButton("提交", func() {
	// 	fyne.CurrentApp().SendNotification(&fyne.Notification{
	// 		Title:   "Title",
	// 		Content: "Content",
	// 	})
	// })
	button := widget.NewButton("Send", func() {
		log.Println("tapped")
	})

	content := fyne.NewContainerWithLayout(
		layout.NewVBoxLayout(),
		output,
		input,
		button,
	)

	return content
}

// text1 := canvas.NewText("1", color.White)
// text2 := canvas.NewText("2", color.White)
// text3 := canvas.NewText("3", color.White)
// grid := fyne.NewContainerWithLayout(layout.NewFixedGridLayout(fyne.NewSize(50, 50)),
// 	text1, text2, text3)
