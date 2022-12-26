package component

import (
	"context"
	"embed"
	"io/fs"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/404name/fyne-demo/global"
	"github.com/404name/fyne-demo/page"
	"github.com/leaanthony/debme"
)

func Srcui(srcs embed.FS) {
	// w := fyne.CurrentApp().NewWindow("sources")
	// w.Resize(fyne.NewSize(900, 600))
	// w.SetContent(widget.NewLabel("this is a window with a label"))
	// w.CenterOnScreen()
	data := make(map[string][]string)
	root, _ := debme.FS(srcs, ".")
	// if txt, err := root.ReadFile("main.go"); err != nil {
	// 	println("err", err)
	// } else {
	// 	println(string(txt))
	// }

	// Fully compatible FS
	fs.WalkDir(root, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		dn := strings.Split(path, "/")
		switch len(dn) {
		case 1:
			data[""] = append(data[""], d.Name())
		case 2:
			data[dn[0]] = append(data[dn[0]], path)
		case 3:
			data[dn[1]] = append(data[dn[1]], path)
		case 4:
			data[dn[2]] = append(data[dn[2]], path)
		}
		// if d.IsDir() || path == d.Name() {
		// 	data[""] = append(data[""], d.Name())
		// } else {
		// 	dn := strings.Split(path, "/")
		// 	if len(dn) == 2 {
		// 		data[dn[0]] = append(data[dn[0]], path)
		// 	}
		// }
		return nil
	})
	//content := widget.NewMultiLineEntry()
	global.TreeData = data
	global.Tree = widget.NewTreeWithStrings(global.TreeData)
	global.Tree.Refresh()

	global.Tree.UpdateNode = func(uid widget.TreeNodeID, branch bool, node fyne.CanvasObject) {
		lbl := node.(*widget.Label)
		dn := strings.Split(uid, "/")
		switch len(dn) {
		case 1:
			lbl.SetText(uid)
		case 2:
			lbl.SetText(dn[1])
		case 3:
			lbl.SetText(dn[2])
		case 4:
			lbl.SetText(dn[3])
		}
		// if branch {
		// 	lbl.SetText(uid)
		// } else {
		// 	dn := strings.Split(uid, "/")
		// 	if len(dn) == 2 {
		// 		lbl.SetText(dn[1])
		// 	} else {
		// 		lbl.SetText(dn[0])
		// 	}
		// }
	}
	global.Tree.OnSelected = func(uid widget.TreeNodeID) {
		if len(global.TreeData[uid]) != 0 {
			// 文件夹不打开
			return
		}
		global.Msg.SetText(uid)
		var str []byte
		ctx := context.Background()
		done := make(chan struct{}, 1)
		global.Glog.Info(context.Background(), "进入页面"+uid)
		go func(ctx context.Context) {
			// 超时任务，如 发送HTTP请求
			str, _ = root.ReadFile(uid)
			//content.SetText(string(str)) //碰到大文件会阻塞界面
			global.TabAppendAndSelect(page.DefaultStringTabItem(uid, string(str)))
			done <- struct{}{}
		}(ctx)
		select {
		case <-done:
			// fmt.Println("call successfully!!!")
			return
		case <-time.After(time.Duration(800 * time.Millisecond)):
			// fmt.Println("timeout!!!")
			//content.SetText("the file is too big")
			global.TabAppendAndSelect(page.DefaultStringTabItem(uid, "the file is too big"))
			return
		}

		// str, _ := root.ReadFile(uid)
		// content.SetText(string(str))
	}
	// cc := container.NewBorder(nil, msg, tree, nil, content)
	// w.SetContent(cc)
	// w.Show()
}
