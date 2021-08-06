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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type QuorumRPC interface {
	PostRequest(method string, id string, response interface{}) error
}

type QuorumRPCClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

func NewQuorumRPCClient(url string) *QuorumRPCClient {
	return &QuorumRPCClient{
		BaseURL: url,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type QuorumRequest struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Id      string `json:"id"`
}

type QuorumDefaultResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  string `json:"result"`
}

type QuorumRaftInfoResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  []struct {
		RaftId     int    `json:"raftId"`
		NodeId     string `json:"nodeId"`
		P2PPort    int    `json:"p2pPort"`
		RaftPort   int    `json:"raftPort"`
		Hostname   string `json:"hostname"`
		Role       string `json:"role"`
		NodeActive bool   `json:"nodeActive"`
	} `json:"result"`
}

func (c QuorumRPCClient) PostRequest(method string, id string, response interface{}) error {

	request := QuorumRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Id:      id,
	}
	reqJson, err := json.Marshal(request)

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(reqJson))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, response)
	if err != nil {
		log.Printf("Reading body failed: %s", err)
		return err
	}

	return nil
}
