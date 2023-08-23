package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "ck8s",
	Short: "Use Ck8s",
	Long: "CLI application for getting k8s object",
}

func Execute(){
	err:= RootCmd.Execute()
	if err!=nil{
		os.Exit(1)
	}
}

