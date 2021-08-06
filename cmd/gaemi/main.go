/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"bitbucket.org/Monitoring/gaemi/core/collector"
	"bitbucket.org/Monitoring/gaemi/core/version"
)

var mainCmd = &cobra.Command{

	Use:   "gaemi",
	Short: "gaemi",
	Long:  "[gaemi] - Ledgermaster's Monitoring proactor",
}

func main() {

	mainCmd.AddCommand(collector.Cmd())
	mainCmd.AddCommand(version.Cmd())

	if err := mainCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
