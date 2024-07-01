package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "mycli",
	Short: "mycli is a tool to do something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from MyCLI!")
	},
}

func main() {
	err := rootCMD.Execute()
	if err != nil {
		log.Panic(err)
	}
}
