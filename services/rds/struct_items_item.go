package rds

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

// ItemsItem is a nested struct in rds response
type ItemsItem struct {
	Uid                   string                `json:"Uid" xml:"Uid"`
	CallerSource          string                `json:"CallerSource" xml:"CallerSource"`
	Status                int                   `json:"Status" xml:"Status"`
	Product               string                `json:"Product" xml:"Product"`
	CallerUid             string                `json:"CallerUid" xml:"CallerUid"`
	InstanceName          string                `json:"InstanceName" xml:"InstanceName"`
	ActionInfo            string                `json:"ActionInfo" xml:"ActionInfo"`
	InstanceType          string                `json:"InstanceType" xml:"InstanceType"`
	ReasonCode            string                `json:"ReasonCode" xml:"ReasonCode"`
	DbType                string                `json:"DbType" xml:"DbType"`
	TaskId                string                `json:"TaskId" xml:"TaskId"`
	StartTime             string                `json:"StartTime" xml:"StartTime"`
	RegionId              string                `json:"RegionId" xml:"RegionId"`
	EndTime               string                `json:"EndTime" xml:"EndTime"`
	RemainTime            int                   `json:"RemainTime" xml:"RemainTime"`
	InstanceId            string                `json:"InstanceId" xml:"InstanceId"`
	CurrentStepName       string                `json:"CurrentStepName" xml:"CurrentStepName"`
	Progress              float64               `json:"Progress" xml:"Progress"`
	DBInstanceId          string                `json:"DBInstanceId" xml:"DBInstanceId"`
	TaskType              string                `json:"TaskType" xml:"TaskType"`
	TaskDetail            string                `json:"TaskDetail" xml:"TaskDetail"`
	ReadDBInstanceNames   []string              `json:"ReadDBInstanceNames" xml:"ReadDBInstanceNames"`
	ReadDelayTimes        []string              `json:"ReadDelayTimes" xml:"ReadDelayTimes"`
	ReadonlyInstanceDelay ReadonlyInstanceDelay `json:"ReadonlyInstanceDelay" xml:"ReadonlyInstanceDelay"`
}
