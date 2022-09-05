package sas

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

// ModifyVpcHoneyPot invokes the sas.ModifyVpcHoneyPot API synchronously
func (client *Client) ModifyVpcHoneyPot(request *ModifyVpcHoneyPotRequest) (response *ModifyVpcHoneyPotResponse, err error) {
	response = CreateModifyVpcHoneyPotResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyVpcHoneyPotWithChan invokes the sas.ModifyVpcHoneyPot API asynchronously
func (client *Client) ModifyVpcHoneyPotWithChan(request *ModifyVpcHoneyPotRequest) (<-chan *ModifyVpcHoneyPotResponse, <-chan error) {
	responseChan := make(chan *ModifyVpcHoneyPotResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVpcHoneyPot(request)
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

// ModifyVpcHoneyPotWithCallback invokes the sas.ModifyVpcHoneyPot API asynchronously
func (client *Client) ModifyVpcHoneyPotWithCallback(request *ModifyVpcHoneyPotRequest, callback func(response *ModifyVpcHoneyPotResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVpcHoneyPotResponse
		var err error
		defer close(result)
		response, err = client.ModifyVpcHoneyPot(request)
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

// ModifyVpcHoneyPotRequest is the request struct for api ModifyVpcHoneyPot
type ModifyVpcHoneyPotRequest struct {
	*requests.RpcRequest
	SourceIp       string `position:"Query" name:"SourceIp"`
	HoneyPotAction string `position:"Query" name:"HoneyPotAction"`
	VpcId          string `position:"Query" name:"VpcId"`
}

// ModifyVpcHoneyPotResponse is the response struct for api ModifyVpcHoneyPot
type ModifyVpcHoneyPotResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyVpcHoneyPotRequest creates a request to invoke ModifyVpcHoneyPot API
func CreateModifyVpcHoneyPotRequest() (request *ModifyVpcHoneyPotRequest) {
	request = &ModifyVpcHoneyPotRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "ModifyVpcHoneyPot", "", "")
	request.Method = requests.POST
	return
}

// CreateModifyVpcHoneyPotResponse creates a response to parse from ModifyVpcHoneyPot response
func CreateModifyVpcHoneyPotResponse() (response *ModifyVpcHoneyPotResponse) {
	response = &ModifyVpcHoneyPotResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
