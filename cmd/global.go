package cmd

import (
	"github.com/TCP404/OneTiny-base/args"
	"github.com/urfave/cli/v2"
)

func newGlobalFlag() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name:        "road",
			Aliases:     []string{"r"},
			Usage:       "指定对外开放的目录`路径`",
			Value:       args.RootPath,
			Required:    false,
			DefaultText: "/",
		},
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			Usage:       "指定开放的`端口`",
			Value:       args.Port,
			Required:    false,
			DefaultText: "9090",
		},
		&cli.BoolFlag{
			Name:        "allow",
			Aliases:     []string{"a"},
			Usage:       "指定`是否`允许访问者上传",
			Value:       args.IsAllowUpload,
			Required:    false,
			DefaultText: "false",
		},
	}
}
