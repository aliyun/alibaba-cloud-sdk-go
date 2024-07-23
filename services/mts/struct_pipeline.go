package mts

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

// Pipeline is a nested struct in mts response
type Pipeline struct {
	SpeedLevel    int64        `json:"SpeedLevel" xml:"SpeedLevel"`
	Name          string       `json:"Name" xml:"Name"`
	State         string       `json:"State" xml:"State"`
	Role          string       `json:"Role" xml:"Role"`
	Speed         string       `json:"Speed" xml:"Speed"`
	QuotaAllocate int64        `json:"QuotaAllocate" xml:"QuotaAllocate"`
	CreationTime  string       `json:"CreationTime" xml:"CreationTime"`
	Id            string       `json:"Id" xml:"Id"`
	ExtendConfig  ExtendConfig `json:"ExtendConfig" xml:"ExtendConfig"`
	NotifyConfig  NotifyConfig `json:"NotifyConfig" xml:"NotifyConfig"`
}
