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
	// "strings"
	"time"
	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
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

type NFSCollector struct {
}

func NewNFSCollector() *NFSCollector {
	return &NFSCollector{}
}

// CollectMetrics collects metrics for testing
func (f *NFSCollector) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	for i := range mts {
		mts[i].Data_ = "Test Value"
		mts[i].Namespace_ = namespacePrefix
		// TODO: Error handling
		mts[i].Source_, _ = os.Hostname()
		mts[i].Timestamp_ = time.Now()
	}
	// return nil, errors.New(fmt.Sprint(mts[0].Data_))
	return mts, nil
}

//GetMetricTypes returns metric types for testing
func (f *NFSCollector) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	mts := []plugin.PluginMetricType{}
	
	for metric := range metricKeys {
		mts = append(mts, plugin.PluginMetricType{Namespace_: append(namespacePrefix, metricKeys[metric]...)})
	}
	return mts, nil
}

//GetConfigPolicy returns a ConfigPolicy for testing
func (f *NFSCollector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	rule, _ := cpolicy.NewStringRule("command", true)
	p := cpolicy.NewPolicyNode()
	p.Add(rule)
	c.Add([]string{"intel", "dummy", "exec"}, p)
	return c, nil
}

//Meta returns meta data for testing
func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(Name, Version, Type, []string{plugin.PulseGOBContentType}, []string{plugin.PulseGOBContentType})
}
