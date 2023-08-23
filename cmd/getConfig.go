package cmd

import (
	

	"github.com/spf13/cobra"
	"github.com/riteshsonawane/ck8s/pkg"
)

var getConfig = &cobra.Command{
	Use:   "config",
	Short: "Get information abt Config file",
	Long:  `Usage is : $ck8s get config`,
	RunE:  GetConfig,
}

func GetConfig(cmd *cobra.Command, args []string) error {
	getpod()
	return nil
}

func getpod(){
	pkg.ConfigFile()
}

func init() {
	GetCmd.AddCommand(getConfig)
}
