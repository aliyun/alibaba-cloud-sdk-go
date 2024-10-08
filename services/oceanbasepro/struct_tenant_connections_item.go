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

// TenantConnectionsItem is a nested struct in oceanbasepro response
type TenantConnectionsItem struct {
	IntranetAddress             string   `json:"IntranetAddress" xml:"IntranetAddress"`
	IntranetPort                int      `json:"IntranetPort" xml:"IntranetPort"`
	InternetAddress             string   `json:"InternetAddress" xml:"InternetAddress"`
	InternetPort                int      `json:"InternetPort" xml:"InternetPort"`
	VpcId                       string   `json:"VpcId" xml:"VpcId"`
	VSwitchId                   string   `json:"VSwitchId" xml:"VSwitchId"`
	IntranetAddressMasterZoneId string   `json:"IntranetAddressMasterZoneId" xml:"IntranetAddressMasterZoneId"`
	IntranetAddressSlaveZoneId  string   `json:"IntranetAddressSlaveZoneId" xml:"IntranetAddressSlaveZoneId"`
	IntranetAddressStatus       string   `json:"IntranetAddressStatus" xml:"IntranetAddressStatus"`
	InternetAddressStatus       string   `json:"InternetAddressStatus" xml:"InternetAddressStatus"`
	TransactionSplit            bool     `json:"TransactionSplit" xml:"TransactionSplit"`
	AddressType                 string   `json:"AddressType" xml:"AddressType"`
	EnableTransactionSplit      bool     `json:"EnableTransactionSplit" xml:"EnableTransactionSplit"`
	ParallelQueryDegree         int64    `json:"ParallelQueryDegree" xml:"ParallelQueryDegree"`
	TenantEndpointId            string   `json:"TenantEndpointId" xml:"TenantEndpointId"`
	MaxConnectionNum            int64    `json:"MaxConnectionNum" xml:"MaxConnectionNum"`
	ConnectionReplicaType       string   `json:"ConnectionReplicaType" xml:"ConnectionReplicaType"`
	ProxyClusterId              string   `json:"ProxyClusterId" xml:"ProxyClusterId"`
	MaxConnectionLimit          int64    `json:"MaxConnectionLimit" xml:"MaxConnectionLimit"`
	InternetMaxConnectionLimit  int64    `json:"InternetMaxConnectionLimit" xml:"InternetMaxConnectionLimit"`
	IntranetRpcPort             int      `json:"IntranetRpcPort" xml:"IntranetRpcPort"`
	InternetMaxConnectionNum    int64    `json:"InternetMaxConnectionNum" xml:"InternetMaxConnectionNum"`
	InternetRpcPort             int      `json:"InternetRpcPort" xml:"InternetRpcPort"`
	IntranetSqlPort             int      `json:"IntranetSqlPort" xml:"IntranetSqlPort"`
	OdpVersion                  string   `json:"OdpVersion" xml:"OdpVersion"`
	ConnectionZones             []string `json:"ConnectionZones" xml:"ConnectionZones"`
	ConnectionLogicalZones      []string `json:"ConnectionLogicalZones" xml:"ConnectionLogicalZones"`
}
