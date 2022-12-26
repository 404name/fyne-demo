package global

import "fyne.io/fyne/v2/container"

func TabAppendAndSelect(view *container.TabItem) {

	for index, item := range Tabs.Items {
		if item.Text == view.Text {
			Tabs.SelectIndex(index)
			return
		}
	}
	Tabs.Append(view)
	Tabs.Select(view)
}
