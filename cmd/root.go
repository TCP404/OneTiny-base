package cmd

//  onetiny
//  ├── -r
//  ├── -p
//  └── -a

import (
	"fmt"
	"io"
	"os"

	"github.com/TCP404/OneTiny-base/args"
	"github.com/urfave/cli/v2"
)

func initCLI() {
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "打印版本信息",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println("当前版本: ", c.App.Version)
		os.Exit(0)
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "打印帮助信息",
	}
	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		cli.HelpPrinterCustom(w, templ, data, nil)
		os.Exit(0)
	}
	cli.ErrWriter = args.Output
	args.RootPath, _ = os.Getwd()
}

// RunCLI 函数作为程序入口，主要负责处理命令和 flag
func CLI() *cli.App {
	initCLI()

	return &cli.App{
		Name:            "OneTiny",
		Usage:           "一个用于局域网内共享文件的FTP程序",
		UsageText:       "onetiny [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS] [参数...]",
		Version:         "v1.0.0",
		Flags:           newGlobalFlag(),
		Authors:         []*cli.Author{{Name: "Boii", Email: "i@tcp404.com"}},
		CommandNotFound: func(c *cli.Context, s string) { cli.ShowAppHelpAndExit(c, 10) },
		Writer:          args.Output,
		ErrWriter:       args.Output,
		Action: func(c *cli.Context) error {
			args.Port = c.Int("port")
			args.RootPath = c.String("road")
			args.IsAllowUpload = c.Bool("allow")
			return nil
		},
	}
}
