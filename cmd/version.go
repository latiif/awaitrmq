package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VERSION is awaitrmq's current version
var VERSION = "unversioned"

// COMMIT is populated upon build
var COMMIT = "unknown commit"

// BRANCH specifies the branch this commit is build from
var BRANCH = "unknown branch"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of awaitrmq",
	Long:  `All software has versions. This is ara's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("awaitrmq", VERSION, BRANCH, COMMIT)
	},
}
