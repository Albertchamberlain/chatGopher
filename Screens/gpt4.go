package UI

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// gpt4Screen loads a panel that shows the gpt4 model available in Fyne
func gpt4Screen(_ fyne.Window) fyne.CanvasObject {
	answer := widget.NewMultiLineEntry()
	answer.Wrapping = fyne.TextWrapBreak
	answer.SetPlaceHolder("准备接收回答ing")
	testText := ""

	answer.TextStyle = fyne.TextStyle{
		TabWidth: 2,
	}
	answer.Text = testText
	answer.Refresh()
	problemEntry := widget.NewMultiLineEntry()
	problemEntry.Text = testText
	problemEntry.Wrapping = fyne.TextWrapBreak
	problemEntry.SetPlaceHolder("请在此输入问题......")
	problemEntry.Refresh()
	//TODO 通知实现
	//TODO 进度条
	// button := widget.NewButton("提交", func() {
	// 	fyne.CurrentApp().SendNotification(&fyne.Notification{
	// 		Title:   "Title",
	// 		Content: "Content",
	// 	})
	// })
	button := widget.NewButton("发送", func() {
		fmt.Println("button tapped!") //发送问题
	})
	problemAndButton := container.NewVSplit(
		problemEntry,
		button,
	)
	problemAndButton.SetOffset(1)
	return container.NewVSplit(
		container.NewVScroll(answer),
		problemAndButton,
	)
}
