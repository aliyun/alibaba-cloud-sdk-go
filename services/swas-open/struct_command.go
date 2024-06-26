package swas_open

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

// Command is a nested struct in swas_open response
type Command struct {
	Name                 string                `json:"Name" xml:"Name"`
	CommandId            string                `json:"CommandId" xml:"CommandId"`
	Timeout              int64                 `json:"Timeout" xml:"Timeout"`
	Username             string                `json:"Username" xml:"Username"`
	WorkingDir           string                `json:"WorkingDir" xml:"WorkingDir"`
	CreationTime         string                `json:"CreationTime" xml:"CreationTime"`
	EnableParameter      bool                  `json:"EnableParameter" xml:"EnableParameter"`
	Provider             string                `json:"Provider" xml:"Provider"`
	CommandName          string                `json:"CommandName" xml:"CommandName"`
	Parameters           string                `json:"Parameters" xml:"Parameters"`
	ResourceGroupId      string                `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InvokeId             string                `json:"InvokeId" xml:"InvokeId"`
	CommandContent       string                `json:"CommandContent" xml:"CommandContent"`
	InvocationStatus     string                `json:"InvocationStatus" xml:"InvocationStatus"`
	Description          string                `json:"Description" xml:"Description"`
	CommandDescription   string                `json:"CommandDescription" xml:"CommandDescription"`
	CommandType          string                `json:"CommandType" xml:"CommandType"`
	Type                 string                `json:"Type" xml:"Type"`
	ParameterNames       []string              `json:"ParameterNames" xml:"ParameterNames"`
	ParameterDefinitions []ParameterDefinition `json:"ParameterDefinitions" xml:"ParameterDefinitions"`
	InvokeInstances      []InvokeInstance      `json:"InvokeInstances" xml:"InvokeInstances"`
	Tags                 []Tag                 `json:"Tags" xml:"Tags"`
}
