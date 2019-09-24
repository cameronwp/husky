package sled

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(pullCmd)
}

// Cmd is the root for husky sled commands
var Cmd = &cobra.Command{
	Use:   "sled",
	Short: "Make the husky use a sled",
	// https://github.com/spf13/cobra#prerun-and-postrun-hooks
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Logging you in...\nok")
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// TALK: business logic goes here
		fmt.Println("What do you want the husky to do with the sled?")
		return cmd.Help()
	},
}
