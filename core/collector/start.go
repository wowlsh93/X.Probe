/*
Copyright Monitoring Corp. All Rights Reserved.

Written by HAMA
*/

package collector

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"bitbucket.org/Monitoring/gaemi/core/collector/config"
	"bitbucket.org/Monitoring/gaemi/core/collector/dummy"
	"bitbucket.org/Monitoring/gaemi/core/collector/proactor"
	"bitbucket.org/Monitoring/gaemi/core/collector/reactor"
	"bitbucket.org/Monitoring/gaemi/core/collector/storage/influxdb"
	"bitbucket.org/Monitoring/gaemi/logging"
)

var configPath string

func startCmd() *cobra.Command {
	collectorStartCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "", "config file")
	collectorStartCmd.PersistentFlags().StringVarP(&configPath, "target", "t", "", "collect type (default is all, proactor or reactor)")
	collectorStartCmd.MarkPersistentFlagRequired("mode")

	return collectorStartCmd
}

var collectorStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start collector..",
	Long:  `start collector..`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 0 {
			return fmt.Errorf("trailing args detected")
		}

		cmd.SilenceUsage = true
		return serve(args)
	},
}

func startUpLogging(conf *config.Configuration) {
	logging.GetLogger().Infof("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	logging.GetLogger().Infof(" Collector version %s Start!!                     ", conf.VERSION)
	logging.GetLogger().Infof("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@")
	logging.GetLogger().Infof("[INFLUXDB]")
	logging.GetLogger().Infof("organizationi: %s", conf.INFLUXDB.ORGANIZATION)
	logging.GetLogger().Infof("bucket: %s", conf.INFLUXDB.BUCKET)
	logging.GetLogger().Infof("db path: %s", conf.INFLUXDB.DBPATH)
	logging.GetLogger().Infof("[TARGET]")
	logging.GetLogger().Infof("[FOLDER WATCHER]")
	logging.GetLogger().Infof("active: %s", conf.TARGET.FOLDER.Active)
	logging.GetLogger().Infof("===================================================")
}

func enddingLogging() {

	logging.GetLogger().Infof("==================================================")
	logging.GetLogger().Infof(" @@@@@@@@@@ Endding by signal @@@@@@@@@@@")
	logging.GetLogger().Infof("==================================================")
}

func serve(args []string) error {
	var conf, err = config.InitConfig(configPath)
	if err != nil {
		logging.GetLogger().Panicf("configuration error : %s", err)
		os.Exit(1)
	}
	logging.InitLog(&conf)
	startUpLogging(&conf)

	done := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	dbconnector := influxdb.NewInfluxdb2Connector(conf.INFLUXDB.DBPATH,
		conf.INFLUXDB.AUTHTOKEN,
		conf.INFLUXDB.ORGANIZATION,
		conf.INFLUXDB.BUCKET)

	go dummy.StartDummyData(conf.TARGET.FOLDER.DataPath, -1)
	go reactor.Start(dbconnector, &conf, done)
	go proactor.Start(dbconnector, &conf, done)

	for {
		select {
		case <-time.After(1 * time.Second):
			//proactor.ProbeAll()
			//case <- resume:
			//case <- suspend:
			//case <- reconfiguring:
			logging.GetLogger().Tracef("just idle......")
		case <-sigs:
			done <- true
			time.Sleep(1000 * time.Millisecond)
			enddingLogging()
			return nil
		}
	}
	return nil
}
