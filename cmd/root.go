/*
Copyright © 2023 David Saah dasorange.hope@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/DaveSaah/days/config"
	"github.com/spf13/cobra"
	// "github.com/spf13/cobra/doc"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "days",
	Short: "Checks the number of days you've spent on earth",
	Run: func(cmd *cobra.Command, _ []string) {
		if show_dob, _ := cmd.Flags().GetBool("show-dob"); show_dob {
			day, mth, yr := config.Cfg.ReadConfig()
			mth_int, _ := strconv.Atoi(mth)
			fmt.Printf("Date of birth: %s %s, %s\n", time.Month(mth_int), day, yr)
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// cobra.CheckErr(doc.GenMarkdownTree(rootCmd, "/home/hesed/Codes/days/docs"))
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().Bool("show-dob", false, "Preview your dob in config file")
}
