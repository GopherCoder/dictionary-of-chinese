package cmd

import "github.com/spf13/cobra"

/*
1. 将 hash 值转变成 set、zset、bitmap、hyperloglog，每种方法给出搜索时间
*/

func init() {
	rootCmd.AddCommand(convertCmd)
}

var convertCmd = &cobra.Command{
	Use: "convert",
	Run: RunFuncConvert,
}

func RunFuncConvert(cmd *cobra.Command, args []string) {}
