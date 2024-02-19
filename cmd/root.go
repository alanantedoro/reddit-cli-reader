package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "redditcli",
	Short: "Reddit-CLI is a CLI tool to browse through reddit.",
}
