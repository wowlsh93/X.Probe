/*
 *
 * Copyright Monitoring Corp. All Rights Reserved.
 *
 * Written by HAMA
 *
 *
 */

package util

import (
	"strconv"
	"strings"
)

func HexaNumberToInt(hexaString string) (int, error) {
	// replace 0x or 0X with empty String
	numberStr := strings.Replace(hexaString, "0x", "", -1)
	numberStr = strings.Replace(numberStr, "0X", "", -1)

	output, err := strconv.ParseInt(numberStr, 16, 64)
	return int(output), err
}
