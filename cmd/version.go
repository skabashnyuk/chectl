package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Che version information",
	Run: func(cmd *cobra.Command, args []string) {
		PrintVersion()

	},
}

func PrintVersion() {
	fmt.Printf("chectl %s\n", "0.1.0")
	fmt.Printf("- os/arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("- go version: %s\n", runtime.Version())
}
