package reactor

import (
	"bitbucket.org/Monitoring/gaemi/core/collector/config"
	"bitbucket.org/Monitoring/gaemi/core/collector/storage/influxdb"
	"bitbucket.org/Monitoring/gaemi/logging"
)

func Start(dbconnector *influxdb.Influxdb2Connector, conf *config.Configuration, done <-chan bool) {
	logging.GetLogger().Infof("==================================================")
	logging.GetLogger().Infof(" Reactor  Start !!!")
	logging.GetLogger().Infof("==================================================")
}
