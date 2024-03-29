package amqp_open

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

// CreateBinding invokes the amqp_open.CreateBinding API synchronously
func (client *Client) CreateBinding(request *CreateBindingRequest) (response *CreateBindingResponse, err error) {
	response = CreateCreateBindingResponse()
	err = client.DoAction(request, response)
	return
}

// CreateBindingWithChan invokes the amqp_open.CreateBinding API asynchronously
func (client *Client) CreateBindingWithChan(request *CreateBindingRequest) (<-chan *CreateBindingResponse, <-chan error) {
	responseChan := make(chan *CreateBindingResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateBinding(request)
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

// CreateBindingWithCallback invokes the amqp_open.CreateBinding API asynchronously
func (client *Client) CreateBindingWithCallback(request *CreateBindingRequest, callback func(response *CreateBindingResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateBindingResponse
		var err error
		defer close(result)
		response, err = client.CreateBinding(request)
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

// CreateBindingRequest is the request struct for api CreateBinding
type CreateBindingRequest struct {
	*requests.RpcRequest
	Argument        string `position:"Body" name:"Argument"`
	DestinationName string `position:"Body" name:"DestinationName"`
	SourceExchange  string `position:"Body" name:"SourceExchange"`
	BindingKey      string `position:"Body" name:"BindingKey"`
	BindingType     string `position:"Body" name:"BindingType"`
	InstanceId      string `position:"Body" name:"InstanceId"`
	VirtualHost     string `position:"Body" name:"VirtualHost"`
}

// CreateBindingResponse is the response struct for api CreateBinding
type CreateBindingResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCreateBindingRequest creates a request to invoke CreateBinding API
func CreateCreateBindingRequest() (request *CreateBindingRequest) {
	request = &CreateBindingRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("amqp-open", "2019-12-12", "CreateBinding", "onsproxy", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateBindingResponse creates a response to parse from CreateBinding response
func CreateCreateBindingResponse() (response *CreateBindingResponse) {
	response = &CreateBindingResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
