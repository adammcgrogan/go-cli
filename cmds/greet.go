package cmds

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-cli",
	Short: "A basic CLI tool created with Go and Cobra",
	Long:  `gocli is a simple application to demonstrate how to build modern command-line interfaces in Go using the Cobra library.`,
	// This function will run if no subcommand is provided.
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from the root command!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
