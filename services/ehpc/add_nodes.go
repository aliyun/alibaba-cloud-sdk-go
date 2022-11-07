package ehpc

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

// AddNodes invokes the ehpc.AddNodes API synchronously
func (client *Client) AddNodes(request *AddNodesRequest) (response *AddNodesResponse, err error) {
	response = CreateAddNodesResponse()
	err = client.DoAction(request, response)
	return
}

// AddNodesWithChan invokes the ehpc.AddNodes API asynchronously
func (client *Client) AddNodesWithChan(request *AddNodesRequest) (<-chan *AddNodesResponse, <-chan error) {
	responseChan := make(chan *AddNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddNodes(request)
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

// AddNodesWithCallback invokes the ehpc.AddNodes API asynchronously
func (client *Client) AddNodesWithCallback(request *AddNodesRequest, callback func(response *AddNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddNodesResponse
		var err error
		defer close(result)
		response, err = client.AddNodes(request)
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

// AddNodesRequest is the request struct for api AddNodes
type AddNodesRequest struct {
	*requests.RpcRequest
	ImageId                 string               `position:"Query" name:"ImageId"`
	SystemDiskLevel         string               `position:"Query" name:"SystemDiskLevel"`
	ClientToken             string               `position:"Query" name:"ClientToken"`
	AllocatePublicAddress   requests.Boolean     `position:"Query" name:"AllocatePublicAddress"`
	InternetMaxBandWidthOut requests.Integer     `position:"Query" name:"InternetMaxBandWidthOut"`
	JobQueue                string               `position:"Query" name:"JobQueue"`
	ImageOwnerAlias         string               `position:"Query" name:"ImageOwnerAlias"`
	SystemDiskType          string               `position:"Query" name:"SystemDiskType"`
	DataDisks               *[]AddNodesDataDisks `position:"Query" name:"DataDisks"  type:"Repeated"`
	MinCount                requests.Integer     `position:"Query" name:"MinCount"`
	SystemDiskSize          requests.Integer     `position:"Query" name:"SystemDiskSize"`
	InstanceType            string               `position:"Query" name:"InstanceType"`
	HostNamePrefix          string               `position:"Query" name:"HostNamePrefix"`
	ComputeSpotPriceLimit   string               `position:"Query" name:"ComputeSpotPriceLimit"`
	AutoRenewPeriod         requests.Integer     `position:"Query" name:"AutoRenewPeriod"`
	Period                  requests.Integer     `position:"Query" name:"Period"`
	Count                   requests.Integer     `position:"Query" name:"Count"`
	ClusterId               string               `position:"Query" name:"ClusterId"`
	ComputeSpotStrategy     string               `position:"Query" name:"ComputeSpotStrategy"`
	HostNameSuffix          string               `position:"Query" name:"HostNameSuffix"`
	Sync                    requests.Boolean     `position:"Query" name:"Sync"`
	VSwitchId               string               `position:"Query" name:"VSwitchId"`
	PeriodUnit              string               `position:"Query" name:"PeriodUnit"`
	ComputeEnableHt         requests.Boolean     `position:"Query" name:"ComputeEnableHt"`
	AutoRenew               string               `position:"Query" name:"AutoRenew"`
	EcsChargeType           string               `position:"Query" name:"EcsChargeType"`
	InternetChargeType      string               `position:"Query" name:"InternetChargeType"`
	CreateMode              string               `position:"Query" name:"CreateMode"`
	ZoneId                  string               `position:"Query" name:"ZoneId"`
	InternetMaxBandWidthIn  requests.Integer     `position:"Query" name:"InternetMaxBandWidthIn"`
}

// AddNodesDataDisks is a repeated param struct in AddNodesRequest
type AddNodesDataDisks struct {
	DataDiskDeleteWithInstance string `name:"DataDiskDeleteWithInstance"`
	DataDiskEncrypted          string `name:"DataDiskEncrypted"`
	DataDiskKMSKeyId           string `name:"DataDiskKMSKeyId"`
	DataDiskSize               string `name:"DataDiskSize"`
	DataDiskCategory           string `name:"DataDiskCategory"`
	DataDiskPerformanceLevel   string `name:"DataDiskPerformanceLevel"`
}

// AddNodesResponse is the response struct for api AddNodes
type AddNodesResponse struct {
	*responses.BaseResponse
	TaskId      string                `json:"TaskId" xml:"TaskId"`
	RequestId   string                `json:"RequestId" xml:"RequestId"`
	InstanceIds InstanceIdsInAddNodes `json:"InstanceIds" xml:"InstanceIds"`
}

// CreateAddNodesRequest creates a request to invoke AddNodes API
func CreateAddNodesRequest() (request *AddNodesRequest) {
	request = &AddNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "AddNodes", "", "")
	request.Method = requests.GET
	return
}

// CreateAddNodesResponse creates a response to parse from AddNodes response
func CreateAddNodesResponse() (response *AddNodesResponse) {
	response = &AddNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
