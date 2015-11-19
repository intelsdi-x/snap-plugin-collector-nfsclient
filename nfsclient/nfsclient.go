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

// import (
//     // "github.com/shirou/gopsutil/net"
//     "os"
// )

var metricKeys = [][]string {
    {"num_connections"},
    {"num_mounts"},
    {"rpc","calls"},
    {"rpc","retransmissions"},
    {"rpc","authrefresh"},
    {"nfsv3","getattr"},
    {"nfsv3","setattr"},
    {"nfsv3","lookup"},
    {"nfsv3","access"},
    {"nfsv3","readlink"},
    {"nfsv3","read"},
    {"nfsv3","write"},
    {"nfsv3","create"},
    {"nfsv3","mkdir"},
    {"nfsv3","remove"},
    {"nfsv3","rmdir"},
    {"nfsv3","rename"},
    {"nfsv3","link"},
    {"nfsv3","readdir"},
    {"nfsv3","readdirplus"},
    {"nfsv3","fsstat"},
    {"nfsv3","fsinfo"},
    {"nfsv3","pathconf"},
}



var nfsstat_positions = map[string]int {
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

var rpc_positions = map[string]int {
    "calls": 1,
    "retransmissions": 2,
    "authrefresh": 3,
}

// var processes = []*(process.NetConnectionStat){}