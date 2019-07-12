package commands

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the api-gateway version information",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`api-gateway:
	version     : 0.0.1
	go version  : %s
	go compiler : %s
	platform    : %s/%s
`, runtime.Version(), runtime.Compiler, runtime.GOOS, runtime.GOARCH)
	},
}
