package resourcesharing

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

// ResourceShare is a nested struct in resourcesharing response
type ResourceShare struct {
	ResourceShareId      string `json:"ResourceShareId" xml:"ResourceShareId"`
	AllowExternalTargets bool   `json:"AllowExternalTargets" xml:"AllowExternalTargets"`
	UpdateTime           string `json:"UpdateTime" xml:"UpdateTime"`
	ResourceGroupId      string `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ResourceShareOwner   string `json:"ResourceShareOwner" xml:"ResourceShareOwner"`
	CreateTime           string `json:"CreateTime" xml:"CreateTime"`
	ResourceShareStatus  string `json:"ResourceShareStatus" xml:"ResourceShareStatus"`
	ResourceShareName    string `json:"ResourceShareName" xml:"ResourceShareName"`
}
