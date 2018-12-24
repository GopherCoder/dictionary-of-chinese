package cmd

import "github.com/spf13/cobra"

/*
1. 使用命令，将数据导出成 csv,json,txt
*/

func init() {
	rootCmd.AddCommand(exportCmd)
}

var exportCmd = &cobra.Command{
	Use: "export",
	Run: RunFuncExport,
}

func RunFuncExport(cmd *cobra.Command, args []string) {}
