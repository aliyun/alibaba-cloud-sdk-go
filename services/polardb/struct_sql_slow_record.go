package polardb

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

// SQLSlowRecord is a nested struct in polardb response
type SQLSlowRecord struct {
	DBName             string `json:"DBName" xml:"DBName"`
	DBNodeId           string `json:"DBNodeId" xml:"DBNodeId"`
	ExecutionStartTime string `json:"ExecutionStartTime" xml:"ExecutionStartTime"`
	HostAddress        string `json:"HostAddress" xml:"HostAddress"`
	LockTimes          int64  `json:"LockTimes" xml:"LockTimes"`
	ParseRowCounts     int64  `json:"ParseRowCounts" xml:"ParseRowCounts"`
	QueryTimeMS        int64  `json:"QueryTimeMS" xml:"QueryTimeMS"`
	QueryTimes         int64  `json:"QueryTimes" xml:"QueryTimes"`
	ReturnRowCounts    int64  `json:"ReturnRowCounts" xml:"ReturnRowCounts"`
	SQLHash            string `json:"SQLHash" xml:"SQLHash"`
	SQLText            string `json:"SQLText" xml:"SQLText"`
}
