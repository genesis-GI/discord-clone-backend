package main

import (
	"os"
	"strings"
)

func checkForParameters(){
	if len(os.Args) > 1 {
		input := os.Args[1]
		input = strings.ToLower(input)

		if input == "production" {
			os.Setenv("RAILWAY_ENVIRONMENT", "production")
		}else {
			errorMesssage("[✗ FAILURE] Invalid argument: "+input)
			os.Exit(1)
		}
	}
}