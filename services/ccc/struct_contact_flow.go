package ccc

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

// ContactFlow is a nested struct in ccc response
type ContactFlow struct {
	ContactFlowId string   `json:"ContactFlowId" xml:"ContactFlowId"`
	CreatedTime   string   `json:"CreatedTime" xml:"CreatedTime"`
	Definition    string   `json:"Definition" xml:"Definition"`
	Description   string   `json:"Description" xml:"Description"`
	DraftId       string   `json:"DraftId" xml:"DraftId"`
	Editor        string   `json:"Editor" xml:"Editor"`
	InstanceId    string   `json:"InstanceId" xml:"InstanceId"`
	Name          string   `json:"Name" xml:"Name"`
	Published     bool     `json:"Published" xml:"Published"`
	Type          string   `json:"Type" xml:"Type"`
	UpdatedTime   string   `json:"UpdatedTime" xml:"UpdatedTime"`
	NumberList    []string `json:"NumberList" xml:"NumberList"`
}
