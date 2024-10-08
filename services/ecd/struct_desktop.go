package ecd

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

// Desktop is a nested struct in ecd response
type Desktop struct {
	ReleaseNote       string   `json:"ReleaseNote" xml:"ReleaseNote"`
	NewAppSize        int64    `json:"NewAppSize" xml:"NewAppSize"`
	StartTime         string   `json:"StartTime" xml:"StartTime"`
	CurrentAppVersion string   `json:"CurrentAppVersion" xml:"CurrentAppVersion"`
	DesktopId         string   `json:"DesktopId" xml:"DesktopId"`
	NewAppVersion     string   `json:"NewAppVersion" xml:"NewAppVersion"`
	DesktopGroupId    string   `json:"DesktopGroupId" xml:"DesktopGroupId"`
	DesktopStatus     string   `json:"DesktopStatus" xml:"DesktopStatus"`
	ConnectionStatus  string   `json:"ConnectionStatus" xml:"ConnectionStatus"`
	ManagementFlag    []string `json:"ManagementFlag" xml:"ManagementFlag"`
}
