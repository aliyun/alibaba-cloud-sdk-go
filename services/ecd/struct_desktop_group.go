package ecd

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

// DesktopGroup is a nested struct in ecd response
type DesktopGroup struct {
	DesktopGroupName        string               `json:"DesktopGroupName" xml:"DesktopGroupName"`
	ImageId                 string               `json:"ImageId" xml:"ImageId"`
	RatioThreshold          float64              `json:"RatioThreshold" xml:"RatioThreshold"`
	SystemDiskSize          int                  `json:"SystemDiskSize" xml:"SystemDiskSize"`
	MaxDesktopsCount        int                  `json:"MaxDesktopsCount" xml:"MaxDesktopsCount"`
	OsType                  string               `json:"OsType" xml:"OsType"`
	StopDuration            int64                `json:"StopDuration" xml:"StopDuration"`
	Version                 int                  `json:"Version" xml:"Version"`
	Creator                 string               `json:"Creator" xml:"Creator"`
	IdleDisconnectDuration  int64                `json:"IdleDisconnectDuration" xml:"IdleDisconnectDuration"`
	BuyDesktopsCount        int                  `json:"BuyDesktopsCount" xml:"BuyDesktopsCount"`
	DirectoryId             string               `json:"DirectoryId" xml:"DirectoryId"`
	VolumeEncryptionEnabled bool                 `json:"VolumeEncryptionEnabled" xml:"VolumeEncryptionEnabled"`
	PayType                 string               `json:"PayType" xml:"PayType"`
	OwnBundleName           string               `json:"OwnBundleName" xml:"OwnBundleName"`
	LoadPolicy              int64                `json:"LoadPolicy" xml:"LoadPolicy"`
	Memory                  int64                `json:"Memory" xml:"Memory"`
	DesktopGroupId          string               `json:"DesktopGroupId" xml:"DesktopGroupId"`
	OfficeSiteId            string               `json:"OfficeSiteId" xml:"OfficeSiteId"`
	EndUserCount            int                  `json:"EndUserCount" xml:"EndUserCount"`
	ResetType               int64                `json:"ResetType" xml:"ResetType"`
	PolicyGroupName         string               `json:"PolicyGroupName" xml:"PolicyGroupName"`
	DirectoryType           string               `json:"DirectoryType" xml:"DirectoryType"`
	DesktopCount            int                  `json:"DesktopCount" xml:"DesktopCount"`
	GpuCount                float64              `json:"GpuCount" xml:"GpuCount"`
	OwnBundleId             string               `json:"OwnBundleId" xml:"OwnBundleId"`
	Cpu                     int                  `json:"Cpu" xml:"Cpu"`
	DataDiskCategory        string               `json:"DataDiskCategory" xml:"DataDiskCategory"`
	GpuSpec                 string               `json:"GpuSpec" xml:"GpuSpec"`
	Comments                string               `json:"Comments" xml:"Comments"`
	ConnectDuration         int64                `json:"ConnectDuration" xml:"ConnectDuration"`
	CreateTime              string               `json:"CreateTime" xml:"CreateTime"`
	OfficeSiteName          string               `json:"OfficeSiteName" xml:"OfficeSiteName"`
	VolumeEncryptionKey     string               `json:"VolumeEncryptionKey" xml:"VolumeEncryptionKey"`
	BindAmount              int64                `json:"BindAmount" xml:"BindAmount"`
	SubnetId                string               `json:"SubnetId" xml:"SubnetId"`
	MinDesktopsCount        int                  `json:"MinDesktopsCount" xml:"MinDesktopsCount"`
	Status                  int                  `json:"Status" xml:"Status"`
	OwnType                 int64                `json:"OwnType" xml:"OwnType"`
	ExpiredTime             string               `json:"ExpiredTime" xml:"ExpiredTime"`
	KeepDuration            int64                `json:"KeepDuration" xml:"KeepDuration"`
	DesktopType             string               `json:"DesktopType" xml:"DesktopType"`
	DataDiskSize            string               `json:"DataDiskSize" xml:"DataDiskSize"`
	OfficeSiteType          string               `json:"OfficeSiteType" xml:"OfficeSiteType"`
	PolicyGroupId           string               `json:"PolicyGroupId" xml:"PolicyGroupId"`
	ProtocolType            string               `json:"ProtocolType" xml:"ProtocolType"`
	SystemDiskCategory      string               `json:"SystemDiskCategory" xml:"SystemDiskCategory"`
	GpuDriverVersion        string               `json:"GpuDriverVersion" xml:"GpuDriverVersion"`
	Tags                    []Tag                `json:"Tags" xml:"Tags"`
	CountPerStatus          []CountPerStatusItem `json:"CountPerStatus" xml:"CountPerStatus"`
}
