package utils

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func GetIniData() *ini.File {
	iniData, err := ini.Load("config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	return iniData
}
