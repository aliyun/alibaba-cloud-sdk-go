package oceanbasepro

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

// DataItemInDescribeOasTopSQLList is a nested struct in oceanbasepro response
type DataItemInDescribeOasTopSQLList struct {
	Executions                  string `json:"Executions" xml:"Executions"`
	RpcCount                    string `json:"RpcCount" xml:"RpcCount"`
	RemotePlans                 string `json:"RemotePlans" xml:"RemotePlans"`
	MissPlans                   string `json:"MissPlans" xml:"MissPlans"`
	MaxElapsedTime              string `json:"MaxElapsedTime" xml:"MaxElapsedTime"`
	TotalWaitTime               string `json:"TotalWaitTime" xml:"TotalWaitTime"`
	ExecPs                      string `json:"ExecPs" xml:"ExecPs"`
	MaxCpuTime                  string `json:"MaxCpuTime" xml:"MaxCpuTime"`
	CpuPercentage               string `json:"CpuPercentage" xml:"CpuPercentage"`
	ClientIp                    string `json:"ClientIp" xml:"ClientIp"`
	UserName                    string `json:"UserName" xml:"UserName"`
	DbName                      string `json:"DbName" xml:"DbName"`
	RetCode4012Count            int64  `json:"RetCode4012Count" xml:"RetCode4012Count"`
	RetCode4013Count            int64  `json:"RetCode4013Count" xml:"RetCode4013Count"`
	RetCode5001Count            int64  `json:"RetCode5001Count" xml:"RetCode5001Count"`
	RetCode5024Count            int64  `json:"RetCode5024Count" xml:"RetCode5024Count"`
	RetCode5167Count            int64  `json:"RetCode5167Count" xml:"RetCode5167Count"`
	RetCode5217Count            int64  `json:"RetCode5217Count" xml:"RetCode5217Count"`
	RetCode6002Count            int64  `json:"RetCode6002Count" xml:"RetCode6002Count"`
	FailPercentage              string `json:"FailPercentage" xml:"FailPercentage"`
	SumWaitTime                 string `json:"SumWaitTime" xml:"SumWaitTime"`
	AvgWaitCount                string `json:"AvgWaitCount" xml:"AvgWaitCount"`
	AvgRpcCount                 string `json:"AvgRpcCount" xml:"AvgRpcCount"`
	LocalPlanPercentage         string `json:"LocalPlanPercentage" xml:"LocalPlanPercentage"`
	RemotePlanPercentage        string `json:"RemotePlanPercentage" xml:"RemotePlanPercentage"`
	DistPlanPercentage          string `json:"DistPlanPercentage" xml:"DistPlanPercentage"`
	SumElapsedTime              string `json:"SumElapsedTime" xml:"SumElapsedTime"`
	AvgNetTime                  string `json:"AvgNetTime" xml:"AvgNetTime"`
	AvgExecutorRpcCount         string `json:"AvgExecutorRpcCount" xml:"AvgExecutorRpcCount"`
	MissPlanPercentage          string `json:"MissPlanPercentage" xml:"MissPlanPercentage"`
	TableScanPercentage         string `json:"TableScanPercentage" xml:"TableScanPercentage"`
	StrongConsistencyPercentage string `json:"StrongConsistencyPercentage" xml:"StrongConsistencyPercentage"`
	WeakConsistencyPercentage   string `json:"WeakConsistencyPercentage" xml:"WeakConsistencyPercentage"`
	MaxAffectedRows             string `json:"MaxAffectedRows" xml:"MaxAffectedRows"`
	MaxReturnRows               string `json:"MaxReturnRows" xml:"MaxReturnRows"`
	MaxWaitTime                 string `json:"MaxWaitTime" xml:"MaxWaitTime"`
	MaxApplicationWaitTime      string `json:"MaxApplicationWaitTime" xml:"MaxApplicationWaitTime"`
	MaxConcurrencyWaitTime      string `json:"MaxConcurrencyWaitTime" xml:"MaxConcurrencyWaitTime"`
	MaxUserIoWaitTime           string `json:"MaxUserIoWaitTime" xml:"MaxUserIoWaitTime"`
	MaxDiskReads                string `json:"MaxDiskReads" xml:"MaxDiskReads"`
	AvgExpectedWorkerCount      string `json:"AvgExpectedWorkerCount" xml:"AvgExpectedWorkerCount"`
	AvgUsedWorkerCount          string `json:"AvgUsedWorkerCount" xml:"AvgUsedWorkerCount"`
	SumLogicalReads             string `json:"SumLogicalReads" xml:"SumLogicalReads"`
	Server                      string `json:"Server" xml:"Server"`
	ServerIp                    string `json:"ServerIp" xml:"ServerIp"`
	ServerPort                  int64  `json:"ServerPort" xml:"ServerPort"`
	SqlTextShort                string `json:"SqlTextShort" xml:"SqlTextShort"`
	SqlType                     string `json:"SqlType" xml:"SqlType"`
	SqlId                       string `json:"SqlId" xml:"SqlId"`
	Inner                       bool   `json:"Inner" xml:"Inner"`
	WaitEvent                   string `json:"WaitEvent" xml:"WaitEvent"`
	AvgAffectedRows             string `json:"AvgAffectedRows" xml:"AvgAffectedRows"`
	AvgReturnRows               string `json:"AvgReturnRows" xml:"AvgReturnRows"`
	AvgPartitionCount           string `json:"AvgPartitionCount" xml:"AvgPartitionCount"`
	FailCount                   string `json:"FailCount" xml:"FailCount"`
	AvgWaitTime                 string `json:"AvgWaitTime" xml:"AvgWaitTime"`
	AvgElapsedTime              string `json:"AvgElapsedTime" xml:"AvgElapsedTime"`
	AvgCpuTime                  string `json:"AvgCpuTime" xml:"AvgCpuTime"`
	AvgNetWaitTime              string `json:"AvgNetWaitTime" xml:"AvgNetWaitTime"`
	AvgQueueTime                string `json:"AvgQueueTime" xml:"AvgQueueTime"`
	AvgDecodeTime               string `json:"AvgDecodeTime" xml:"AvgDecodeTime"`
	AvgGetPlanTime              string `json:"AvgGetPlanTime" xml:"AvgGetPlanTime"`
	AvgExecuteTime              string `json:"AvgExecuteTime" xml:"AvgExecuteTime"`
	AvgApplicationWaitTime      string `json:"AvgApplicationWaitTime" xml:"AvgApplicationWaitTime"`
	AvgConcurrencyWaitTime      string `json:"AvgConcurrencyWaitTime" xml:"AvgConcurrencyWaitTime"`
	AvgUserIoWaitTime           string `json:"AvgUserIoWaitTime" xml:"AvgUserIoWaitTime"`
	AvgScheduleTime             string `json:"AvgScheduleTime" xml:"AvgScheduleTime"`
	AvgRowCacheHit              string `json:"AvgRowCacheHit" xml:"AvgRowCacheHit"`
	AvgBloomFilterCacheHit      string `json:"AvgBloomFilterCacheHit" xml:"AvgBloomFilterCacheHit"`
	AvgBlockCacheHit            string `json:"AvgBlockCacheHit" xml:"AvgBlockCacheHit"`
	AvgBlockIndexCacheHit       string `json:"AvgBlockIndexCacheHit" xml:"AvgBlockIndexCacheHit"`
	AvgDiskReads                string `json:"AvgDiskReads" xml:"AvgDiskReads"`
	RetryCount                  string `json:"RetryCount" xml:"RetryCount"`
	AvgMemstoreReadRows         string `json:"AvgMemstoreReadRows" xml:"AvgMemstoreReadRows"`
	AvgSsstoreReadRows          string `json:"AvgSsstoreReadRows" xml:"AvgSsstoreReadRows"`
	AvgLogicalReads             string `json:"AvgLogicalReads" xml:"AvgLogicalReads"`
}
