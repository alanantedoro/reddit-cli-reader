package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "redditcli",
	Short: "Reddit-CLI is a CLI tool to browse through reddit.",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("Welcome to Reddit.")

	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	for {
	// 		fmt.Print("> ")
	// 		if !scanner.Scan() {
	// 			break
	// 		}

	// 		input := scanner.Text()
	// 		input = strings.TrimSpace(input)

	// 		switch input {
	// 		case "help":
	// 			fmt.Println("Available commands:")
	// 		case "exit":
	// 			fmt.Println("Exiting RedditCLI...")
	// 			return
	// 		case "me":
	// 			fmt.Println("dentro de me")
	// 			cmd.RootCmd.MeCmd.Execute()
	// 		}

	// 		fmt.Println("Command not recognized. Type 'help' for a list of commands.")
	// 	}

	// 	if err := scanner.Err(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// },
}
