package main

import (
	"gator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Can't read config file %v\n", err)
	}

	err = cfg.SetUser("matt")
	if err != nil {
		fmt.Printf("Error setting user %v\n", err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("Error reading file a second time %v\n", err)
	}
	fmt.Printf("Config: \n%v\n", cfg)
}
