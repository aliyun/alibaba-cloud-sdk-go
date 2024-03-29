package sddp

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

// ExecDatamask invokes the sddp.ExecDatamask API synchronously
func (client *Client) ExecDatamask(request *ExecDatamaskRequest) (response *ExecDatamaskResponse, err error) {
	response = CreateExecDatamaskResponse()
	err = client.DoAction(request, response)
	return
}

// ExecDatamaskWithChan invokes the sddp.ExecDatamask API asynchronously
func (client *Client) ExecDatamaskWithChan(request *ExecDatamaskRequest) (<-chan *ExecDatamaskResponse, <-chan error) {
	responseChan := make(chan *ExecDatamaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecDatamask(request)
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

// ExecDatamaskWithCallback invokes the sddp.ExecDatamask API asynchronously
func (client *Client) ExecDatamaskWithCallback(request *ExecDatamaskRequest, callback func(response *ExecDatamaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecDatamaskResponse
		var err error
		defer close(result)
		response, err = client.ExecDatamask(request)
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

// ExecDatamaskRequest is the request struct for api ExecDatamask
type ExecDatamaskRequest struct {
	*requests.RpcRequest
	Data        string           `position:"Query" name:"Data"`
	FeatureType requests.Integer `position:"Query" name:"FeatureType"`
	TemplateId  requests.Integer `position:"Query" name:"TemplateId"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Lang        string           `position:"Query" name:"Lang"`
}

// ExecDatamaskResponse is the response struct for api ExecDatamask
type ExecDatamaskResponse struct {
	*responses.BaseResponse
	Data      string `json:"Data" xml:"Data"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateExecDatamaskRequest creates a request to invoke ExecDatamask API
func CreateExecDatamaskRequest() (request *ExecDatamaskRequest) {
	request = &ExecDatamaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sddp", "2019-01-03", "ExecDatamask", "sddp", "openAPI")
	request.Method = requests.POST
	return
}

// CreateExecDatamaskResponse creates a response to parse from ExecDatamask response
func CreateExecDatamaskResponse() (response *ExecDatamaskResponse) {
	response = &ExecDatamaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
