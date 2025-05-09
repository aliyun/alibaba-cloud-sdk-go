package dysmsapi

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

// DeleteSmsSign invokes the dysmsapi.DeleteSmsSign API synchronously
func (client *Client) DeleteSmsSign(request *DeleteSmsSignRequest) (response *DeleteSmsSignResponse, err error) {
	response = CreateDeleteSmsSignResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSmsSignWithChan invokes the dysmsapi.DeleteSmsSign API asynchronously
func (client *Client) DeleteSmsSignWithChan(request *DeleteSmsSignRequest) (<-chan *DeleteSmsSignResponse, <-chan error) {
	responseChan := make(chan *DeleteSmsSignResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSmsSign(request)
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

// DeleteSmsSignWithCallback invokes the dysmsapi.DeleteSmsSign API asynchronously
func (client *Client) DeleteSmsSignWithCallback(request *DeleteSmsSignRequest, callback func(response *DeleteSmsSignResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSmsSignResponse
		var err error
		defer close(result)
		response, err = client.DeleteSmsSign(request)
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

// DeleteSmsSignRequest is the request struct for api DeleteSmsSign
type DeleteSmsSignRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SignName             string           `position:"Query" name:"SignName"`
}

// DeleteSmsSignResponse is the response struct for api DeleteSmsSign
type DeleteSmsSignResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	SignName  string `json:"SignName" xml:"SignName"`
}

// CreateDeleteSmsSignRequest creates a request to invoke DeleteSmsSign API
func CreateDeleteSmsSignRequest() (request *DeleteSmsSignRequest) {
	request = &DeleteSmsSignRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dysmsapi", "2017-05-25", "DeleteSmsSign", "dysms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSmsSignResponse creates a response to parse from DeleteSmsSign response
func CreateDeleteSmsSignResponse() (response *DeleteSmsSignResponse) {
	response = &DeleteSmsSignResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
