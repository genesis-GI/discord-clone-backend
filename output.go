package main

import (
	"github.com/fatih/color"
)

func warn(message string) {
	color.Yellow("[WARN] " + message)
}

func info(message string) {
	color.Cyan("[INFO] " + message)
}

func neutral(message string) {
	color.White("[OUTPUT] " + message)
}

func success(message string) {
	color.Green("[SUCCESS] " + message)
}

func errorMesssage(message string) {
	color.Red("[ERROR] " + message)
}

func environment(message string) {
	color.Magenta("[ENVIRONMENT] " + message)
}
