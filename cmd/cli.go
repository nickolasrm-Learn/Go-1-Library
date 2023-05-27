package cmd

import (
	"fmt"
	"os"

	"github.com/nickolasrm-Learn/Go-2-Library/cmd/product"
	"github.com/nickolasrm-Learn/Go-2-Library/cmd/user"
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

func Init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(product.Product())
	rootCmd.AddCommand(user.User())
}

func Execute() {
	Init()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
