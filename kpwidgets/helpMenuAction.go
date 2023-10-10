package kpwidgets

import (
	"github.com/skratchdot/open-golang/open"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

func InitHelpMenu(menuBar *widgets.QMenuBar, window *widgets.QMainWindow) {
	// Create the help menu
	helpMenu := menuBar.AddMenu2("Help")

	// Create the "About" action for the help menu
	helpAction := helpMenu.AddAction("Help Contents")
	//helpAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Alt+S", gui.QKeySequence__NativeText))
	helpAction.SetShortcut(gui.NewQKeySequence5(gui.QKeySequence__HelpContents))
	helpAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Toggle_Log.png"))

	// Connect the "About" action to its triggered event
	helpAction.ConnectTriggered(func(checked bool) {

		open.Run("D:\\Program Files (x86)\\KeePass Password Safe 2\\KeePass.chm")
		/*dialog := widgets.NewQDialog(window, 0)
		dialog.SetWindowTitle("QDialogButtonBox Example")

		// Create the button box
		buttonBox := widgets.NewQDialogButtonBox(dialog)
		okButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Ok)
		cancelButton := buttonBox.AddButton3(widgets.QDialogButtonBox__Cancel)

		// Connect the button box's accepted signal
		buttonBox.ConnectAccepted(func() {
			fmt.Println("OK button clicked")
			dialog.Accept()
		})

		// Connect the button box's rejected signal
		buttonBox.ConnectRejected(func() {
			fmt.Println("Cancel button clicked")
			dialog.Reject()
		})

		// Set the button text

		// Create the layout
		layout := widgets.NewQVBoxLayout2(dialog)
		layout.AddWidget(buttonBox, 0, core.Qt__AlignCenter)

		// Set the layout for the dialog
		dialog.SetLayout(layout)

		// Show the main window and the dialog

		dialog.Exec()
		*/

	})

	helpSourceAction := helpMenu.AddAction("Help Source")
	helpSourceAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_KOrganizer.png"))
	helpSourceAction.ConnectTriggered(func(checked bool) {

	})

	helpMenu.AddSeparator()
	webSiteAction := helpMenu.AddAction("KeePass Website")
	webSiteAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Folder_Home.png"))
	webSiteAction.ConnectTriggered(func(checked bool) {
		gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput("https://keepass.info/"))
	})

	donateAction := helpMenu.AddAction("Donate...")
	donateAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Identity.png"))
	donateAction.ConnectTriggered(func(checked bool) {
		gui.QDesktopServices_OpenUrl(core.QUrl_FromUserInput("https://keepass.info/donate.html"))
	})

	helpMenu.AddSeparator()
	checkForUpdatesAction := helpMenu.AddAction("Check for Updates")
	checkForUpdatesAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_FTP.png"))
	checkForUpdatesAction.ConnectTriggered(func(checked bool) {

	})

	helpMenu.AddSeparator()
	// Create the "About" action for the help menu
	aboutAction := helpMenu.AddAction("About KeePassGo")
	aboutAction.SetIcon(gui.NewQIcon5("Resources/Nuvola/B16x16_Help.png"))
	// Connect the "About" action to its triggered event
	aboutAction.ConnectTriggered(func(checked bool) {
		// 创建自定义对话框
		dialog := widgets.NewQDialog(nil, 0)
		dialog.SetWindowTitle("About")

		// 创建版权文本 QLabel
		copyrightLabel := widgets.NewQLabel2("Copyright (c) 2003-2017 Dominik Reil.", nil, 0)
		keePassLabel := widgets.NewQLabel2("KeePass is OSI Certified Open Source Software.", nil, 0)
		licenceLabel := widgets.NewQLabel2("The program is distributed under the terms of the GNU General Public Licence v2 or later.", nil, 0)

		// 创建 KeePass 官网链接 QLabel
		websiteLabel := widgets.NewQLabel2("<a href=\"https://keepass.info\">KeePass website</a>", nil, 0)
		websiteLabel.SetOpenExternalLinks(true)

		acknowledgementsLabel := widgets.NewQLabel2("<a href=\"https://keepass.info\">KeePass website</a>", nil, 0)
		acknowledgementsLabel.SetOpenExternalLinks(true)

		helpLabel := widgets.NewQLabel2("<a href=\"https://keepass.info/help/base/index.html\">Help</a>", nil, 0)
		helpLabel.SetOpenExternalLinks(true)

		licenseLabel := widgets.NewQLabel2("<a href=\"https://keepass.info/donate.html\">License</a>", nil, 0)
		licenseLabel.SetOpenExternalLinks(true)

		donateLabel := widgets.NewQLabel2("<a href=\"https://keepass.info/donate.html\">Donate</a>", nil, 0)
		donateLabel.SetOpenExternalLinks(true)

		// 创建 Component/Status/Version 表格
		table := widgets.NewQTableWidget2(2, 3, nil)
		table.SetHorizontalHeaderLabels([]string{"Component", "Status", "Version"})
		table.SetItem(0, 0, widgets.NewQTableWidgetItem2("KeePass", 0))
		table.SetItem(0, 1, widgets.NewQTableWidgetItem2("Status 1", 0))
		table.SetItem(0, 2, widgets.NewQTableWidgetItem2("2.37", 0))
		table.SetItem(1, 0, widgets.NewQTableWidgetItem2("keePassLicC", 0))
		table.SetItem(1, 1, widgets.NewQTableWidgetItem2("Status 2", 0))
		table.SetItem(1, 2, widgets.NewQTableWidgetItem2("1.34", 0))

		// 创建 OK 按钮
		okButton := widgets.NewQPushButton2("OK", nil)

		// 连接 OK 按钮的点击事件
		okButton.ConnectClicked(func(bool) {
			dialog.Close()
		})

		// 创建主布局
		layout := widgets.NewQVBoxLayout()

		// 添加部件到主布局
		layout.AddWidget(copyrightLabel, 0, core.Qt__AlignLeft)
		layout.AddSpacing(10)
		layout.AddWidget(keePassLabel, 0, core.Qt__AlignLeft)
		layout.AddSpacing(10)
		layout.AddWidget(licenceLabel, 0, core.Qt__AlignLeft)

		layout.AddSpacing(10)
		//layout.AddWidget(websiteLabel, 0, core.Qt__AlignCenter)
		//layout.AddSpacing(10)
		gridLayout2 := widgets.NewQGridLayout(nil)
		gridLayout2.AddWidget2(websiteLabel, 0, 0, core.Qt__AlignLeft)
		gridLayout2.AddWidget2(acknowledgementsLabel, 0, 1, core.Qt__AlignLeft)
		gridLayout2.AddWidget2(helpLabel, 0, 2, core.Qt__AlignLeft)
		gridLayout2.AddWidget2(licenseLabel, 1, 0, core.Qt__AlignLeft)
		gridLayout2.AddWidget2(donateLabel, 1, 1, core.Qt__AlignLeft)
		layout.AddLayout(gridLayout2, 0)

		layout.AddWidget(table, 0, 0)

		// 创建网格布局
		/*gridLayout := widgets.NewQGridLayout(nil)

		// 门户网站连接列表
		websites := []string{
			"https://www.google.com",
			"https://www.github.com",
			"https://www.microsoft.com",
			"https://www.apple.com",
			"https://www.openai.com",
			"https://www.spotify.com",
			"https://www.amazon.com",
			"https://www.facebook.com",
			"https://www.twitter.com",
		}

		// 创建随机数生成器
		rand.Seed(time.Now().UnixNano())

		// 在网格布局中随机放置门户网站连接
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				// 从门户网站连接列表中随机选择一个连接
				randomWebsite := websites[rand.Intn(len(websites))]

				// 创建 QLabel 和 QUrl
				label := widgets.NewQLabel(nil, 0)
				url := core.NewQUrl3(randomWebsite, core.QUrl__TolerantMode)

				// 设置 QLabel 的文本和打开外部链接
				label.SetText(randomWebsite)
				label.SetOpenExternalLinks(true)

				// 将 QLabel 添加到网格布局的指定位置
				gridLayout.AddWidget2(label, i, j, core.Qt__AlignLeft)

				// 释放 QUrl
				url.DestroyQUrl()
			}
		}*/

		// 创建底部布局
		bottomLayout := widgets.NewQHBoxLayout2(nil)
		bottomLayout.AddSpacing(10)
		bottomLayout.AddStretch(1)
		bottomLayout.AddWidget(okButton, 0, core.Qt__AlignRight|core.Qt__AlignBottom)

		// 添加底部布局到主布局

		layout.AddSpacing(10)
		layout.AddLayout(bottomLayout, 0)
		layout.AddSpacing(10)
		//layout.AddLayout(gridLayout, 0)
		// 设置主布局为对话框的布局
		dialog.SetLayout(layout)

		// 显示对话框
		dialog.Exec()
	})

}
