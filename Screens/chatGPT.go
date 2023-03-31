package Screens

import (
	"fmt"

	Database "chatGopher/Database"
	Models "chatGopher/Models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func chatGPTScreen(_ fyne.Window) fyne.CanvasObject {
	answer := widget.NewMultiLineEntry()
	answer.Wrapping = fyne.TextWrapBreak
	answer.SetPlaceHolder("Ready to answering")
	testText := ""

	answer.TextStyle = fyne.TextStyle{
		TabWidth: 2,
	}
	answer.Text = testText
	answer.Refresh()
	problemEntry := widget.NewMultiLineEntry()
	problemEntry.Text = testText
	problemEntry.Wrapping = fyne.TextWrapBreak
	problemEntry.SetPlaceHolder("PLZ input you problem here......")
	problemEntry.Refresh()
	//TODO 通知实现
	//TODO 进度条
	// button := widget.NewButton("提交", func() {
	// 	fyne.CurrentApp().SendNotification(&fyne.Notification{
	// 		Title:   "Title",
	// 		Content: "Content",
	// 	})
	// })

	//pre get
	chatGPTKey := Database.GetValueByKey("gpt3")
	fmt.Println(chatGPTKey)
	Models.NewChatGPT(chatGPTKey)
	button := widget.NewButton("Sent", func() {
		fmt.Println("button tapped!") //发送问题
		answer.Refresh()
		answer.Text = Models.ChatGPTNormalMode(problemEntry.Text)
		answer.Refresh()
		fmt.Println(problemEntry.Text)
		fmt.Println(answer.Text)
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
