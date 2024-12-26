package schedulerx3

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

// Data is a nested struct in schedulerx3 response
type Data struct {
	DesignateType    int                    `json:"DesignateType" xml:"DesignateType"`
	StatisticDetail  map[string]interface{} `json:"StatisticDetail" xml:"StatisticDetail"`
	OrderId          int64                  `json:"OrderId" xml:"OrderId"`
	Transferable     bool                   `json:"Transferable" xml:"Transferable"`
	IntranetDomain   string                 `json:"IntranetDomain" xml:"IntranetDomain"`
	ChargeType       string                 `json:"ChargeType" xml:"ChargeType"`
	CreateTime       string                 `json:"CreateTime" xml:"CreateTime"`
	Spm              int                    `json:"Spm" xml:"Spm"`
	PageSize         int                    `json:"PageSize" xml:"PageSize"`
	MaxJobNum        int                    `json:"MaxJobNum" xml:"MaxJobNum"`
	ProductType      int                    `json:"ProductType" xml:"ProductType"`
	AccessToken      string                 `json:"AccessToken" xml:"AccessToken"`
	ClusterName      string                 `json:"ClusterName" xml:"ClusterName"`
	WorkerNum        int                    `json:"WorkerNum" xml:"WorkerNum"`
	JobDescription   string                 `json:"JobDescription" xml:"JobDescription"`
	PageNumber       int                    `json:"PageNumber" xml:"PageNumber"`
	VpcId            string                 `json:"VpcId" xml:"VpcId"`
	EngineType       string                 `json:"EngineType" xml:"EngineType"`
	InternetDomain   string                 `json:"InternetDomain" xml:"InternetDomain"`
	JobId            int64                  `json:"JobId" xml:"JobId"`
	ClusterId        string                 `json:"ClusterId" xml:"ClusterId"`
	Status           int                    `json:"Status" xml:"Status"`
	EngineVersion    string                 `json:"EngineVersion" xml:"EngineVersion"`
	ClusterSpec      string                 `json:"ClusterSpec" xml:"ClusterSpec"`
	JobNum           int                    `json:"JobNum" xml:"JobNum"`
	KubeConfig       string                 `json:"KubeConfig" xml:"KubeConfig"`
	AppGroupId       int64                  `json:"AppGroupId" xml:"AppGroupId"`
	EndTime          string                 `json:"EndTime" xml:"EndTime"`
	JobExecutionId   string                 `json:"JobExecutionId" xml:"JobExecutionId"`
	Total            int64                  `json:"Total" xml:"Total"`
	Zones            []string               `json:"Zones" xml:"Zones"`
	RootProgress     RootProgress           `json:"RootProgress" xml:"RootProgress"`
	TotalProgress    TotalProgress          `json:"TotalProgress" xml:"TotalProgress"`
	VSwitches        []VswitchesItem        `json:"VSwitches" xml:"VSwitches"`
	Records          []Record               `json:"Records" xml:"Records"`
	TaskProgress     []TaskProgressItem     `json:"TaskProgress" xml:"TaskProgress"`
	WorkerProgress   []WorkerProgressItem   `json:"WorkerProgress" xml:"WorkerProgress"`
	ShardingProgress []ShardingProgressItem `json:"ShardingProgress" xml:"ShardingProgress"`
}
