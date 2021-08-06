/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package collector

import (
	"fmt"
	"github.com/spf13/cobra"
)

const (
	collectorFuncName = "collector"
	collectorCmdDes   = "collector start"
)

func Cmd() *cobra.Command {
	scannerCmd.AddCommand(startCmd())
	return scannerCmd
}

var scannerCmd = &cobra.Command{
	Use:   collectorFuncName,
	Short: fmt.Sprint(collectorCmdDes),
	Long:  fmt.Sprint(collectorCmdDes),
}
