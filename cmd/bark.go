package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra#create-additional-commands

var loud bool

var barkCmd = &cobra.Command{
	Use:   "bark",
	Short: "Make the husky say something",
	Long:  "Husky ASCII art from: https://www.asciiart.eu/animals/dogs.",
	// https://github.com/spf13/cobra#positional-and-custom-arguments
	Args: cobra.MinimumNArgs(1),
	RunE: runE,
}

func init() {
	// https://github.com/spf13/cobra#local-flags
	barkCmd.Flags().BoolVarP(&loud, "loud", "l", false, "bark with more volume")
}

func runE(cmd *cobra.Command, args []string) error {
	huskyASCII := huskyASCII1 + "`" + huskyASCII2

	words := args[0]
	if len(args) > 1 {
		words = strings.Join(args, " ")
	}

	if loud {
		words = strings.ToUpper(words)
	}

	huskysay := fmt.Sprintf(huskyASCII, words)
	fmt.Println(huskysay)
	return nil
}

// we have to break up the raw string literal because the ASCII art contains a
// backtick
const huskyASCII1 = `
     |\_/|
     | @ @   %s
     |   <>              _
     |  _/\------____ ((| |))
     |               `

const huskyASCII2 = `--' |
 ____|_       ___|   |___.'
/_/_____/____/_______|
`
