package proactor

import (
	"bitbucket.org/Monitoring/gaemi/core/collector/config"
	"bitbucket.org/Monitoring/gaemi/core/collector/proactor/folder_watcher"
	"bitbucket.org/Monitoring/gaemi/core/collector/proactor/quorum_watcher"
	"bitbucket.org/Monitoring/gaemi/core/collector/storage/influxdb"
	"bitbucket.org/Monitoring/gaemi/logging"
	"time"
)

type Proactor struct {
	watchers []Watcher
}

func NewProactor() *Proactor {
	return &Proactor{
		watchers: []Watcher{},
	}
}

func (pro *Proactor) Registor(watcher Watcher) {
	pro.watchers = append(pro.watchers, watcher)
}

func (pro *Proactor) ProbeAll() {
	logging.GetLogger().Infof(" Probe All !")
	for _, watcher := range pro.watchers {
		watcher.Probe()
	}
}

func (pro *Proactor) ProbeAllAsync() {
	for _, watcher := range pro.watchers {
		go watcher.Probe()
	}
}

func Start(dbconnector *influxdb.Influxdb2Connector, conf *config.Configuration, done <-chan bool) {
	logging.GetLogger().Infof("==================================================")
	logging.GetLogger().Infof(" Proactor  Start !!!")
	logging.GetLogger().Infof("==================================================")

	proactor := NewProactor()
	// register folder watcher
	if conf.TARGET.FOLDER.Active == "on" || conf.TARGET.FOLDER.Active == "1" {
		fwinserter := folder_watcher.NewFolderWatcherInserter(dbconnector, conf.TARGET.FOLDER.MEASUREMENT)
		fw := folder_watcher.NewFolderWatcher(fwinserter, conf.TARGET.FOLDER.DataPath)
		proactor.Registor(fw)
	}

	// register quorum watcher
	if conf.TARGET.QUORUM.Active == "on" || conf.TARGET.FOLDER.Active == "1" {
		qwinserter := quorum_watcher.NewQuorumWatcherInserter(dbconnector, conf.TARGET.QUORUM.MEASUREMENT)
		qw := quorum_watcher.NewQuorumWatcher(qwinserter, quorum_watcher.NewQuorumRPCClient(conf.TARGET.QUORUM.URL))
		proactor.Registor(qw)
	}

	for {
		select {
		case <-time.After(1 * time.Second):
			proactor.ProbeAll()

		case <-done:
			logging.GetLogger().Infof(" Proactor  Finished !!!")

			return
		}
	}
}
