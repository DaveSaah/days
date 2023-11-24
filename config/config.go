package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Config defines the structure of the config file
type Config struct {
	Day   string
	Month string
	Year  string
}

var Cfg = Config{}

// Cfg_path is the path to the config file
var cfg_path string

// InitConfig initializes the config file
func (c *Config) InitConfig() {
	// find home directory
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// create config directory if it does not exist
	path := home + "/.config/days/"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}

	cfg_path = path + "config"

	// create config file if it does not exist
	if _, err := os.Stat(cfg_path); os.IsNotExist(err) {
		os.Create(cfg_path)
		fmt.Println("Config file created at ", cfg_path)
	}
}

// WriteConfig writes the date of birth to the config file
func (c *Config) WriteConfig(date string) {
	// open config file
	file, _ := os.OpenFile(cfg_path, os.O_RDWR, 0644)
	defer file.Close()

	// populate config struct
	dob := strings.Split(date, "-")
	c.Day = dob[0]
	c.Month = dob[1]
	c.Year = dob[2]

	// write config struct to config file
	encoder := json.NewEncoder(file)
	encoder.Encode(c)

	// display message
	fmt.Println("Date of birth saved successfully")
}

// ReadConfig reads the date of birth from the config file
// Returns the day, month and year
func (c *Config) ReadConfig() (string, string, string) {
	// open config file
	file, _ := os.OpenFile(cfg_path, os.O_RDWR, 0644)
	defer file.Close()

	// check if file is empty
	f_info, _ := file.Stat()
	if f_info.Size() == 0 {
		fmt.Println("Error: Config file is empty")
		os.Exit(1)
	}

	// decode config file into config struct
	decoder := json.NewDecoder(file)
	decoder.Decode(c)

	return c.Day, c.Month, c.Year
}

func (c *Config) DeleteConfig() {
	os.Remove(cfg_path)
	fmt.Println("Config file deleted successfully")
}
