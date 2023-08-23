package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)



var GetCmd= &cobra.Command{
	Use: "get",
	Short: "ck8s get",
	Long: "ck8s cmd to get k8s resources $ck8s get pod",
	RunE: getfunc,
}

func getfunc(cmd *cobra.Command,args []string)error{

	fmt.Println("This is a Get cmd")
	return nil
}

func init(){
	RootCmd.AddCommand(GetCmd)
}
