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
	// "fmt"
	"os"
	"strings"
	"time"

	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
)

const (
	// Name of plugin
	Name = "nfsclient"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

var namespacePrefix = []string{"intel", "nfs", "client"}

type nfsCollector struct {
	stats getNFSStats
	//TODO: Mockout proc reader
}

// NewNFSCollector returns and nfsCollector implementation
func NewNFSCollector(g getNFSStats) *nfsCollector {
	return &nfsCollector{
		g,
	}
}

type getNFSStats interface {
	getNFSMetric(string, string) int
	getRPCMetric(string) int
	getNumConnections(int64) int
	computeMounts() int
	getMetricKeys() [][]string
	regenerate()
}

// CollectMetrics collects metrics for testing
func (f *nfsCollector) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	if len(mts) == 0 {
		return nil, nil
	}

	//Find a way to regenerate the data on each task run automatically. We shouldn't do this manually
	f.stats.regenerate()

	for i := range mts {
		//This throws away the common namespace prefix and returns only them important parts
		importantNamespace := mts[i].Namespace_[len(namespacePrefix):]
		if namespaceContains("nfs", importantNamespace) {
			mts[i].Data_ = f.stats.getNFSMetric(importantNamespace[0], importantNamespace[1])
		} else if namespaceContains("rpc", importantNamespace) {
			mts[i].Data_ = f.stats.getRPCMetric(importantNamespace[1])
		} else if namespaceContains("num_connections", importantNamespace) { //Then it is one of the top level
			mts[i].Data_ = f.stats.getNumConnections(int64(2049))
		} else if namespaceContains("num_mounts", importantNamespace) {
			mts[i].Data_ = f.stats.computeMounts()
		}
		// TODO: Error handling
		mts[i].Source_, _ = os.Hostname()
		mts[i].Timestamp_ = time.Now()
	}
	// return nil, errors.New(fmt.Sprint(mts[0].Data_))
	return mts, nil
}

//GetMetricTypes returns metric types
func (f *nfsCollector) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mts := []plugin.PluginMetricType{}

	for metric := range f.stats.getMetricKeys() {
		mts = append(mts, plugin.PluginMetricType{Namespace_: append(namespacePrefix, metricKeys[metric]...)})
	}
	return mts, nil
}

//GetConfigPolicy returns a ConfigPolicy for testing
func (f *nfsCollector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	rule, _ := cpolicy.NewStringRule("command", true)
	p := cpolicy.NewPolicyNode()
	p.Add(rule)
	c.Add([]string{"intel", "dummy", "exec"}, p)
	return c, nil
}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.SnapGOBContentType}, []string{plugin.SnapGOBContentType})
}

func namespaceContains(element string, slice []string) bool {
	for _, v := range slice {
		if strings.Contains(v, element) {
			return true
		}
	}
	return false
}
