/*
Copyright © 2023 David Saah dasorange.hope@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/DaveSaah/days/config"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
// init command Initializes dob and stores it in a config file
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the date you were born",
	Long:  `days init <dd-mm-yyyy> - Initialize the date you were born`,
	Run: func(cmd *cobra.Command, args []string) {
		// check if args is empty
		if len(args) == 0 {
			cmd.Help()
			os.Exit(1)
		}

		// check for date errors
		dob := strings.Split(args[0], "-")

		switch {
		case len(dob) != 3:
			fmt.Printf("Error: Invalid date format\n\n")
			cmd.Help()
			os.Exit(1)
		case len(dob[0]) != 2 || len(dob[1]) != 2 || len(dob[2]) != 4:
			fmt.Printf("Error: Invalid date format\n\n")
			cmd.Help()
			os.Exit(1)
		}

		config.Cfg.WriteConfig(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
