package global

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/glog"
)

var (
	Name       string // 当前数据库名称
	Db         gdb.DB
	CurrentTbl string
	// Tbls       []string
	Tbls *garray.StrArray // 当前数据库的所有表名
	// Tblsm     map[string]*schemas.Table
	Ctx       context.Context
	BdNumTbls binding.Int
	DbPath    binding.String
	App       fyne.App
	Window    fyne.Window
	DataList  *widget.List
	Msg       *widget.Label
	Tree      *widget.Tree
	TreeData  map[string][]string
	Tabs      *container.DocTabs
	Err       error
	Glog      = glog.New()
)
