/*
Copyright © 2023 David Saah dasorange.hope@gmail.com
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/DaveSaah/days/config"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
// check command calculates the number of days spent on earth
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Calculates the number of days you've spent on earth",
	Run: func(_ *cobra.Command, _ []string) {
		day, mth, yr := config.Cfg.ReadConfig()
		date := fmt.Sprintf("%s-%s-%s", yr, mth, day)
		days := daysSinceBirth(date)
		fmt.Printf("You've been alive for %d days\n", int(days))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func daysSinceBirth(birthdate string) int {
	// Parse the birthdate string
	layout := time.DateOnly
	birth, _ := time.Parse(layout, birthdate)

	// Calculate the difference in days
	days := int(time.Since(birth).Hours() / 24)

	return days
}
