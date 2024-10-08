package pvtz

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

// Zone is a nested struct in pvtz response
type Zone struct {
	UpdateTime       string       `json:"UpdateTime" xml:"UpdateTime"`
	ProxyPattern     string       `json:"ProxyPattern" xml:"ProxyPattern"`
	Remark           string       `json:"Remark" xml:"Remark"`
	CreateTimestamp  int64        `json:"CreateTimestamp" xml:"CreateTimestamp"`
	DnsGroup         string       `json:"DnsGroup" xml:"DnsGroup"`
	CreatorType      string       `json:"CreatorType" xml:"CreatorType"`
	CreateTime       string       `json:"CreateTime" xml:"CreateTime"`
	RecordCount      int          `json:"RecordCount" xml:"RecordCount"`
	CreatorSubType   string       `json:"CreatorSubType" xml:"CreatorSubType"`
	UpdateTimestamp  int64        `json:"UpdateTimestamp" xml:"UpdateTimestamp"`
	DnsGroupChanging bool         `json:"DnsGroupChanging" xml:"DnsGroupChanging"`
	ResourceGroupId  string       `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ZoneId           string       `json:"ZoneId" xml:"ZoneId"`
	ZoneTag          string       `json:"ZoneTag" xml:"ZoneTag"`
	Creator          string       `json:"Creator" xml:"Creator"`
	ZoneName         string       `json:"ZoneName" xml:"ZoneName"`
	IsPtr            bool         `json:"IsPtr" xml:"IsPtr"`
	ZoneType         string       `json:"ZoneType" xml:"ZoneType"`
	Vpcs             Vpcs         `json:"Vpcs" xml:"Vpcs"`
	ResourceTags     ResourceTags `json:"ResourceTags" xml:"ResourceTags"`
}
