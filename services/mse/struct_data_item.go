package mse

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DataItem is a nested struct in mse response
type DataItem struct {
	Name                     string                   `json:"Name" xml:"Name"`
	NewVersionPublishingFlag bool                     `json:"NewVersionPublishingFlag" xml:"NewVersionPublishingFlag"`
	Phase                    int                      `json:"Phase" xml:"Phase"`
	WasmLang                 int                      `json:"WasmLang" xml:"WasmLang"`
	IP                       string                   `json:"IP" xml:"IP"`
	SyncType                 string                   `json:"SyncType" xml:"SyncType"`
	LocalName                string                   `json:"LocalName" xml:"LocalName"`
	ClusterSpecificationName string                   `json:"ClusterSpecificationName" xml:"ClusterSpecificationName"`
	ClusterNamePrefix        string                   `json:"clusterNamePrefix" xml:"clusterNamePrefix"`
	Mode                     int                      `json:"Mode" xml:"Mode"`
	GmtCreate                string                   `json:"GmtCreate" xml:"GmtCreate"`
	OriginInstanceName       string                   `json:"OriginInstanceName" xml:"OriginInstanceName"`
	UserId                   string                   `json:"UserId" xml:"UserId"`
	ServiceName              string                   `json:"ServiceName" xml:"ServiceName"`
	TargetClusterUrl         string                   `json:"TargetClusterUrl" xml:"TargetClusterUrl"`
	ProjectDesc              string                   `json:"ProjectDesc" xml:"ProjectDesc"`
	App                      string                   `json:"App" xml:"App"`
	ShowName                 string                   `json:"ShowName" xml:"ShowName"`
	PodName                  string                   `json:"podName" xml:"podName"`
	ClusterType              string                   `json:"ClusterType" xml:"ClusterType"`
	CpuCapacity              string                   `json:"CpuCapacity" xml:"CpuCapacity"`
	Priority                 int                      `json:"Priority" xml:"Priority"`
	Code                     string                   `json:"Code" xml:"Code"`
	Port                     string                   `json:"Port" xml:"Port"`
	Summary                  string                   `json:"Summary" xml:"Summary"`
	Version                  string                   `json:"Version" xml:"Version"`
	TargetClusterName        string                   `json:"TargetClusterName" xml:"TargetClusterName"`
	GmtModified              string                   `json:"GmtModified" xml:"GmtModified"`
	Id                       int64                    `json:"Id" xml:"Id"`
	Agent                    string                   `json:"Agent" xml:"Agent"`
	MaxVersion               string                   `json:"MaxVersion" xml:"MaxVersion"`
	PublishState             int                      `json:"PublishState" xml:"PublishState"`
	Cluster                  string                   `json:"Cluster" xml:"Cluster"`
	TargetInstanceId         string                   `json:"TargetInstanceId" xml:"TargetInstanceId"`
	OriginInstanceNamespace  string                   `json:"OriginInstanceNamespace" xml:"OriginInstanceNamespace"`
	PrimaryUser              string                   `json:"PrimaryUser" xml:"PrimaryUser"`
	ZoneId                   string                   `json:"ZoneId" xml:"ZoneId"`
	ConfigCheck              string                   `json:"ConfigCheck" xml:"ConfigCheck"`
	WasmFile                 string                   `json:"WasmFile" xml:"WasmFile"`
	NamespaceId              string                   `json:"NamespaceId" xml:"NamespaceId"`
	Addr                     string                   `json:"Addr" xml:"Addr"`
	Category                 int                      `json:"Category" xml:"Category"`
	OriginInstanceAddress    string                   `json:"OriginInstanceAddress" xml:"OriginInstanceAddress"`
	Status                   string                   `json:"Status" xml:"Status"`
	MemoryCapacity           string                   `json:"MemoryCapacity" xml:"MemoryCapacity"`
	Values                   []map[string]interface{} `json:"values" xml:"values"`
}
