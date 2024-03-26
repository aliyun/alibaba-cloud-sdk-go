package cbn

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

// ResolveAndRouteServiceInCen invokes the cbn.ResolveAndRouteServiceInCen API synchronously
func (client *Client) ResolveAndRouteServiceInCen(request *ResolveAndRouteServiceInCenRequest) (response *ResolveAndRouteServiceInCenResponse, err error) {
	response = CreateResolveAndRouteServiceInCenResponse()
	err = client.DoAction(request, response)
	return
}

// ResolveAndRouteServiceInCenWithChan invokes the cbn.ResolveAndRouteServiceInCen API asynchronously
func (client *Client) ResolveAndRouteServiceInCenWithChan(request *ResolveAndRouteServiceInCenRequest) (<-chan *ResolveAndRouteServiceInCenResponse, <-chan error) {
	responseChan := make(chan *ResolveAndRouteServiceInCenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResolveAndRouteServiceInCen(request)
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

// ResolveAndRouteServiceInCenWithCallback invokes the cbn.ResolveAndRouteServiceInCen API asynchronously
func (client *Client) ResolveAndRouteServiceInCenWithCallback(request *ResolveAndRouteServiceInCenRequest, callback func(response *ResolveAndRouteServiceInCenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResolveAndRouteServiceInCenResponse
		var err error
		defer close(result)
		response, err = client.ResolveAndRouteServiceInCen(request)
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

// ResolveAndRouteServiceInCenRequest is the request struct for api ResolveAndRouteServiceInCen
type ResolveAndRouteServiceInCenRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	CenId                string           `position:"Query" name:"CenId"`
	Description          string           `position:"Query" name:"Description"`
	UpdateInterval       requests.Integer `position:"Query" name:"UpdateInterval"`
	Host                 string           `position:"Query" name:"Host"`
	HostRegionId         string           `position:"Query" name:"HostRegionId"`
	HostVpcId            string           `position:"Query" name:"HostVpcId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	AccessRegionIds      *[]string        `position:"Query" name:"AccessRegionIds"  type:"Repeated"`
}

// ResolveAndRouteServiceInCenResponse is the response struct for api ResolveAndRouteServiceInCen
type ResolveAndRouteServiceInCenResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateResolveAndRouteServiceInCenRequest creates a request to invoke ResolveAndRouteServiceInCen API
func CreateResolveAndRouteServiceInCenRequest() (request *ResolveAndRouteServiceInCenRequest) {
	request = &ResolveAndRouteServiceInCenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ResolveAndRouteServiceInCen", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateResolveAndRouteServiceInCenResponse creates a response to parse from ResolveAndRouteServiceInCen response
func CreateResolveAndRouteServiceInCenResponse() (response *ResolveAndRouteServiceInCenResponse) {
	response = &ResolveAndRouteServiceInCenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
