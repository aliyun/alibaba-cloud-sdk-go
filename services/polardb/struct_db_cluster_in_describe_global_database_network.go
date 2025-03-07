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

// DBClusterInDescribeGlobalDatabaseNetwork is a nested struct in polardb response
type DBClusterInDescribeGlobalDatabaseNetwork struct {
	ReplicaLag           string   `json:"ReplicaLag" xml:"ReplicaLag"`
	ExpireTime           string   `json:"ExpireTime" xml:"ExpireTime"`
	Expired              string   `json:"Expired" xml:"Expired"`
	DBNodeClass          string   `json:"DBNodeClass" xml:"DBNodeClass"`
	PayType              string   `json:"PayType" xml:"PayType"`
	DBType               string   `json:"DBType" xml:"DBType"`
	RegionId             string   `json:"RegionId" xml:"RegionId"`
	DBVersion            string   `json:"DBVersion" xml:"DBVersion"`
	DBClusterId          string   `json:"DBClusterId" xml:"DBClusterId"`
	DBClusterStatus      string   `json:"DBClusterStatus" xml:"DBClusterStatus"`
	StorageUsed          string   `json:"StorageUsed" xml:"StorageUsed"`
	DBClusterDescription string   `json:"DBClusterDescription" xml:"DBClusterDescription"`
	Category             string   `json:"Category" xml:"Category"`
	Role                 string   `json:"Role" xml:"Role"`
	ServerlessType       string   `json:"ServerlessType" xml:"ServerlessType"`
	MemberStatus         string   `json:"MemberStatus" xml:"MemberStatus"`
	DBNodes              []DBNode `json:"DBNodes" xml:"DBNodes"`
}
