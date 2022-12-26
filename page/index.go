package page

import (
	"fyne.io/fyne/v2/widget"
)

func Index() *widget.Entry {
	indexView := widget.NewMultiLineEntry()
	indexView.SetText("这是一个首页0.0")
	return indexView
}
