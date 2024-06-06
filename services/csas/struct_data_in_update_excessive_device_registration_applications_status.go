package csas

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

// DataInUpdateExcessiveDeviceRegistrationApplicationsStatus is a nested struct in csas response
type DataInUpdateExcessiveDeviceRegistrationApplicationsStatus struct {
	DeviceTag     string `json:"DeviceTag" xml:"DeviceTag"`
	DeviceType    string `json:"DeviceType" xml:"DeviceType"`
	Hostname      string `json:"Hostname" xml:"Hostname"`
	Username      string `json:"Username" xml:"Username"`
	SaseUserId    string `json:"SaseUserId" xml:"SaseUserId"`
	Department    string `json:"Department" xml:"Department"`
	Mac           string `json:"Mac" xml:"Mac"`
	IsUsed        bool   `json:"IsUsed" xml:"IsUsed"`
	Status        string `json:"Status" xml:"Status"`
	CreateTime    string `json:"CreateTime" xml:"CreateTime"`
	Description   string `json:"Description" xml:"Description"`
	ApplicationId string `json:"ApplicationId" xml:"ApplicationId"`
}
