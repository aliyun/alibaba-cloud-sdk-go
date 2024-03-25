package oceanbasepro

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

// DataInDescribeInstances is a nested struct in oceanbasepro response
type DataInDescribeInstances struct {
	VpcId                           string                  `json:"VpcId" xml:"VpcId"`
	CommodityCode                   string                  `json:"CommodityCode" xml:"CommodityCode"`
	ExpireTime                      string                  `json:"ExpireTime" xml:"ExpireTime"`
	State                           string                  `json:"State" xml:"State"`
	InstanceClass                   string                  `json:"InstanceClass" xml:"InstanceClass"`
	CreateTime                      string                  `json:"CreateTime" xml:"CreateTime"`
	DeployMode                      string                  `json:"DeployMode" xml:"DeployMode"`
	CpuArchitecture                 string                  `json:"CpuArchitecture" xml:"CpuArchitecture"`
	MaintainTime                    string                  `json:"MaintainTime" xml:"MaintainTime"`
	DeployType                      string                  `json:"DeployType" xml:"DeployType"`
	PayType                         string                  `json:"PayType" xml:"PayType"`
	DiskSize                        string                  `json:"DiskSize" xml:"DiskSize"`
	DiskType                        string                  `json:"DiskType" xml:"DiskType"`
	InstanceId                      string                  `json:"InstanceId" xml:"InstanceId"`
	ExpireSeconds                   int                     `json:"ExpireSeconds" xml:"ExpireSeconds"`
	Mem                             int64                   `json:"Mem" xml:"Mem"`
	EnableUpgradeNodes              bool                    `json:"EnableUpgradeNodes" xml:"EnableUpgradeNodes"`
	Cpu                             int                     `json:"Cpu" xml:"Cpu"`
	Version                         string                  `json:"Version" xml:"Version"`
	InstanceName                    string                  `json:"InstanceName" xml:"InstanceName"`
	Series                          string                  `json:"Series" xml:"Series"`
	UsedDiskSize                    int64                   `json:"UsedDiskSize" xml:"UsedDiskSize"`
	ResourceGroupId                 string                  `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceType                    string                  `json:"InstanceType" xml:"InstanceType"`
	InstanceRole                    string                  `json:"InstanceRole" xml:"InstanceRole"`
	InTempCapacityStatus            bool                    `json:"InTempCapacityStatus" xml:"InTempCapacityStatus"`
	EnableReadOnlyReplicaManagement bool                    `json:"EnableReadOnlyReplicaManagement" xml:"EnableReadOnlyReplicaManagement"`
	AvailableZones                  []string                `json:"AvailableZones" xml:"AvailableZones"`
	Resource                        Resource                `json:"Resource" xml:"Resource"`
	DataDiskAutoScaleConfig         DataDiskAutoScaleConfig `json:"DataDiskAutoScaleConfig" xml:"DataDiskAutoScaleConfig"`
}
