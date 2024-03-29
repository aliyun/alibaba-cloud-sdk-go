package oos

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

// Template is a nested struct in oos response
type Template struct {
	TemplateId          string                 `json:"TemplateId" xml:"TemplateId"`
	Category            string                 `json:"Category" xml:"Category"`
	VersionName         string                 `json:"VersionName" xml:"VersionName"`
	TotalExecutionCount int                    `json:"TotalExecutionCount" xml:"TotalExecutionCount"`
	TemplateName        string                 `json:"TemplateName" xml:"TemplateName"`
	Hash                string                 `json:"Hash" xml:"Hash"`
	CreatedBy           string                 `json:"CreatedBy" xml:"CreatedBy"`
	Publisher           string                 `json:"Publisher" xml:"Publisher"`
	TemplateVersion     string                 `json:"TemplateVersion" xml:"TemplateVersion"`
	TemplateFormat      string                 `json:"TemplateFormat" xml:"TemplateFormat"`
	UpdatedBy           string                 `json:"UpdatedBy" xml:"UpdatedBy"`
	ResourceGroupId     string                 `json:"ResourceGroupId" xml:"ResourceGroupId"`
	Constraints         string                 `json:"Constraints" xml:"Constraints"`
	CreatedDate         string                 `json:"CreatedDate" xml:"CreatedDate"`
	HasTrigger          bool                   `json:"HasTrigger" xml:"HasTrigger"`
	IsFavorite          bool                   `json:"IsFavorite" xml:"IsFavorite"`
	TemplateType        string                 `json:"TemplateType" xml:"TemplateType"`
	Description         string                 `json:"Description" xml:"Description"`
	UpdatedDate         string                 `json:"UpdatedDate" xml:"UpdatedDate"`
	Tags                map[string]interface{} `json:"Tags" xml:"Tags"`
	Popularity          int                    `json:"Popularity" xml:"Popularity"`
	ShareType           string                 `json:"ShareType" xml:"ShareType"`
}
