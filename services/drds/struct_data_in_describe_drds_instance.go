package drds

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

// DataInDescribeDrdsInstance is a nested struct in drds response
type DataInDescribeDrdsInstance struct {
	Type                  string                                      `json:"Type" xml:"Type"`
	Status                string                                      `json:"Status" xml:"Status"`
	CreateTime            int64                                       `json:"CreateTime" xml:"CreateTime"`
	VersionAction         string                                      `json:"VersionAction" xml:"VersionAction"`
	StorageType           string                                      `json:"StorageType" xml:"StorageType"`
	NetworkType           string                                      `json:"NetworkType" xml:"NetworkType"`
	Label                 string                                      `json:"Label" xml:"Label"`
	MysqlVersion          int                                         `json:"MysqlVersion" xml:"MysqlVersion"`
	InstanceSpec          string                                      `json:"InstanceSpec" xml:"InstanceSpec"`
	VpcCloudInstanceId    string                                      `json:"VpcCloudInstanceId" xml:"VpcCloudInstanceId"`
	Description           string                                      `json:"Description" xml:"Description"`
	Version               int64                                       `json:"Version" xml:"Version"`
	ExpireDate            int64                                       `json:"ExpireDate" xml:"ExpireDate"`
	MasterInstanceId      string                                      `json:"MasterInstanceId" xml:"MasterInstanceId"`
	CommodityCode         string                                      `json:"CommodityCode" xml:"CommodityCode"`
	MachineType           string                                      `json:"MachineType" xml:"MachineType"`
	InstanceSeries        string                                      `json:"InstanceSeries" xml:"InstanceSeries"`
	ProductVersion        string                                      `json:"ProductVersion" xml:"ProductVersion"`
	RegionId              string                                      `json:"RegionId" xml:"RegionId"`
	ResourceGroupId       string                                      `json:"ResourceGroupId" xml:"ResourceGroupId"`
	DrdsInstanceId        string                                      `json:"DrdsInstanceId" xml:"DrdsInstanceId"`
	ZoneId                string                                      `json:"ZoneId" xml:"ZoneId"`
	InstRole              string                                      `json:"InstRole" xml:"InstRole"`
	OrderInstanceId       string                                      `json:"OrderInstanceId" xml:"OrderInstanceId"`
	ReadOnlyDBInstanceIds ReadOnlyDBInstanceIdsInDescribeDrdsInstance `json:"ReadOnlyDBInstanceIds" xml:"ReadOnlyDBInstanceIds"`
	Vips                  VipsInDescribeDrdsInstance                  `json:"Vips" xml:"Vips"`
}
