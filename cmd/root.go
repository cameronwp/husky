package cmd

import (
	"fmt"
	"os"

	"github.com/cameronwp/husky/cmd/sled"
	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra#getting-started

func init() {
	rootCmd.AddCommand(barkCmd)
	rootCmd.AddCommand(sled.Cmd)
}

var rootCmd = &cobra.Command{
	Use:   "husky",
	Short: "Husky helps you get around in the snow",
	Long:  `Huskies are loyal, lovable companions. They'll do anything you need them to do in the wintertime because they love the cold like arctic psychopaths.`,
	// TALK: I prefer RunE over Run, especially when creating commands for technical users. It's easier to surface errors
	RunE: func(cmd *cobra.Command, args []string) error {
		// https://github.com/spf13/cobra#help-command
		return cmd.Help()
	},
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
