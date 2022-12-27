package main

import (
	"context"
	"embed"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/404name/fyne-demo/component"
	"github.com/404name/fyne-demo/global"
	"github.com/404name/fyne-demo/page"
	themee "github.com/404name/fyne-demo/theme"
	octicons "github.com/lusingander/fyne-octicons"
)

//go:embed component global page theme main.go go.mod go.sum
var srcs embed.FS

func main() {
	global.App = app.NewWithID("io.fyne.demo")
	global.App.NewWindow("Database Manager")
	global.App.Settings().SetTheme(&themee.MyTheme{})
	global.App.SetIcon(octicons.DatabaseIcon())

	component.SetGlog()
	//global.A.Tblsm = make(map[string]*schemas.Table)
	// bdddl = binding.NewString()

	global.Msg = widget.NewLabel("")
	component.Srcui(srcs)
	// menu := make(map[string][]string)
	// menu[""] = []string{"首页", "功能列表", "系统"}
	// menu["功能列表"] = []string{"测试1", "功能2", "功能3"}
	// menu["系统"] = []string{"设置", "关于", "检查更新", "编译信息"}
	// global.Tree = widget.NewTreeWithStrings(menu)
	// global.Tree.Refresh()
	// global.Tree.OnSelected = func(uid widget.TreeNodeID) {
	// 	//v, _ := garray.NewStrArrayFrom(strings.Split(uid, "/")).PopRight()
	// 	if uid == "编译信息" {
	// 		// 这边应该是一个渲染组件
	// 		global.TabAppendAndSelect(page.BuildInfoTabItem())
	// 	} else if len(menu[uid]) == 0 {
	// 		global.Msg.SetText("进入页面" + uid)
	// 		global.Glog.Info(context.Background(), "进入页面"+uid)
	// 		global.TabAppendAndSelect(page.DefaultStringTabItem(uid, uid))
	// 	} else {
	// 		global.Msg.SetText("你点击一级目录" + uid)
	// 		global.Glog.Info(context.Background(), "你点击一级目录"+uid)
	// 	}
	// }

	global.Tabs = container.NewDocTabs()
	IndexView := container.NewTabItemWithIcon("首页",
		octicons.NoteIcon(),
		page.Index())
	global.Tabs.Append(IndexView)
	global.Tabs.Select(IndexView)
	cc := container.NewHSplit(global.Tree, container.NewBorder(nil, component.LogEntry, nil, nil, global.Tabs))
	cc.Offset = 0.2
	tb := widget.NewToolbar(
		widget.NewToolbarAction(theme.GridIcon(), func() {

			component.Srcui(srcs)
			global.Msg.SetText("切換文件预览模式")
			global.Glog.Info(context.Background(), "切換文件预览模式")
		}),
		widget.NewToolbarAction(theme.ContentAddIcon(), func() {
			global.Msg.SetText("添加tab")
			global.Glog.Info(context.Background(), "添加tab")
			global.TabAppendAndSelect(page.DefaultTabItem(rand.Intn(100)))
		}),
	)

	btn := widget.NewButton("关闭全部tab", func() {
		global.Msg.SetText("你点击关闭全部tab")
		global.Glog.Info(context.Background(), "你点击Connetct")
		global.Tabs.Items = nil
		global.Tabs.Refresh()
	})
	top := container.NewVBox(tb,
		btn,
	)

	c := container.NewBorder(top, global.Msg, nil, nil, cc)
	global.Window = fyne.CurrentApp().NewWindow("")
	global.Window.SetContent(c)
	global.Window.SetMainMenu(component.MainMenu(srcs, global.App))
	global.Window.Resize(fyne.NewSize(900, 700))
	global.Window.SetMaster()
	global.Window.SetCloseIntercept(func() {
		dialog.ShowConfirm("Closing...", "Are you sure to close?",
			func(b bool) {
				if b {
					global.Window.Close()
				}
			}, global.Window)
	})
	global.Window.ShowAndRun()
	defer func() {
		// if newDb != nil {
		// 	newDb.Close(newDb.GetCtx())
		// }
	}()
}
