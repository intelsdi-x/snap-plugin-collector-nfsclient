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

package samba

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"errors"
	"encoding/json"

	"github.com/intelsdi-x/pulse/control/plugin"
	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"
)

const (
	// Name of plugin
	Name = "nfsclient"
	// Version of plugin
	Version = 1
	// Type of plugin
	Type = plugin.CollectorPluginType
)

var namespace_prefix = []string{"intel", "nfs", "client"}

type NFSCollector struct {
}

func NewNFSCollector() *NFSCollector {
	return &NFSCollector{}
}

// CollectMetrics collects metrics for testing
func (f *NFSCollector) CollectMetrics(mts []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	for i := range mts {
		if command, ok := mts[i].Config().Table()["command"]; ok {
			str_command := command.(ctypes.ConfigValueStr)
			parsed_command := strings.Split(str_command.Value, " ")
			main_command := parsed_command[0]
			var args []string
			if len(parsed_command) > 1 {
				args = parsed_command[1:len(parsed_command)]
			} else {
				args = make([]string, 0)
  			}
			out, error := exec.Command(main_command, args...).Output()
			if error != nil {
				return nil, errors.New(fmt.Sprint("Oh noes! An error ", error))
			}
			var decoded_data interface{}
			json.Unmarshal(out, &decoded_data)
			json_data := decoded_data.(map[string]interface{})
			mts[i].Data_ = json_data["value"].(float64);
			// Change the Namespace so it matches the namespace returned from the script
			mts[i].Namespace_ = []string{"intel", "exec", json_data["namespace"].(string)}
		}
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
	if _, ok := cfg.Table()["test-fail"]; ok {
		return mts, fmt.Errorf("testing")
	}
	if _, ok := cfg.Table()["test"]; ok {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "dummy", "test"}})
	}

	mts = append(mts, plugin.PluginMetricType{Namespace_: []string{"intel", "exec"}})

	// TODO: specify multiple commands/namespaces in a single config?
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
