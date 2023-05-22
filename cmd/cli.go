package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "library",
	Short: "Library JSON manipulator",
	Long:  `An application that manipulates the 'library.json' file populating it with the required methods`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the app version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("0.0.1")
	},
}

var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User manipulation commands",
}

var addUser = &cobra.Command{
	Use:   "add",
	Short: "Register an user to the library",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Here")
	},
}

func Init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(userCmd)
	userCmd.AddCommand(addUser)
}

func Execute() {
	Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
