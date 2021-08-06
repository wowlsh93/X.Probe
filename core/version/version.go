/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

const ProgramName = "gaemi"

func Cmd() *cobra.Command {
	return cobraCommand
}

var cobraCommand = &cobra.Command{
	Use:   "version",
	Short: "Print current version of the gaemi",
	Long:  `Print current version of the gaemi`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("unnessary trailing args detected")
		}

		cmd.SilenceUsage = true
		fmt.Print(GetInfo())
		return nil
	},
}

func GetInfo() string {

	return fmt.Sprintf("%s:\n  Version: %s\n  Go version: %s\n"+
		"  OS/Arch: %s\n",
		ProgramName, "v0.0.1", runtime.Version(),
		fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH))

}
