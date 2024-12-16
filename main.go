package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	amiCopy "github.com/bdwyertech/packer-plugin-aws-ami/post-processor/ami-copy"
	amiDelete "github.com/bdwyertech/packer-plugin-aws-ami/post-processor/ami-delete"
	"github.com/bdwyertech/packer-plugin-aws-ami/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(amiCopy.PostProcessor))
	pps.RegisterPostProcessor(plugin.DEFAULT_NAME, new(amiDelete.PostProcessor))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
