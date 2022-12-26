package component

import (
	"bytes"

	"fyne.io/fyne/v2/widget"
	"github.com/404name/fyne-demo/global"
)

var LogEntry *widget.Entry
var LogBuf *bytes.Buffer

type MyWriter struct {
	Msg *widget.Entry
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	LogBuf.Write(p)
	w.Msg.SetText(LogBuf.String())
	return len(p), nil
}
func SetGlog() {
	LogEntry = widget.NewMultiLineEntry()
	LogBuf = bytes.NewBuffer(nil)
	global.Glog.SetWriter(&MyWriter{Msg: LogEntry})
}
