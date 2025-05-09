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

// CreateDBCluster invokes the polardb.CreateDBCluster API synchronously
func (client *Client) CreateDBCluster(request *CreateDBClusterRequest) (response *CreateDBClusterResponse, err error) {
	response = CreateCreateDBClusterResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBClusterWithChan invokes the polardb.CreateDBCluster API asynchronously
func (client *Client) CreateDBClusterWithChan(request *CreateDBClusterRequest) (<-chan *CreateDBClusterResponse, <-chan error) {
	responseChan := make(chan *CreateDBClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBCluster(request)
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

// CreateDBClusterWithCallback invokes the polardb.CreateDBCluster API asynchronously
func (client *Client) CreateDBClusterWithCallback(request *CreateDBClusterRequest, callback func(response *CreateDBClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBClusterResponse
		var err error
		defer close(result)
		response, err = client.CreateDBCluster(request)
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

// CreateDBClusterRequest is the request struct for api CreateDBCluster
type CreateDBClusterRequest struct {
	*requests.RpcRequest
	ResourceOwnerId                        requests.Integer      `position:"Query" name:"ResourceOwnerId"`
	DBClusterDescription                   string                `position:"Query" name:"DBClusterDescription"`
	ProxyClass                             string                `position:"Query" name:"ProxyClass"`
	ProxyType                              string                `position:"Query" name:"ProxyType"`
	ScaleMax                               string                `position:"Query" name:"ScaleMax"`
	StorageType                            string                `position:"Query" name:"StorageType"`
	CreationCategory                       string                `position:"Query" name:"CreationCategory"`
	ResourceGroupId                        string                `position:"Query" name:"ResourceGroupId"`
	DBNodeClass                            string                `position:"Query" name:"DBNodeClass"`
	CreationOption                         string                `position:"Query" name:"CreationOption"`
	Tag                                    *[]CreateDBClusterTag `position:"Query" name:"Tag"  type:"Repeated"`
	SourceResourceId                       string                `position:"Query" name:"SourceResourceId"`
	ScaleMin                               string                `position:"Query" name:"ScaleMin"`
	BackupRetentionPolicyOnClusterDeletion string                `position:"Query" name:"BackupRetentionPolicyOnClusterDeletion"`
	Period                                 string                `position:"Query" name:"Period"`
	OwnerId                                requests.Integer      `position:"Query" name:"OwnerId"`
	VSwitchId                              string                `position:"Query" name:"VSwitchId"`
	SecurityIPList                         string                `position:"Query" name:"SecurityIPList"`
	DBMinorVersion                         string                `position:"Query" name:"DBMinorVersion"`
	ProvisionedIops                        requests.Integer      `position:"Query" name:"ProvisionedIops"`
	AutoRenew                              requests.Boolean      `position:"Query" name:"AutoRenew"`
	HotStandbyCluster                      string                `position:"Query" name:"HotStandbyCluster"`
	StoragePayType                         string                `position:"Query" name:"StoragePayType"`
	ZoneId                                 string                `position:"Query" name:"ZoneId"`
	StorageAutoScale                       string                `position:"Query" name:"StorageAutoScale"`
	TDEStatus                              requests.Boolean      `position:"Query" name:"TDEStatus"`
	AllowShutDown                          string                `position:"Query" name:"AllowShutDown"`
	LowerCaseTableNames                    string                `position:"Query" name:"LowerCaseTableNames"`
	StorageEncryption                      requests.Boolean      `position:"Query" name:"StorageEncryption"`
	ScaleRoNumMax                          string                `position:"Query" name:"ScaleRoNumMax"`
	StandbyAZ                              string                `position:"Query" name:"StandbyAZ"`
	ClientToken                            string                `position:"Query" name:"ClientToken"`
	DefaultTimeZone                        string                `position:"Query" name:"DefaultTimeZone"`
	ClusterNetworkType                     string                `position:"Query" name:"ClusterNetworkType"`
	StorageEncryptionKey                   string                `position:"Query" name:"StorageEncryptionKey"`
	ParameterGroupId                       string                `position:"Query" name:"ParameterGroupId"`
	Engine                                 string                `position:"Query" name:"Engine"`
	GDNId                                  string                `position:"Query" name:"GDNId"`
	LooseXEngine                           string                `position:"Query" name:"LooseXEngine"`
	LoosePolarLogBin                       string                `position:"Query" name:"LoosePolarLogBin"`
	Architecture                           string                `position:"Query" name:"Architecture"`
	ResourceOwnerAccount                   string                `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount                           string                `position:"Query" name:"OwnerAccount"`
	LooseXEngineUseMemoryPct               string                `position:"Query" name:"LooseXEngineUseMemoryPct"`
	UsedTime                               string                `position:"Query" name:"UsedTime"`
	BurstingEnabled                        string                `position:"Query" name:"BurstingEnabled"`
	TargetMinorVersion                     string                `position:"Query" name:"TargetMinorVersion"`
	DBNodeNum                              requests.Integer      `position:"Query" name:"DBNodeNum"`
	StorageUpperBound                      requests.Integer      `position:"Query" name:"StorageUpperBound"`
	VPCId                                  string                `position:"Query" name:"VPCId"`
	ScaleRoNumMin                          string                `position:"Query" name:"ScaleRoNumMin"`
	DBType                                 string                `position:"Query" name:"DBType"`
	DBVersion                              string                `position:"Query" name:"DBVersion"`
	StrictConsistency                      string                `position:"Query" name:"StrictConsistency"`
	CloneDataPoint                         string                `position:"Query" name:"CloneDataPoint"`
	PayType                                string                `position:"Query" name:"PayType"`
	StorageSpace                           requests.Integer      `position:"Query" name:"StorageSpace"`
	ServerlessType                         string                `position:"Query" name:"ServerlessType"`
}

// CreateDBClusterTag is a repeated param struct in CreateDBClusterRequest
type CreateDBClusterTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// CreateDBClusterResponse is the response struct for api CreateDBCluster
type CreateDBClusterResponse struct {
	*responses.BaseResponse
	DBClusterId     string `json:"DBClusterId" xml:"DBClusterId"`
	OrderId         string `json:"OrderId" xml:"OrderId"`
	RequestId       string `json:"RequestId" xml:"RequestId"`
	ResourceGroupId string `json:"ResourceGroupId" xml:"ResourceGroupId"`
}

// CreateCreateDBClusterRequest creates a request to invoke CreateDBCluster API
func CreateCreateDBClusterRequest() (request *CreateDBClusterRequest) {
	request = &CreateDBClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("polardb", "2017-08-01", "CreateDBCluster", "polardb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateDBClusterResponse creates a response to parse from CreateDBCluster response
func CreateCreateDBClusterResponse() (response *CreateDBClusterResponse) {
	response = &CreateDBClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
