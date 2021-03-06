package main

import (
	"log"
	"os"

	"github.com/appilon/tfplugin/cmd/docs"
	"github.com/appilon/tfplugin/cmd/schema"
	"github.com/appilon/tfplugin/cmd/status"
	"github.com/appilon/tfplugin/cmd/upgrade/golang"
	"github.com/appilon/tfplugin/cmd/upgrade/modules"
	"github.com/appilon/tfplugin/cmd/upgrade/pr"
	"github.com/appilon/tfplugin/cmd/upgrade/sdk"
	"github.com/mitchellh/cli"
)

func main() {
	log.SetFlags(log.Llongfile)
	c := cli.NewCLI("tfplugin", "0.2.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		schema.CommandName:  schema.CommandFactory,
		docs.CommandName:    docs.CommandFactory,
		golang.CommandName:  golang.CommandFactory,
		sdk.CommandName:     sdk.CommandFactory,
		modules.CommandName: modules.CommandFactory,
		pr.CommandName:      pr.CommandFactory,
		status.CommandName:  status.CommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
