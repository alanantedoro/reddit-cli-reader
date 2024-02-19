package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Shows user identity and information.",
	Run: func(cmd *cobra.Command, args []string) {
		// query de identity
		fmt.Println("me exc")
	},
}

func init() {
	RootCmd.AddCommand(meCmd)
}
