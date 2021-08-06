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
	"bitbucket.org/Monitoring/gaemi/core/collector/util"
	"bitbucket.org/Monitoring/gaemi/logging"
)

type QuorumWatcher struct {
	inserter  QuorumWatcherInserter
	rpcClient QuorumRPC
}

func NewQuorumWatcher(inserter QuorumWatcherInserter, client QuorumRPC) *QuorumWatcher {
	return &QuorumWatcher{
		inserter:  inserter,
		rpcClient: client,
	}
}

func (qw *QuorumWatcher) Probe() {

	//qw.getRaftCluster()
	qw.getRaftLeader()
	//qw.getPeerCount()

}

func (qw *QuorumWatcher) getRaftCluster() {
	var response = QuorumRaftInfoResponse{}
	var err = qw.rpcClient.PostRequest("raft_cluster", "1", &response)
	check(err)
	//qw.inserter.WriteQuorumReaderInfo(readerinfo)
}

func (qw *QuorumWatcher) getRaftLeader() {
	var response = QuorumDefaultResponse{}
	var err = qw.rpcClient.PostRequest("raft_leader", "1", &response)
	check(err)
	qw.inserter.WriteQuorumReaderInfo(response.Result)
}

func (qw *QuorumWatcher) getPeerCount() {
	var response = QuorumDefaultResponse{}
	var err = qw.rpcClient.PostRequest("net_peerCount", "1", &response)
	count, err := util.HexaNumberToInt(response.Result)

	check(err)
	qw.inserter.WriteQuorumPeerCount(count)
}

func check(e error) {
	if e != nil {
		logging.GetLogger().Panicf("Panic!! %s", e.Error())
		panic(e)
	}
}
