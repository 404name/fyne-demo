package component

import (
	"embed"

	"fyne.io/fyne/v2"
	"github.com/404name/fyne-demo/global"
	"github.com/404name/fyne-demo/page"
)

func MainMenu(srcs embed.FS, a fyne.App) *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("Tools", fyne.NewMenuItem(
			"Icons", func() {
				global.Msg.SetText("你点击了Tools")
			},
		), fyne.NewMenuItem(
			"Json2Struct", func() {
				global.Msg.SetText("你点击了Json2Struct")
			},
		),
			fyne.NewMenuItem(
				"Html2Txt", func() {
					w := a.NewWindow("Log")
					w.Resize(fyne.NewSize(800, 500))
					w.Show()
				},
			),
			fyne.NewMenuItem(
				"json filter", func() {
					global.Msg.SetText("你点击了json")
				},
			),
		),

		fyne.NewMenu("About",
			fyne.NewMenuItem(
				"Log", func() {
					wlog := a.NewWindow("Log")
					wlog.Resize(fyne.NewSize(800, 500))
					wlog.SetContent(LogEntry)
					wlog.Show()
				},
			), fyne.NewMenuItem(
				"Clear Log", func() {
					LogBuf.Reset()
					LogEntry.SetText("")
				},
			), fyne.NewMenuItem(
				"Sources", func() {
					Srcui(srcs)
				}), fyne.NewMenuItem(
				"Build Info", func() {
					global.TabAppendAndSelect(page.BuildInfoTabItem())
				}), fyne.NewMenuItem(
				"dasel", func() {
					global.Msg.SetText("你点击了dasel")
				}),
		))
}
