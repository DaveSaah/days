/*
Copyright © 2023 David Saah
*/
package main

import (
	"github.com/DaveSaah/days/cmd"
	"github.com/DaveSaah/days/config"
)

func main() {
	config.Cfg.InitConfig() // makes sure config file exists
	cmd.Execute()
}
