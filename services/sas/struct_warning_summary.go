package sas

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

// WarningSummary is a nested struct in sas response
type WarningSummary struct {
	UniqueInfo            string       `json:"UniqueInfo" xml:"UniqueInfo"`
	DataSource            string       `json:"DataSource" xml:"DataSource"`
	ContainerId           string       `json:"ContainerId" xml:"ContainerId"`
	OccurrenceTime        string       `json:"OccurrenceTime" xml:"OccurrenceTime"`
	LastTime              string       `json:"LastTime" xml:"LastTime"`
	HasTraceInfo          bool         `json:"HasTraceInfo" xml:"HasTraceInfo"`
	IntranetIp            string       `json:"IntranetIp" xml:"IntranetIp"`
	K8sClusterName        string       `json:"K8sClusterName" xml:"K8sClusterName"`
	LowWarningCount       int          `json:"LowWarningCount" xml:"LowWarningCount"`
	DisplaySandboxResult  bool         `json:"DisplaySandboxResult" xml:"DisplaySandboxResult"`
	DatabaseRisk          bool         `json:"DatabaseRisk" xml:"DatabaseRisk"`
	K8sClusterId          string       `json:"K8sClusterId" xml:"K8sClusterId"`
	EventStatus           int          `json:"EventStatus" xml:"EventStatus"`
	RiskName              string       `json:"RiskName" xml:"RiskName"`
	InstanceName          string       `json:"InstanceName" xml:"InstanceName"`
	Id                    int64        `json:"Id" xml:"Id"`
	WarningMachineCount   int          `json:"WarningMachineCount" xml:"WarningMachineCount"`
	AlarmEventNameDisplay string       `json:"AlarmEventNameDisplay" xml:"AlarmEventNameDisplay"`
	Advanced              bool         `json:"Advanced" xml:"Advanced"`
	HighWarningCount      int          `json:"HighWarningCount" xml:"HighWarningCount"`
	LastTimeStamp         int64        `json:"LastTimeStamp" xml:"LastTimeStamp"`
	ContainerImageName    string       `json:"ContainerImageName" xml:"ContainerImageName"`
	SubTypeAlias          string       `json:"SubTypeAlias" xml:"SubTypeAlias"`
	TypeAlias             string       `json:"TypeAlias" xml:"TypeAlias"`
	AlarmEventTypeDisplay string       `json:"AlarmEventTypeDisplay" xml:"AlarmEventTypeDisplay"`
	EventSubType          string       `json:"EventSubType" xml:"EventSubType"`
	CanBeDealOnLine       bool         `json:"CanBeDealOnLine" xml:"CanBeDealOnLine"`
	OperateTime           int64        `json:"OperateTime" xml:"OperateTime"`
	MediumWarningCount    int          `json:"MediumWarningCount" xml:"MediumWarningCount"`
	Name                  string       `json:"Name" xml:"Name"`
	SaleVersion           string       `json:"SaleVersion" xml:"SaleVersion"`
	Desc                  string       `json:"Desc" xml:"Desc"`
	K8sNodeName           string       `json:"K8sNodeName" xml:"K8sNodeName"`
	AlarmUniqueInfo       string       `json:"AlarmUniqueInfo" xml:"AlarmUniqueInfo"`
	RiskId                int64        `json:"RiskId" xml:"RiskId"`
	ContainHwMode         bool         `json:"ContainHwMode" xml:"ContainHwMode"`
	AlarmEventType        string       `json:"AlarmEventType" xml:"AlarmEventType"`
	Uuid                  string       `json:"Uuid" xml:"Uuid"`
	SecurityEventIds      string       `json:"SecurityEventIds" xml:"SecurityEventIds"`
	OccurrenceTimeStamp   int64        `json:"OccurrenceTimeStamp" xml:"OccurrenceTimeStamp"`
	MarkMisRules          string       `json:"MarkMisRules" xml:"MarkMisRules"`
	InstanceId            string       `json:"InstanceId" xml:"InstanceId"`
	CheckCount            int          `json:"CheckCount" xml:"CheckCount"`
	LastFoundTime         string       `json:"LastFoundTime" xml:"LastFoundTime"`
	CanCancelFault        bool         `json:"CanCancelFault" xml:"CanCancelFault"`
	OperateMsg            string       `json:"OperateMsg" xml:"OperateMsg"`
	InternetIp            string       `json:"InternetIp" xml:"InternetIp"`
	Level                 string       `json:"Level" xml:"Level"`
	CheckExploit          bool         `json:"CheckExploit" xml:"CheckExploit"`
	AlarmEventName        string       `json:"AlarmEventName" xml:"AlarmEventName"`
	Stages                string       `json:"Stages" xml:"Stages"`
	MaliciousRuleStatus   string       `json:"MaliciousRuleStatus" xml:"MaliciousRuleStatus"`
	ContainerImageId      string       `json:"ContainerImageId" xml:"ContainerImageId"`
	K8sPodName            string       `json:"K8sPodName" xml:"K8sPodName"`
	K8sNamespace          string       `json:"K8sNamespace" xml:"K8sNamespace"`
	K8sNodeId             string       `json:"K8sNodeId" xml:"K8sNodeId"`
	AutoBreaking          bool         `json:"AutoBreaking" xml:"AutoBreaking"`
	AppName               string       `json:"AppName" xml:"AppName"`
	OperateErrorCode      string       `json:"OperateErrorCode" xml:"OperateErrorCode"`
	TacticItems           []TacticItem `json:"TacticItems" xml:"TacticItems"`
	Details               []QuaraFile  `json:"Details" xml:"Details"`
	EventNotes            []EventNote  `json:"EventNotes" xml:"EventNotes"`
}
