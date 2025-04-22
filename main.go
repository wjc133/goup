package main

import (
	"github.com/sirupsen/logrus"
	"github.com/wjc133/goup/internal/commands"
)

func main() {
	rootCmd := commands.NewCommand()
	if err := rootCmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
