package cmd

import (
	"dictionary-of-chinese/cmd/idioms"
	"dictionary-of-chinese/cmd/proverb"
	"dictionary-of-chinese/cmd/words"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(importCmd)
}

const (
	command = "import"
)

var importCmd = &cobra.Command{
	Use: command,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			help()
			return
		}
		switch args[0] {

		case "idioms":
			idioms.Start()
		case "proverb":
			proverb.Start()
		case "words":
			words.Start()
		default:
			help()
		}
	},
}

func help() {
	fmt.Println(fmt.Sprintf("args should choose one of them: [%s, %s, %s]",
		strconv.Quote("idioms"),
		strconv.Quote("proverb"),
		strconv.Quote("words")))
}
