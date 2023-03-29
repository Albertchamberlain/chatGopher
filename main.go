// Package main provides various examples of Fyne API capabilities.
package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"chaGopher/DB"
	UI "chaGopher/UI"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/cmd/fyne_settings/settings"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const preferenceCurrentTutorial = "currentTutorial"

var topWindow fyne.Window

func main() {
	os.Setenv("FYNE_FONT", "./Font/LXGWWENKAIGBSCREENR.TTF") //加载字体
	a := app.NewWithID("github.albertchamberlain.chatGopher")
	a.SetIcon(theme.FyneLogo())
	makeTray(a)
	logLifecycle(a)
	w := a.NewWindow("chatGopher")
	topWindow = w

	w.SetMainMenu(makeMenu(a, w))
	w.SetMaster()

	content := container.NewMax()
	title := widget.NewLabel("Component name")
	intro := widget.NewLabel("An introduction would probably go\nhere, as well as a")
	intro.Wrapping = fyne.TextWrapWord
	setUI := func(t UI.AiModel) {
		if fyne.CurrentDevice().IsMobile() {
			child := a.NewWindow(t.Title)
			topWindow = child
			child.SetContent(t.View(topWindow))
			child.Show()
			child.SetOnClosed(func() {
				topWindow = w
			})
			return
		}

		title.SetText(t.Title)
		intro.SetText(t.Intro)

		content.Objects = []fyne.CanvasObject{t.View(w)}
		content.Refresh()
	}

	UI := container.NewBorder(
		container.NewVBox(title, widget.NewSeparator(), intro), nil, nil, nil, content)
	if fyne.CurrentDevice().IsMobile() {
		w.SetContent(makeNav(setUI, false))
	} else {
		split := container.NewHSplit(makeNav(setUI, true), UI)
		split.Offset = 0.2
		w.SetContent(split)
	}
	w.Resize(fyne.NewSize(900, 900))
	w.ShowAndRun()
}

//打印窗口生命周期
func logLifecycle(a fyne.App) {
	a.Lifecycle().SetOnStarted(func() {
		log.Println("Lifecycle: Started")
	})
	a.Lifecycle().SetOnStopped(func() {
		log.Println("Lifecycle: Stopped")
	})
	a.Lifecycle().SetOnEnteredForeground(func() {
		log.Println("Lifecycle: Entered Foreground")
	})
	a.Lifecycle().SetOnExitedForeground(func() {
		log.Println("Lifecycle: Exited Foreground")
	})
}

//菜单栏
func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	newItem := fyne.NewMenuItem("Add OpenAI Key", nil) //添加open ai key

	checkedNotificationItem := fyne.NewMenuItem("Open Notifications", nil)
	checkedNotificationItem.Checked = true

	checkedAutoCopyItem := fyne.NewMenuItem("Auto Copy", nil)
	checkedAutoCopyItem.Checked = true

	// disabledItem := fyne.NewMenuItem("Disabled", nil)
	// disabledItem.Disabled = true

	// ListKeyItem := fyne.NewMenuItem("", nil)

	// mailItem := fyne.NewMenuItem("Mail", func() { fmt.Println("Menu New->Other->Mail") })
	// mailItem.Icon = theme.MailComposeIcon()

	// otherItem.ChildMenu = fyne.NewMenu("",
	// 	fyne.NewMenuItem("Project", func() { fmt.Println("Menu New->Other->Project") }),
	// 	mailItem,
	// )
	addGPT3KeyWindow := func() {
		w := a.NewWindow("Add GPT-3 Key") //另开一个窗口
		inputkeyEntry := widget.NewPasswordEntry()
		inputkeyEntry.SetPlaceHolder("Enter GPT-3 Key")
		inputkeyEntryItem := container.NewVBox(inputkeyEntry, widget.NewButton("Save", func() {
			fmt.Println(inputkeyEntry.Text)
			DB.SetKeyAndValue("gpt3", []byte(inputkeyEntry.Text)) //插入到DB中
		}))
		w.SetContent(inputkeyEntryItem)
		w.Resize(fyne.NewSize(400, 100))
		w.Show()
	}
	addGPT3KeyItem := fyne.NewMenuItem("GPT-3", addGPT3KeyWindow)
	addGPT3KeyItem.Icon = theme.AccountIcon()

	addGPT4KeyWindow := func() {
		w := a.NewWindow("Add GPT-4 Key") //另开一个窗口
		inputkeyEntry := widget.NewPasswordEntry()
		inputkeyEntry.SetPlaceHolder("Enter GPT-4 Key")
		inputkeyEntryItem := container.NewVBox(inputkeyEntry, widget.NewButton("Save", func() {
			fmt.Println(inputkeyEntry.Text)
			DB.SetKeyAndValue("gpt4", []byte(inputkeyEntry.Text)) //插入到DB中
		}))
		w.SetContent(inputkeyEntryItem)
		w.Resize(fyne.NewSize(400, 100))
		w.Show()
	}
	addGPT4KeyItem := fyne.NewMenuItem("GPT-4", addGPT4KeyWindow)
	addGPT4KeyItem.Icon = theme.AccountIcon()

	ListAllKeyWindow := func() {
		w := a.NewWindow("List All Key") //另开一个窗口
		kvs := DB.ListAllApi()
		fmt.Println(kvs)
		var counter float32
		counter = 1
		outputkeyEntryItem := container.NewVBox()
		for k, v := range kvs {
			outputkeyEntry := widget.NewPasswordEntry()
			kvitems := k + ":" + v
			outputkeyEntry.SetText(kvitems)
			fmt.Println(outputkeyEntry.Text)
			outputkeyEntryItem.Add(outputkeyEntry)
			w.Resize(fyne.NewSize(400, 100*counter))
			counter++
		}
		w.SetContent(outputkeyEntryItem)
		//w.SetContent(outputkeyEntryItem)
		//w.Resize(fyne.NewSize(400, 100))
		w.Show()
	}
	listAllKeyItem := fyne.NewMenuItem("List All Key", ListAllKeyWindow)
	listAllKeyItem.Icon = theme.ComputerIcon()

	// gpt4keyitem := fyne.NewMenuItem("gpt4", func() { fmt.Println("Menu New->File") })
	// gpt4keyitem.Icon = theme.AccountIcon()

	// dirItem := fyne.NewMenuItem("Directory", func() { fmt.Println("Menu New->Directory") })
	// dirItem.Icon = theme.FolderIcon()
	newItem.ChildMenu = fyne.NewMenu("",
		addGPT3KeyItem,
		addGPT4KeyItem,
		listAllKeyItem,
	)

	openSettings := func() {
		w := a.NewWindow("Appearance Settings") //另开一个窗口
		w.SetContent(settings.NewSettings().LoadAppearanceScreen(w))
		w.Resize(fyne.NewSize(480, 480))
		w.Show()
	}
	settingsItem := fyne.NewMenuItem("Appearance", openSettings)
	settingsShortcut := &desktop.CustomShortcut{KeyName: fyne.KeyComma, Modifier: fyne.KeyModifierShortcutDefault}
	settingsItem.Shortcut = settingsShortcut
	w.Canvas().AddShortcut(settingsShortcut, func(shortcut fyne.Shortcut) {
		openSettings()
	})

	cutShortcut := &fyne.ShortcutCut{Clipboard: w.Clipboard()}
	cutItem := fyne.NewMenuItem("Cut", func() {
		shortcutFocused(cutShortcut, w)
	})
	cutItem.Shortcut = cutShortcut
	copyShortcut := &fyne.ShortcutCopy{Clipboard: w.Clipboard()}
	copyItem := fyne.NewMenuItem("Copy", func() {
		shortcutFocused(copyShortcut, w)
	})
	copyItem.Shortcut = copyShortcut

	pasteShortcut := &fyne.ShortcutPaste{Clipboard: w.Clipboard()}
	pasteItem := fyne.NewMenuItem("Paste", func() {
		shortcutFocused(pasteShortcut, w)
	})
	pasteItem.Shortcut = pasteShortcut
	performFind := func() { fmt.Println("Menu Find") }

	findItem := fyne.NewMenuItem("Find", performFind)
	findItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyF, Modifier: fyne.KeyModifierShortcutDefault | fyne.KeyModifierAlt | fyne.KeyModifierShift | fyne.KeyModifierControl | fyne.KeyModifierSuper}
	w.Canvas().AddShortcut(findItem.Shortcut, func(shortcut fyne.Shortcut) {
		performFind()
	})

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("Documentation", func() {
			u, _ := url.Parse("https://github.com/Albertchamberlain/chatGopher")
			a.OpenURL(u)
		}),
		fyne.NewMenuItem("Support", func() {
			u, _ := url.Parse("https://github.com/Albertchamberlain/chatGopher")
			a.OpenURL(u)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Sponsor", func() {
			u, _ := url.Parse("https://github.com/Albertchamberlain/chatGopher")
			a.OpenURL(u)
		}))

	// a quit item will be appended to our first (File) menu
	file := fyne.NewMenu("setting", newItem, checkedAutoCopyItem, checkedNotificationItem)
	device := fyne.CurrentDevice()
	if !device.IsMobile() && !device.IsBrowser() { //PC
		file.Items = append(file.Items, fyne.NewMenuItemSeparator(), settingsItem)
	}

	mainMean := fyne.NewMainMenu(
		file,
		fyne.NewMenu("Edit", cutItem, copyItem, pasteItem, fyne.NewMenuItemSeparator(), findItem),
		helpMenu,
	)

	checkedNotificationItem.Action = func() {
		checkedNotificationItem.Checked = !checkedNotificationItem.Checked
		mainMean.Refresh()
		fmt.Println(checkedAutoCopyItem.Checked)
	}
	checkedAutoCopyItem.Action = func() {
		checkedAutoCopyItem.Checked = !checkedAutoCopyItem.Checked
		mainMean.Refresh()
		fmt.Println(checkedAutoCopyItem.Checked)
	}

	return mainMean
}

func makeTray(a fyne.App) {
	if desk, ok := a.(desktop.App); ok {
		h := fyne.NewMenuItem("Hello", func() {})
		h.Icon = theme.HomeIcon()
		menu := fyne.NewMenu("Hello World", h)
		h.Action = func() {
			log.Println("System tray menu tapped")
			h.Label = "Welcome"
			menu.Refresh()
		}
		desk.SetSystemTrayMenu(menu)
	}
}

func unsupportedUI(t UI.AiModel) bool {
	return !t.SupportWeb && fyne.CurrentDevice().IsBrowser()
}

//侧面导航栏
func makeNav(setModelUI func(modelUi UI.AiModel), loadPrevious bool) fyne.CanvasObject {
	a := fyne.CurrentApp()

	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			return UI.AiModelsIndex[uid]
		},
		IsBranch: func(uid string) bool {
			children, ok := UI.AiModelsIndex[uid]
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			t, ok := UI.AiModels[uid]
			if !ok {
				fyne.LogError("Missing ui panel: "+uid, nil)
				return
			}
			obj.(*widget.Label).SetText(t.Title)
			if unsupportedUI(t) {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{Italic: true}
			} else {
				obj.(*widget.Label).TextStyle = fyne.TextStyle{}
			}
		},
		OnSelected: func(uid string) {
			if t, ok := UI.AiModels[uid]; ok {
				if unsupportedUI(t) {
					return
				}
				a.Preferences().SetString(preferenceCurrentTutorial, uid)
				setModelUI(t)
			}
		},
	}

	if loadPrevious {
		currentPref := a.Preferences().StringWithFallback(preferenceCurrentTutorial, "welcome")
		tree.Select(currentPref)
	}

	themes := container.NewGridWithColumns(2,
		widget.NewButton("Light", func() {
			a.Settings().SetTheme(theme.LightTheme())
		}),
		widget.NewButton("Dark", func() {
			a.Settings().SetTheme(theme.DarkTheme())
		}),
	)

	return container.NewBorder(nil, themes, nil, nil, tree)
}

//剪切板相关
func shortcutFocused(s fyne.Shortcut, w fyne.Window) {
	switch sh := s.(type) {
	case *fyne.ShortcutCopy:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutCut:
		sh.Clipboard = w.Clipboard()
	case *fyne.ShortcutPaste:
		sh.Clipboard = w.Clipboard()
	}
	if focused, ok := w.Canvas().Focused().(fyne.Shortcutable); ok {
		focused.TypedShortcut(s)
	}
}
