/*
 *
 * Copyright Monitoring Corp. All Rights Reserved.
 *
 * Written by HAMA
 *
 *
 */

package quorum_watcher

import (
	connector "bitbucket.org/Monitoring/gaemi/core/collector/storage/influxdb"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"time"
)

type QuorumWatcherInserter struct {
	connector   *connector.Influxdb2Connector
	measurement string
}

func NewQuorumWatcherInserter(connector *connector.Influxdb2Connector,
	measurement string) QuorumWatcherInserter {

	return QuorumWatcherInserter{
		connector:   connector,
		measurement: measurement,
	}
}

func (inserter *QuorumWatcherInserter) WriteQuorumReaderInfo(enodeid string) {
	p := influxdb2.NewPointWithMeasurement(inserter.measurement).
		AddTag("keyword", "raft").
		AddField("raft_reader_enodeid", enodeid).
		SetTime(time.Now())

	inserter.connector.Write(p)
}

func (inserter *QuorumWatcherInserter) WriteQuorumPeerCount(count int) {
	p := influxdb2.NewPointWithMeasurement(inserter.measurement).
		AddTag("keyword", "net").
		AddField("peer_count", count).
		SetTime(time.Now())

	inserter.connector.Write(p)
}
