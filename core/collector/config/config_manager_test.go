/*
 *
 * Copyright Monitoring Corp. All Rights Reserved.
 *
 * Written by HAMA
 *
 *
 */

package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigurationInit(t *testing.T) {

	_, err := InitConfig("")
	assert.NoError(t, err, "configuration error!!")
}
