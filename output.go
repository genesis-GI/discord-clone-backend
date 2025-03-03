package main

import (
	"github.com/fatih/color"
)

func warn(message string) {
	color.Yellow("[WARN] " + message)
}
