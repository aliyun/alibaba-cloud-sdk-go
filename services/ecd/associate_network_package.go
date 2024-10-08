package ecd

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

// AssociateNetworkPackage invokes the ecd.AssociateNetworkPackage API synchronously
func (client *Client) AssociateNetworkPackage(request *AssociateNetworkPackageRequest) (response *AssociateNetworkPackageResponse, err error) {
	response = CreateAssociateNetworkPackageResponse()
	err = client.DoAction(request, response)
	return
}

// AssociateNetworkPackageWithChan invokes the ecd.AssociateNetworkPackage API asynchronously
func (client *Client) AssociateNetworkPackageWithChan(request *AssociateNetworkPackageRequest) (<-chan *AssociateNetworkPackageResponse, <-chan error) {
	responseChan := make(chan *AssociateNetworkPackageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateNetworkPackage(request)
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

// AssociateNetworkPackageWithCallback invokes the ecd.AssociateNetworkPackage API asynchronously
func (client *Client) AssociateNetworkPackageWithCallback(request *AssociateNetworkPackageRequest, callback func(response *AssociateNetworkPackageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateNetworkPackageResponse
		var err error
		defer close(result)
		response, err = client.AssociateNetworkPackage(request)
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

// AssociateNetworkPackageRequest is the request struct for api AssociateNetworkPackage
type AssociateNetworkPackageRequest struct {
	*requests.RpcRequest
	OfficeSiteId     string `position:"Query" name:"OfficeSiteId"`
	NetworkPackageId string `position:"Query" name:"NetworkPackageId"`
}

// AssociateNetworkPackageResponse is the response struct for api AssociateNetworkPackage
type AssociateNetworkPackageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAssociateNetworkPackageRequest creates a request to invoke AssociateNetworkPackage API
func CreateAssociateNetworkPackageRequest() (request *AssociateNetworkPackageRequest) {
	request = &AssociateNetworkPackageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ecd", "2020-09-30", "AssociateNetworkPackage", "gwsecd", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAssociateNetworkPackageResponse creates a response to parse from AssociateNetworkPackage response
func CreateAssociateNetworkPackageResponse() (response *AssociateNetworkPackageResponse) {
	response = &AssociateNetworkPackageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
