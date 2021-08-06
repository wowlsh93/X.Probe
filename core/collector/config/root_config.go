/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package config

import "bitbucket.org/Monitoring/gaemi/core/collector/config/configurations"

type Configuration struct {
	VERSION  string
	INFLUXDB configurations.InfluxDBConfiguration
	TARGET   configurations.TargetConfiguration
	LOG      configurations.LogConfiguration
}
