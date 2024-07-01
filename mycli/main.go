package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	gender string
)

var rootCMD = &cobra.Command{
	Use:   "mycli",
	Short: "mycli is a tool to do something",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from MyCLI!")
	},
}

var hiCMD = &cobra.Command{
	Use:   "hi [name]",
	Short: "hi is a command to say Hi!",
	// Untuk validasi argument
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("please input argument")
		}
		return nil
	},
	// Untuk handle ketika enter "command"
	Run: func(cmd *cobra.Command, args []string) {
		log.Println(args[1])
		title := "Mr. "
		if gender == "female" {
			title = "Ms. "
		}
		fmt.Printf("Hi %s%s!\n", title, args[0])
	},
}

var byeCMD = &cobra.Command{
	Use:   "bye",
	Short: "bye is a command to say Bye!",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bye!")
	},
}

func main() {
	// set flag for command "hi"
	hiCMD.Flags().StringVarP(&gender, "gender", "g", "male", "gender is a flag to detemine gender")
	// Init command
	rootCMD.AddCommand(hiCMD)
	rootCMD.AddCommand(byeCMD)

	err := rootCMD.Execute()
	if err != nil {
		log.Panic(err)
	}
}
