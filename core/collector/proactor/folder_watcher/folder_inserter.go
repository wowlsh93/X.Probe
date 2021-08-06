/*
 *
 * Copyright Monitoring Corp. All Rights Reserved.
 *
 * Written by HAMA
 *
 *
 */

package folder_watcher

import (
	connector "bitbucket.org/Monitoring/gaemi/core/collector/storage/influxdb"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type FolderWatcherInserter struct {
	connector   *connector.Influxdb2Connector
	measurement string
}

func NewFolderWatcherInserter(connector *connector.Influxdb2Connector,
	measurement string) *FolderWatcherInserter {

	return &FolderWatcherInserter{
		connector:   connector,
		measurement: measurement,
	}
}

func (inserter *FolderWatcherInserter) WriteFolderStatus(tag string, currentFolderSize int64) {
	p := influxdb2.NewPointWithMeasurement(inserter.measurement).
		AddTag("tag", tag).
		AddField("size", currentFolderSize).
		SetTime(time.Now())

	inserter.connector.Write(p)
}
