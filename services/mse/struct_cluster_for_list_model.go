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

// ClusterForListModel is a nested struct in mse response
type ClusterForListModel struct {
	AppVersion       string `json:"AppVersion" xml:"AppVersion"`
	ClusterAliasName string `json:"ClusterAliasName" xml:"ClusterAliasName"`
	InternetAddress  string `json:"InternetAddress" xml:"InternetAddress"`
	IntranetAddress  string `json:"IntranetAddress" xml:"IntranetAddress"`
	ClusterId        string `json:"ClusterId" xml:"ClusterId"`
	InstanceId       string `json:"InstanceId" xml:"InstanceId"`
	ChargeType       string `json:"ChargeType" xml:"ChargeType"`
	InternetDomain   string `json:"InternetDomain" xml:"InternetDomain"`
	CreateTime       string `json:"CreateTime" xml:"CreateTime"`
	InitStatus       string `json:"InitStatus" xml:"InitStatus"`
	ClusterType      string `json:"ClusterType" xml:"ClusterType"`
	IntranetDomain   string `json:"IntranetDomain" xml:"IntranetDomain"`
	EndDate          string `json:"EndDate" xml:"EndDate"`
}
