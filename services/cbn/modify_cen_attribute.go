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

// ModifyCenAttribute invokes the cbn.ModifyCenAttribute API synchronously
func (client *Client) ModifyCenAttribute(request *ModifyCenAttributeRequest) (response *ModifyCenAttributeResponse, err error) {
	response = CreateModifyCenAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCenAttributeWithChan invokes the cbn.ModifyCenAttribute API asynchronously
func (client *Client) ModifyCenAttributeWithChan(request *ModifyCenAttributeRequest) (<-chan *ModifyCenAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyCenAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCenAttribute(request)
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

// ModifyCenAttributeWithCallback invokes the cbn.ModifyCenAttribute API asynchronously
func (client *Client) ModifyCenAttributeWithCallback(request *ModifyCenAttributeRequest, callback func(response *ModifyCenAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCenAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyCenAttribute(request)
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

// ModifyCenAttributeRequest is the request struct for api ModifyCenAttribute
type ModifyCenAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                string           `position:"Query" name:"CenId"`
	Ipv6Level            string           `position:"Query" name:"Ipv6Level"`
	Description          string           `position:"Query" name:"Description"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Version              string           `position:"Query" name:"Version"`
	ProtectionLevel      string           `position:"Query" name:"ProtectionLevel"`
	Name                 string           `position:"Query" name:"Name"`
}

// ModifyCenAttributeResponse is the response struct for api ModifyCenAttribute
type ModifyCenAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyCenAttributeRequest creates a request to invoke ModifyCenAttribute API
func CreateModifyCenAttributeRequest() (request *ModifyCenAttributeRequest) {
	request = &ModifyCenAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ModifyCenAttribute", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyCenAttributeResponse creates a response to parse from ModifyCenAttribute response
func CreateModifyCenAttributeResponse() (response *ModifyCenAttributeResponse) {
	response = &ModifyCenAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
