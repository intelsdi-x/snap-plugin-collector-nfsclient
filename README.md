DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.
<!--
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
-->
# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


# Snap collector plugin - nfsclient

This plugin collects NFS client statistics from any system that has NFS tools installed

Used for monitoring the rate of change in NFS ops 

The intention for this plugin is to identify which hosts could be having slowness issues or problems with NFS.

This plugin is used in the [Snap framework] (http://github.com/intelsdi-x/snap).


1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

In order to use this plugin you need "nfs-client" to be installed on a Linux target host.

### System Requirements

* Linux OS
* [nfs-client package](#installation)
* [golang 1.6+](https://golang.org/dl/)

### Installation

#### Install nfs-client package:
To install sysstat package from the official repositories simply use:
- For Ubuntu, Debian: `sudo apt-get install nfs-client`
- For CentOS, Fedora: `sudo yum install nfs-client`

#### To build the plugin binary:
Get the source by running a `go get` to fetch the code:
```
$ go get github.com/intelsdi-x/snap-plugin-collector-nfsclient
```

Build the plugin by running make within the cloned repo:
```
$ cd $GOPATH/src/github.com/intelsdi-x/snap-plugin-collector-nfsclient && make
```
This builds the plugin in `./build`

#### Builds
You can also download prebuilt binaries for OS X and Linux (64-bit) at the [releases](https://github.com/intelsdi-x/snap-plugin-collector-nfsclient/releases) page

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

By default iostat executable binary are searched in the directories named by the PATH environment. 

## Documentation

To learn more about this plugin:

* [Snap nfsclient examples](#examples)

### Collected Metrics
This plugin has the ability to gather the following metrics:

**NFS Statistics**

This collector has support for NFS v2, v3, and v4. These are all counters.

Metric namespace prefix: /intel/nfs/client/nfsv{nfs_version_number}

Name |
------------ |
getattr|
setattr|
lookup|
access|
readlink|
read|
write|
create|
mkdir|
remove|
rmdir|
rename|
link|
readdir|
readdirplus|
fsstat|
fsinfo|
pathconf|



**RPC statistics**

NOTE: These are all counters

Metric namespace prefix: /intel/nfs/client/rpc

Name |
------------ |
calls |
retransmission|
authrefresh|

**Other**

Metric namespace prefix: /intel/nfs/client

Name | Description
------------ | -------------
num_connections | The number of NFS connections from the client to NFS servers
num_mounts | The number of NFS mounts on the client

*Notes:*

By default metrics are gathered once per second.

### Examples
This is still in progress

### Roadmap
This plugin is still in active development. As we launch this plugin, we have a few items in mind for the next few releases:
- [ ] Add support for non-default NFS ports
- [ ] Additional error handling

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-nfsclient/issues) 
and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-nfsclient/pulls).

## Community Support
This repository is one of **many** plugins in the **Snap**, a powerful telemetry agent framework. See the full project at
http://github.com/intelsdi-x/snap. To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).


## Contributing
We love contributions! :heart_eyes:

There is more than one way to give back, from examples to blogs to code updates.

## License

[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).


## Acknowledgements

* Author: [Taylor Thomas](https://github.com/thomastaylor312)
* Contributor: [Esteban Martinez](https://github.com/ecmartz)
