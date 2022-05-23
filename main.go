package main

import (
	"fmt"
	"os"

	"github.com/TCP404/OneTiny-base/args"
	"github.com/TCP404/OneTiny-base/cmd"
	"github.com/TCP404/OneTiny-base/include"

	"github.com/TCP404/OneTiny-core/gateway/vo"
	"github.com/fatih/color"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			fmt.Println(color.RedString("%v", err.Error()))
		}
	}()

	if err = cmd.CLI().Run(os.Args); err != nil {
		return
	}

	// coreClient := include.Core

	// if err = coreClient.NewVerifyChain().
	// 	AddToHead(coreClient.NewPortHandler(args.Port)).
	// 	AddToHead(coreClient.NewPathHandler(args.RootPath)).
	// 	Iterator(); err != nil {
	// 	return
	// }

	// if err = coreClient.New(vo.Release).
	// 	SetConfig(args.IP, args.Port, args.RootPath, vo.IsAllowUpload(args.IsAllowUpload), vo.MaxLevel(100)).
	// 	PrintInfo(vo.PAddr, vo.PPath, vo.PAllow).
	// 	Run(); err != nil {
	// 	return
	// }
}
