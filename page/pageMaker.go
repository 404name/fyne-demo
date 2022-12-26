package page

import (
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	octicons "github.com/lusingander/fyne-octicons"
)

func DefaultTabItem(randomNum int) *container.TabItem {
	return container.NewTabItemWithIcon("Tab"+strconv.Itoa(randomNum),
		octicons.NoteIcon(),
		DefaultPage(randomNum))
}

func DefaultPage(num int) *widget.Entry {
	pageContent := widget.NewMultiLineEntry()
	pageContent.SetText("这是一个页面" + strconv.Itoa(num))
	return pageContent
}

func DefaultStringTabItem(name string, context string) *container.TabItem {
	return container.NewTabItemWithIcon(name,
		octicons.NoteIcon(),
		DefaultStringPage(context))
}

func DefaultStringPage(context string) *widget.Entry {
	pageContent := widget.NewMultiLineEntry()
	pageContent.SetText(context)
	return pageContent
}
