/*
 *
 * Copyright Monitoring Corp. All Rights Reserved.
 *
 * Written by HAMA
 *
 *
 */

package dummy

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStartDummyData(t *testing.T) {

	//given or ready
	Clear("/home/hama/monitoring/data")

	//when
	StartDummyData("/home/hama/monitoring", 10)

	//then
	_, err := os.Stat("/home/hama/monitoring/data")
	assert.NoError(t, err, "file dose not exits")

}

func TestClearDummyData(t *testing.T) {

	//given or ready
	StartDummyData("/home/hama/monitoring", 10)

	//when
	Clear("/home/hama/monitoring/data")

	//then
	_, err := os.Stat("/home/hama/monitoring/data")
	assert.Error(t, err, "file exits")
}
