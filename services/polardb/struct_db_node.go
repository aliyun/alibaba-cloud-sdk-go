package polardb

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

// DBNode is a nested struct in polardb response
type DBNode struct {
	DBNodeClass             string `json:"DBNodeClass" xml:"DBNodeClass"`
	DBNodeDescription       string `json:"DBNodeDescription" xml:"DBNodeDescription"`
	MasterId                string `json:"MasterId" xml:"MasterId"`
	RemoteMemorySize        string `json:"RemoteMemorySize" xml:"RemoteMemorySize"`
	HotReplicaMode          string `json:"HotReplicaMode" xml:"HotReplicaMode"`
	MultiMasterPrimaryNode  string `json:"MultiMasterPrimaryNode" xml:"MultiMasterPrimaryNode"`
	ServerlessType          string `json:"ServerlessType" xml:"ServerlessType"`
	Serverless              string `json:"Serverless" xml:"Serverless"`
	MultiMasterLocalStandby string `json:"MultiMasterLocalStandby" xml:"MultiMasterLocalStandby"`
	MaxConnections          int    `json:"MaxConnections" xml:"MaxConnections"`
	RegionId                string `json:"RegionId" xml:"RegionId"`
	ServerWeight            string `json:"ServerWeight" xml:"ServerWeight"`
	MirrorInsName           string `json:"MirrorInsName" xml:"MirrorInsName"`
	DBNodeId                string `json:"DBNodeId" xml:"DBNodeId"`
	SccMode                 string `json:"SccMode" xml:"SccMode"`
	ImciSwitch              string `json:"ImciSwitch" xml:"ImciSwitch"`
	Tair                    string `json:"Tair" xml:"Tair"`
	MaxIOPS                 int    `json:"MaxIOPS" xml:"MaxIOPS"`
	FailoverPriority        int    `json:"FailoverPriority" xml:"FailoverPriority"`
	DBNodeRole              string `json:"DBNodeRole" xml:"DBNodeRole"`
	DBNodeStatus            string `json:"DBNodeStatus" xml:"DBNodeStatus"`
	CpuCores                string `json:"CpuCores" xml:"CpuCores"`
	SubCluster              string `json:"SubCluster" xml:"SubCluster"`
	CreationTime            string `json:"CreationTime" xml:"CreationTime"`
	AddedCpuCores           string `json:"AddedCpuCores" xml:"AddedCpuCores"`
	ZoneId                  string `json:"ZoneId" xml:"ZoneId"`
	Orca                    string `json:"Orca" xml:"Orca"`
	MemorySize              string `json:"MemorySize" xml:"MemorySize"`
}
