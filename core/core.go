package core

import (
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/TCP404/OneTiny-base/args"
	"github.com/TCP404/OneTiny-base/core/middleware"
	"github.com/TCP404/OneTiny-base/core/routes"
	"github.com/TCP404/OneTiny-base/util"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// RunCore 函数负责启动 gin 实例，开始提供 HTTP 服务
func RunCore() error {
	r := StartUpGin()

	printInfo()

	err := r.Run(":" + strconv.Itoa(args.Port))
	if _, ok := err.(*net.OpError); ok {
		return errors.New(color.RedString("指定的 %d 端口已被占用，请换一个端口号。", args.Port))
	}
	return err
}

func StartUpGin() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	middleware.Setup(r)
	routes.Setup(r)
	return r
}

// printInfo 会在程序启动后打印本机 IP、共享目录、是否允许上传的信息
func printInfo() {
	log.SetOutput(color.Output)
	ip, err := util.GetIP()
	args.IP = ip
	// Print IP infomation
	if args.IP == "" || err != nil {
		log.Printf("%s", color.YellowString("Warning: [ 暂时获取不到您的IP，可以打开新的命令行窗口输入 ->  ipconfig , 查看您的IP。]"))
	} else {
		log.Printf("Run on   [ %s ]", color.BlueString("http://%s:%d", args.IP, args.Port))
	}

	// Print RootPath infomation
	log.Printf("Run with [ %s ]", color.BlueString("%s", args.RootPath))

	// Print Max allow access level
	log.Printf("Allow access level: [ %s ]", color.BlueString("%d", args.MaxLevel))

	// Print Allow upload Status
	status := color.RedString("%t", args.IsAllowUpload)
	if args.IsAllowUpload {
		status = color.GreenString("%t", args.IsAllowUpload)
	}
	log.Printf("Allow upload: [ %s ]", status)
}
