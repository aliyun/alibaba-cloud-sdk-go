package sophonsoar

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

// Action is a nested struct in sophonsoar response
type Action struct {
	TriggerType    string `json:"TriggerType" xml:"TriggerType"`
	TriggerDataId  string `json:"TriggerDataId" xml:"TriggerDataId"`
	Component      string `json:"Component" xml:"Component"`
	Action         string `json:"Action" xml:"Action"`
	StartTime      int64  `json:"StartTime" xml:"StartTime"`
	EndTime        int64  `json:"EndTime" xml:"EndTime"`
	Status         string `json:"Status" xml:"Status"`
	RequestUuid    string `json:"RequestUuid" xml:"RequestUuid"`
	TaskUuid       string `json:"TaskUuid" xml:"TaskUuid"`
	SrcHostName    string `json:"SrcHostName" xml:"SrcHostName"`
	TriggerUser    string `json:"TriggerUser" xml:"TriggerUser"`
	TaskName       string `json:"TaskName" xml:"TaskName"`
	TaskStatus     string `json:"TaskStatus" xml:"TaskStatus"`
	ActionUuid     string `json:"ActionUuid" xml:"ActionUuid"`
	DataSourceName string `json:"DataSourceName" xml:"DataSourceName"`
	AssetName      string `json:"AssetName" xml:"AssetName"`
	NodeName       string `json:"NodeName" xml:"NodeName"`
	EventId        string `json:"EventId" xml:"EventId"`
	NodeId         string `json:"NodeId" xml:"NodeId"`
	NodeType       string `json:"NodeType" xml:"NodeType"`
	TaskTenantId   string `json:"TaskTenantId" xml:"TaskTenantId"`
	TaskType       string `json:"TaskType" xml:"TaskType"`
	TenantId       string `json:"TenantId" xml:"TenantId"`
}
