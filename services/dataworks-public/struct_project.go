package dataworks_public

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

// Project is a nested struct in dataworks_public response
type Project struct {
	ClusterId                      string `json:"ClusterId" xml:"ClusterId"`
	ProjectOwnerBaseId             string `json:"ProjectOwnerBaseId" xml:"ProjectOwnerBaseId"`
	ProjectIdentifier              string `json:"ProjectIdentifier" xml:"ProjectIdentifier"`
	DisableDevelopment             bool   `json:"DisableDevelopment" xml:"DisableDevelopment"`
	ProjectStatusCode              string `json:"ProjectStatusCode" xml:"ProjectStatusCode"`
	DbType                         string `json:"DbType" xml:"DbType"`
	ResourceManagerResourceGroupId string `json:"ResourceManagerResourceGroupId" xml:"ResourceManagerResourceGroupId"`
	IsDefault                      int    `json:"IsDefault" xml:"IsDefault"`
	ProjectName                    string `json:"ProjectName" xml:"ProjectName"`
	ProjectDescription             string `json:"ProjectDescription" xml:"ProjectDescription"`
	ProjectStatus                  int    `json:"ProjectStatus" xml:"ProjectStatus"`
	UseProxyOdpsAccount            bool   `json:"UseProxyOdpsAccount" xml:"UseProxyOdpsAccount"`
	ProjectId                      int64  `json:"ProjectId" xml:"ProjectId"`
	TablePrivacyMode               int    `json:"TablePrivacyMode" xml:"TablePrivacyMode"`
	Tags                           []Tag  `json:"Tags" xml:"Tags"`
}
