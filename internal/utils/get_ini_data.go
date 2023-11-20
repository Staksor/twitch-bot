package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func GetIniData() *ini.File {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("config.ini wasn't found")
		os.Exit(1)
	}

	return iniData
}
