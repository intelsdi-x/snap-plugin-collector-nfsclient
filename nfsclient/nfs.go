/*
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nfsclient

import (
    "os"
    "bufio"
    "strings"
    "strconv"
)

type nfsClient struct {
	data map[string][]string
}

func New() *nfsClient {
    return &nfsClient{
        data: generate(),
    }
}

var nfsValues = []string {
    "getattr",
    "setattr",
    "lookup",
    "access",
    "readlink",
    "read",
    "write",
    "create",
    "mkdir",
    "remove",
    "rmdir",
    "rename",
    "link",
    "readdir",
    "readdirplus",
    "fsstat",
    "fsinfo",
    "pathconf",
}
var metricKeys = [][]string {
    {"num_connections"},
    {"num_mounts"},
    {"rpc","calls"},
    {"rpc","retransmissions"},
    {"rpc","authrefresh"},
}

var nfsstatPositions = map[string]int {
    "getattr": 3,
    "setattr": 4,
    "lookup": 5,
    "access": 6,
    "readlink": 7,
    "read": 8,
    "write": 9,
    "create": 10,
    "mkdir": 11,
    "remove": 14,
    "rmdir": 15,
    "rename": 16,
    "link": 17,
    "readdir": 18,
    "readdirplus": 19,
    "fsstat": 20,
    "fsinfo": 21,
    "pathconf": 22,
}

var rpcPositions = map[string]int {
    "calls": 1,
    "retransmissions": 2,
    "authrefresh": 3,
}

var nfsFileMapping = map[string]string {
    "net": "net",
    "rpc": "rpc",
    "proc2": "nfsv2",
    "proc3": "nfsv3",
    "proc4": "nfsv4",
}

func (n *nfsClient) getMetricKeys() [][]string {
    // This just creates all the same measurements for nfsv2,3,and 4. They all have the same measurement values
    for proto := 2; proto < 5; proto++ {
        for i := range nfsValues {
            var value = []string {"nfsv" + strconv.Itoa(proto), nfsValues[i]}
            metricKeys = append(metricKeys, value)
        }   
    }
    return metricKeys
}

func computeConnections() int {
    count := 0
    file, _ := os.Open("/proc/net/tcp")
    scanner := bufio.NewScanner(bufio.NewReader(file))
    for scanner.Scan() {
        //NFS port in hex is 0801 (2049 in decimal), we can change this to be flexible for other ports later
        if strings.Contains(scanner.Text(), ":0801") {
            count++
        }
    }
    return count
}

func computeMounts() int {
    count := 0
    file, _ := os.Open("/proc/mounts")
    scanner := bufio.NewScanner(bufio.NewReader(file))
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), " nfs ") {
            count++
        }
    }
    return count
}

func (n *nfsClient) getNFSMetric(nfsType string, statName string) int  {
    // Throw away the error
    value, _ := strconv.Atoi(n.data[nfsType][nfsstatPositions[statName]])
    return value
}

func (n *nfsClient) getRPCMetric(statName string) int {
    // Throw away the error
    value, _ := strconv.Atoi(n.data["rpc"][rpcPositions[statName]])
    return value
}

// var connections []*(process.NetConnectionStat)
func (n *nfsClient) getOtherMetric(statName string) int  {
    var value int
    switch statName {
    case "num_conditions": value = computeConnections()
    case "num_mounts": value = computeMounts()
    //Handle a default case?
    }
    return value
}

func generate() map[string][]string {
    nfsStats := make(map[string][]string)
    file, _ := os.Open("/proc/net/rpc/nfs")
    scanner := bufio.NewScanner(bufio.NewReader(file))
    for scanner.Scan() {
        processedLine := strings.Split(scanner.Text(), " ")
        // Get the line name
        lineName := processedLine[0]
        nfsStats[nfsFileMapping[lineName]] = processedLine
    }
    return nfsStats
}