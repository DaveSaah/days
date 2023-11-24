/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/DaveSaah/days/config"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
// clean command deletes the stored config file
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Delete config file",
	Run: func(_ *cobra.Command, _ []string) {
		config.Cfg.DeleteConfig()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
