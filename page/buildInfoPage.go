package page

import (
	"bytes"
	"runtime/debug"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/v2/util/gutil"
	octicons "github.com/lusingander/fyne-octicons"
)

func BuildInfoTabItem() *container.TabItem {
	return container.NewTabItemWithIcon("编译信息",
		octicons.NoteIcon(),
		BuildInfoPage())
}

func BuildInfoPage() *widget.Entry {
	info, ok := debug.ReadBuildInfo()
	txt := widget.NewMultiLineEntry()
	if ok {
		bb := bytes.NewBuffer(nil)
		gutil.DumpTo(bb, info, gutil.DumpOption{})
		txt.SetText(bb.String())
		bb.Reset()
		// txt.SetText(gconv.String(info))
	}
	return txt
}
