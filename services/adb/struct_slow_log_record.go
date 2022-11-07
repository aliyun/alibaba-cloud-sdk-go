package adb

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

// SlowLogRecord is a nested struct in adb response
type SlowLogRecord struct {
	ScanRows           int64  `json:"ScanRows" xml:"ScanRows"`
	User               string `json:"User" xml:"User"`
	QueueTime          int64  `json:"QueueTime" xml:"QueueTime"`
	QueryTime          int64  `json:"QueryTime" xml:"QueryTime"`
	HostAddress        string `json:"HostAddress" xml:"HostAddress"`
	ScanTime           int64  `json:"ScanTime" xml:"ScanTime"`
	PlanningTime       int64  `json:"PlanningTime" xml:"PlanningTime"`
	ScanSize           string `json:"ScanSize" xml:"ScanSize"`
	ParseRowCounts     int64  `json:"ParseRowCounts" xml:"ParseRowCounts"`
	SQLText            string `json:"SQLText" xml:"SQLText"`
	OutputSize         string `json:"OutputSize" xml:"OutputSize"`
	ExecuteTime        string `json:"ExecuteTime" xml:"ExecuteTime"`
	ExecutionStartTime string `json:"ExecutionStartTime" xml:"ExecutionStartTime"`
	ProcessID          string `json:"ProcessID" xml:"ProcessID"`
	State              string `json:"State" xml:"State"`
	ReturnRowCounts    int64  `json:"ReturnRowCounts" xml:"ReturnRowCounts"`
	PeakMemoryUsage    string `json:"PeakMemoryUsage" xml:"PeakMemoryUsage"`
	TotalTime          string `json:"TotalTime" xml:"TotalTime"`
	SQLType            string `json:"SQLType" xml:"SQLType"`
	DBName             string `json:"DBName" xml:"DBName"`
	UserName           string `json:"UserName" xml:"UserName"`
	WallTime           int64  `json:"WallTime" xml:"WallTime"`
	Succeed            string `json:"Succeed" xml:"Succeed"`
	ConnId             string `json:"ConnId" xml:"ConnId"`
}
