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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDBClusterAttribute invokes the polardb.DescribeDBClusterAttribute API synchronously
func (client *Client) DescribeDBClusterAttribute(request *DescribeDBClusterAttributeRequest) (response *DescribeDBClusterAttributeResponse, err error) {
	response = CreateDescribeDBClusterAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDBClusterAttributeWithChan invokes the polardb.DescribeDBClusterAttribute API asynchronously
func (client *Client) DescribeDBClusterAttributeWithChan(request *DescribeDBClusterAttributeRequest) (<-chan *DescribeDBClusterAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDBClusterAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBClusterAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDBClusterAttributeWithCallback invokes the polardb.DescribeDBClusterAttribute API asynchronously
func (client *Client) DescribeDBClusterAttributeWithCallback(request *DescribeDBClusterAttributeRequest, callback func(response *DescribeDBClusterAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBClusterAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBClusterAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDBClusterAttributeRequest is the request struct for api DescribeDBClusterAttribute
type DescribeDBClusterAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DescribeType         string           `position:"Query" name:"DescribeType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId          string           `position:"Query" name:"DBClusterId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDBClusterAttributeResponse is the response struct for api DescribeDBClusterAttribute
type DescribeDBClusterAttributeResponse struct {
	*responses.BaseResponse
	DeletionLock                 int               `json:"DeletionLock" xml:"DeletionLock"`
	Category                     string            `json:"Category" xml:"Category"`
	ResourceGroupId              string            `json:"ResourceGroupId" xml:"ResourceGroupId"`
	DataLevel1BackupChainSize    int64             `json:"DataLevel1BackupChainSize" xml:"DataLevel1BackupChainSize"`
	DBClusterId                  string            `json:"DBClusterId" xml:"DBClusterId"`
	DBType                       string            `json:"DBType" xml:"DBType"`
	DBClusterNetworkType         string            `json:"DBClusterNetworkType" xml:"DBClusterNetworkType"`
	IsLatestVersion              bool              `json:"IsLatestVersion" xml:"IsLatestVersion"`
	HasCompleteStandbyRes        bool              `json:"HasCompleteStandbyRes" xml:"HasCompleteStandbyRes"`
	HotStandbyClusterStatus      string            `json:"HotStandbyClusterStatus" xml:"HotStandbyClusterStatus"`
	HotStandbyCluster            string            `json:"HotStandbyCluster" xml:"HotStandbyCluster"`
	DataSyncMode                 string            `json:"DataSyncMode" xml:"DataSyncMode"`
	StandbyHAMode                string            `json:"StandbyHAMode" xml:"StandbyHAMode"`
	CompressStorageMode          string            `json:"CompressStorageMode" xml:"CompressStorageMode"`
	StorageMax                   int64             `json:"StorageMax" xml:"StorageMax"`
	DBVersion                    string            `json:"DBVersion" xml:"DBVersion"`
	ZoneIds                      string            `json:"ZoneIds" xml:"ZoneIds"`
	MaintainTime                 string            `json:"MaintainTime" xml:"MaintainTime"`
	Engine                       string            `json:"Engine" xml:"Engine"`
	RequestId                    string            `json:"RequestId" xml:"RequestId"`
	VPCId                        string            `json:"VPCId" xml:"VPCId"`
	DBClusterStatus              string            `json:"DBClusterStatus" xml:"DBClusterStatus"`
	VSwitchId                    string            `json:"VSwitchId" xml:"VSwitchId"`
	DBClusterDescription         string            `json:"DBClusterDescription" xml:"DBClusterDescription"`
	Expired                      string            `json:"Expired" xml:"Expired"`
	PayType                      string            `json:"PayType" xml:"PayType"`
	StoragePayType               string            `json:"StoragePayType" xml:"StoragePayType"`
	LockMode                     string            `json:"LockMode" xml:"LockMode"`
	StorageUsed                  int64             `json:"StorageUsed" xml:"StorageUsed"`
	CompressStorageUsed          int64             `json:"CompressStorageUsed" xml:"CompressStorageUsed"`
	StorageSpace                 int64             `json:"StorageSpace" xml:"StorageSpace"`
	DBVersionStatus              string            `json:"DBVersionStatus" xml:"DBVersionStatus"`
	CreationTime                 string            `json:"CreationTime" xml:"CreationTime"`
	SQLSize                      int64             `json:"SQLSize" xml:"SQLSize"`
	InodeTotal                   int64             `json:"InodeTotal" xml:"InodeTotal"`
	InodeUsed                    int64             `json:"InodeUsed" xml:"InodeUsed"`
	BlktagTotal                  int64             `json:"BlktagTotal" xml:"BlktagTotal"`
	BlktagUsed                   int64             `json:"BlktagUsed" xml:"BlktagUsed"`
	RegionId                     string            `json:"RegionId" xml:"RegionId"`
	ExpireTime                   string            `json:"ExpireTime" xml:"ExpireTime"`
	SubCategory                  string            `json:"SubCategory" xml:"SubCategory"`
	DeployUnit                   string            `json:"DeployUnit" xml:"DeployUnit"`
	IsProxyLatestVersion         bool              `json:"IsProxyLatestVersion" xml:"IsProxyLatestVersion"`
	StorageType                  string            `json:"StorageType" xml:"StorageType"`
	ServerlessType               string            `json:"ServerlessType" xml:"ServerlessType"`
	StrictConsistency            string            `json:"StrictConsistency" xml:"StrictConsistency"`
	ProxyCpuCores                string            `json:"ProxyCpuCores" xml:"ProxyCpuCores"`
	ProxyStandardCpuCores        string            `json:"ProxyStandardCpuCores" xml:"ProxyStandardCpuCores"`
	ProxyType                    string            `json:"ProxyType" xml:"ProxyType"`
	ProxyStatus                  string            `json:"ProxyStatus" xml:"ProxyStatus"`
	FeatureHTAPSupported         string            `json:"FeatureHTAPSupported" xml:"FeatureHTAPSupported"`
	ProxyServerlessType          string            `json:"ProxyServerlessType" xml:"ProxyServerlessType"`
	Architecture                 string            `json:"Architecture" xml:"Architecture"`
	AiType                       string            `json:"AiType" xml:"AiType"`
	ProvisionedIops              string            `json:"ProvisionedIops" xml:"ProvisionedIops"`
	HotStandbyHealthy            bool              `json:"HotStandbyHealthy" xml:"HotStandbyHealthy"`
	HotStandbyWhiteListSwitch    bool              `json:"HotStandbyWhiteListSwitch" xml:"HotStandbyWhiteListSwitch"`
	StorageTypeWhiteListSwitch   bool              `json:"StorageTypeWhiteListSwitch" xml:"StorageTypeWhiteListSwitch"`
	AiFreeMode                   string            `json:"AiFreeMode" xml:"AiFreeMode"`
	AiCreatingTime               string            `json:"AiCreatingTime" xml:"AiCreatingTime"`
	SupportInstantSwitchWithImci string            `json:"SupportInstantSwitchWithImci" xml:"SupportInstantSwitchWithImci"`
	Orca                         string            `json:"Orca" xml:"Orca"`
	SourceDBCluster              string            `json:"SourceDBCluster" xml:"SourceDBCluster"`
	RestoreType                  string            `json:"RestoreType" xml:"RestoreType"`
	RestoreDataPoint             string            `json:"RestoreDataPoint" xml:"RestoreDataPoint"`
	SourceRegionId               string            `json:"SourceRegionId" xml:"SourceRegionId"`
	ImciAutoIndex                string            `json:"ImciAutoIndex" xml:"ImciAutoIndex"`
	BurstingEnabled              string            `json:"BurstingEnabled" xml:"BurstingEnabled"`
	RowCompression               string            `json:"RowCompression" xml:"RowCompression"`
	ImperceptibleSwitch          string            `json:"ImperceptibleSwitch" xml:"ImperceptibleSwitch"`
	RelatedAPInstance            RelatedAPInstance `json:"RelatedAPInstance" xml:"RelatedAPInstance"`
	DBNodes                      []DBNode          `json:"DBNodes" xml:"DBNodes"`
	Tags                         []Tag             `json:"Tags" xml:"Tags"`
}

// CreateDescribeDBClusterAttributeRequest creates a request to invoke DescribeDBClusterAttribute API
func CreateDescribeDBClusterAttributeRequest() (request *DescribeDBClusterAttributeRequest) {
	request = &DescribeDBClusterAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "DescribeDBClusterAttribute", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeDBClusterAttributeResponse creates a response to parse from DescribeDBClusterAttribute response
func CreateDescribeDBClusterAttributeResponse() (response *DescribeDBClusterAttributeResponse) {
	response = &DescribeDBClusterAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
