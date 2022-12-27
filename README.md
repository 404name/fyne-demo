# 预览
![image](https://user-images.githubusercontent.com/56631419/209564097-e7c82c22-7a1a-44a3-81b7-09b4d0949173.png)

# 启动项目
1. 克隆到本地
2. go mod tidy
3. go run . (必须文件夹下的一起编译)

# 进度

## TODO
1， 文件夹排序 & 文件
2。 左边第一个menu去了

log控件保证滚动条是最新

## 初始化项目
1. 创建UI组件
2. 创建layout(侧边栏，顶部，tab视图，view界面，底部信息栏)
3. 创建log组件并且接入系统日志输出
4. 读取系统配置
5. 读取main文件路径下的所有文件，并渲染到界面

```
读取使用embed项目，使用指令go:embed，指定目录下文件及目录读取
//go:embed component global page theme main.go go.mod go.sum
var srcs embed.FS
```
