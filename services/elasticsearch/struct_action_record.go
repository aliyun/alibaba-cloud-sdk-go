package elasticsearch

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

// ActionRecord is a nested struct in elasticsearch response
type ActionRecord struct {
	ActionName             string                 `json:"ActionName" xml:"ActionName"`
	ActionParams           map[string]interface{} `json:"ActionParams" xml:"ActionParams"`
	EndTime                int64                  `json:"EndTime" xml:"EndTime"`
	InstanceId             string                 `json:"InstanceId" xml:"InstanceId"`
	Process                string                 `json:"Process" xml:"Process"`
	RecordDiff             map[string]interface{} `json:"RecordDiff" xml:"RecordDiff"`
	RequestId              string                 `json:"RequestId" xml:"RequestId"`
	StartTime              int64                  `json:"StartTime" xml:"StartTime"`
	StateType              string                 `json:"StateType" xml:"StateType"`
	UserId                 string                 `json:"UserId" xml:"UserId"`
	UserType               string                 `json:"UserType" xml:"UserType"`
	OwnerId                string                 `json:"OwnerId" xml:"OwnerId"`
	UserInfo               string                 `json:"UserInfo" xml:"UserInfo"`
	MetaNow                string                 `json:"MetaNow" xml:"MetaNow"`
	MetaOld                string                 `json:"MetaOld" xml:"MetaOld"`
	ActionResultAccessList []string               `json:"ActionResultAccessList" xml:"ActionResultAccessList"`
	RecordIds              []string               `json:"RecordIds" xml:"RecordIds"`
	StatusInfo             []StatusInfoItem       `json:"StatusInfo" xml:"StatusInfo"`
}
