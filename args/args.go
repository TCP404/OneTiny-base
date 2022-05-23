package args

import (
	"io"
	"os"
	"runtime"

	"github.com/fatih/color"
)

var (
	Output io.Writer = If(runtime.GOOS == "windows", color.Output, os.Stderr).(io.Writer)
	OS     string    = runtime.GOOS // 程序所在的操作系统，默认值 linux
	IP     string    = "127.0.0.1"  // 本机局域网IP
	Pwd    string    = "/"

	// 各个参数的原厂默认值
	RootPath      string = "/"   // 共享目录的根路径，默认值：当前目录
	MaxLevel      uint8  = 0     // 允许访问的最大层级，默认值  0
	Port          int    = 9090  // 指定的服务端口，默认值 9090
	IsAllowUpload bool   = false // 是否允许上传，默认值：否
)

func If(cond bool, t any, f any) any {
	if cond {
		return t
	}
	return f
}
