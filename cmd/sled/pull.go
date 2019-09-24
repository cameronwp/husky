package sled

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Make the husky pull something on a sled",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		msg := fmt.Sprintf("A CLI won't help you pull a %s, go here instead:\n\thttps://www.rover.com/blog/train-dog-pull-sled/", args[0])
		fmt.Println(msg)
		return nil
	},
}
